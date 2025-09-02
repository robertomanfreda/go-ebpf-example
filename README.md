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
├── gen/                 # Auto-generated Go bindings and object files
├── build/               # Compiled binary output
├── main.go              # Entry point for Go app
├── gen.go               # go:generate directive to build eBPF code
├── Makefile             # Commands for build, run, and generate
├── go.mod/sum           # Go dependencies
├── README.md            # This file
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
git clone https://github.com/your-username/go-ebpf-example
cd go-ebpf-example
```

### 2. Generate, build and execute

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

## 📂 Useful Resources

* [eBPF.io](https://ebpf.io)
* [Cilium eBPF Library](https://github.com/cilium/ebpf)
* [XDP Project & Tutorials](https://xdp-project.net)
* [Go eBPF](https://ebpf-go.dev/guides/getting-started)

---

## 📄 License

This project is licensed under **Dual MIT/GPL**.
