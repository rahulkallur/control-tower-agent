package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"example.com/control-tower-agent/internal/client"
	"example.com/control-tower-agent/internal/event"
	"example.com/control-tower-agent/internal/hub"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/go-uuid"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var h *hub.Hub

func serveWs(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	// Check for Auth Header
	if r.Header.Get("Authorization") == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Header.Get("Authorization") != os.Getenv("SECRET_KEY") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	clientID, _ := uuid.GenerateUUID()
	client := &client.Client{
		ID:   clientID,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	h.Register <- client

	go client.WritePump()
	go client.ReadPump(h.Unregister, h.Incoming)
}

func main() {
	r := gin.New()
	r.Use(cors.Default())
	port, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))

	if err != nil {
		panic("Port is not set or is invalid")
	}

	if ginMode := os.Getenv("GIN_MODE"); ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Register routes
	r.GET("/events", func(c *gin.Context) {
		serveWs(h, c.Writer, c.Request)
	})

	h = hub.NewHub()
	go h.Run()

	// Goroutine to handle incoming events
	go func() {
		for msg := range h.Incoming {
			event.HandleEvent(msg.Message, msg.ClientID, h)
		}
	}()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: r,
	}

	// Start Server in goroutine
	go func() {
		fmt.Printf("🚀 Server started on http://localhost:%d\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("❌ Server error: %s\n", err)
		}
	}()

	// Graceful shutdown on interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Server forced to shutdown: %s\n", err)
	}

	fmt.Println("✅ Server exiting")

}
