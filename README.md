# 📋 OpenList

OpenList is an ultra-lightweight, self-hosted shared checklist application. Designed to run on anything (even a potato 🥔), it's perfect for managing family grocery lists or shared tasks without the weight of modern bloatware, tracking, or absurd subscriptions.

---

## 🎯 The Philosophy: "No Black Box"

OpenList was born from a desire to learn Go by building something genuinely useful. It follows a strict minimalist approach:

- **Zero-Bloat Frontend:** Powered by Pico.css, no heavy frameworks, no external fonts. Just HTML, Vanilla JS, and optimized SVGs.
- **High Performance:** Backend in Go with an optimized SQLite database (WAL Mode).
- **Full Control:** Written line by line to understand exactly what happens under the hood.

---

## 🧱 Optimized Build & Footprint

One of the main challenges of OpenList is to achieve the smallest possible binary footprint.

| Build Type       | Size    | Status        |
|------------------|---------|---------------|
| Standard Build   | ~15MB   | ❌ Too heavy  |
| Optimized Build  | 6.5MB   | ⚠️ Current    |
| Target Build     | < 1.5MB | 🎯 Goal       |

### Compilation Magic

We use advanced Go compiler flags to strip unnecessary metadata and ensure portability:

```bash
# To get the optimized production binary:
make
```

**Flags used:**

- `CGO_ENABLED=0` — Produces a 100% static binary (runs on any Linux distro).
- `-tags "netgo osusergo static_build"` — Forces Go-native implementations.
- `-trimpath` — Removes local filesystem paths for privacy and reproducibility.
- `-ldflags "-s -w"` — Strips debug information and symbol tables.
- **UPX Compression** — Automatically applies `--ultra-brute` compression if UPX is installed.

---

## 🚀 Getting Started

```bash
git clone https://github.com/coockieHunt/OpenList.git
cd OpenList
cp .env.exemple .env
go mod tidy
go run main.go
```

- API runs on `http://localhost:8080`
- Web UI runs on `http://localhost:4000`

---

## 📡 API Endpoints Summary

### Lists

| Method | Route         | Description      |
|--------|---------------|------------------|
| GET    | /api/list     | Get all lists    |
| POST   | /api/list     | Create a new list|
| DELETE | /api/list/:id | Delete a list    |

### Items

| Method | Route                        | Description              |
|--------|------------------------------|--------------------------|
| POST   | /api/item/:idList            | Add item                 |
| PUT    | /api/item/:idList/:idItem    | Toggle validation (Sync) |
| DELETE | /api/item/:idList/:idItem    | Remove item              |

---

## 🗺️ Roadmap: The Journey to 1MB

- [x] CRUD API & Embedded SQLite (WAL Mode)
- [x] Simple Web Interface (Pico.css)
- [ ] Ultra-basic User System: Just a password + persistent token.
- [ ] Family Sharing: Secure list sharing via unique hash links.
- [ ] Docker "One-Command": Multi-stage build based on scratch image.
- [ ] PWA Support: Installable app with manifest and service worker.
- [ ] ORM Removal: Switching from GORM to pure SQL to halve the binary size.

---

## 🛠️ Tech Stack

| Layer     | Technology                                        |
|-----------|---------------------------------------------------|
| Backend   | Go (Fast, statically typed)                       |
| Framework | Gin (Lightweight routing)                         |
| Database  | SQLite with WAL mode (Zero-config, high-performance) |
| Frontend  | Pico.css (Minimalist CSS) & Vanilla JS            |
| Icons     | Tabler Icons (Only 3 custom SVGs)                 |

---

## 🤝 Contributing

Contributions are more than welcome! Whether it's for SQL optimization, UI/UX improvements for "Supermarket usage" (large touch targets), or Dockerization, feel free to open a PR.

---

**License:** MIT  
**Hand-crafted by** Jonathan (Coockie)
