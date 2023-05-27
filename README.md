# Go OAuth2 Providers
go-oauth2-providers is a Go library for OAuth2 authentication integration. It provides a unified interface to create OAuth2 provider instances and retrieve user information, and supports multiple providers, including Google, Facebook, and GitHub. It's extensively tested and can be used in projects of any size.

## Table of Contents
- [Installation](#installation)
- [Supported Providers](#supported-providers)
- [Components](#components)
- [Example](#example)
- [License](#license)

## Installation

To install the `go-oauth2-providers` package, you can use the `go get` command:

```shell
go get github.com/Zhima-Mochi/go-oauth2-providers
```

## Supported Providers

The `go-oauth2-providers` package currently supports the following OAuth2 providers:

- Facebook
- GitHub
- Google
- Line

Each provider implementation includes files such as `provider_xxx.go`, `config_provider.go`, `options_provider.go`, `authCodeOption.go`, `token.go`, and `user.go`, along with other utility files. These files provide the necessary logic and structures for interacting with the respective OAuth2 provider.

## Components

The `go-oauth2-providers` package includes the following components:

- `Auth` - The `Auth` component is the main component of the `go-oauth2-providers` package. It contains all the necessary information for interacting with an OAuth2 provider, including retrieving the authorization URL, exchanging the authorization code for an access token, and retrieving the user information.

- `Provider` - The `Provider` component is the factory for creating `Auth` instances. It contains the necessary information for creating an `Auth` instance, including the provider's name, configuration, and options.

## Example

See [Example](./example/main.go) for a simple example of how to use the `go-oauth2-providers` package.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
