# AI Note Taking Backend — Go

Backend **AI Note Taking** yang dibangun ulang menggunakan Go.

Project ini sekaligus menjadi **learning project Go** dengan target 30–60 menit per hari.

## Tujuan

- Belajar Go melalui project nyata.
- Memahami idiomatic Go, bukan hanya syntax.
- Membangun REST API dengan Fiber.
- Menggunakan PostgreSQL.
- Menerapkan arsitektur yang sederhana dan scalable.
- Mengintegrasikan AI, embedding, vector search, dan RAG.
- Menerapkan asynchronous processing dan worker.

## Learning Approach

Setiap hari memiliki:
1. Satu branch Git (`day-XX`)
2. Satu dokumentasi (`docs/learning/day-XX.md`)
3. 1–2 konsep utama
4. Latihan kecil
5. Hasil yang bisa dijalankan
6. Commit yang jelas

## Learning Progress

### Phase 1 — Go Fundamentals

- [ ] Day 01 — Go Project, Module & Hello World
- [ ] Day 02 — Package, Folder & Import
- [ ] Day 03 — Struct & Method
- [ ] Day 04 — Function, Pointer & Value
- [ ] Day 05 — Interface
- [ ] Day 06 — Error Handling
- [ ] Day 07 — Context
- [ ] Day 08 — Goroutine
- [ ] Day 09 — Channel

### Phase 2 — Go Backend

- [ ] Day 10 — Project Structure
- [ ] Day 11 — Fiber
- [ ] Day 12 — Routing
- [ ] Day 13 — Middleware
- [ ] Day 14 — DTO & Request/Response
- [ ] Day 15 — Validation
- [ ] Day 16 — PostgreSQL
- [ ] Day 17 — Repository
- [ ] Day 18 — Service
- [ ] Day 19 — Dependency Injection
- [ ] Day 20 — CRUD API

### Phase 3 — AI Note Taking

- [ ] Day 21 — AI Client
- [ ] Day 22 — Embedding
- [ ] Day 23 — Vector Search
- [ ] Day 24 — RAG
- [ ] Day 25 — Event
- [ ] Day 26 — Queue
- [ ] Day 27 — Worker
- [ ] Day 28 — AI Chat
- [ ] Day 29 — Testing
- [ ] Day 30 — Docker & Production

> Jadwal fleksibel. Jika satu konsep membutuhkan lebih dari satu hari, kita pecah.

## Target Architecture

Menggunakan **Modular Monolith**.

```text
HTTP
  ↓
Handler
  ↓
Service
  ↓
Repository / External Client
  ↓
Database / External Service
```

Untuk proses asynchronous:

```text
API
 ↓
Event / Queue
 ↓
Worker
 ↓
Embedding / AI Processing
```

Target struktur:

```text
ai-notetaking-go/
├── cmd/
│   ├── api/
│   └── worker/
├── internal/
│   ├── note/
│   ├── notebook/
│   ├── embedding/
│   ├── ai/
│   ├── auth/
│   ├── event/
│   └── infrastructure/
├── migrations/
├── docs/
│   ├── learning/
│   ├── architecture/
│   └── decisions/
├── go.mod
└── README.md
```

Struktur tersebut adalah **target**, bukan sesuatu yang harus langsung diisi pada Day 01.

## Design Principles

- Simple over complex.
- Idiomatic Go.
- Manual Dependency Injection.
- Separation of concerns.
- Interface hanya ketika memang diperlukan.
- Business logic tidak bergantung langsung pada Fiber.
- External AI/embedding diisolasi melalui client/service.
- Async processing untuk pekerjaan yang tidak perlu menahan HTTP request.
- Testability tanpa over-engineering.

### Pattern yang direncanakan

| Pattern | Penggunaan |
|---|---|
| Modular Monolith | Struktur aplikasi |
| Layered Architecture | Handler → Service → Repository |
| Repository Pattern | Akses database |
| Dependency Injection | Wiring dependency |
| DTO | API contract |
| Adapter | External AI/embedding |
| Event-driven | Async processing |
| Strategy | Jika ada beberapa AI provider |

Pattern tidak diterapkan hanya demi menggunakan pattern.

## Learning Documentation

Dokumentasi harian berada di `docs/learning/`.

Format setiap hari:

```text
Goal
Concepts
Project / Code
What I Learned
Why?
Exercise
Definition of Done
Notes
Git
Old Code → New Code
```

## Architecture Decisions

Keputusan arsitektur penting disimpan sebagai ADR di `docs/decisions/`.

Contoh:

```text
001-modular-monolith.md
002-manual-dependency-injection.md
003-repository-pattern.md
```

Tujuannya mencatat **apa yang dipilih dan mengapa**.

## Git Workflow

Setiap hari menggunakan branch:

```bash
git checkout -b day-01
```

Contoh commit:

```bash
git add .
git commit -m "day-01: initialize go project"
git push origin day-01
```

Branch menjadi checkpoint pembelajaran.

```text
main
 ├── day-01
 ├── day-02
 ├── day-03
 └── ...
```

`main` diusahakan selalu stabil.

## Daily Learning Rule

Target: **30–60 menit per hari**.

Satu hari dianggap berhasil jika:
- memahami konsep utama;
- menyelesaikan latihan;
- project dapat dijalankan jika relevan;
- dokumentasi selesai;
- perubahan di-commit.

**Pemahaman lebih penting daripada jumlah kode.**

## Rebuild Strategy

Source code lama `ai-notetaking-be` digunakan sebagai referensi.

```text
Old Code
   ↓
Understand
   ↓
Identify Problem / Pattern
   ↓
Learn Go Concept
   ↓
Design New Version
   ↓
Implement
   ↓
Test
   ↓
Document
```

Tujuannya bukan menerjemahkan kode lama satu-per-satu, tetapi membangun ulang dengan pendekatan yang lebih idiomatik Go.

## Starting Point

**Day 01 — Go Project, Module & Hello World**

Day 01 sengaja sederhana:

```text
ai-notetaking-go/
├── go.mod
└── main.go
```

Struktur aplikasi lengkap akan tumbuh setelah konsep dasar Go dipahami.

## Status

**Current Day:** Day 01  
**Status:** Project initialized
