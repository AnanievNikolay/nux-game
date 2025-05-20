# 🎮 Nux Game Project

![Status](https://img.shields.io/badge/status-active-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)
![Docker](https://img.shields.io/badge/docker-ready-blue)
![Go Version](https://img.shields.io/badge/go-1.23-blue)

A lightweight Go-based backend for the Nux Game project, ready to run locally or inside Docker.

---

## ⚙️ Configuration

Before running the project, make sure to review and adjust the configuration in the `config.json` file.

### Example `config.json`:

```json
{
    "logger": {
        "filename": "./logs/app.log",
        "maxsize": 10,
        "maxage": 180,
        "maxbackups": 100,
        "compress": true,
        "localtime": true
    },
    "delivery": {
        "http": {
            "host": "0.0.0.0",
            "port": 8080
        }
    },
    "db": {
        "sqlite": {
            "migrate_path": "./migration/sqlite",
            "file_name": "./_nux_game.db",
            "file_folder": "./sqlite-db",
            "ttl": 600
        }
    },
    "service": {
        "token": {
            "ttl": 604800
        }
    }
}
```

> 💡 If you change the port, remember to update `docker-compose.yml` accordingly.

---

## 🐋 Run in Docker

```
chmod +x container_start.sh
./container_start.sh
```

---

## 💻 Run Locally

```
chmod +x local_start.sh
./local_start.sh
```

---

## 📦 Dependencies

```
- Go 1.23+
- Docker
- docker-compose
- (Optional) SQLite for inspecting the DB
```

---

## 🧪 Scripts

```
container_start.sh → Run the app inside Docker  
local_start.sh     → Run the app locally  
config.json        → Main configuration file
```

---

## ✅ Tips

```
- Ensure config.json exists and is correctly configured
- Grant execution rights to shell scripts:
  chmod +x container_start.sh local_start.sh
- Check Docker port bindings if running in container
```