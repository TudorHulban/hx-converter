package hxconverter

import (
	"golang.org/x/net/html"
)

type NodeFunc func(*html.Node) string
