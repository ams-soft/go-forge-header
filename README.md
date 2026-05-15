# go-forge-header

A Go library for printing a standardized startup header banner for services and CLI applications.

## Installation

```bash
go get github.com/ams-soft/go-forge-header
```

## Usage

```go
package main

import (
    header "github.com/ams-soft/go-forge-header"
)

func main() {
    header.Print(nil, "my service", "1.0.0", "production", false)
}
```

### Output

```
--------------------------------------------------------------------------------
⫽⫽⫽⫽⫽AMS SOFT - MY SERVICE
--------------------------------------------------------------------------------
                                                                                
--------------------------------------------------------------------------------
SYSTEM INFORMATION
--------------------------------------------------------------------------------
VERSION:           1.0.0
MODE:              PRODUCTION
DATE:              2026-05-15T12:00:00-03:00
                                                                                
```

## API

### `Print(vendor *string, name, version, env string, utc bool)`

Prints the full startup header to stdout.

| Parameter | Type      | Description                                                                 |
|-----------|-----------|-----------------------------------------------------------------------------|
| `vendor`  | `*string` | Custom vendor name. Pass `nil` to use the default AMS Soft branded logo.    |
| `name`    | `string`  | Service or application name. Displayed in uppercase.                        |
| `version` | `string`  | Application version string. Displayed in uppercase.                         |
| `env`     | `string`  | Runtime environment (e.g. `"prod"`, `"staging"`). Displayed in uppercase.   |
| `utc`     | `bool`    | If `true`, the timestamp is printed in UTC; otherwise uses local time.      |

**Custom vendor example:**

```go
vendor := "MY COMPANY"
header.Print(&vendor, "my service", "2.1.0", "staging", true)
```

### `AmsLogo() string`

Returns the AMS Soft logo string (colored or plain depending on the environment). Useful when you need the logo string without printing the full header.

## Environment Variables

| Variable                    | Default | Description                                                         |
|-----------------------------|---------|---------------------------------------------------------------------|
| `GO_FORGE_HEADER_NO_COLOR`  | `false` | Set to `true` to disable colored output and use plain text instead. |

```bash
GO_FORGE_HEADER_NO_COLOR=true ./my-service
```

## Requirements

- Go 1.22+

## License

MIT — see [LICENSE](LICENSE).  
Copyright (c) 2026 AMS TECNOLOGIA E SERVICOS FINANCEIROS LTDA
