package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Breakout")

	rl.SetTargetFPS(60)

	var size int32 = 100

	for !rl.WindowShouldClose() {


		if(size <= 200){
			size = size + 10;
		}else{
			size = 100;
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("A rectangle:", 20, 20, 20, rl.DarkGray)

		rl.DrawLine(18, 42, screenWidth-18, 42, rl.Black)


		rl.DrawRectangle(screenWidth/4*2-60, 100, size, 60, rl.Red)


		rl.EndDrawing()
	}

	rl.CloseWindow()
}
