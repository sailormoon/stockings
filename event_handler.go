package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

// HandleEvent turns a termbox.Event into a possible action.
func HandleEvent(event termbox.Event) {
	if event.Type == termbox.EventKey {
		switch event.Key {
		case termbox.KeyEsc, 'q', 'Q':
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
