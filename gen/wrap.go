//go:build linux

package gen

import (
	"fmt"
	"net"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// Counter incapsula gli oggetti eBPF generati e il link XDP.
type Counter struct {
	objs counterObjects
	l    link.Link
	ifi  *net.Interface
}

// Start inizializza rlimit, carica gli oggetti eBPF e attacca il programma XDP
// all'interfaccia specificata. Restituisce un handle Counter che puoi usare
// per leggere il contatore o chiudere/sganciare tutto.
func Start(ifname string) (*Counter, error) {
	// Per kernel < 5.11
	if err := rlimit.RemoveMemlock(); err != nil {
		return nil, fmt.Errorf("remove memlock: %w", err)
	}

	ifi, err := net.InterfaceByName(ifname)
	if err != nil {
		return nil, fmt.Errorf("get interface %s: %w", ifname, err)
	}

	var objs counterObjects
	if err := loadCounterObjects(&objs, nil); err != nil {
		return nil, fmt.Errorf("load eBPF objects: %w", err)
	}

	l, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.CountPackets,
		Interface: ifi.Index,
	})
	if err != nil {
		objs.Close()
		return nil, fmt.Errorf("attach XDP: %w", err)
	}

	return &Counter{objs: objs, l: l, ifi: ifi}, nil
}

// Count legge il valore corrente dalla mappa PktCount (chiave fissa 0).
func (c *Counter) Count() (uint64, error) {
	var count uint64
	if err := c.objs.PktCount.Lookup(uint32(0), &count); err != nil {
		return 0, fmt.Errorf("map lookup: %w", err)
	}
	return count, nil
}

// Interface restituisce l'interfaccia su cui è agganciato XDP.
func (c *Counter) Interface() *net.Interface {
	return c.ifi
}

// Close sgancia XDP e rilascia tutte le risorse eBPF.
func (c *Counter) Close() error {
	var firstErr error
	if c.l != nil {
		if err := c.l.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	// Close rilascia mappe/programmi generati.
	if err := c.objs.Close(); err != nil && firstErr == nil {
		firstErr = err
	}
	return firstErr
}
