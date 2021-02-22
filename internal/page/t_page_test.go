package page

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	p, _ := NewPage(nil)

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Page Start",
		HTML: PageStart,
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Head",
		HTML: HEAD,
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Body Start",
		HTML: BODYStart,
	}))

	level := p.GetCurrentPos() + 1
	require.Nil(t, p.Add(level, &Node{
		Name: "2a",
		HTML: "y",
	}))

	require.Nil(t, p.Add(level, &Node{
		Name: "2b",
		HTML: "z",
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Body End",
		HTML: BODYEnd,
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Footer",
		HTML: FOOTER,
	}))

	require.Nil(t, p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Page End",
		HTML: PageEnd,
	}))

	//require.Equal(t, len(p.Nodes), 4)

	fmt.Println("Page:")
	fmt.Println(p.GetString())
}
