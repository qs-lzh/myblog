package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/handlers"
)

func main() {
	router := httprouter.New()

	router.GET("/about", showAbout)
}
