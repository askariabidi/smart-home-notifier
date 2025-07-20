# Smart Home Event Notification System

A distributed system that simulates a smart home environment with event notifications using **RabbitMQ**, **Go**, **SQLite**, and **REST APIs**. Sensor data is processed in real-time and made available via APIs and a web dashboard.

---

## 📌 Features

- 🔌 Simulated sensors (motion, temperature)
- 🐇 RabbitMQ message queue for event delivery
- 🛠️ Backend service in Go
- 💾 SQLite database for event persistence
- 🌐 RESTful APIs for retrieving logs and system status
- 📊 Web dashboard and simulator interface using `html/template`

---

## 🧱 Technologies Used

- **Go** (Golang)
- **RabbitMQ** (via Docker)
- **SQLite** (using `mattn/go-sqlite3`)
- **HTML Templates** (Go’s `html/template`)
- **Docker** *(optional, for deployment)*

---

## 🧠 Architecture

```
[Sensor Simulator]
        ↓
[RabbitMQ Queue]
        ↓
[Go Consumer Service] ──→ [SQLite DB]
        ↓
  [REST API + Dashboard]
```

---

## 🚀 Getting Started

### 1. Prerequisites

- Go 1.18+ installed
- Docker installed (for RabbitMQ)

### 2. Run RabbitMQ via Docker

```bash
docker run -d --hostname smart-home-rabbit --name smart-home-rabbit \
  -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Visit [http://localhost:15672](http://localhost:15672)  
**Username:** `guest` • **Password:** `guest`

---

### 3. Clone the Repository

```bash
git clone https://github.com/askariabidi/smart-home-notifier.git
cd smart-home-notifier
go mod tidy
```

---

### 4. Run the App

```bash
go run ./cmd
```

---

## 🌐 Available Endpoints

| Route           | Method | Description                         |
|-----------------|--------|-------------------------------------|
| `/`             | GET    | Welcome message                     |
| `/logs`         | GET    | All stored sensor event logs        |
| `/status`       | GET    | Latest reading per sensor           |
| `/dashboard`    | GET    | HTML dashboard (table view)         |
| `/simulate`     | GET    | HTML form to simulate sensor events |

---

## 📸 Screenshots

*Add screenshots of your dashboard and simulate page here*

---

## 👨‍💻 Author

**Syed Mohammad Askari Abidi**  
Master’s in Software Science and Technology  
University of Florence, 2024–2025

GitHub: [@askariabidi](https://github.com/askariabidi)

---

## ⚖️ License

This project is licensed under the MIT License – see the `LICENSE` file for details.
