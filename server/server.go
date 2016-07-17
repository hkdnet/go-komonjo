package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hkdnet/go-komonjo/api"
	"github.com/nlopes/slack"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

// Option is for server.
type Option struct {
	Port int
}

type errorResponse struct {
	status  string
	message string
}

func newUnexepectedError() []byte {
	res := errorResponse{
		status:  "error",
		message: "An unexpected error ocurred.",
	}
	b, _ := json.Marshal(res)
	return b
}

// Up runs a http server.
func Up(opt Option) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("data/index.html")
		if err != nil {
			fmt.Printf("%v\n", err)
			w.Write(newUnexepectedError())
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
	http.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("data/index.css")
		if err != nil {
			fmt.Printf("%v\n", err)
			w.Write(newUnexepectedError())
			return
		}
		w.Header().Set("Content-Type", "text/css")
		w.Write(data)
	})
	http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("data/index.js")
		if err != nil {
			fmt.Printf("%v\n", err)
			w.Write(newUnexepectedError())
			return
		}
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(data)
	})
	http.HandleFunc("/api/channels", func(w http.ResponseWriter, r *http.Request) {
		client := api.NewClient()
		channels, err := client.GetChannels(false)
		if err != nil {
			w.Write(newUnexepectedError())
			return
		}
		filter := r.FormValue("name")
		if filter != "" {
			channels = api.FilterChannel(channels, filter)
		}
		b, _ := json.Marshal(channels)
		w.Write(b)
	})
	http.HandleFunc("/api/histories", func(w http.ResponseWriter, r *http.Request) {
		channelID := r.FormValue("channelID")
		channelName := r.FormValue("channelName")
		if channelID == "" && channelName == "" {
			w.Write(newUnexepectedError())
			return
		}
		countStr := r.FormValue("count")
		count := 100
		if tmp, err := strconv.Atoi(countStr); err == nil {
			count = tmp
		}
		client := api.NewClient()
		if channelID == "" {
			channel, err := client.GetChannel(channelName)
			if err != nil {
				w.Write(newUnexepectedError())
				return
			}
			channelID = channel.ID
		}
		history, err := client.GetChannelHistory(channelID, slack.HistoryParameters{Count: count})
		if err != nil {
			w.Write(newUnexepectedError())
			return
		}
		b, _ := json.Marshal(history)
		w.Write(b)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", opt.Port), nil)
}
