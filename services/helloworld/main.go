package main

import (
	"net"
	"strconv"

	"goby/pkg/conf"
	"goby/pkg/httpx"

	"golang.org/x/net/netutil"

	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	svc    = conf.GetService("helloworld")
	prefix = svc.HTTPEndpoint.RoutePrefixes[0]
)

var (
	// Path of static resources of the file server.
	staticResourcePathWebapp = (svc.Extra["static_resource_path_webapp"]).(string)
)

func routing(router *httpx.Router) {
	fsWebapp := http.FileServer(http.Dir(staticResourcePathWebapp))
	router.Router.PathPrefix("/static/").Handler(fsWebapp)

	router.Group("/", false, false, func() {
		router.Get("/", indexPageHandler)
	})

	router.Group(prefix, false, false, func() {

		// Say hello
		router.Get("/hello", helloHandler)

		// Save advice
		router.Post("/advice/save", saveAdviceHandler)

		// Get advice list
		router.Get("/advice/list", getAdviceListHandler)
	})
}

func main() {
	port := svc.HTTPEndpoint.Port
	log.Printf("helloworld is running on port => %d\n\n", port)
	log.Println("prefix => ", prefix)
	log.Println("static_resource_path_webapp => ", staticResourcePathWebapp)

	router := &httpx.Router{
		Router: mux.NewRouter().StrictSlash(true),
	}

	routing(router)

	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	l = netutil.LimitListener(l, 50)

	// log.Fatal(http.Serve(l, cors.AllowAll().Handler(router)))
	log.Fatal(http.Serve(l, router))
}
