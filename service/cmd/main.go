package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/spf13/pflag"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iceEvening/me/service/db"
	"github.com/iceEvening/me/service/router"
)

func init() {
	isDebug := os.Getenv("DEBUG")
	if isDebug == "1" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func setConfig() {
	flag.String("host", "127.0.0.1", "database host")
	flag.Int("port", 7001, "database port")
	flag.String("database", "me", "database name")
	flag.String("username", "postgres", "user name")
	flag.String("password", "postgres", "user password")
	flag.Int("maxconn", 50, "max conns")
	flag.Int("maxidle", 20, "max idle conns")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setupRouter() *router.Router {
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://www.sanghongwei.com",
			"http://sanghongwei.com",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Origin",
			"Accept",
			"X-Requested-With",
			"Content-Type",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
			"access-control-allow-origin",
			"X-USER-TOKEN",
		},
		AllowMethods: []string{
			"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"X-USER-TOKEN",
		},
	}))
	r := router.GetRouter(e, db.GetDB(
		viper.GetString("host"),
		viper.GetInt("port"),
		viper.GetString("database"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetInt("maxconn"),
		viper.GetInt("maxidle"),
	))
	r.RegisterTestRouter()
	r.RegisterUserRouter()
	r.RegisterArticleRouter()
	return r
}

func main() {
	setConfig()
	r := setupRouter()
	serv := &http.Server{
		Addr:         ":8080",
		Handler:      r.E,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	serv.ListenAndServe()
}
