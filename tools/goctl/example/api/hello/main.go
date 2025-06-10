// Code scaffolded by goctl. Safe to edit.
// goctl <no value>

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"hello/internal/config"
	"hello/internal/handler"
	"hello/internal/svc"
)

// 环境默认是[dev]
var envMode = "dev"

func main() {
	readEnv()
	conf := readConfig()
	startServer(conf)
}

func readEnv() {
	value := os.Getenv("XEnvMode")
	if value != "" {
		envMode = value
	}
}

func readConfig() config.Config {
	data, err := os.ReadFile(fmt.Sprintf("etc/config-%s.yaml", envMode))
	if err != nil {
		panic(err)
	}
	var conf config.Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func startServer(conf config.Config) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// cors in dev mode
	corsInDevMode(engine)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        engine,
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}
	// init service context
	svcCtx := svc.NewServiceContext(conf)
	// setup routes
	handler.RegisterHandlers(engine, svcCtx)
	// start server
	fmt.Println("[gin] ==> ready to work!")
	_ = s.ListenAndServe()
}

func corsInDevMode(engine *gin.Engine) {
	if envMode == "dev" {
		engine.Use(func(c *gin.Context) {
			method := c.Request.Method
			origin := c.Request.Header.Get("Origin")
			if origin != "" {
				c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
				c.Header("Access-Control-Allow-Headers",
					"Origin, X-Requested-With, Content-Type, Accept, Authorization")
				c.Header("Access-Control-Expose-Headers",
					"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			if method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
			c.Next()
		})
	}
}
