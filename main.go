package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

var (
	name string
)

func onSubmit() {
	fmt.Println("Name:", name)
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Can Data Be Queer?"),
		giu.InputText(&name),
		giu.Button("submit").OnClick(onSubmit),
	)
}

func main() {
	window := giu.NewMasterWindow("[Data Privacy] Can Data Be Queer?", 800, 600, giu.MasterWindowFlagsNotResizable)
	window.Run(loop)
}
