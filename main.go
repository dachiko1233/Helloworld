package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Hello World</title>
  <style>
    * { margin: 0; padding: 0; box-sizing: border-box; }

    body {
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      background: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%);
      font-family: 'Segoe UI', sans-serif;
    }

    .card {
      text-align: center;
      padding: 60px 80px;
      background: rgba(255, 255, 255, 0.05);
      border: 1px solid rgba(255, 255, 255, 0.1);
      border-radius: 24px;
      backdrop-filter: blur(12px);
      box-shadow: 0 24px 60px rgba(0, 0, 0, 0.4);
    }

    h1 {
      font-size: 4rem;
      font-weight: 700;
      background: linear-gradient(90deg, #e94560, #0f3460, #53d8fb);
      background-size: 200%;
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      animation: shimmer 3s linear infinite;
    }

    p {
      margin-top: 16px;
      color: rgba(255, 255, 255, 0.5);
      font-size: 1rem;
      letter-spacing: 0.1em;
      text-transform: uppercase;
    }

    @keyframes shimmer {
      0%   { background-position: 0% }
      100% { background-position: 200% }
    }
  </style>
</head>
<body>
  <div class="card">
    <h1>Hello, World!</h1>
    <p>Welcome to my Go web server</p>
  </div>
</body>
</html>`

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(indexHTML))
	})

	router.Run(":5000")
}
