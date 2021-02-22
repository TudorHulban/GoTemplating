package page

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	p, _ := NewPage(nil)

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "1",
		HTML: "x",
	}))

	level := p.GetCurrentPos() + 1
	require.Nil(t, p.Add(level, &Node{
		Name: "2a",
		HTML: "y",
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "3",
		HTML: "xxxxxxxx",
	}))

	require.Nil(t, p.Add(level, &Node{
		Name: "2b",
		HTML: "z",
	}))

	require.Equal(t, len(p.Nodes), 3)

	fmt.Println("Page:")
	fmt.Println(p.GetString())
}
