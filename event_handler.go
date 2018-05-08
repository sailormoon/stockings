package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

// HandleEvent turns a termbox.Event into a possible action.
func HandleEvent(event termbox.Event) {
	if event.Type == termbox.EventKey {
		if event.Key == termbox.KeyEsc || event.Ch == 'q' || event.Ch == 'Q' {
			os.Exit(0)
		}
	}
}

// PollEvents continuously polls for events and sends them down the channel.
func PollEvents(eventQueue chan<- termbox.Event) {
	for {
		eventQueue <- termbox.PollEvent()
	}
}
