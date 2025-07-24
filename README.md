# 🪙 Gold Price Fetcher (Go)

A simple Go CLI app that fetches and displays the **current gold price in USD** using the [GoldAPI](https://www.goldapi.io/).

---

## 📦 Features

- Fetches real-time gold price (`USD`)
- Uses HTTP GET request to GoldAPI
- Handles API errors and invalid responses

---

## 🚀 Getting Started

### 🔑 API Key

This app uses [GoldAPI](https://www.goldapi.io/) to fetch real-time gold prices.

By default, a demo API key is included. For higher limits or production use:

1. Go to https://www.goldapi.io/
2. Sign up and get your free API key
3. Set it in the code :

```go
    const apiKey = your_key_here
```

## 🔧 Installation

Clone the repo:

```bash
git clone https
go run goldPri.go
