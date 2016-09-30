package main

import (
	"fmt"
	//"io"
	//"log"
	//"net/http"
	//"os"
	//"strings"
	//"bytes"
	//"io/ioutil"
	//"time"
	"encoding/json"
	"time"
	"net/http"
	"io/ioutil"
	"strings"
)

const NUMBER_OF_EVENTS = 360
const EVENT_NAME  = "e3600"

type Event struct {
	User_Id string     `json:"user_id"`
	Event_name string  `json:"event_name"`
	Hour int	   `json:"hour"`
	Minute int	   `json:"minute"`
	Second int	   `json:"second"`
	Event_Ts int64 `json:"event_time_ts"`
}

type Events struct {
	Events []Event `json:"events"`
}

func main() {
	now := time.Now()
	event1:= Event{User_Id: "1",Event_name:EVENT_NAME,Event_Ts:now.Unix()*1000,Hour:now.Hour(),Minute:now.Minute(),Second:now.Second()}
	event2:= Event{User_Id: "2",Event_name:EVENT_NAME,Event_Ts:now.Unix()*1000,Hour:now.Hour(),Minute:now.Minute(),Second:now.Second()}
	event3:= Event{User_Id: "3",Event_name:EVENT_NAME,Event_Ts:now.Unix()*1000,Hour:now.Hour(),Minute:now.Minute(),Second:now.Second()}
	events := Events{Events:[]Event{event1,event2,event3}}
	b, err := json.Marshal(events)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	data:="data=" + string(b)
	url:="https://api.cooladata.com/v1/4b59crbv2opribp8gudgiean0nvngwo3/track"
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(1 * time.Second),

	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}



