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
		fmt.Fprintln(w, "Hello Http")
	})
	http.Handle("/websocket", websocket.Handler(onWebSocket))

	http.ListenAndServe(":8080", nil)
}
func onWebSocket(ws *websocket.Conn) {
	var message = "Hello Web Socket"
	websocket.JSON.Send(ws, message)

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
	return "Hello Native Bridge"
}
