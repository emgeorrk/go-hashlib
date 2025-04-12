# go-hashlib 🚀

A collection of lightweight and efficient hash function implementations in Go, designed for use in IoT environments. These functions are optimized for minimal memory usage, low power consumption, and fast computation, making them ideal for embedded systems and resource-constrained devices.

## ✨ Features

- Multiple hash algorithms optimized for lightweight use:
    - `HashOne`
    - `Photon`
    - `Quark`
    - `Spongent`
    - `Tjuilik`
- Simple and consistent API
- Suitable for embedded systems and IoT devices
- Pure Go implementation, no dependencies

## 📦 Installation

Use `go get` to add the module to your project:

```bash
go get github.com/emgeorrk/go-hashlib
```

Then import it in your code:

```go
import "github.com/emgeorrk/go-hashlib"
```

## 🚀 Example Usage

Here's how to hash a string using the `HashOne` algorithm:

```go
package main

import (
    "fmt"

    "github.com/emgeorrk/go-hashlib"
)

func main() {
    h := hashlib.NewHashOne()

    str := "Hello World!"

    hashed := h.Hash([]byte(str))

    fmt.Printf("%s -> %v\n", str, hashed)
}
```

All hash functions implement the same interface, so you can easily switch between them.

## 🧪 Available Hash Functions

Each function implements a `Hash([]byte) []byte` method:

- `HashOne` – a minimalistic hash for tiny data
- `Photon` – compact and secure sponge-based construction
- `DQuark` – lightweight permutation-based hash
- `Spongent` – ultra-light sponge construction for constrained environments
- `Tjuilik` – inspired by Saturnin hash designs

## 📁 Project Structure

```
.
├── hash/                # Core implementations of hash functions
├── examples/            # Example usages and tests
├── tests/               # Unit tests for each algorithm
├── internal/utils/      # Bit-level utilities
└── api.go               # Shared interfaces and types
```

## 🛠 Use Cases

- Message integrity in sensor networks
- Lightweight authentication for edge devices
- Efficient hashing for telemetry data

## 📄 License

[MIT](LICENSE)

---
