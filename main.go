package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var windowX int32 = 800
	var windowY int32 = 450
	screenWidth := int32(windowX)
	screenHeight := int32(windowY)

	rl.InitWindow(screenWidth, screenHeight, "Breakout")

	rl.SetTargetFPS(60)

	var paddlePos = windowX/2 - 20
	var ballX = windowX/2
	var ballY int32 = 450-30
	var ballMoveX int32 = 0
	var ballMoveY int32 = -5

	for !rl.WindowShouldClose() {

		if(rl.IsKeyDown(rl.KeyRight)){
			paddlePos = paddlePos + 10
		} else if(rl.IsKeyDown(rl.KeyLeft)){
			paddlePos = paddlePos - 10
		}

		if(paddlePos < -5){
			paddlePos = windowX+5
		} else if (paddlePos > windowX+5){
			paddlePos = -5
		}

		if(ballY >= (windowY-25) && ballX >= paddlePos && ballX <= (paddlePos+50)){
			if(ballX >= paddlePos+40) {
				ballMoveX = -ballMoveY
				ballMoveY = -ballMoveX
			} else if(ballX >= paddlePos+30) {
				ballMoveX = -ballMoveY
				ballMoveY = -ballMoveX
			}else{
				ballMoveY = -ballMoveY
				ballMoveX = -ballMoveX
			}
		}else if(ballY<=5 || ballX >= (windowX-5) || ballX <=5){
			var temp int32 = ballMoveX
			ballMoveX = ballMoveY
			ballMoveY = -temp

		}

		ballX += ballMoveX
		ballY += ballMoveY


		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("A rectangle:", 20, 20, 20, rl.DarkGray)

		rl.DrawLine(18, 42, screenWidth-18, 42, rl.Black)


		rl.DrawRectangle(paddlePos, 430, 50, 10, rl.Red)

		rl.DrawCircle(ballX, ballY, 10, rl.Red)


		rl.EndDrawing()
	}

	rl.CloseWindow()
}
