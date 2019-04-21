# pwaudit

A simple CLI that will check your passwords against the known pwned passwords on [haveibeenpwned](https://haveibeenpwned.com/).

## How does it work?

Your plaintext passwords are safe (as long as you plaintext file) and are not transmitted over the internet. The beginning of the sha1 hash of your password is sent over the internet to find potential matches, but the full sha1 is never sent over the internet; all full hash comparison is done locally. [Read more about the haveibeenpwned api](https://haveibeenpwned.com/API/v2#SearchingPwnedPasswordsByRange).

## Install

If you have a working golang installation:

```
go get -u github.com/scnewma/pwaudit
```

Otherwise you can download on of the pre-compiled releases from the [releases page](https://github.com/scnewma/pwaudit/releases).

## Usage

```
$ pwaudit -h
Usage: pwaudit [options]

  Check the provided passwords to see if they have been compromised.

Input Options:

  --lastpass-csv=path       Path to a LastPass exported CSV file. You can export
                            this file in the browser extension via
                            More Options > Advanced > Export

  --passwords=path|-        Path to a line-delimited password dump or '-'. If '-'
                            is provided then the passwords are read from stdin.

Output Options:

  --show-passwords          Print the plaintext passwords that were checked as well
                            as the password description. The description and the
                            password will be the same for some inputs.

  --show-all                Print both compromised and non-compromised passwords
                            to the screen.

Other Options:

  -v,--version              Print the version.
```
