// Copyright (c) 2025 The bel2 developers

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jsonrpc "github.com/BeL2Labs/Arbiter_Signer/app/web/json_rpc"
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func getEvents(url string) ([]events.EventInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	eventsResp := jsonrpc.EventInfoResponse{}
	if err := json.Unmarshal(body, &eventsResp); err != nil {
		return nil, err
	}
	fmt.Println("eventResp:", eventsResp)
	if eventsResp.Code != 0 {
		return nil, fmt.Errorf("failed to get events: %s", eventsResp.Message)
	}

	return eventsResp.Data.Events, nil
}

func indexHandler(r *ghttp.Request) {
	r.Response.WriteTpl("static/index.html")
}
func apiHandler(r *ghttp.Request) {
	action := r.Get("action")
	var url string
	switch action {
	case "getSucceedEvents":
		url = "http://127.0.0.1:8000/succeed_events"
	case "getFailedEvents":
		url = "http://127.0.0.1:8000/failed_events"
	case "getRequiredEvents":
		url = "http://127.0.0.1:8000/required_events"
	case "getAllEvents":
		url = "http://127.0.0.1:8000/events"
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
