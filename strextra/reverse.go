/*
Package strextra implements additional functions to manipulate UTF-8
encoded strings, beyond what is provided in the standard "strings" package.

Every package should have a package comment, a block comment preceding the package clause.
The package comment should introduce the package and provide information relevant to the package
as a whole.

If the package is simple, the package comment can be brief.

    cmd:
        $ go doc -all regexp | grep -i <word>
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
*/
package strextra

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
