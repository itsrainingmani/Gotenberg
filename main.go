package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/thoj/go-ircevent"
)

const channel = "#bookz"
const server = "172.106.10.18:6667"
const dict = "abcdefghijklmnopqrstuvwxyz"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = dict[rand.Intn(len(dict))]
	}
	return string(b)
}

func main() {
	ircnick := "blaok12"
	irccon := irc.IRC(ircnick, "asalks")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.PingFreq = time.Second * 5

	// teststr := randStr(20)
	// testmsgok := make(chan bool, 1)

	// irccon.UseTLS = true
	// irccon.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) { irccon.Privmsgf(channel, "%s\n", "@search Ready Player One") })
	irccon.AddCallback("NOTICE", func(e *irc.Event) {
		if e.Nick == "SearchOok" {
			fmt.Println("THIS IS THE RAW MESSAGE - ", e.Raw)
		}
	})
	err := irccon.Connect(server)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
