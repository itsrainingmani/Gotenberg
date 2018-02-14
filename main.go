package main

import (
	"crypto/tls"
	"fmt"
	"github.com/thoj/go-ircevent"
)

const channel = "#bookz"
const serverssl = "irc.undernet.org:6667"

func main() {
	ircnick1 := "zeeus___"
	irccon := irc.IRC(ircnick1, "zeeus___")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
