package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-pusher/result"
	"github.com/pusher/pusher-http-go"
)

type ArgsData struct {
	AppId   string `json:"appid"`
	Cluster string `json:"cluster"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type Data struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func SendMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var SECRET = os.Getenv("SECRET")
	var KEY = os.Getenv("KEY")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
		return
	}

	client := pusher.Client{
		AppId:   argsdata.AppId,
		Key:     KEY,
		Secret:  SECRET,
		Cluster: argsdata.Cluster,
		Secure:  true,
	}

	data := Data{
		Title:   argsdata.Title,
		Message: argsdata.Message,
	}
	_, triggerErr := client.Trigger(argsdata.Channel, argsdata.Event, data)
	if triggerErr != nil {
		result.WriteErrorResponse(responseWriter, triggerErr)
		return
	}

	message := Message{"true", "notification sent", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
