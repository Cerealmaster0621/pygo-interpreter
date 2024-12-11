package tokenizer

import (
	"regexp"
	"strings"
)

var (
	whitespace = `^[ \f\t]*`
	comment    = `#[^\r\n]*`
	ignore     = whitespace + any(`\\\r?\n`+whitespace) + maybe(comment)
	name       = `\w+`

	hexnumber   = `0[xX](?:_?[0-9a-fA-F])`
	binnumber   = `0[bB](?:_?[01])+`
	octnumber   = `0[oO](?:_?[0-7])+`
	decnumber   = `(?:0(?:_?0)*|[1-9](?:_?[0-9])*)`
	intnumber   = group(hexnumber, binnumber, octnumber, decnumber)
	exponent    = `[eE][-+]?[0-9](?:_?[0-9])*`
	pointfloat  = group(`[0-9](?:_?[0-9])*\.(?:[0-9](?:_?[0-9])*)?`, `\.[0-9](?:_?[0-9])*`+maybe(exponent))
	expfloat    = `[0-9](?:_?[0-9])*` + exponent
	floatnumber = group(pointfloat, expfloat)
	imagnumber  = group(`[0-9](?:_?[0-9])*[jJ]`, floatnumber+`[jJ]`)
	number      = group(imagnumber, floatnumber, intnumber)

	single  = `[^'\\]*(?:\\.[^'\\]*)*'`
	double  = `[^"\\]*(?:\\.[^"\\]*)*"`
	single3 = `[^'\\]*(?:(?:\\.|'(?!''))[^'\\]*)*'''`
	double3 = `[^"\\]*(?:(?:\\.|"(?!""))[^"\\]*)*"""`
	triple  = group()
)

// Regex expressions
var (
	whitespaceRegex  = regexp.MustCompile(whitespace)  // find whitespaces
	commentRegex     = regexp.MustCompile(comment)     // find Comments # in python
	hexnumberRegex   = regexp.MustCompile(hexnumber)   // find hexnumber
	binnumberRegex   = regexp.MustCompile(binnumber)   // find binnumber
	octnumberRegex   = regexp.MustCompile(octnumber)   // find octnumber
	decnumberRegex   = regexp.MustCompile(decnumber)   // find decnumber
	IntnumberRegex   = regexp.MustCompile(intnumber)   // find intnumber
	ExponentRegex    = regexp.MustCompile(exponent)    // find exponent
	PointfloatRegex  = regexp.MustCompile(pointfloat)  // find pointfloat
	expfloatRegex    = regexp.MustCompile(expfloat)    // find expfloat
	floatnumberRegex = regexp.MustCompile(floatnumber) // find floatnumber
	imagnumberRegex  = regexp.MustCompile(imagnumber)  // find i number
	numberRegex      = regexp.MustCompile(number)      // find number
	opRegex          = regexp.MustCompile(`^[+\-*/=<>!]+`)
	nameRegex        = regexp.MustCompile(`^[A-Za-z_]\w*`)
)

// group binds multiple regex expressions
func group(expressions ...string) string {
	return "(?:" + strings.Join(expressions, "|") + ")"
}

// any creates a regex that matches 0 OR MORE occurences of the expressions
// ex - any("a", "b") -> "" , "a", "b", "aa", "bb", "ab"
func any(expressions ...string) string {
	return group(expressions...) + `*`
}

// maybe creates a regex that matches 0 OR 1 occurrences of the given expressions
// ex - maybe("a", "b) -> "", "a", "b"
func maybe(expressions ...string) string {
	return group(expressions...) + `?`
}
