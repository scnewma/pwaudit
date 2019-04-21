package lastpass

import "net/url"

// Entry represents a set of values in LastPass (website, secure note, etc)
type Entry struct {
	URL      url.URL
	Username string
	Password string
	Extra    string
	Name     string
	Fav      int
}
