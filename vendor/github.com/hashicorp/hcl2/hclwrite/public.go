package hclwrite

import (
	"bytes"

	"github.com/hashicorp/hcl2/hcl"
)

// ParseConfig interprets the given source bytes into a *hclwrite.File. The
// resulting AST can be used to perform surgical edits on the source code
// before turning it back into bytes again.
func ParseConfig(src []byte, filename string, start hcl.Pos) (*File, hcl.Diagnostics) {
	return parse(src, filename, start)
}

// Format takes source code and performs simple whitespace changes to transform
// it to a canonical layout style.
//
// Format skips constructing an AST and works directly with tokens, so it
// is less expensive than formatting via the AST for situations where no other
// changes will be made. It also ignores syntax errors and can thus be applied
// to partial source code, although the result in that case may not be
// desirable.
func Format(src []byte) []byte {
	tokens := lexConfig(src)
	format(tokens)
	buf := &bytes.Buffer{}
	(&TokenSeq{tokens}).WriteTo(buf)
	return buf.Bytes()
}
