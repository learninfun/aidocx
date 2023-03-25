package convert

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func JsonToString(content any) string {
	byteSlice, err := json.Marshal(content)
	checkErr(err)
	return string(byteSlice)
}

func JsonToStringBeauty(content any) string {
	byteSlice, err := json.MarshalIndent(content, "", "  ")
	checkErr(err)
	return string(byteSlice)
}

func MarkdownToHTML(markdown string) string {
	// Create a new Goldmark object with some extensions
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)

	// Convert the markdown to HTML
	var buf bytes.Buffer
	err := md.Convert([]byte(markdown), &buf)
	checkErr(err)

	return buf.String()
}
