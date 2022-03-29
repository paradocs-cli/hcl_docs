package hcl_docs

import (
	"fmt"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"strconv"
)

type CommentData struct {
	Filename string
	Line     string
	Text     string
}

//ParseHclBlocks takes an input of string and parses HCL blocks and comments
func ParseHclBlocks(s string) (ast.File, error) {
	parse, err := hcl.Parse(s)
	if err != nil {
		return *parse, fmt.Errorf(err.Error())
	}
	return *parse, nil
}

//ReadBlockComments takes an input of type ast.File and returns data about comments in the present hcl blocks
func ReadBlockComments(a ast.File) []CommentData {
	var comms []CommentData
	for _, v := range a.Comments {
		var c CommentData
		for _, z := range v.List {
			c.Filename = z.Pos().Filename
			c.Line = strconv.Itoa(z.Pos().Line)
			c.Text = z.Text
			comms = append(comms, c)
		}
	}
	return comms
}
func ParseVarBlocks()     {}
func ParseVarDefsBlocks() {}
