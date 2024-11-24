package main

import (
	"github.com/hhenrichsen/tapgosdk"
	"github.com/hhenrichsen/tapgosdk/events"
)

type onTapped struct{}

func (ot *onTapped) Handle(tap uint8) {
	println("tapdata: ", tap)
}

func main() {
	tm := tapgosdk.NewTapManager()
	onTapped := onTapped{}
	events.TappedData.Register(&onTapped)
	tm.Start()

	select {}
}
