package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var windowX int32 = 800
var windowY int32 = 450
var screenWidth = int32(windowX)
var screenHeight = int32(windowY)
var paddlePos = windowX/2 - 20
var ballX = windowX/2
var ballY = windowY-30
var ballMoveX int32 = 0
var ballMoveY int32 = 0
var launching = false



func checkRebound(){
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
		ballMoveY = -ballMoveY
		ballMoveX = -ballMoveX
	}
}

func movePaddle(){
	if(rl.IsKeyDown(rl.KeyRight)){
		paddlePos = paddlePos + 10
	} else if(rl.IsKeyDown(rl.KeyLeft)){
		paddlePos = paddlePos - 10
	} else if(rl.IsKeyDown(rl.KeySpace)){
		launching = true
	}

	if(paddlePos < -5){
		paddlePos = windowX+5
	} else if (paddlePos > windowX+5){
		paddlePos = -5
	}
}
func launchBall(){
	var launchX int32 = 0
	var launchY int32 = -5
	var ballX = paddlePos + 25
	movePaddle()
	rl.DrawCircle(ballX, ballY, 10, rl.Red)

	rl.DrawLineEx(rl.NewVector2(float32(ballX), float32(ballY)),
		rl.NewVector2(float32(ballX+launchX*5), float32(ballY+launchY*5)),5, rl.Blue)
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Breakout")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		checkRebound()

		if(rl.IsKeyDown(rl.KeyR)){
			ballX = windowX/2
			ballY = windowY-25
			ballMoveX = 0
			ballMoveY = -5
		}


		if(!launching) {
			launchBall()
		} else {
			movePaddle()
			ballX += ballMoveX
			ballY += ballMoveY
			rl.DrawCircle(ballX, ballY, 10, rl.Red)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangle(paddlePos, 430, 50, 10, rl.Red)


		rl.EndDrawing()
	}

	rl.CloseWindow()
}
