package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	common "github.com/life1347/go-http-streaming-example/common"
)

func streamingEchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		msgBytes, _ := json.Marshal(common.Message{
			Timestamp: time.Now().Nanosecond(),
			Msg:       fmt.Sprintf("[%d] this is a dummy streaming message\n", time.Now().Nanosecond()),
		})
		// 1. fmt.Fprintf
		// fmt.Fprintf(w, "%s\n", msgBytes)

		// 2. io.Copy
		// io.Copy(w, bytes.NewBufferString(fmt.Sprintf("%s\n", msgBytes)))

		// 3. http.ResponseWriter
		w.Write([]byte(fmt.Sprintf("%s\n", msgBytes))) // append "\n" as line delimiter
	}
}

func main() {
	http.HandleFunc("/", streamingEchoHandler)
	http.ListenAndServe("127.0.0.1:1234", nil)
}
