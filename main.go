package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	width, height := 20, 10

	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}

	screen.Clear()

	vx, vy := 1, 1
	x, y := 0, 0

	buffer := make([]rune, 0, width*height+height)

	for {
		screen.MoveTopLeft()

		// Topun konumunu güncelle
		x += vx
		y += vy

		// Duvara çarptı mı?
		if x >= width-1 || x <= 0 {
			vx = -vx
		}
		if y >= height-1 || y <= 0 {
			vy = -vy
		}

		// Board'a topu yerleştir
		board[y][x] = true

		// Buffer'ı temizle
		buffer = buffer[:0]

		// Board'u buffer'a çiz
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				if board[j][i] {
					buffer = append(buffer, '⚽')
				} else {
					buffer = append(buffer, ' ')
				}
			}
			buffer = append(buffer, '\n')
		}

		fmt.Print(string(buffer))

		// Topu eski konumdan sil
		board[y][x] = false

		time.Sleep(time.Second / 20)
	}
}
