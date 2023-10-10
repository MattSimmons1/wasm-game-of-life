//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
	"time"
	"wasm-game-of-life/life"
)

func init() {
	register_callbacks()
	select {} // block forever
}

func register_callbacks() {
	js.Global().Get("wasm").Set("startGame", js.FuncOf(startGame))
}

func startGame(this js.Value, p []js.Value) interface{} {
	doc := js.Global().Get("document")
	gameElement := doc.Call("querySelector", "#game")
	grid := gameElement.Get("innerText")

	gridString := grid.String()

	game := life.CreateWorld(gridString)
	go func() {
		for {
			// 25 generations per rendered 'frame' for super speed
			//for i := 0; i < 25; i++ {
			for i := 0; i < 1; i++ {
				game.NextGeneration()
			}
			gameElement.Set("innerText", game.Print())
			time.Sleep(time.Millisecond * 10) // limited to 100 fps
		}
	}()
	return nil
}

func main() {

}
