package main

import (
	"fmt"
	"net/http"
	"os"

	"linehook/models"
	"linehook/services/line"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()
	lineClient := line.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))

	e.POST("/webhook", func(c echo.Context) error {
		var req models.EventMessage
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		if len(req.Events) > 0 {
			if req.Events[0].Type == "beacon" {
				switch req.Events[0].Beacon.Type {
				case "enter":
					lineClient.Reply(req.Events[0].ReplyToken, "เข้าสู่ระยะของ Beacon แล้วจ้า")
				case "leave":
					lineClient.Reply(req.Events[0].ReplyToken, "ออกจากระยะ Beacon แล้วจ้า")
				}
			}
		}

		return c.String(http.StatusOK, "POST webook Hello, World!")
	}, lineClient.MessageValidator)

	e.Logger.Fatal(e.Start(":1323"))
}
