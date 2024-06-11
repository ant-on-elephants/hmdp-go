package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"
	"hmdp-go/common/setting"
	"hmdp-go/common/validator"
	"hmdp-go/routers"
)

func main() {
	binding.Validator = new(validator.DefaultValidator)
	router := routers.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
