package main

import (
	msg "github.com/jcocozza/text_reminder/messaging"
)

func main() {
	msg.SendMessage("SUBJECT", "MESSAGE")
}
