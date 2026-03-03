# 📋 OpenList

A lightweight API to manage shared checklists  like grocery lists   with your family.  
Built with **Go**, **Gin** and **SQLite**.

---

## 🎯 Why this project?

OpenList is a **learning project**. The main goal is to explore and get comfortable with the **basics of the Go language**   a fast, simple, and statically typed language that I wanted to understand by building something real and useful.

Go intrigued me because of its simplicity, its performance, and its opinionated approach to writing clean code. Rather than following tutorials, I decided to learn by doing: building a small REST API from scratch with a real use case. This project covers core Go concepts like structs, interfaces, packages, error handling, and working with external libraries.

If you're also learning Go, feel free to explore the code   it's intentionally kept simple and readable.

---

## 🚀 Getting started

```bash
git clone https://github.com/your-user/OpenList.git
cd OpenList
cp .env.exemple .env
go mod tidy
go run main.go
```

Server runs on `http://localhost:8080`

### Smallest build

```bash
make
```

This build uses:
- `-tags "netgo osusergo static_build"` (fully static binary, no CGO)
- `-trimpath`
- `-ldflags "-s -w"`
- `CGO_ENABLED=0`

The `compress` step runs automatically after build (if UPX is installed):

```bash
make build   # build only
make         # build + UPX compression (ultra-brute + lzma)
```

> **Note:** UPX must be installed for compression. If not found, only the uncompressed binary is produced.

---

## ⚙️ Configuration

After copying `.env.exemple` to `.env`, you can configure:

| Variable | Default | Description |
|---|---|---|
| `API_PORT` | `8080` | Port used by the API server |
| `WEB_PORT` | `4000` | Port used by the web UI server |
| `API_URL` | `http://localhost:8080/api` | Base URL used by frontend API calls |

---

## 📁 Project structure

```
OpenList/
├── main.go          # Entry point & routes
├── routes/          # HTTP handlers
├── sqlite/          # Models & database init
└── README.md
```

---

## 🗄️ Models

### List
| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Unique identifier |
| `title` | string | List title |

### Item
| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Unique identifier |
| `list_id` | int | Parent list |
| `name` | string | Item name |
| `quantity` | int | Item quantity |
| `validated` | bool | Checked or not |

---

## 📡 API Endpoints

### Lists

| Method | Route | Description |
|--------|-------|-------------|
| `GET` | `/api/list` | Get all lists |
| `GET` | `/api/list/:idList` | Get a list by ID |
| `POST` | `/api/list` | Create a new list |
| `DELETE` | `/api/list/:idList` | Delete a list |

### Items

| Method | Route | Description |
|--------|-------|-------------|
| `POST` | `/api/item/:idList` | Add an item to a list |
| `DELETE` | `/api/item/:idList/:idItem` | Delete an item |
| `PUT` | `/api/item/:idList/:idItem` | Check / uncheck an item |

---

## 🧪 Usage examples

**Create a list**
```bash
curl -X POST http://localhost:8080/api/list \
  -H "Content-Type: application/json" \
  -d '{"title": "Weekend groceries"}'
```

**Add an item**
```bash
curl -X POST http://localhost:8080/api/item/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Milk", "quantity": 1, "validated": false}'
```

**Check an item**
```bash
curl -X PUT http://localhost:8080/api/item/1/2
```

**Get all lists**
```bash
curl http://localhost:8080/api/list
```

---

## 🗺️ Roadmap

- [x] CRUD API for lists & items
- [x] Embedded SQLite database
- [x] Simple web interface (html/js) 
- [ ] Share lists with family members
- [ ] GlowUp ui web interface (html/js) 
- [ ] Real-time notifications (WebSocket)

---

## 🛠️ Tech stack

- [Go](https://go.dev/)   backend language
- [Gin](https://github.com/gin-gonic/gin)   HTTP framework
- [GORM](https://gorm.io/)  SQLite   embedded database

---

## 🚀 Project Status & Contributions
This project is currently under active development (Work In Progress). We are focusing on building a high-performance, lightweight architecture with Go and Vanilla JS.

Contributions are more than welcome! Pull Requests (PRs) are open for anyone looking to help with code optimization, UI/UX improvements with Pico.css, or general refactoring. Feel free to fork the repo and submit your ideas!

Current Footprint: 6.5M (build) (Targeting BUILD < THE SMALLEST POSSIBLE)

## powered by
- picos css : https://picocss.com/
- icon : https://tabler.io/icons

## 📄 License

MIT