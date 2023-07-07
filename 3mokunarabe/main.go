package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	windowWidth  = 320
	windowHeight = windowWidth
	cellSize     = windowWidth / 3
)

var (
	board       [3][3]int
	currentTurn = 1
	gameOver    bool
)

func (g *Game) Draw(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !gameOver {
		x, y := ebiten.CursorPosition()

		col := x / cellSize
		row := y / cellSize

		if isValidMove(row, col) {
			board[row][col] = currentTurn

			if checkWin(currentTurn) {
				gameOver = true
			} else if isBoardFull() {
				gameOver = true
			} else {
				currentTurn *= -1
			}
		}
	}

	if gameOver && ebiten.IsKeyPressed(ebiten.KeySpace) {
		resetGame()
	}

	screen.Fill(color.White)

	// 盤面を描画
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			x := col * cellSize
			y := row * cellSize

			if board[row][col] == 1 {
				vector.StrokeCircle(screen, float32(x)+cellSize/2, float32(y)+cellSize/2, cellSize/3, 3, color.RGBA{0, 200, 0, 255}, true)

			} else if board[row][col] == -1 {

				vector.StrokeLine(screen, float32(x+cellSize/4), float32(y+cellSize/4), float32(x+cellSize*3/4), float32(y+cellSize*3/4), 3, color.RGBA{255, 0, 0, 255}, true)
				vector.StrokeLine(screen, float32(x+cellSize*3/4), float32(y+cellSize/4), float32(x+cellSize/4), float32(y+cellSize*3/4), 3, color.RGBA{255, 0, 0, 255}, true)
			}
		}
	}

	// ゲーム終了時のメッセージを描画
	if gameOver {
		message := ""
		if checkWin(1) {
			message = "Player 1 Wins!"
		} else if checkWin(-1) {
			message = "Player 2 Wins!"
		} else {
			message = "Draw!"
		}
		ebitenutil.DebugPrint(screen, message+"\nPress SPACE to restart")
	}

}

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowWidth
}

func (g *Game) Update() error {
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(windowWidth, windowWidth)
	ebiten.SetWindowTitle("Ebiten Tic-Tac-Toe")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func isValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false
	}
	if board[row][col] != 0 {
		return false
	}
	return true
}

func checkWin(player int) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func isBoardFull() bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

func resetGame() {
	board = [3][3]int{}
	currentTurn = 1
	gameOver = false
}
