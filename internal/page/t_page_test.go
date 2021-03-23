package page

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	p, _ := NewPage(l)

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
		HTML: "x",
	}))

	require.Nil(t, p.Add(level, &Node{
		Name: "2b",
		HTML: "y",
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

	info := SiteInfo{
		Title:    "Title",
		Subtitle: "Subtitle",
	}

	require.Nil(t, p.Render("/home/tudi/ram/xxxx", info))
}
