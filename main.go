package main

import (
	"flag"
	"fmt"
	"github.com/qa-dev/wda-inspector/handlers"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
	"github.com/qa-dev/go-core/middleware"
)

func main() {
	var iHost, iPort, bundleId string
	flag.StringVar(&iHost, "h", "", "WDA host")
	flag.StringVar(&iPort, "p", "8100", "WDA port, 8100 by default")
	flag.StringVar(&bundleId, "bundleId", "", "Bundle Id, default 'ru.avito.services.dev'")
	flag.Parse()
	if iHost == "" {
		log.Fatal("PLease set inspectors host, for example '-h=10.10.10.48'")
	}
	if bundleId == "" {
		log.Fatal("Please set bundle id, for example '-bundleId=com.apple.mobilesafari'")
	}
	requester := wda.NewHttpClient(iHost, iPort)
	iClient := wda.NewClient(requester, bundleId)
	mux := http.NewServeMux()
	setHandlers(mux, iClient)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8888), mux))
}

func setHandlers(mux *http.ServeMux, iClient *wda.Client) {
	mux.Handle("/static/", handlers.NewStaticHandler())
	mux.Handle("/screenshot", middleware.New(handlers.NewScreenshotHandler(iClient)))
	mux.Handle("/source", middleware.New(handlers.NewSourceHandler(iClient)))
	mux.Handle("/find", middleware.New(handlers.NewFindHandler(iClient)))
	mux.Handle("/getText", middleware.New(handlers.NewGetTextHandler(iClient)))
	mux.Handle("/", middleware.New(handlers.NewIndexPage()))
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
