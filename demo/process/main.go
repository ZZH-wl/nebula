package main

import (
	"context"
	"github.com/Wall-js/nebula"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func process(ctx context.Context, cancel context.CancelFunc) error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	// wait on kill signal
	case sig := <-ch:
		log.Printf("Received signal %s\n", sig)
	// wait on context cancel
	case <-ctx.Done():
		log.Printf("Received context shutdown")
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
	return nil
}

func main() {
	nebula.SetName("com.nebula.api.greeter")
	nebula.RunProcess(process)
}
