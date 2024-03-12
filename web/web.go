package web

import "regexp"

var (
	isSchemeRegExp = regexp.MustCompile(`^[^:]+://`)
	// Ref: https://github.com/git/git/blob/master/Documentation/urls.txt#L37
	isScpUrlRegExp = regexp.MustCompile(`^(?:(?P<user>[^@]+)@)?(?P<host>[^:\s]+):(?:(?P<port>[0-9]{1,5}):)?(?P<path>[^\\].*)$`)
)

func MatchesScheme(url string) bool {
	return isSchemeRegExp.MatchString(url)
}

func MatchesScp(url string) bool {
	return isScpUrlRegExp.MatchString(url)
}

// FindScpComponents returns the user, host, port and path of the given SCP URI.
func FindScpComponents(url string) (user, host, port, path string) {
	m := isScpUrlRegExp.FindStringSubmatch(url)
	return m[1], m[2], m[3], m[4]
}

// IsLocalEndpoint returns true if the given URL string specifies a local file endpoint. For example, on a Linux machine,
// `/home/user/src/inocencio/grinvon` matches as a local endpoint, but `https://github.com/inocencio/grinvon` wouldn't.
func IsLocalEndpoint(url string) bool {
	return !MatchesScheme(url) && !MatchesScp(url)
}
