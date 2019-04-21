package haveibeenpwned

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/scnewma/pwaudit/pkg/pw"
)

type PasswordChecker struct {
	Client *Client
}

func (c *PasswordChecker) Check(password pw.Password) (pw.CheckedPassword, error) {
	client := c.client()

	hash := sha1hash(password.Plaintext)
	hashPrefix := hash[:5]

	checked := pw.CheckedPassword{
		Password:    password,
		Compromised: false,
	}

	pwnedPasswords, err := client.PwnedPasswordsByRange(hashPrefix)
	if err != nil {
		return checked, err
	}

	for _, p := range pwnedPasswords {
		pHash := hashPrefix + p.HashSuffix

		if strings.ToUpper(pHash) == strings.ToUpper(hash) {
			checked.Compromised = true
			return checked, nil
		}
	}

	return checked, nil
}

func (c *PasswordChecker) client() *Client {
	if c.Client != nil {
		return c.Client
	}
	return DefaultClient
}

func sha1hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
