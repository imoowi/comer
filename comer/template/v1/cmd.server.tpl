/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "{{.moduleName}}/apps"
	"{{.moduleName}}/frame"
	"{{.moduleName}}/global"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "开始api服务",
	Example:      "{{.moduleName}} server -c config/settings.yml",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		setUp()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func setUp() {
	initServerConfig()
}
func initServerConfig() {
	global.Config.SetDefault("server.host", "0.0.0.0")
	global.Config.SetDefault("server.port", 8001)
}

func run() error {
	startServer()
	return nil
}

func startServer() {
	mode := global.Config.GetString("application.mode")
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}

	// 初始化全局资源
	global.Boot()

	// 初始化路由
	r := router.InitRouter()

	port := global.Config.GetUint16("server.port")
	host := global.Config.GetString("server.host")
	s := &http.Server{
		Addr:           fmt.Sprintf(`%s:%d`, host, port),
		Handler:        r,
		ReadTimeout:    time.Duration(global.Config.GetInt("server.readtimeout")) * time.Second,
		WriteTimeout:   time.Duration(global.Config.GetInt("server.writertimeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println(`server port: `, port)

	go func() {
		// 延迟100ms后初始化定时任务
		time.Sleep(100 * time.Millisecond)
		global.App.InitEngineCrontab()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(`Server Shutdown:`, err)
	}
	log.Println(`Server exiting`)
}
