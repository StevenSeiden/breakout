package main

import (
	"github.com/dogboy21/go-discord-rp/connection"
	"github.com/gen2brain/raylib-go/raylib"
)

var windowX int32 = 807
var windowY int32 = 450
var blockWidth int32 = 30
var blockHeight int32 = 10
var blockRow int32 = 15
var paddlePos = windowX/2 - 20
var ballX = windowX / 2
var ballY = windowY - 30
var launchAngle int32 = 0
var ballMoveX int32 = 0
var ballMoveY int32 = 0
var playing = false
var movingLeft = false


type Bricks [][]int32

func init() {
	connection.OpenSocket("564965758178820146")
	connection.SetActivity("State", "Details", "pixel_large", "Small Text.", "pixel_large", "BIGGER TEXT.")
}

func checkRebound(bricks Bricks) Bricks  {
	if ballY >= (windowY-25) && ballX >= paddlePos && ballX <= (paddlePos+50) {
		ballMoveY = -ballMoveY
	} else if ballY <= 5 {
		ballMoveY = -ballMoveY
	} else if ballX >= (windowX-5) || ballX <= 5 {
		ballMoveX = -ballMoveX
	} else if ballY > windowY{
		return reset()
	} else {
		for i := 0; i <= len(bricks)-1; i++ {
			if ballX - 20 >= bricks[i][0] && ballX - 20 <= bricks[i][0] + blockWidth && ballY - 20 >= bricks[i][1] && ballY - 20 <= bricks[i][1] {
				ballMoveY = -ballMoveY;
				return append(bricks[:i], bricks[i+1:]...)
			}
		}
	}
	return bricks
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
	rl.DrawCircle(ballX, ballY, 10, rl.Purple)

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

func reset() Bricks{
	playing = false
	ballX = windowX / 2
	ballY = windowY - 30
	ballMoveX = launchAngle
	ballMoveY = -5
	paddlePos = windowX/2 - 20
	return genBricks()
}

func drawBricks(bricks Bricks){
	for i := 0; i <= len(bricks)-1; i++ {
		rl.DrawRectangle(bricks[i][0], bricks[i][1], blockWidth, blockHeight, rl.Red)
	}
}

func genBricks() Bricks{
	bricks := Bricks{}
	for i  := int32(1); i <= windowX; i = i + blockWidth + 1{
		for j := int32(0); j<= blockRow; j++ {
			bricks = append(bricks, []int32{i, int32(11)*j})
		}
	}
	return bricks
}

func drawBoard(bricks Bricks){
	drawBricks(bricks)
	rl.DrawRectangle(paddlePos, 430, 50, 10, rl.Red)
}

func main() {
	rl.InitWindow(windowX, windowY, "Breakout")
	rl.SetTargetFPS(60)
	bricks := genBricks()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		bricks = checkRebound(bricks)
		if rl.IsKeyDown(rl.KeyR) {
			bricks = reset()
		}
		movePaddle()

		if playing {
			ballX += ballMoveX
			ballY += ballMoveY
			rl.DrawCircle(ballX, ballY, 10, rl.Purple)
		} else {
			launchBall()
		}
		 /*bool CheckCollisionPointRec(Vector2 point, Rectangle rec);  // Check if point is inside rectangle*/
		drawBoard(bricks)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
