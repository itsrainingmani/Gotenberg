package main

import (
	"fmt"

	"github.com/thoj/go-ircevent"
)

const channel = "#bookz"
const serverssl = "172.106.10.18:6667"

func main() {
	ircnick1 := "blatiblat"
	irccon := irc.IRC(ircnick1, "IRCTestSSL")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	// irccon.UseTLS = true
	// irccon.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
