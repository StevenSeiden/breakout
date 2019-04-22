package main

import (
	"fmt"
	"github.com/dogboy21/go-discord-rp/connection"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

var windowX int32 = 807
var windowY int32 = 450

const blockWidth int32 = 30
const blockHeight int32 = 10

var blockRow int32 = 15
var paddlePos = windowX/2 - paddleWidth/2

const ballSize float32 = 10

var ballX = windowX/2 + 10/2
var ballY = windowY - 30
var launchAngle float64 = 0
var ballMoveX int32 = 0
var ballMoveY int32 = 0
var playing = false
var movingLeft = false
var paddleWidth int32 = 80
var score = 0
var startGame = true

const debugMode = false

type Bricks [][]int32

func init() {
	connection.OpenSocket("564965758178820146")
	connection.SetActivity("Playing", "Score: "+fmt.Sprintf("%d", score), "pixel_large", "Beta", "logo_pixelated", "This is a picture of Taylor.")
}

func checkRebound(bricks Bricks) Bricks {
	if ballY >= (windowY - 25) {
		if ballX >= paddlePos && ballX <= (paddlePos+paddleWidth) {
			ballMoveY = -ballMoveY
		} else {
			return reset()
		}
	} else if ballY <= int32(ballSize) {
		ballMoveY = -ballMoveY
	} else if ballX >= (windowX-int32(ballSize)) || ballX <= int32(ballSize) {
		ballMoveX = -ballMoveX
	} else if ballY > windowY-20 {
		return reset()
	} else { //Checking for brick collisions
		for i := 0; i <= len(bricks)-1; i++ {
			if debugMode {
				fmt.Println("Checking brick #" + fmt.Sprintf("%d", i))
			}
			if ballX+int32(ballSize) >= bricks[i][0] && ballX-int32(ballSize) <= bricks[i][0]+blockWidth &&
				ballY+int32(ballSize) >= bricks[i][1] && ballY-int32(ballSize) <= bricks[i][1]+blockHeight {
				ballMoveY = -ballMoveY
				score = score + 100
				connection.SetActivity("Playing", "Score: "+fmt.Sprintf("%d", score), "pixel_large", "Beta", "logo_pixelated", "This is a picture of Taylor.")
				if debugMode {
					fmt.Println("COLLISION with brick #" + fmt.Sprintf("%d", i))
				}
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
		ballMoveX = int32(launchAngle)
		ballMoveY = int32(math.Abs(launchAngle)) - 8
	}

	if paddlePos < -5 {
		paddlePos = windowX + 5
	} else if paddlePos > windowX+5 {
		paddlePos = -5
	}
}
func launchBall() {
	ballX = paddlePos + paddleWidth/2
	rl.DrawCircle(ballX, ballY, ballSize, rl.DarkPurple)

	rl.DrawLineEx(rl.NewVector2(float32(ballX), float32(ballY)),
		rl.NewVector2(float32(ballX+int32(launchAngle)*5), float32(ballY-(25-int32(math.Abs(launchAngle))))), 5, rl.Blue)
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

func reset() Bricks {
	playing = false
	ballX = windowX / 2
	ballY = windowY - 30
	ballMoveX = int32(launchAngle)
	ballMoveY = -5
	paddlePos = windowX/2 - 20
	score = 0
	return genBricks()
}

func drawBricks(bricks Bricks) {
	for i := 0; i <= len(bricks)-1; i++ {
		rl.DrawRectangle(bricks[i][0], bricks[i][1], blockWidth, blockHeight, rl.Red)
		if debugMode {
			rl.DrawText(fmt.Sprintf("%d", i), bricks[i][0], bricks[i][1], 7, rl.White)
		}
	}

}

func genBricks() Bricks {
	bricks := Bricks{}
	for i := int32(1); i <= windowX; i = i + blockWidth + 1 {
		for j := int32(0); j <= blockRow; j++ {
			bricks = append(bricks, []int32{i, int32(11) * j})
		}
	}
	return bricks
}

func drawBoard(bricks Bricks) {
	drawBricks(bricks)
	rl.DrawRectangle(paddlePos, windowY-20, paddleWidth, 10, rl.Red)
	rl.DrawText(fmt.Sprintf("%v",score), 750,windowY-40 , 10,rl.Red)
}

func main() {
	rl.InitWindow(windowX, windowY, "Breakout")
	rl.SetTargetFPS(60)
	rl.BeginDrawing()
	bricks := genBricks()

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)
		if !startGame {
			bricks = checkRebound(bricks)
			if rl.IsKeyDown(rl.KeyR) {
				bricks = reset()
			}
			movePaddle()

			if playing {
				ballX += ballMoveX
				ballY += ballMoveY
				rl.DrawCircle(ballX, ballY, ballSize, rl.DarkPurple)
			} else {
				launchBall()
			}
			/*bool CheckCollisionPointRec(Vector2 point, Rectangle rec);  // Check if point is inside rectangle*/
			drawBoard(bricks)

			rl.EndDrawing()

			if debugMode {
				paddlePos = ballX - 35
			} //Enable to automate the game
		}else {
			rl.DrawRectangle(10,10,10,10,rl.DarkGreen)
			if rl.IsKeyDown(rl.KeyS){
				startGame = false
			}
		}
	}

	rl.CloseWindow()
}
