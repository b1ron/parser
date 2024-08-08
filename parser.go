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

type tree struct {
	element *node
	next    *node
}

type node struct {
	name  string
	value string
	next  *node
}

type Decoder struct {
	s    scanner.Scanner
	b    bytes.Buffer
	tree *node // DOM tree
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		s:    scanner.Scanner{},
		b:    bytes.Buffer{},
		tree: &node{},
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
			d.s.Scan()
			switch d.s.TokenText() {
			case HTML:
			case BASE:
			case HEAD:
			case LINK:
			case META:
				// <meta name="twitter:card" content="summary">
				// <meta name="twitter:site" content="@golang">
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
