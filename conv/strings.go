package conv

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
