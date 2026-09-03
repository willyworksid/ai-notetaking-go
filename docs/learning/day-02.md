# Day 02 — Package, Folder & Import

**Status:** Completed

## 🎯 Goal

Memahami konsep dasar `package`, folder, `import`, serta perbedaan **exported** dan **unexported** pada Go.

Hari ini mulai belajar bagaimana Go memisahkan kode ke dalam package yang memiliki tanggung jawab tertentu.

## 🧠 Concepts

- `package`
- `package main`
- package biasa
- folder sebagai lokasi package
- `import`
- exported identifier
- unexported identifier
- aturan huruf kapital pada identifier

## 📁 Project Structure

```text
ai-notetaking-go/
├── greeting/
│   └── greeting.go
├── main.go
├── go.mod
└── ...
```

## 💻 Code

### `greeting/greeting.go`

```go
package greeting

func Hello(name string) string {
	return "Hello, " + name + "!"
}
```

### `main.go`

```go
package main

import (
	"fmt"

	"github.com/willyworksid/ai-notetaking-go/greeting"
)

func main() {
	fmt.Println(greeting.Hello("Willy"))
}
```

Output:

```text
Hello, Willy!
```

## 📦 Apa itu Package?

Package adalah cara Go mengelompokkan kode yang memiliki hubungan atau tanggung jawab tertentu.

Contoh:

```text
note/
auth/
embedding/
ai/
```

Konsep sederhananya:

```text
package main
      │
      │ import
      ▼
package greeting
      │
      ▼
    Hello()
```

## 🔗 Apa itu Import?

`import` digunakan untuk menggunakan package dari package lain.

```go
import "github.com/willyworksid/ai-notetaking-go/greeting"
```

Kemudian function exported dapat dipanggil:

```go
greeting.Hello("Willy")
```

Format sederhananya:

```text
namaPackage.namaIdentifier
```

## 🔓 Exported vs Unexported

Go tidak menggunakan keyword `public`, `private`, atau `protected` seperti PHP/Java/C#.

Go menggunakan huruf pertama identifier.

### Exported

```go
func Hello() {}
```

Huruf pertama `H` kapital, sehingga dapat digunakan dari package lain.

```go
greeting.Hello()
```

✅ Bisa.

### Unexported

```go
func secret() {}
```

Huruf pertama `s` kecil, sehingga hanya dapat digunakan dari package yang sama.

```go
greeting.secret()
```

❌ Tidak bisa dari `package main`.

## 🔐 Unexported ≈ Private

Secara konsep, **unexported mirip dengan private**, tetapi batas aksesnya berbeda.

Go menggunakan **package sebagai batas**, bukan class.

Misalnya:

```text
greeting/
├── greeting.go
├── validation.go
└── helper.go
```

Jika semuanya menggunakan:

```go
package greeting
```

maka function unexported seperti:

```go
func secret() {}
```

tetap dapat digunakan oleh file-file lain dalam package `greeting`.

Jadi:

> **Unexported = hanya terlihat di dalam package yang sama.**

Bukan hanya di file yang sama.

## ❓ Kenapa `greeting.secret()` Error?

Misalnya:

```go
package greeting

func secret() string {
	return "This is secret"
}
```

Kemudian dari `main`:

```go
fmt.Println(greeting.secret())
```

akan error karena `secret` adalah **unexported**.

Sebaliknya:

```go
func Hello() string {
	return "Hello!"
}
```

bisa dipanggil:

```go
greeting.Hello()
```

karena `Hello` adalah **exported**.

## 🧩 `package main`

`package main` memiliki fungsi khusus.

Package yang menggunakan:

```go
package main
```

dan memiliki:

```go
func main()
```

dapat menghasilkan executable program.

Sedangkan package seperti:

```go
package greeting
```

biasanya digunakan sebagai package yang dipakai oleh package lain.

## 🤔 Why?

### Kenapa Go menggunakan package?

Agar kode tidak menumpuk dalam satu file atau satu tempat.

Daripada semua fitur berada di `main.go`, kode dapat dipisah berdasarkan tanggung jawab.

Contoh target project kita nantinya:

```text
note/
auth/
embedding/
ai/
```

### Kenapa menggunakan huruf kapital?

Go membuat aturan exported/unexported sangat sederhana:

```go
Hello()  // exported
hello()  // unexported
```

Tidak perlu keyword `public`, `private`, atau `protected`.

## 🧪 Exercise

### Exercise 1 — Greeting

Buat:

```go
func Hello(name string) string {
	return "Hello, " + name + "!"
}
```

Kemudian panggil dari `main`.

Target:

```text
Hello, Willy!
```

### Exercise 2 — Unexported Function

Tambahkan:

```go
func secret() string {
	return "This is secret"
}
```

Coba panggil dari `main`:

```go
greeting.secret()
```

Amati error yang muncul, kemudian hapus percobaan tersebut.

### Exercise 3 — Pemahaman

Jawab dengan kata-kata sendiri:

1. Apa itu package?
2. Apa fungsi `import`?
3. Apa perbedaan `Hello()` dan `hello()`?
4. Apakah unexported berarti hanya bisa digunakan dalam satu file?
5. Apa yang membuat `package main` berbeda?

## ✅ Definition of Done

- [x] Memahami package
- [x] Memahami folder package
- [x] Memahami import
- [x] Bisa membuat package `greeting`
- [x] Bisa menggunakan function dari package lain
- [x] Memahami exported
- [x] Memahami unexported
- [x] Memahami bahwa batas unexported adalah package
- [x] Memahami dasar `package main`

## 📝 Notes

Konsep penting Day 02:

```text
Huruf kapital
     ↓
Exported
     ↓
Bisa digunakan package lain
```

```text
Huruf kecil
     ↓
Unexported
     ↓
Hanya bisa digunakan package yang sama
```

Cara berpikir sederhana:

> **Package adalah boundary kode di Go.**

## 🔄 Old Code → New Code

### Day 01

```text
ai-notetaking-go/
├── main.go
└── go.mod
```

Semua kode masih sederhana dan berada di `package main`.

### Day 02

```text
ai-notetaking-go/
├── greeting/
│   └── greeting.go
├── main.go
└── go.mod
```

Sekarang sudah ada package terpisah dan `main` menggunakan package tersebut melalui `import`.

## 🌱 Hubungan dengan Project AI Note Taking

Konsep package yang dipelajari hari ini akan menjadi dasar ketika project mulai berkembang.

Nantinya kita dapat memiliki package seperti:

```text
note/
auth/
embedding/
ai/
```

Namun **belum perlu membuat semuanya sekarang**. Kita akan memperkenalkannya sedikit demi sedikit.

## 🔀 Git

Branch:

```bash
git checkout -b day-02
```

Setelah latihan selesai:

```bash
gofmt -w main.go greeting/greeting.go
go run .
git status
```

Commit:

```bash
git add .
git commit -m "day-02: learn package and import"
```

Push:

```bash
git push -u origin day-02
```

Setelah Day 02 benar-benar selesai, branch dapat di-merge ke `main`.

## 🎯 Kesimpulan

Day 02 memperkenalkan salah satu konsep paling penting dalam Go:

> **Package adalah batas organisasi dan akses kode.**

Dengan aturan sederhana:

```go
Hello() // exported
hello() // unexported
```

kita sudah memiliki mekanisme dasar untuk menentukan kode mana yang boleh digunakan oleh package lain.
