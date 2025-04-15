package hxconverter

import (
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	"golang.org/x/net/html"
)

func processElement(elName string, node *html.Node) string {
	return hxhelpers.Sprintf(
		"\n%s.%s(\n%s,\n),",

		_PackagePrimitives,
		elName,
		mapNodesAttributes(
			extractAttributes(node),
		),
	)
}

func processAnchor(node *html.Node) string {
	return processElement(
		"A",
		node,
	)
}

func processDiv(node *html.Node) string {
	return processElement(
		"Div",
		node,
	)
}

func processH1(node *html.Node) string {
	return processElement(
		"H1",
		node,
	)
}

func processH2(node *html.Node) string {
	return processElement(
		"H2",
		node,
	)
}

func processH3(node *html.Node) string {
	return processElement(
		"H3",
		node,
	)
}

func processH4(node *html.Node) string {
	return processElement(
		"H4",
		node,
	)
}

func processH5(node *html.Node) string {
	return processElement(
		"H5",
		node,
	)
}

func processH6(node *html.Node) string {
	return processElement(
		"H6",
		node,
	)
}

func processImg(node *html.Node) string {
	return processElement(
		"Img",
		node,
	)
}

func processListItem(node *html.Node) string {
	return processElement(
		"Li",
		node,
	)
}

func processP(node *html.Node) string {
	return processElement(
		"P",
		node,
	)
}

func processSpan(node *html.Node) string {
	return processElement(
		"Span",
		node,
	)
}

func processNav(node *html.Node) string {
	return processElement(
		"Nav",
		node,
	)
}

func processOrderedList(node *html.Node) string {
	return processElement(
		"Ol",
		node,
	)
}
