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
go mod tidy
go run main.go
```

Server runs on `http://localhost:8080`

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
| `name` | string | List name |

### Item
| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Unique identifier |
| `list_id` | int | Parent list |
| `name` | string | Item name |
| `checked` | bool | Checked or not |

---

## 📡 API Endpoints

### Lists

| Method | Route | Description |
|--------|-------|-------------|
| `GET` | `/list` | Get all lists |
| `GET` | `/list/:id` | Get a list by ID |
| `POST` | `/list` | Create a new list |
| `DELETE` | `/list/:id` | Delete a list |

### Items

| Method | Route | Description |
|--------|-------|-------------|
| `POST` | `/item/:id` | Add an item to a list |
| `DELETE` | `/item/:id` | Delete an item |
| `PUT` | `/item/:id` | Check / uncheck an item |

---

## 🧪 Usage examples

**Create a list**
```bash
curl -X POST http://localhost:8080/list \
  -H "Content-Type: application/json" \
  -d '{"name": "Weekend groceries"}'
```

**Add an item**
```bash
curl -X POST http://localhost:8080/item/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Milk"}'
```

**Check an item**
```bash
curl -X PUT http://localhost:8080/item/1
```

**Get all lists**
```bash
curl http://localhost:8080/list
```

---

## 🗺️ Roadmap

- [x] CRUD API for lists & items
- [x] Embedded SQLite database
- [ ] Share lists with family members
- [ ] Simple web interface (React / Vue) 
- [ ] Real-time notifications (WebSocket)

---

## 🛠️ Tech stack

- [Go](https://go.dev/)   backend language
- [Gin](https://github.com/gin-gonic/gin)   HTTP framework
- [GORM](https://gorm.io/) / SQLite   embedded database

---

## 📄 License

MIT