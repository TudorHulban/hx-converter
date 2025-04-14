package hxconverter

import (
	"strings"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	"golang.org/x/net/html"
)

func processAnchor(node *html.Node) string {
	return hxhelpers.Sprintf(
		"\nhx.A(\n%s ,\nhx.Text(%s),\n)\n",

		mapNodesAttributes(
			extractAttributes(node),
		),
		`"`+strings.TrimSpace(extractText(node))+`"`,
	)
}

func processListItem(node *html.Node) string {
	return hxhelpers.Sprintf(
		"\nhx.Li(\n%s ,\nhx.Text(%s),\n)\n",

		mapNodesAttributes(
			extractAttributes(node),
		),
		`"`+strings.TrimSpace(extractText(node))+`"`,
	)
}

func processSpan(node *html.Node) string {
	return hxhelpers.Sprintf(
		"\nhx.Span(\n%s ,\nhx.Text(%s),\n)\n",

		mapNodesAttributes(
			extractAttributes(node),
		),
		`"`+strings.TrimSpace(extractText(node))+`"`,
	)
}

func processNav(node *html.Node) string {
	return hxhelpers.Sprintf(
		"\nhx.Nav(\n%s ,\nhx.Text(%s),\n)\n",

		mapNodesAttributes(
			extractAttributes(node),
		),
		`"`+strings.TrimSpace(extractText(node))+`"`,
	)
}

func processOrderedList(node *html.Node) string {
	return hxhelpers.Sprintf(
		"\nhx.Ol(\n%s ,\nhx.Text(%s),\n)\n",

		mapNodesAttributes(
			extractAttributes(node),
		),
		`"`+strings.TrimSpace(extractText(node))+`"`,
	)
}
