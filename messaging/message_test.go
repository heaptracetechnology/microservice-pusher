package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestArgsData struct {
	AppId           string      `json:"appid"`
	Cluster         string      `json:"cluster"`
	Data            interface{} `json:"data"`
	Channel			string		`json:"channel"`
	Event			string		`json:"event"`
}

var _ = Describe("Pusher messaging", func() {

	testmessage := TestArgsData{
		AppId: "728602",
		Cluster: "ap2",
		Channel :"my-channel",
		Event:"my-event",
		Data : map[string]string{"message": "hello world"},
	}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	os.Setenv("SECRET", "14b80ade884bc6e20261")
	os.Setenv("KEY", "53961c904521c1807997")

	req, err := http.NewRequest("POST", "/send", reqbody)
	if err != nil {
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
		AppId: "728602",
		Cluster: "ap2",
		Channel :"my-channel",
		Event:"my-event",
		Data : map[string]string{"message": "hello world"},
	}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	os.Setenv("SECRET", "14b80ade884bc6e20261")
	os.Setenv("KEY", "53961c904521c180799709800")

	req, err := http.NewRequest("POST", "/send", reqbody)
	if err != nil {
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