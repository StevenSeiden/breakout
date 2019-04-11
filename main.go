package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var windowX int32 = 800
var windowY int32 = 450
var screenWidth = int32(windowX)
var screenHeight = int32(windowY)
var paddlePos = windowX/2 - 20
var ballX = windowX / 2
var ballY = windowY - 30
var launchAngle int32 = 0
var ballMoveX int32 = 0
var ballMoveY int32 = 0
var playing = false
var movingLeft = false



func checkRebound() {
	if ballY >= (windowY-25) && ballX >= paddlePos && ballX <= (paddlePos+50) {
		ballMoveY = -ballMoveY
	} else if ballY <= 5 {
		ballMoveY = -ballMoveY
	} else if ballX >= (windowX-5) || ballX <= 5 {
		ballMoveX = -ballMoveX
	} else if ballY > screenHeight{
		reset()
	}
}

func movePaddle() {
	if rl.IsKeyDown(rl.KeyRight) {
		paddlePos = paddlePos + 10
	} else if rl.IsKeyDown(rl.KeyLeft) {
		paddlePos = paddlePos - 10
	} else if rl.IsKeyDown(rl.KeySpace) {
		playing = true
		ballMoveX = launchAngle
		ballMoveY = -5
	}

	if paddlePos < -5 {
		paddlePos = windowX + 5
	} else if paddlePos > windowX+5 {
		paddlePos = -5
	}
}
func launchBall() {

	ballX = paddlePos + 25
	rl.DrawCircle(ballX, ballY, 10, rl.Red)

	rl.DrawLineEx(rl.NewVector2(float32(ballX), float32(ballY)),
		rl.NewVector2(float32(ballX+launchAngle*5), float32(ballY-25)), 5, rl.Blue)
	if launchAngle == -8 && movingLeft {
		movingLeft = false
	} else if movingLeft {
		launchAngle--
	} else if launchAngle < 8 && !movingLeft {
		launchAngle++
	} else if launchAngle == 8 {
		movingLeft = true
	}

}

func reset(){
	playing = false
	ballX = windowX / 2
	ballY = windowY - 30
	ballMoveX = launchAngle
	ballMoveY = -5
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Breakout")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		checkRebound()
		if rl.IsKeyDown(rl.KeyR) {
			reset()
		}
		movePaddle()
		if playing {
			ballX += ballMoveX
			ballY += ballMoveY
			rl.DrawCircle(ballX, ballY, 10, rl.Red)
		} else {
			launchBall()
		}
		 /*bool CheckCollisionPointRec(Vector2 point, Rectangle rec);  // Check if point is inside rectangle*/

		rl.DrawRectangle(paddlePos, 430, 50, 10, rl.Red)
		/* blocks should have 5 pixals between each other*/
		//for loop to draw all the rectangle
		for i := 800; i >= 20; i = i - 45 {
			
			rl.DrawRectangle(int32(i), 0, 40, 10, rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
