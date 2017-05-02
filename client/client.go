package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	common "github.com/life1347/go-http-streaming-example/common"
)

func streamingClient(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Failed to connect streaming server: %v\n", err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		// remove string delimiter
		line = bytes.TrimRight(line, "\n")
		msg := common.Message{}
		if err := json.Unmarshal(line, &msg); err != nil {
			fmt.Printf("Failed to parse message: %v", err)
			break
		}
		io.Copy(os.Stdout, bytes.NewBufferString(msg.Msg))
	}
}

func main() {
	streamingClient("http://127.0.0.1:1234/")
}
