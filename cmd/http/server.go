package http

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func startServer(r *gin.Engine, port string) {
	fmt.Println("Server is starting at port:", port)
	graceful(r, port)
}

// handles the graceful shutdown of a server.
func graceful(engine *gin.Engine, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: ", err)
	}

	fmt.Println("Server exiting")
}
