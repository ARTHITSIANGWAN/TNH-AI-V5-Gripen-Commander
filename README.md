# 🏰 TNH GRIPEN SQUADRON ENGINE (V84.9.2)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square)
![Security](https://img.shields.io/badge/security-HMAC--SHA256-blue?style=flat-square)

**TNH Gripen Squadron Engine** is the core backend automation service for the ThitNueaHub (TNH) AI architecture. Developed in Go, this engine acts as an asynchronous task executor with strict HMAC-SHA256 security validation to prevent unauthorized access.

## 🚀 Architecture Overview

*   **L9 Shield (Security):** Intercepts incoming requests and validates the HMAC signature.
*   **L11 Swarm (Executor):** Utilizes Go concurrency (`goroutine`) to execute system tasks asynchronously without blocking the main HTTP thread.
*   **API Protocol:** Listens on port `2026` via HTTP POST requests.

## 🛠 Prerequisites

*   [Go (Golang)](https://golang.org/doc/install) 1.21 or higher installed on your system (or Termux environment).

## ⚙️ Installation & Usage

1.  **Clone the repository / Setup local environment:**
    ```bash
    git clone [https://github.com/your-repo/tnh-gripen-engine.git](https://github.com/your-repo/tnh-gripen-engine.git)
    cd tnh-gripen-engine
    ```

2.  **Run the Engine:**
    ```bash
    go run main.go
    ```
    *You should see the following output:*
    `⚡ [Go Engine Sovereign]: ล็อกตำแหน่งที่พอร์ต :2026`

## 📡 API Documentation

### 1. Status Check
*   **Endpoint:** `/`
*   **Method:** `GET`
*   **Description:** Verifies if the engine is online.

### 2. Squadron Launch (Task Execution)
*   **Endpoint:** `/api/v84/squadron/launch`
*   **Method:** `POST`
*   **Headers:** `Content-Type: application/json`

#### Payload Schema (JSON)
```json
{
  "command_id": "CMD-001",
  "action": "clear_cache",
  "squadron": "alpha",
  "timestamp": 1716500000,
  "signature": "<HMAC-SHA256-HEX-STRING>"
}
