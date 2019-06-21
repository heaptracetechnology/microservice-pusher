package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestArgsData struct {
	AppID   string `json:"appid"`
	Cluster string `json:"cluster"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

var (
	secret  = os.Getenv("PUSHER_SECRET")
	key     = os.Getenv("PUSHER_KEY")
	appID   = os.Getenv("PUSHER_APP_ID")
	cluster = os.Getenv("PUSHER_CLUSTER")
	channel = os.Getenv("PUSHER_CHANNEL")
	event   = os.Getenv("PUSHER_EVENT")
	mockKey = "53961c904521c180799709800"
)

var _ = Describe("Pusher messaging", func() {

	testmessage := TestArgsData{
		AppID:   appID,
		Cluster: cluster,
		Channel: channel,
		Event:   event,
		Title:   "Sports",
		Message: "Cricket",
	}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	os.Setenv("SECRET", secret)
	os.Setenv("KEY", key)

	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Pusher messaging negative", func() {

	testmessage := TestArgsData{
		AppID:   appID,
		Cluster: cluster,
		Channel: channel,
		Event:   event,
		Title:   "Sports",
		Message: "Cricket",
	}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	os.Setenv("SECRET", secret)
	os.Setenv("KEY", mockKey)

	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Pusher messaging negative unmarshal ", func() {

	testmessage := []byte(`{"status":false}`)

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	os.Setenv("SECRET", secret)
	os.Setenv("KEY", mockKey)

	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
