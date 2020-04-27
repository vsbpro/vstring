package vstring

import "fmt"

// Token struct to encapsulate a token encountered in a string.
type Token struct {
	// StartDelimiter is the delimiter on the left.
	StartDelimiter rune
	// Value between the two delimiters.
	Value string
	// EndDelimiter is the delimiter on the right.
	EndDelimiter rune
}

func (t *Token) String() string {
	var start, end string
	if t.StartDelimiter == Nothing {
		start = "Nothing"
	}else {
		start = fmt.Sprintf("%c",t.StartDelimiter)
	}
	if t.EndDelimiter == Nothing {
		end = "Nothing"
	}else {
		end = fmt.Sprintf("%c",t.EndDelimiter)
	}
	return fmt.Sprintf("Start: %s, Value: %s, End: %s\n",start,t.Value,end)
}

const Nothing = int32(-9999)

// SplitByMultipleDelimiters scans a string and look for the delimiters which are the single characters in the 'delimiters' string.
// It returns the array to the pointers of Token.
func SplitByMultipleDelimiters(s string, delimiters string) []*Token {
	var tokens []*Token
	l := len(s)
	ld := len(delimiters)
	d1 := Nothing // Start delimiter for the token.
	d2 := Nothing // End delimiter for the token.
	lastValueOffset := 0
	for i, v := range s[0:l] {
		for _ ,d := range delimiters[0:ld] {
			if v == d {
				// Is it the first time delimiter encountered?
				if d2 == Nothing {
					// Create token only if there are some contents for value.
					if i > 0 {
						t := &Token{
							StartDelimiter: Nothing,
							Value:          s[0:i],
							EndDelimiter:   d,
						}
						lastValueOffset = i
						tokens = append(tokens,t)
					}
					// Assign d to d2 and break the inner loop.
					d2 = d
					break
				}
				d1 = d2
				d2 = d
				t := &Token{StartDelimiter: d1,	Value: s[lastValueOffset+1:i], EndDelimiter:   d2}
				lastValueOffset = i
				tokens = append(tokens,t)
				break
			}
		}
	}
	if lastValueOffset > 0 && lastValueOffset != l {
		t := &Token{StartDelimiter: d2,	Value: s[lastValueOffset+1:], EndDelimiter:   Nothing}
		tokens = append(tokens,t)
	}
	return tokens
}