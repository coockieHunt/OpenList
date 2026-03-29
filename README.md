# 📋 OpenList

> **Ultra-lightweight, self-hosted shared checklist — because your grocery list doesn't need a SaaS subscription.**

---

## 🎯 Philosophy : "No Black Box"

OpenList est né d'une volonté d'apprendre Go en construisant quelque chose de réellement utile. Le projet suit une approche minimaliste stricte :

- **Zéro-bloat frontend** Propulsé par [Pico.css](https://picocss.com/), sans framework lourd, sans polices externes. Juste du HTML, du Vanilla JS et des SVGs optimisés.
- **Performance** Backend Go avec SQLite optimisé en mode WAL.
- **Contrôle total** Écrit ligne par ligne pour comprendre exactement ce qui se passe sous le capot.

---

## 🧱 Build & Empreinte optimisés

Chaque octet compte.

| Type de build | Taille | Statut |
| --- | --- | --- |
| Build standard | ~25 MB | ❌ Trop lourd |
| Optimisé + UPX | 6.5 MB | ✅ Actuel |

### Magie de compilation

Le `Makefile` compresse le binaire :

```bash
make
```

**Ce qui se passe:**

- `CGO_ENABLED=0` & `static_build` Garantit que le binaire tourne sur n'importe quelle distro Linux sans dépendances.
- `-ldflags "-s -w"` Supprime les infos de debug et les tables de symboles.
- **UPX Ultra-Brute** Compresse l'exécutable à son minimum absolu.

---

## 🚀 Démarrage rapide

### Installation

```bash
git clone https://github.com/coockieHunt/OpenList.git
cd OpenList
cp .env.exemple .env
go mod tidy
make
```

### Administration

Pour définir ou changer le mot de passe administrateur :

```bash
./Openlist setPassword "votre_nouveau_mot_de_passe"
```

### Lancer l'application

```bash
./Openlist
```

| Service | URL |
| --- | --- |
| API | http://localhost:8080 |
| Web UI | http://localhost:4000 |

---

## 📡 API Endpoints

### 🔐 Authentification

| Méthode | Route | Description |
| --- | --- | --- |
| POST | `/auth/login` | Connexion administrateur |
| POST | `/auth/logout` | Déconnexion |
| GET | `/auth/status` | Vérifier le statut de session |
| POST | `/auth/change-password` | Changer le mot de passe |

**Login** `POST /auth/login`

```json
{
  "username": "admin",
  "password": "adminpassword"
}
```

**Change password** `POST /auth/change-password`

```json
{
  "current_password": "currentpassword",
  "new_password": "newpassword"
}
```

### 📋 Listes

| Méthode | Route | Description |
| --- | --- | --- |
| GET | `/api/list` | Récupérer toutes les listes |
| POST | `/api/list` | Créer une nouvelle liste |
| DELETE | `/api/list/:id` | Supprimer une liste |

### ✅ Items

| Méthode | Route | Description |
| --- | --- | --- |
| POST | `/api/item/:idList` | Ajouter un item |
| PUT | `/api/item/:idList/:idItem` | Basculer la validation (Sync) |
| DELETE | `/api/item/:idList/:idItem` | Supprimer un item |

---

## 🗺️ Roadmap : La Route vers 1 MB

- [x] API CRUD & SQLite embarqué (mode WAL)
- [x] Interface Web simple (Pico.css)
- [x] Sécurisée par un utilisateur unique
- [x] CLI Admin (commande `setPassword`)
- [ ] **Family Sharing** Partage de listes sécurisé via des liens hash uniques
- [ ] **Docker "One-Command"** Build multi-stage basé sur une image `scratch`
- [ ] **Suppression ORM** Passage de GORM à du SQL pur pour atteindre la cible 1 MB

---

## 🛠️ Stack Technique

| Couche | Technologie |
| --- | --- |
| Backend | Go (rapide, statiquement typé) |
| Framework | Gin (routing léger) |
| Base de données | SQLite + mode WAL (zéro config, haute perf) |
| Frontend | Pico.css (CSS minimaliste) & Vanilla JS |
| Icônes | Tabler Icons (SVGs triés sur le volet) |

---

## 🤝 Contribuer

Les contributions sont les bienvenues ! Que ce soit pour l'optimisation SQL, des améliorations UI/UX pour une utilisation en supermarché (grandes zones tactiles) ouvrez une PR.

---

**Licence : MIT**  
Fait à la main par Jonathan (Coockie) 🍪

---

---

# 📋 OpenList

> **Ultra-lightweight, self-hosted shared checklist — because your grocery list doesn't need a SaaS subscription.**

---

## 🎯 Philosophy : "No Black Box"

OpenList was born from a desire to learn Go by building something genuinely useful. The project follows a strict minimalist approach:

- **Zero-bloat frontend** Powered by [Pico.css](https://picocss.com/), no heavy frameworks, no external fonts. Just HTML, Vanilla JS and optimized SVGs.
- **Performance** Go backend with SQLite optimized in WAL mode.
- **Full control** Written line by line to understand exactly what happens under the hood.

---

## 🧱 Optimized Build & Footprint

Every byte counts.

| Build type | Size | Status |
| --- | --- | --- |
| Standard build | ~25 MB | ❌ Too heavy |
| Optimized + UPX | 6.5 MB | ✅ Current |

### Compilation Magic

The custom `Makefile` strips metadata and compresses the binary:

```bash
make
```

**What happens under the hood:**

- `CGO_ENABLED=0` & `static_build` Ensures the binary runs on any Linux distro without dependencies.
- `-ldflags "-s -w"` Strips debug info and symbol tables.
- **UPX Ultra-Brute** Compresses the executable to its absolute minimum.

---

## 🚀 Getting Started

### Installation

```bash
git clone https://github.com/coockieHunt/OpenList.git
cd OpenList
cp .env.exemple .env
go mod tidy
make
```

### Administration

To set or change the administrator password:

```bash
./Openlist setPassword "your_new_password"
```

### Running the app

```bash
./Openlist
```

| Service | URL |
| --- | --- |
| API | http://localhost:8080 |
| Web UI | http://localhost:4000 |

---

## 📡 API Endpoints

### 🔐 Authentication

| Method | Route | Description |
| --- | --- | --- |
| POST | `/auth/login` | Admin login |
| POST | `/auth/logout` | Logout |
| GET | `/auth/status` | Check session status |
| POST | `/auth/change-password` | Change password |

**Login** `POST /auth/login`

```json
{
  "username": "admin",
  "password": "adminpassword"
}
```

**Change password** `POST /auth/change-password`

```json
{
  "current_password": "currentpassword",
  "new_password": "newpassword"
}
```

### 📋 Lists

| Method | Route | Description |
| --- | --- | --- |
| GET | `/api/list` | Get all lists |
| POST | `/api/list` | Create a new list |
| DELETE | `/api/list/:id` | Delete a list |

### ✅ Items

| Method | Route | Description |
| --- | --- | --- |
| POST | `/api/item/:idList` | Add an item |
| PUT | `/api/item/:idList/:idItem` | Toggle validation (Sync) |
| DELETE | `/api/item/:idList/:idItem` | Remove an item |

---

## 🗺️ Roadmap : The Road to 1 MB

- [x] CRUD API & embedded SQLite (WAL mode)
- [x] Simple Web Interface (Pico.css)
- [x] Secured by a single user
- [x] Admin CLI (`setPassword` command)
- [ ] **Family Sharing** Secure list sharing via unique hash links
- [ ] **Docker "One-Command"** Multi-stage build based on a `scratch` image
- [ ] **ORM Removal** Switching from GORM to pure SQL to reach the 1 MB target

---

## 🛠️ Tech Stack

| Layer | Technology |
| --- | --- |
| Backend | Go (fast, statically typed) |
| Framework | Gin (lightweight routing) |
| Database | SQLite + WAL mode (zero-config, high-perf) |
| Frontend | Pico.css (minimalist CSS) & Vanilla JS |
| Icons | Tabler Icons (hand-picked SVGs) |

---

## 🤝 Contributing

Contributions are more than welcome! Whether it's for SQL optimization, UI/UX improvements for supermarket usage (large touch targets), or Dockerization feel free to open a PR.

---

**License: MIT**  
Hand-crafted by Jonathan (Coockie) 🍪