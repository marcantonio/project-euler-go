
package main

import (
	"fmt"
	"math"
)

// Must be odd.
const Size = 1001

type Direction int
const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	value, x, y int
	direction Direction
}

func (p *Position) next() {
	switch p.direction {
	case Up:
		p.y -= 1
	case Down:
		p.y += 1
	case Left:
		p.x -= 1
	case Right:
		p.x += 1
	}
	p.value++
}

func (p *Position) changeDirection() {
	switch p.direction {
	case Up:
		p.direction = Right
	case Down:
		p.direction = Left
	case Left:
		p.direction = Up
	case Right:
		p.direction = Down
	}
}

func (p Position) String() string {
	return fmt.Sprintf("[%v %v,%v %v]", p.value, p.x, p.y, p.direction)
}

func main() {
	square := make([][]int, Size, Size)
	for i := range square {
		square[i] = make([]int, Size, Size)
	}
	
	p := Position{1, mid(Size), mid(Size), Right}
	for n, c := 1, 0; n <= Size; {
		for m := 0; m < n; m++ {
			square[p.y][p.x] = p.value
			p.next()
		}
		p.changeDirection()
		c++
		
		// The pattern breaks here.  The last number should only go once.
		if (c == 1) && (n == Size) {
			break
		}
		
		if c == 2 {
			n++
			c = 0
		}
	}
	
//	for y := range square {
//		fmt.Print("[ ")
//		for x := range square[y] {
//			fmt.Printf("%2d ", square[y][x])
//		}
//		fmt.Println("]")
//	}
	
	sum := 0
	for n := 0; n < len(square); n++ {
		if n == mid(Size) {
			sum += 1
			continue
		}
		sum += square[n][n]
		sum += square[n][len(square)-1-n]
	}
	fmt.Println(sum)
}

func mid(x int) int {
	return int(math.Floor(float64(x) / 2))
}
