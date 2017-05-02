package common

type Message struct {
	Timestamp int    `json:timestamp`
	Msg       string `json:msg`
}
