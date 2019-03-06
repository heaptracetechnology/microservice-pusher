package messaging

import ("github.com/pusher/pusher-http-go"
"github.com/heaptracetechnology/microservice-pusher/result"
"net/http"
"os"
"io/ioutil"
"encoding/json"
)

type ArgsData struct {
	AppId           string      `json:"appid"`
	Cluster         string      `json:"cluster"`
	Channel			string		`json:"channel"`
	Event			string		`json:"event"`
	Title			string		`json:"title"`
	Message			string		`json:"message"`
}

type Message struct {
    Success string `json:"success"`
    Message string `json:"message"`
	StatusCode int `json:"statuscode"`
}

type Data struct {
    Title 	string `json:"title"`
    Message string `json:"message"`
}

func SendMessage(w http.ResponseWriter, r *http.Request){

	var SECRET = os.Getenv("SECRET")
	var KEY = os.Getenv("KEY")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var argsdata ArgsData
	err = json.Unmarshal(body, &argsdata)

	
  	client := pusher.Client{
		AppId: argsdata.AppId,
		Key: KEY,
		Secret: SECRET,
		Cluster: argsdata.Cluster,
		Secure: true,
	  }	
	  
	data  := Data{
		Title	: argsdata.Title,
		Message : argsdata.Message,
	}
	_, err = client.Trigger(argsdata.Channel, argsdata.Event, data)

	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	message := Message{"true","notification sent",http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
	return

}