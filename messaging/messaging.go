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
	Data            interface{} `json:"data"`
	Channel			string		`json:"channel"`
	Event			string		`json:"event"`
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

  data := argsdata.Data
  response, err := client.Trigger(argsdata.Channel, argsdata.Event, data)
  if err != nil {
	result.WriteErrorResponse(w, err)
}
bytes, _ := json.Marshal(response)
result.WriteJsonResponse(w, bytes, http.StatusOK)

}