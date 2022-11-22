package main

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/zaffka/design/internal/handler"
	"github.com/zaffka/design/pkg/log"
)

var (
	httpPort              = ":8080"
	httpShutdownTimeout   = 500 * time.Millisecond
	httpReadHeaderTimeout = 500 * time.Millisecond
)

func main() {
	rootCtx, rootCtxCancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer rootCtxCancelFn()

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.Orders)

	hServer := http.Server{
		Addr:              httpPort,
		ReadHeaderTimeout: httpReadHeaderTimeout,
	}

	go func() {
		log.Infof("starting HTTP server at %s", httpPort)

		err := hServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		if err != nil {
			log.Errorf("failed to run HTTP server at %s, %s", httpPort, err)
			rootCtxCancelFn()
		}
	}()

	<-rootCtx.Done()

	shutCtx, shutCancelFn := context.WithTimeout(context.Background(), httpShutdownTimeout)
	defer shutCancelFn()

	if err := hServer.Shutdown(shutCtx); err != nil {
		log.Errorf("failed to gracefully shut HTTP server, %s", err)
	} else {
		log.Infof("server is shut down gracefully")
	}
}
