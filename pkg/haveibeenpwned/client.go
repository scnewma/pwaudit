package haveibeenpwned

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type PwnedPassword struct {
	HashSuffix string
	Count      int
}

const (
	defaultAPI       = "https://api.pwnedpasswords.com"
	defaultUserAgent = "pwchecker"
	hashLen          = 5
)

type Client struct {
	API        string
	HttpClient *http.Client
	UserAgent  string
}

var DefaultClient = &Client{}

func (c *Client) PwnedPasswordsByRange(hashPrefix string) ([]PwnedPassword, error) {
	if len(hashPrefix) != hashLen {
		return nil, fmt.Errorf("hashPrefix must be %d chars", hashLen)
	}
	body, err := c.get(fmt.Sprintf("/range/%s", hashPrefix))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\r\n")
	return parsePwnedPasswords(lines)
}

func PwnedPasswordsByRange(hash string) ([]PwnedPassword, error) {
	return DefaultClient.PwnedPasswordsByRange(hash)
}

func (c *Client) get(path string) ([]byte, error) {
	endpoint := c.api() + path
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent())

	resp, err := c.httpClient().Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) api() string {
	if c.API != "" {
		return c.API
	}
	return defaultAPI
}

func (c *Client) httpClient() *http.Client {
	if c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}

func (c *Client) userAgent() string {
	if c.UserAgent != "" {
		return c.UserAgent
	}
	return defaultUserAgent
}

func parsePwnedPasswords(lines []string) ([]PwnedPassword, error) {
	pwned := make([]PwnedPassword, len(lines))
	for i, line := range lines {
		pwnedPassword, err := parsePwnedPasswordLine(line)
		if err != nil {
			return nil, err
		}
		pwned[i] = pwnedPassword
	}
	return pwned, nil
}

func parsePwnedPasswordLine(line string) (PwnedPassword, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return PwnedPassword{}, fmt.Errorf("invalid pwned password response: %v", line)
	}

	hashSuffix := parts[0]
	count, err := strconv.Atoi(parts[1])
	if err != nil {
		return PwnedPassword{}, fmt.Errorf("invalid pwned password response: %v", line)
	}

	return PwnedPassword{HashSuffix: hashSuffix, Count: count}, nil
}
