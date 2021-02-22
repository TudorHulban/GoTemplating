package page

import (
	"html/template"
	"os"
	"strings"

	"github.com/TudorHulban/log"
)

type SiteInfo struct {
	Title    string
	Subtitle string
}

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
		p.l.Debug("Creating new page with new logger.")

		return &p, nil
	}

	return &Page{
		Nodes: [][]*Node{},
		l:     l,
	}, nil
}

// Add Method adds node.
func (p *Page) Add(pos uint, n *Node) error {
	if pos <= uint(len(p.Nodes)) {
		p.l.Infof("Adding node at level %d", pos)
		p.Nodes[pos-1] = append(p.Nodes[pos-1], n)
		return nil
	}

	p.l.Infof("Adding new level %d", pos)
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

func (p *Page) GetString() string {
	return strings.Join(p.GetHTML(), "\n")
}

func (p *Page) Render(renderTo string, model SiteInfo) error {
	funcs := template.FuncMap{"join": strings.Join}

	t, errParse := template.New("").Funcs(funcs).Parse(p.GetString())
	if errParse != nil {
		p.l.Warn("errParse: ", errParse)
		return errParse
	}

	f, errCreate := os.Create(renderTo)
	if errCreate != nil {
		p.l.Warn("errCreate: ", errCreate)
		return errCreate
	}
	defer f.Close()

	if errExec := t.Execute(f, model); errExec != nil {
		p.l.Warn("errExec: ", errExec)
		return errExec
	}

	return nil
}
