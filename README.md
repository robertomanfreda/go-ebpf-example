# go-ebpf-example

> Count incoming packets on a network interface using eBPF/XDP and Go

## ✨ Overview

This project demonstrates how to use **eBPF** in combination with **XDP (eXpress Data Path)** and the **Go programming
language** to count packets received on a specific network interface.

It leverages the [cilium/ebpf](https://github.com/cilium/ebpf) library to load, manage, and interact with eBPF programs
and maps from Go code.

---

## ⚙️ Features

* Loads an XDP eBPF program that increments a counter for every received packet
* Reads the packet counter in real time using Go
* Automatically cleans up eBPF resources on exit
* Uses `go generate` to compile eBPF C code into Go bindings and objects

---

## 📁 Project Structure

```
.
├── bpf/                 # eBPF C code (XDP program)
├── gen/                 # Auto-generated Go bindings, object files (available after `make`) and wrapper
├── build/               # Compiled binary output (available after `make`)
├── cmd/                 # Entry point for Go application (main.go)
├── gen.go               # go:generate directive to build eBPF code
├── Makefile             # Commands for build, run, and generate
├── go.mod/sum           # Go dependencies
```

---

## 🔧 Requirements

* Linux with kernel **>= 5.4** (preferably >= 5.11)
* **Go 1.20+**
* **clang/llvm** for compiling eBPF C code
* `sudo` privileges to attach XDP programs

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/robertomanfreda/go-ebpf-example
cd go-ebpf-example
```

### 2. Generate

```bash
make generate
```

### 3. Build

```bash
make build
```

### 4. Execute

```bash
make run
```

### 5. All in one

```bash
make
```

You should see output like:

```
Counting incoming packets on enp3s0..
Received 52 packets
Received 78 packets
...
```

> ℹ The network interface is currently hardcoded as `enp3s0` in `main.go`. You can change it as needed.

---

## 🧠 How It Works

1. **XDP Program (`bpf/counter.bpf.c`)**: Attached at a low level to the NIC, increments a map every time a packet is
   received.

2. **Go Code (`main.go`)**:

    * Loads the eBPF program onto the network interface
    * Reads the packet count every second
    * Logs the number of received packets
    * Handles graceful exit on `SIGINT`

3. **eBPF Map**:

    * Type: `BPF_MAP_TYPE_ARRAY`
    * Key: always `0`
    * Value: `__u64` packet counter

---

## 🔹 Customization

To use a different network interface, modify this line in `main.go`:

```go
c, err := gen.Start("enp3s0")

```

--- 

## 🗒️ Notes

As reported on [ebpf-go.dev](https://ebpf-go.dev/guides/getting-started/#whats-next)

1. Use clang --version to check which version of LLVM you have installed. Refer to your distribution's package index to
   finding the right packages to install, as this tends to vary wildly across distributions. Some distributions ship
   clang
   and llvm-strip in separate packages.
2. For Debian/Ubuntu, you'll typically need libbpf-dev. On Fedora, it's libbpf-devel.
3. On AMD64 Debian/Ubuntu, install linux-headers-amd64. On Fedora, install kernel-devel.
   On Debian, you may also need ln -sf /usr/include/asm-generic/ /usr/include/asm since the example expects to find <
   asm/types.h>.

---

## 📂 Useful Resources

* [eBPF.io](https://ebpf.io)
* [Cilium eBPF Library](https://github.com/cilium/ebpf)
* [XDP Project & Tutorials](https://xdp-project.net)
* [Go eBPF](https://ebpf-go.dev/guides/getting-started)

---

## 📄 License

This project is licensed under **MIT License**.
