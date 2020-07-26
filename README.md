# Incognito
A wrapper script for retrieving tokens from AWS Cognito


## Installation
`go get github.com/highwingio/incognito`

## Usage

Once installed, you'll need to run `incognito login` once with your username,
password, user pool ID, and client ID, which will confirm those credentials with
Cognito then store them securely using
[Keyring](https://github.com/99designs/keyring).

After that, you can run `incognito generate` anytime to generate a new Cognito
token to be used elsewhere. You can use something like the following to
automatically copy it to your clipboard:

```bash
$ incognito generate | pbcopy
```

## Development
Incognito is written in Go.

We use [Cobra](https://github.com/spf13/cobra) to structure/manage the CLI
commands.

There are no tests currently, but PRs are welcome!
