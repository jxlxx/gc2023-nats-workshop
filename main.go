package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

var faves = map[string]string{
	"drink": "beer",
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected to nats")

	config := micro.Config{
		Name:        "jxlxx",
		Description: "Returns a list of jxlxx's faves",
		Version:     "0.0.99",
	}

	service, err := micro.AddService(nc, config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.AddEndpoint("list_favourites", micro.HandlerFunc(func(r micro.Request) {
		// biz
		r.RespondJSON(faves)
	}), micro.WithEndpointSubject("jxlxx.favourites"), micro.WithEndpointMetadata(map[string]string{
		"descrippie": "swag",
	})); err != nil {
		log.Fatalln(err)
	}

	if err := service.AddEndpoint("get_favourite", micro.HandlerFunc(func(r micro.Request) {
		fave, ok := faves[string(r.Data())]
		if !ok {
			r.Error("404", "no fave", nil)
		}
		// biz
		r.RespondJSON(fave)
	}), micro.WithEndpointSubject("jxlxx.favourite"), micro.WithEndpointMetadata(map[string]string{
		"descrippie": "swag",
	})); err != nil {
		log.Fatalln(err)
	}
	log.Println("serving")

	runtime.Goexit()
}
