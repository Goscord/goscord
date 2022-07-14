#  Goscord

Goscord is a [Go](https://golang.org/) package that provides high level 
bindings to the [Discord](https://discord.com/) API.

This project remains a hobby so it is not maintained continuously.

## Getting Started

### Installing

```sh
go get -u github.com/Goscord/goscord
```

### Usage

Import the package into your project.

```go
import "github.com/Goscord/goscord"
```

Construct a new Discord client which can be used to access the variety of 
Discord API functions and to set callback functions for Discord events.

```go
client := goscord.New(&gateway.Options{ Token: "token" })
```

See documentation for more detailed information.


## Documentation

[Click here](https://goscord.dev/documentation)

## Examples

**Coming soon**

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
