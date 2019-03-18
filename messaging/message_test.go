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
	AppId   string `json:"appid"`
	Cluster string `json:"cluster"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

var _ = Describe("Pusher messaging", func() {

	testmessage := TestArgsData{
		AppId:   "728602",
		Cluster: "ap2",
		Channel: "my-channel1",
		Event:   "my-event",
		Title:   "Sports",
		Message: "Cricket",
	}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	os.Setenv("SECRET", "14b80ade884bc6e20261")
	os.Setenv("KEY", "53961c904521c1807997")

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
		AppId:   "728602",
		Cluster: "ap2",
		Channel: "my-channel1",
		Event:   "my-event",
		Title:   "Sports",
		Message: "Cricket",
	}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(testmessage)
	if errr != nil {
		log.Fatal(errr)
	}

	os.Setenv("SECRET", "14b80ade884bc6e20261")
	os.Setenv("KEY", "53961c904521c180799709800")

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

	os.Setenv("SECRET", "14b80ade884bc6e20261")
	os.Setenv("KEY", "53961c904521c180799709800")

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
