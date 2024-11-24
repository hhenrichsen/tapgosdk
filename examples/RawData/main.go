package main

import (
	"github.com/hhenrichsen/tapgosdk"
	"github.com/hhenrichsen/tapgosdk/events"
)

type onRaw struct{}

func (ot *onRaw) Handle(tap []byte) {
	println("tapdata: ", tap)
}

func main() {
	tm := tapgosdk.NewTapManager()
	tm.SetDefaultIputMode(tapgosdk.RawSensors)
	onRaw := onRaw{}
	events.RawData.Register(&onRaw)
	tm.Start()

	select {}
}
