package hxconverter

import (
	"strings"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
)

func mapNodeAttribute(nodeKey, nodeValue string) string {
	switch nodeKey {
	case "class":
		return hxhelpers.Sprintf(
			`hx.AttrClass("%s")`,
			nodeValue,
		)

	case "href":
		return hxhelpers.Sprintf(
			`hx.Href("%s")`,
			nodeValue,
		)

	default:
		return hxhelpers.Sprintf(
			`hx.AttrWithValue("%s","%s")`,
			nodeKey,
			nodeValue,
		)
	}
}

func mapNodesAttributes(data map[string]string) string {
	result := make([]string, 0)

	for node, value := range data {
		result = append(
			result,
			mapNodeAttribute(
				node,
				value,
			),
		)
	}

	return strings.Join(
		result,
		",\n",
	)
}
