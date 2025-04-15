package hxconverter

import (
	"strings"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
)

func mapNodeAttribute(nodeKey, nodeValue string) string {
	switch nodeKey {
	case "class":
		return hxhelpers.Sprintf(
			`%s.AttrClass("%s")`,

			_PackagePrimitives,
			nodeValue,
		)

	case "href":
		return hxhelpers.Sprintf(
			`%s.Href("%s")`,

			_PackageHTML,
			nodeValue,
		)

	default:
		return hxhelpers.Sprintf(
			`%s.AttrWithValue("%s","%s")`,

			_PackagePrimitives,
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
