package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -tags=linux -output-dir=gen -go-package=gen counter bpf/counter.bpf.c
