package main

import (
	"github.com/gofiber/fiber/v2"
)

// net/http --> gin,echo,gorilla/mux web implamantasyonu popüler olan.
// valyala/fasthttp --> fiber, en hızlı request handler eden web frameworktur.

//curl -v http://localhost:300/v1/d4d63bce-1809-4250-91ac-0470aa392ca5 mocki.io üzerinden oluşturulan  linke request atıyoruz.

//Yanlış yönlendirmeler cacheleme ve sürekli isteklere karşı koruma oluşturduğumuz bir http server.
//@GoTurkey Http-103 Eğitimi

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	//https://docs.gofiber.io/api/middware/proxy
	//https://docs.gofiber.io/guide/routing
	//Make request within handler

	//Key sonrası ne olursa olsun handlere düşsün.
	app.Get("/:key/*", ProxyHandler)

	app.Listen(":3000")
}
