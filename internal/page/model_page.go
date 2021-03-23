package page

import (
	"html/template"
	"os"
	"strings"

	"github.com/rs/zerolog"
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
	l     zerolog.Logger
}

func NewPage(l zerolog.Logger) (*Page, error) {
	return &Page{
		Nodes: [][]*Node{},
		l:     l,
	}, nil
}

// Add Method adds node.
func (p *Page) Add(pos uint, n *Node) error {
	if pos <= uint(len(p.Nodes)) {
		p.l.Info().Msgf("Adding node at level %d", pos)

		p.Nodes[pos-1] = append(p.Nodes[pos-1], n)
		return nil
	}

	p.l.Info().Msgf("Adding new level %d", pos)
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
	t, errParse := template.New("").Parse(p.GetString())
	if errParse != nil {
		p.l.Warn().Str("errParse", errParse.Error()).Msg("")
		return errParse
	}

	f, errCreate := os.Create(renderTo)
	if errCreate != nil {
		p.l.Warn().Msgf("error creating file into which to render: %s", errCreate.Error())
		return errCreate
	}
	defer f.Close()

	if errExec := t.Execute(f, model); errExec != nil {
		p.l.Warn().Msgf("error parsing template: %s", errExec.Error())
		return errExec
	}

	return nil
}
