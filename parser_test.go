package parser

import (
	"strings"
	"testing"
)

var blob = `<!DOCTYPE html>
<html lang="en" data-theme="auto">
<head>

<link rel="preconnect" href="https://www.googletagmanager.com">
<script >(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00add8">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons">
<link rel="stylesheet" href="/css/styles.css">
<link rel="icon" href="/images/favicon-gopher.png" sizes="any">
<link rel="apple-touch-icon" href="/images/favicon-gopher-plain.png"/>
<link rel="icon" href="/images/favicon-gopher.svg" type="image/svg+xml">
<link rel="me" href="https://hachyderm.io/@golang">

  
  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<script src="/js/site.js"></script>
<meta name="og:url" content="https://go.dev/play/">
<meta name="og:title" content="Go Playground - The Go Programming Language">
<title>Go Playground - The Go Programming Language</title>

<meta name="og:image" content="https://go.dev/doc/gopher/gopher5logo.jpg">
<meta name="twitter:image" content="https://go.dev/doc/gopher/gopherbelly300.jpg">
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@golang">
</head>`

var simple = `<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@golang">`

var sibling = `<div><p>Paragraph 1</p><p>Paragraph 2</p></div>`

func TestNewDecoder(t *testing.T) {
	decoder := NewDecoder(strings.NewReader(simple))

	if decoder == nil {
		t.Error("NewDecoder should not return nil")
	}

	decoder.parse()
}
