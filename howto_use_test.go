package hxconverter

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHowToUse(t *testing.T) {
	htmlData := `
	<li class="breadcrumb-item">
	<a rel="noopener noreferrer" href="#" title="Back to homepage" class="breadcrumb-link">Home</a>
	</li>
	<li class="breadcrumb-item">
	<span class="breadcrumb-separator">/</span>
	<a rel="noopener noreferrer" href="#" class="breadcrumb-link">Parent</a>
	</li>
	`

	doc, errParse := html.Parse(strings.NewReader(htmlData))
	require.NoError(t, errParse)
	require.NotNil(t, doc)
	require.NotEmpty(t, doc)

	nodes := traverseAST(doc)
	require.NotEmpty(t, nodes)

	fmt.Println(
		nodes,
	)
}

func TestHowToUseWithURL(t *testing.T) {
	testURL := "https://example.com"

	parsedURL, err := url.ParseRequestURI(testURL)
	require.NoError(t, err)
	require.NotEmpty(t, parsedURL)

	resp, err := http.Get(parsedURL.String())
	require.NoError(t, err)

	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	doc, errParse := html.Parse(resp.Body)
	require.NoError(t, errParse)
	require.NotNil(t, doc)
	require.NotEmpty(t, doc)

	nodes := traverseAST(doc)
	require.NotEmpty(t, nodes)

	fmt.Println(
		nodes,
	)
}
