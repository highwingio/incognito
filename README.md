# Incognito
A wrapper script for generating tokens from AWS Cognito


## Installation
Since `go get` doesn't support modules just yet, it's easiest to install using `https://github.com/FiloSottile/homebrew-gomod`:

```bash
brew install FiloSottile/gomod/brew-gomod
brew gomod github.com/highwingio/incognito
export AWS_REGION=<region>
```

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
