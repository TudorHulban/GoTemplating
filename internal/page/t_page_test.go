package page

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	p, _ := NewPage(nil)

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "1",
		HTML: "x",
	})

	fmt.Println("Page:")
	fmt.Println(p.GetHTML())
}
