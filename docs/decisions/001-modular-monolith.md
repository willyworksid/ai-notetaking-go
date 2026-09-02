# ADR-001 — Modular Monolith

## Status

Accepted

## Context

Project dikembangkan oleh tim kecil dan membutuhkan REST API, AI, embedding,
search, dan asynchronous processing.

## Decision

Menggunakan **Modular Monolith** sebagai architecture awal.

## Why

- Sederhana untuk tim kecil.
- Satu codebase mudah dikembangkan dan di-debug.
- Deployment lebih sederhana.
- Domain tetap dapat dipisahkan secara modular.
- Worker dapat dijalankan sebagai process terpisah.
- Tidak memaksakan kompleksitas microservices sejak awal.

## Consequences

Jika suatu domain nantinya membutuhkan scaling atau deployment independen,
domain tersebut dapat dipisahkan kemudian.
