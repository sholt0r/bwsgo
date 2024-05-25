# Bitwarden Secret Manager for Go (bwsgo)

This is a Go package that provides a function to retrieve secrets from Bitwarden secret manager.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go programming language
- [Bitwarden Secrets CLI](https://github.com/bitwarden/sdk/releases)
- [jq command-line JSON processor](https://jqlang.github.io/jq/)

### Installing

1. Get the module from Github
```go
go get github.com/sholt0r/bwsgo
```
2. Import it
```go
import "github.com/sholt0r/bwsgo"
```

## Usage

To use the `GetSecret` function, pass the name of the environment variable that holds the secret key as an argument. Here's an example:

```go
import "github.com/sholt0r/bwsgo"
secret := bws.GetSecret("UUID_ENVIRONMENT_VAR")
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/sholt0r/bws-go/LICENSE.md) file for details
