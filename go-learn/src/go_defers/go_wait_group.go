package go_defers
//https://golangbot.com/defer/
import (
	"sync"
	"fmt"
)

type rect struct {
	length int
	width int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v length less then 0\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v width less then 0\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v area %d\n", r, area )
}

func WaitGroupDemo()  {
	var wg sync.WaitGroup
	r1 := rect{-1, 100}
	r2 := rect{5, -10}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}