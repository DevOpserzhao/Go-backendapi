package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

/**
1. gin Web
2. 日志
3. 链路追踪
4. 监控
5. 平滑启动
**/

type App struct {
	Addr   string
	Pprof  string
	Mode   string
	Secret string
}

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger(), NoCache(), Cors(), Secure(), RequestID())
	return engine
}

func (s *App) Run(e *gin.Engine) {
	srv := &http.Server{
		Addr:         ":" + s.Addr,
		Handler:      e,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	go func() {
		log.Printf("\033[1;32;32m App Server Running: [%s] \033[0m", s.Addr)
		s.system()
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 second")
	}
	log.Println("Server exiting ...")
}

func (s *App) PProf() {
	log.Printf("\033[1;32;32m Pprof Running: [%s] \033[0m", s.Pprof)
	log.Fatal(http.ListenAndServe(":"+s.Pprof, nil))
}

func (s *App) system() {
	var mem runtime.MemStats
	log.Printf("\033[1;32;32m Go: %s | CPU: %d | MEM: %d | GOOS: %s | DATE: %s\033[0m",
		runtime.Version(),
		runtime.NumCPU(),
		&mem.Sys,
		runtime.GOOS, time.Now().Format("2006-01-02 15:02:01"))
}
