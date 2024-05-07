package tree_sitter_bibtex_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-bibtex"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_bibtex.Language())
	if language == nil {
		t.Errorf("Error loading Bibtex grammar")
	}
}
