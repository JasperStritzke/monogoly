<br />
<div align="center">
<h3 align="center">Monogoly</h3>

  <p align="center">
    A PayPal like application to play monopoly without the paper cash.
    <br />
    <a href="#">View Demo</a>
    ·
    <a href="./issues">Report Bug</a>
    ·
    <a href="./issues">Request Feature</a>
  </p>
</div>

---

## About the project

---
### Built with

<img src="https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D" alt="Vue"/>
<img src="https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="TailwindCSS"/>
<img src="https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white" alt="Vite"/>
<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Golang"/>
<img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"/>

---
### Routes

``POST /api/game/create``

```json
{
  "currency": "USD",
  "amountGo": 200,
  "initialBalance": 2000,
  "password": "1234"
}
```

```json
{
  "id": "4839",
  "password": "1234"
}
```

<br>

``POST /api/game/join``

```json
{
  "id": "4839",
  "password": "1234",
  "name": "Thor Odinson"
}
```

```json
{
  "currency": "USD",
  "players": [
    {
      "id": 1,
      "name": "Tony Stark",
      "balance": 1500,
      "host": true
    },
    {
      "id": 2,
      "name": "Peter Parker",
      "balance": 2500
    }
  ],
  "you": {
    "id": 3,
    "name": "Thor Odinson",
    "balance": 2000
  }
}
```
