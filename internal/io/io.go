package io

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"gopkg.in/yaml.v3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FileCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// exists returns whether the given file or directory exists
func FileExists(path string) bool {
	_, error := os.Stat(path)
	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func StringToFile(filePath, content string) {
	fo, err := os.Create(filePath)
	checkErr(err)

	fo.WriteString(content)

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
}

func FileToString(filePath string) string {
	dat, err := os.ReadFile(filePath)
	checkErr(err)
	return string(dat)
}

func FileNameNoExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func YamlFileToStruct(filePath string, out interface{}) {
	// read YAML file
	data, err := os.ReadFile(filePath)
	checkErr(err)

	// parse YAML
	err = yaml.Unmarshal(data, out)
	checkErr(err)
}

type TreeNode struct {
	Name     string
	Parent   *TreeNode
	Children []*TreeNode
}

func MdListToTreeNode(filePath string) *TreeNode {
	inputStr := FileToString(filePath)
	fmt.Println(inputStr)

	source := []byte(inputStr)
	reader := text.NewReader(source)

	mdParser := goldmark.DefaultParser()
	node := mdParser.Parse(reader)

	rootTreeNode := &TreeNode{Name: "Dummy"}
	//child := &TreeNode{Name: "Child", Parent: rootNode}
	//rootNode.Children = append(rootNode.Children, child)

	currentTreeNode := rootTreeNode
	//level := 0
	// Traverse the Markdown AST and find all list nodes
	// Traverse the Markdown AST and find all list nodes
	//resultPath := outputFolder
	Walk(node, func(node ast.Node, entering bool, idx int) (ast.WalkStatus, error) {
		if listItem, ok := node.(*ast.ListItem); ok {
			listText := string(listItem.FirstChild().Text(source))
			if entering {
				childTreeNode := &TreeNode{Name: listText, Parent: currentTreeNode}
				currentTreeNode.Children = append(currentTreeNode.Children, childTreeNode)
				currentTreeNode = childTreeNode

			} else {
				//resultPath = resultPath[:len(resultPath)-1-len(listText)]
				currentTreeNode = currentTreeNode.Parent
			}
		}

		return ast.WalkContinue, nil
	}, 0)

	fmt.Println(currentTreeNode)

	return rootTreeNode
}

type Walker func(n ast.Node, entering bool, idx int) (ast.WalkStatus, error)

// Walk walks a AST tree by the depth first search algorithm.
func Walk(n ast.Node, walker Walker, idx int) error {
	_, err := walkHelper(n, walker, idx)
	return err
}

func walkHelper(n ast.Node, walker Walker, idx int) (ast.WalkStatus, error) {
	status, err := walker(n, true, idx)
	if err != nil || status == ast.WalkStop {
		return status, err
	}
	if status != ast.WalkSkipChildren {
		var lIdx = 1
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			if st, err := walkHelper(c, walker, lIdx); err != nil || st == ast.WalkStop {
				return ast.WalkStop, err
			}

			lIdx++
		}
	}
	status, err = walker(n, false, idx)
	if err != nil || status == ast.WalkStop {
		return ast.WalkStop, err
	}
	return ast.WalkContinue, nil
}
