package http

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/ports"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/logger"

	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"flag"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Adapter struct {
	s    *http.Server
	writer *kafka.Writer
	l    net.Listener
	tasks ports.Tasks
	isProd bool
}

type Config struct {
	Port int `env:"HTTP_PORT" envDefault:"3000"`
}

var async = flag.Bool("a", false, "use async")

func New(tasks ports.Tasks, log logger.Logger, isProd bool) (*Adapter, error) {
	flag.Parse()
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse server http adapter configuration failed: %w", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("server start failed: %w", err)
	}

	router := gin.Default()
	server := http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     []string{"127.0.0.1:29092", "127.0.0.1:39092", "127.0.0.1:49092"},
		Topic:       "demo",
		Async:       *async,
		BatchSize:   2000,

		// CompressionCodec: &compress.Lz4Codec,

		Balancer: &SimpleBalancer{},
	})

	a := Adapter{
		s:    &server,
		writer: writer,
		l:    l,
		tasks: tasks,
		isProd: isProd,
	}
	initRouter(&a, router, log)

	return &a, nil
}

func (a *Adapter) Start() error {
	var err error
	go func() {
		err = a.s.Serve(a.l)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			err = fmt.Errorf("server start failed: %w", err)
		}
		err = nil
	}()

	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) Stop(ctx context.Context) error {
	var (
		err  error
		once sync.Once
	)
	once.Do(func() {
		err = a.s.Shutdown(ctx)
	})
	return err
}

type UserProjection struct {
	Email string	
}

func (a *Adapter) AuthRequired() gin.HandlerFunc {
	client := &http.Client{Timeout: 10 * time.Second}

	return func(ctx *gin.Context) {
		if !a.isProd {
			ctx.Set("email", "email1")
			ctx.Next()
			return
		}
		endpoint := "http://localhost:8080/auth/api/v1/verify"
		
		accessCookie, err := ctx.Request.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "verification failed",
			})
			return
		}

		refreshCookie, err := ctx.Request.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "verification failed",
			})
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, nil)
		if err != nil {
			a.BindError(ctx, err)
			return
		}
		
		req.AddCookie(accessCookie)
		req.AddCookie(refreshCookie)

		resp, err := client.Do(req)
		if err != nil {
			a.BindError(ctx, err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != http.StatusOK {
			ctx.JSON(resp.StatusCode, gin.H{
				"error": resp.Status,
			})
			return
		}

		user := UserProjection{}
		json.NewDecoder(resp.Body).Decode(&user)
	
		ctx.Set("email", user.Email)
		ctx.Next()
	}
}
