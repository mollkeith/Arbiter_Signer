// Copyright (c) 2025 The bel2 developers

package main

import (
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Event struct {
	EventID string `json:"event_id"`
	Height  int64  `json:"height"`
	OrderID string `json:"order_id"`
}

func getEvents(url string) ([]Event, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var events []Event
	if err := gconv.Scan(resp.Body, &events); err != nil {
		return nil, err
	}
	return events, nil
}
func indexHandler(r *ghttp.Request) {
	r.Response.WriteTpl("static/index.html")
}
func apiHandler(r *ghttp.Request) {
	action := r.Get("action")
	var url string
	switch action {
	case "getSucceedEvents":
		url = "https://127.0.0.1:8000/getSucceedEvents"
	case "getFailedEvents":
		url = "https://127.0.0.1:8000/getFailedEvents"
	case "getAllEvents":
		url = "https://127.0.0.1:8000/getAllEvents"
	default:
		r.Response.WriteJson(g.Map{"error": "invalid action"})
		return
	}
	events, err := getEvents(url)
	if err != nil {
		r.Response.WriteJson(g.Map{"error": err.Error()})
		return
	}
	r.Response.WriteJson(events)
}
func main() {
	s := g.Server()
	s.BindHandler("/", indexHandler)
	s.BindHandler("/api", apiHandler)
	s.SetPort(8080)
	s.Run()
}
