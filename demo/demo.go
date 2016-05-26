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
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 2)
			fmt.Fprintln(w, "Hello Http")
		})
		http.Handle("/websocket", websocket.Handler(onWebSocket))

		err := http.ListenAndServe(":8080", nil)
		if(err != nil){
			panic(err)
		}
	}()
}
func onWebSocket(ws *websocket.Conn) {
	message := "Hello Web Socket"
	websocket.Message.Send(ws, message)

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
