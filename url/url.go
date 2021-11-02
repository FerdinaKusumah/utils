package url

import "net/url"

// parseUrl ...
func parseUrl(text string) (*url.URL, error) {
	return url.Parse(text)
}

// EscapeStringFromUrl ...
func EscapeStringFromUrl(text string) string {
	return url.QueryEscape(text)
}

// UnescapeStringFromUrl ...
func UnescapeStringFromUrl(text string) string {
	u, err := url.QueryUnescape(text)
	if err != nil {
		return ""
	}
	return u
}

// IsUrlValid ...
func IsUrlValid(text string) bool {
	_, err := parseUrl(text)
	return err == nil
}

// GetHostFromUrl ...
func GetHostFromUrl(text string) string {
	q, err := parseUrl(text)
	if err != nil {
		return ""
	}
	return q.Host
}

// GetSchemeFromUrl ...
func GetSchemeFromUrl(text string) string {
	q, err := parseUrl(text)
	if err != nil {
		return ""
	}
	return q.Scheme
}

// GetPathFromUrl ...
func GetPathFromUrl(text string) string {
	q, err := parseUrl(text)
	if err != nil {
		return ""
	}
	return q.Path
}

// GetRawPathFromUrl ...
func GetRawPathFromUrl(text string) string {
	q, err := parseUrl(text)
	if err != nil {
		return ""
	}
	return q.RawPath
}

// GetRawQueryFromUrl ...
func GetRawQueryFromUrl(text string) string {
	q, err := parseUrl(text)
	if err != nil {
		return ""
	}
	return q.RawQuery
}
