# Day 01 — Go Project, Module & Hello World

> Status: Completed

## 🎯 Goal

Memahami struktur paling dasar project Go dan bagaimana sebuah program Go dijalankan.

Pada hari pertama, fokus bukan membuat API, tetapi memahami lifecycle sederhana sebuah project Go:

```text
go.mod
  ↓
package main
  ↓
func main()
  ↓
go run / go build
```

## 📚 Concepts

* Go Module
* `go.mod`
* Package
* `package main`
* `func main()`
* `import`
* Standard Library
* `go run`
* `go build`

## 💻 Project / Code

Struktur awal:

```text
ai-notetaking-go/
├── go.mod
├── main.go
├── README.md
├── cmd/
├── internal/
├── migrations/
└── docs/
```

Untuk Day-01, struktur aplikasi belum digunakan. Fokus hanya pada:

```text
ai-notetaking-go/
├── go.mod
└── main.go
```

### `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("AI Note Taking Backend")
    fmt.Println("Learning Go Day 01")
    fmt.Println("Built with Go")
}
```

Program dijalankan dengan:

```bash
go run .
```

Untuk membuat executable:

```bash
go build .
```

## 🔍 What I Learned

### 1. Go Module

`go.mod` mendefinisikan module Go dan menjadi dasar dependency management project.

Contoh:

```go
module github.com/username/ai-notetaking-go
```

### 2. Package

Setiap file Go berada di dalam sebuah package.

Pada executable utama digunakan:

```go
package main
```

### 3. `func main()`

`main()` adalah entry point untuk executable Go.

Program mulai dieksekusi dari:

```go
func main() {
    ...
}
```

### 4. Import

Package lain dapat digunakan melalui `import`.

Contohnya:

```go
import "fmt"
```

`fmt` merupakan bagian dari standard library Go.

### 5. `go run`

```bash
go run .
```

digunakan untuk menjalankan aplikasi secara langsung melalui proses build/run.

### 6. `go build`

```bash
go build .
```

digunakan untuk melakukan build dan menghasilkan executable.

## ❓ Why?

### Kenapa `package main`?

Karena project Day-01 merupakan executable application, bukan library.

### Kenapa ada `func main()`?

Go membutuhkan entry point untuk menentukan dari mana executable mulai berjalan.

### Kenapa menggunakan `go.mod`?

Aplikasi Go menggunakan module sebagai identitas dan dasar pengelolaan dependency.

### Kenapa belum menggunakan Fiber?

Karena tujuan Day-01 adalah memahami dasar project Go terlebih dahulu.

Fiber akan diperkenalkan setelah konsep dasar bahasa Go dan package dipahami.

## 🧪 Exercise

### Exercise 1

Buat program yang menghasilkan:

```text
AI Note Taking Backend
Learning Go Day 01
Built with Go
```

### Exercise 2

Buat variable:

```go
name := "Willy"
```

kemudian tampilkan:

```text
Hello, Willy!
```

### Exercise 3

Jalankan:

```bash
go run .
```

kemudian:

```bash
go build .
```

dan jalankan executable hasil build.

## ✅ Definition of Done

* [x] Go terinstall dan dapat dijalankan melalui terminal
* [x] Git repository dibuat
* [x] Branch `day-01` dibuat
* [x] `go.mod` berhasil dibuat
* [x] `main.go` berhasil dibuat
* [x] Memahami `package main`
* [x] Memahami `func main()`
* [x] Memahami dasar `import`
* [x] `go run .` berhasil
* [x] `go build .` berhasil
* [x] Exercise selesai
* [x] Dokumentasi Day-01 diperbarui
* [x] Perubahan di-commit ke Git

## 📝 Notes

Hal terpenting dari Day-01 bukan syntax `Hello World`, tetapi memahami bahwa aplikasi Go dimulai dari sebuah **module** dan executable menggunakan `package main` dengan `func main()` sebagai entry point.

Project sengaja belum menggunakan struktur arsitektur final.

Struktur akan berkembang secara bertahap sesuai materi yang dipelajari.

## 🔀 Git

Branch:

```bash
git checkout -b day-01
```

Commit:

```bash
git add .
git commit -m "day-01: initialize go project"
```

Checkpoint:

```text
day-01
└── Go project dapat dijalankan
```

## 🔄 Old Code → New Code

Belum ada pada Day-01.

Perbandingan dengan source `ai-notetaking-be` akan mulai dilakukan ketika materi sudah masuk ke struktur aplikasi, Fiber, service, repository, dan dependency injection.
