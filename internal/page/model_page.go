package page

import (
	"os"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

type Node struct {
	Name string
	HTML string
}

type Page struct {
	Nodes [][]*Node
	l     *log.LogInfo
}

func NewPage(l *log.LogInfo) (*Page, error) {
	if l == nil {
		p := Page{
			Nodes: [][]*Node{},
			l:     log.New(log.DEBUG, os.Stdout, true),
		}
		p.l.Debug("creating new page with new logger")

		return &p, nil
	}

	return &Page{
		Nodes: [][]*Node{},
		l:     l,
	}, nil
}

// Add Method adds node.
func (p *Page) Add(pos uint, n *Node) error {
	if pos > uint(len(p.Nodes)) {
		return errors.New("invalid position for node adding")
	}

	if pos <= uint(len(p.Nodes)) {
		p.Nodes[pos] = append(p.Nodes[pos], n)
		return nil
	}

	p.Nodes = append(p.Nodes, []*Node{n})

	return nil
}

func (p *Page) GetCurrentPos() uint {
	return uint(len(p.Nodes))
}

func (p *Page) GetHTML() []string {
	var result []string

	for _, nodeLevel := range p.Nodes {
		for _, node := range nodeLevel {
			result = append(result, node.HTML)
		}
	}

	return result
}
