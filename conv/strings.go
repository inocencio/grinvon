package conv

import "strings"

type SearchResult struct {
	Text     string
	Elements []string
	Tokens   []*Token
}

type Token struct {
	StartPos int
	EndPos   int
	Value    string
}

// IsEqualsValues checks if two string values are equal ignoring leading and trailing whitespaces,
// and case insensitive. It returns true if the values are equal, otherwise false.
// Leading and trailing whitespaces as well as null and carriage return characters are trimmed
// before comparing the values. The comparison is case-insensitive.
func IsEqualsValues(v1, v2 string) bool {
	v1 = strings.Trim(v1, "\n\r\x00")
	v2 = strings.Trim(v2, "\n\r\x00")

	v1 = strings.ToLower(strings.TrimSpace(v1))
	v2 = strings.ToLower(strings.TrimSpace(v2))

	return strings.ToLower(v1) == strings.ToLower(v2)
}

func (sr *SearchResult) HasElements() bool {
	if len(sr.Tokens) > 0 {
		return true
	}

	return false
}

func (sr *SearchResult) HasAllElements() bool {
	var items = make([]string, 0)

	for _, e := range sr.Elements {
		for _, token := range sr.Tokens {

			if token.Value == e {

				found := false
				if len(items) > 0 {
					for _, v := range items {
						if v == e {
							found = true
						}
					}

					if !found {
						items = append(items, e)
					}
				} else {
					items = append(items, e)
				}

			}

		}
	}

	if len(items) == len(sr.Elements) {
		return true
	}

	return false
}

func SearchAll(text string, elements ...string) {

}
