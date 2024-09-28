package css_parser

import (
	"strings"

	"github.com/reesericci/esbuild/pkg/css_ast"
	"github.com/reesericci/esbuild/pkg/css_lexer"
)

func (p *parser) mangleFontWeight(token css_ast.Token) css_ast.Token {
	if token.Kind != css_lexer.TIdent {
		return token
	}

	switch strings.ToLower(token.Text) {
	case "normal":
		token.Text = "400"
		token.Kind = css_lexer.TNumber
	case "bold":
		token.Text = "700"
		token.Kind = css_lexer.TNumber
	}

	return token
}
