package hxconverter

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var nodeProcessorMap = map[string]NodeFunc{
	"a":    processAnchor,
	"ol":   processOrderedList,
	"li":   processListItem,
	"nav":  processNav,
	"span": processSpan,
}

func traverseAST(node *html.Node) []string {
	var result []string

	switch node.Type {
	case html.ElementNode:
		if processFunc, exists := nodeProcessorMap[node.Data]; exists {
			nodeContent := processFunc(node)

			var childrenContent []string

			// Recursively process children and inject into the current node
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				childrenContent = append(
					childrenContent,
					traverseAST(child)...,
				)
			}

			// inject any children before the closing parentheses
			if len(childrenContent) > 0 {
				// Insert children content into the node string
				// Find the position of the closing parentheses for insertion
				injectionPoint := strings.LastIndex(nodeContent, "\n)")
				if injectionPoint != -1 {
					// Inject the children before the closing parentheses
					nodeContent = nodeContent[:injectionPoint] +
						strings.Join(childrenContent, "") +
						nodeContent[injectionPoint:]
				}
			}

			// Append the node with its children injected
			result = append(result, nodeContent)
		} else {
			// If node type is not handled, traverse children
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				result = append(
					result,
					traverseAST(child)...,
				)
			}
		}

	case html.TextNode:
		// Handle text nodes if needed (e.g., add to the result)

	case html.DocumentNode:
		// Recursively process document children
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			result = append(result, traverseAST(child)...)
		}

	default:
		// Non-element nodes can be ignored or logged
		fmt.Println("Encountered non-element node:", node.Type)
	}

	return result
}

func extractText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.FirstChild != nil {
		return extractText(node.FirstChild)
	}

	return ""
}
