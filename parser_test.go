package parser

import (
	"strings"
	"testing"
)

var simple = `
<!DOCTYPE html>
<html>
</html>
`

var simple2 = `
<!DOCTYPE html>
<html>
   <head>
      <title>Test</title>
   </head>
   <body>
      <h1>Test</h1>
   </body>
</html>
`

var withAttributes = `
<!DOCTYPE html>
<html lang="en">
   <head>
      <title>Test</title>
   </head>
   <body>
      <h1>Test</h1>
   </body>
</html>
`

var withLinks = `
<!DOCTYPE html>
<html lang="en">
   <head>
      <title>Test</title>
      <link rel="stylesheet" href="style.css">
   </head>
   <body>
      <h1>Test</h1>
   </body>
</html>
`

func TestParser(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"simple", simple},
		{"simple2", simple2},
		{"withAttributes", withAttributes},
		{"withLinks", withLinks},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newParser(strings.NewReader(tt.in))
			tr := p.parse()
			for _, child := range tr.root.children {
				if child.Type() == "element" {
					t.Logf("Element: %s", child.Data())
				}
			}
		})
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"simple2", simple2},
	}
	for _, tt := range tests {
		want := []string{"html", "head", "title", "body", "h1"} // FIXME node elements that close should be not be included
		t.Log(want)
		t.Run(tt.name, func(t *testing.T) {
			p := newParser(strings.NewReader(tt.in))
			tr := p.parse()
			for _, child := range tr.root.children {
				if child.Type() == "element" {
					t.Logf("Element: %s", child.Data())
				}
			}
		})
	}
}
