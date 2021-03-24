package page

import (
	"html/template"
	"os"
	"strings"

	"github.com/TudorHulban/GoTemplating/internal/article"
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
	l     *log.Logger
}

type Option func(p *Page) error

func NewPage(l *log.Logger, options ...Option) (*Page, error) {
	result := &Page{
		Nodes: [][]*Node{},
		l:     l,
	}

	for _, opt := range options {
		if err := opt(result); err != nil {
			return nil, err
		}
	}

	return result, nil
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

// GetHTML Method returns page HTML as slice of string.
func (p *Page) GetHTML() []string {
	var result []string

	for _, nodeLevel := range p.Nodes {
		for _, node := range nodeLevel {
			result = append(result, node.HTML)
		}
	}

	return result
}

// GetHTML Method returns page HTML as string.
func (p *Page) GetString() string {
	return strings.Join(p.GetHTML(), "\n")
}

func (p *Page) RenderArticle(a article.Article) error {
	t, errParse := template.New("").Parse(p.GetString())
	if errParse != nil {
		p.l.Warn("errParse", errParse)
		return errParse
	}

	f, errCreate := os.Create(a.CODE + ".html")
	if errCreate != nil {
		p.l.Warnf("error creating file into which to render: %s", errCreate.Error())
		return errCreate
	}
	defer f.Close()

	// model contains site and article information
	model := struct {
		string
	}{
		"x",
	}

	if errExec := t.Execute(f, model); errExec != nil {
		p.l.Warnf("error parsing template: %s", errExec.Error())
		return errExec
	}

	return nil
}
