package demo

import (
	//"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

func StartListening() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 2)
		fmt.Fprintln(w, "Hello For the Last Time")
	})
	http.Handle("/json", websocket.Handler(onWebSocket))

	http.ListenAndServe(":8080", nil)
}
func onWebSocket(ws *websocket.Conn) {
	var page = ReactComponent{
		Type: "View",
		Props: struct {
			Style interface{} `json:"style"`
		}{
			struct {
				BackgroundColor string `json:"backgroundColor"`
				Flex int `json:"flex"`
				JustifyContent string `json:"justifyContent"`
				AlignItems string `json:"alignItems"`
			}{
				"green",
				1,
				"center",
				"center",
			},
		},
		Children: []interface{}{
			ReactComponent{
				Type:  "Text",
				Props: nil,
				Children: []interface{}{
					"Hello World From Go",
				},
			},
		},
	}
	websocket.JSON.Send(ws, page)

}
func WriteFile(baseDir string) {
	d1 := []byte("hello go\n")
	file := filepath.Join(baseDir, "test.txt")
	ioutil.WriteFile(file, d1, 0644)
}
func ReadFile(baseDir string) string {
	file := filepath.Join(baseDir, "test.txt")
	bytes, _ := ioutil.ReadFile(file)
	return string(bytes)
}
func HelloWorld() string {
	return "Hello Again Cruel World!"
}
