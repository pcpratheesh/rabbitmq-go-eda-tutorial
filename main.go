package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pcpratheesh/rabbitmq-go-eda-tutorial/cmd/consumer"
	"github.com/pcpratheesh/rabbitmq-go-eda-tutorial/cmd/producer"
	"github.com/pcpratheesh/rabbitmq-go-eda-tutorial/models"
	"github.com/r3labs/sse"
)

//go:embed assets/* templates/*
var f embed.FS

var socketDataPaylod = make(chan models.WebsocketDataPayload, 1)

func main() {

	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tpl"))
	router.SetHTMLTemplate(templ)
	router.StaticFS("/public", http.FS(f))

	// SSE endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tpl", nil)
	})

	server := sse.New()
	defer server.Close()

	router.GET("/event", func(ctx *gin.Context) {
		consumer.CloseAllConsumers()

		// Set the SSE headers
		ctx.Header("Content-Type", "text/event-stream")
		ctx.Header("Cache-Control", "no-cache")
		ctx.Header("Connection", "keep-alive")

		// Create a channel to signal when the client disconnects
		clientClosed := ctx.Writer.CloseNotify()

		// Loop to send events
		for {
			select {
			case data := <-socketDataPaylod:
				// Send an event to the client
				ctx.SSEvent("event", data)
				ctx.Writer.Flush()
			case <-clientClosed:
				// Client disconnected, exit the loop
				return
			}
		}
	})

	router.POST("/consumer", AddConsumer)
	router.POST("/producer", RunProducer)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func AddConsumer(ctx *gin.Context) {
	var reqest struct {
		Name string `json:"name"`
	}
	if err := ctx.BindJSON(&reqest); err != nil {
		ctx.JSON(http.StatusBadRequest, "unable to process your request")
		return
	}

	// run a new consumer
	data := models.WebsocketDataPayload{
		Type: "consumer",
		Name: reqest.Name,
	}

	socketDataPaylod <- data

	go consumer.LaunchConsumer(reqest.Name, socketDataPaylod)

	ctx.JSON(http.StatusOK, "addded")
}

func RunProducer(ctx *gin.Context) {
	go producer.RunProducer()
	ctx.JSON(http.StatusOK, "producer running")
}
