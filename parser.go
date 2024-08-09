package parser

/* HTML parser */

import (
	"bytes"
	"fmt"
	"io"
	"text/scanner"
)

// HTML elements
const (
	HTML    = "html"
	BASE    = "base"
	HEAD    = "head"
	LINK    = "link"
	META    = "meta"
	STYLE   = "style"
	TITLE   = "title"
	BODY    = "body"
	ADDRESS = "address"
	ARTICLE = "article"
	ASIDE   = "aside"
	FOOTER  = "footer"
	HEADER  = "header"
	H1      = "h1"
	H2      = "h2"
	H3      = "h3"
	H4      = "h4"
	H5      = "h5"
	H6      = "h6"
	HGROUP  = "hgroup"
	MAIN    = "main"
	NAV     = "nav"
	SECTION = "section"
	SEARCH  = "search"
)

// DOM tree
type tree struct {
	element *node
	next    *node
}

type node struct {
	name  string
	value string
	next  *node
}

// a stack to keep track of the DOM tree structure
type stack struct {
	b []rune
}

func (s *stack) push(r rune) {
	s.b = append(s.b, r)
}

func (s *stack) pop() rune {
	if len(s.b) == 0 {
		return '0'
	}
	r := s.b[len(s.b)-1]
	s.b = s.b[:len(s.b)-1]
	return r
}

func (s *stack) peek() rune {
	if len(s.b) == 0 {
		return '0'
	}
	return s.b[len(s.b)-1]
}

func (s *stack) empty() bool {
	return len(s.b) == 0
}

type Decoder struct {
	s     scanner.Scanner
	stack stack
	b     bytes.Buffer
	tree  *node
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		s:     scanner.Scanner{},
		stack: stack{},
		b:     bytes.Buffer{},
		tree:  &node{},
	}

	d.s.Init(r)
	return d
}

func (d *Decoder) String() string {
	return d.b.String()
}

func (d *Decoder) parse() error {
	for tok := d.s.Scan(); tok != scanner.EOF; tok = d.s.Scan() {
		switch tok {
		case '<':
			// start tag
			d.stack.push(tok) // TODO: we need to push and pop full tag names
			d.s.Scan()
			switch d.s.TokenText() {
			case HTML:
			case BASE:
			case HEAD:
			case LINK:
			case META:
				// <meta name="twitter:card" content="summary">
				// <meta name="twitter:site" content="@golang">
				// <div><p>Paragraph 1</p><p>Paragraph 2</p></div>
				d.tree.next = &node{name: d.s.TokenText()}
				fmt.Println(d.tree.name, d.tree.next.name)
			case STYLE:
			case TITLE:
			case BODY:
			case ADDRESS:
			case ARTICLE:
			case ASIDE:
			case FOOTER:
			case HEADER:
			case H1:
			case H2:
			case H3:
			case H4:
			case H5:
			case H6:
			case HGROUP:
			case MAIN:
			case NAV:
			case SECTION:
			case SEARCH:
			}
		case '>':
			// end tag
		default:
			s := d.s.TokenText()
			if s == "name" {
				d.b.WriteString(d.s.TokenText())
			}
		}
	}
	return nil
}
