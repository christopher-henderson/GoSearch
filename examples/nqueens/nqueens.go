package main

import (
	"log"
	"time"
	"runtime"
)

type Queen struct {
	Column int
	Row int
}

func NQueens(N int) {
	in := make(chan int, 0)
	out := make(chan int, 0)
	go func() {
		winners := 0
		for range in {
			winners += 1
		}
		out <- winners
	}()
	search Queen{0, 0}; Queen; runtime.NumCPU() {
	children:
		column := node.Column + 1
		c := make(chan Queen, 0)
		// If the parent is in the final column
		// then there are no children.
		if column > N {
			close(c)
			return c
		}
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				c <- Queen{column, r}
			}
		}()
		return c
	accept:
		if len(solution) == N {
			// log.Println(solution)
			in <- 1
			return true
		}
		return false
	reject:
		row, column := node.Row, node.Column
		for _, q := range solution {
			r, c := q.Row, q.Column
			if row == r ||
				column == c ||
				row+column == r+c ||
				row-column == r-c {
				return true
			}
		}
		return false
	}
	close(in)
	log.Println(<-out)
	close(out)
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := time.Now()
	NQueens(8)
	log.Println(time.Now().Sub(s))
}
