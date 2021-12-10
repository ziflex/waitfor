# waitfor
> Test and wait on the availability of a remote resource.

# DEPRECATED
The project moved to [go-waitfor](https://github.com/go-waitfor) organization.
- [CLI](https://github.com/go-waitfor/cli)
- [lib](https://github.com/go-waitfor/waitfor)

## Features
- Parallel availability tests
- Exponential backoff
- Different types of remote resource (http(s), proc, postgres, mysql)

## Resources
- File (``file://``)
- OS Process (``proc://``)
- HTTP(S) Endpoint (``http://`` & ``https://``)
- MongoDB (``mongodb://``)
- Postgres (``postgres://``)
- MySQL/MariaDB (``mysql://`` & ``mariadb://``)

## Resource URLs
All resource locations start with url schema type e.g. ``file://./myfile`` or ``postgres://locahost:5432/mydb?user=user&password=test``

## Library

### Test resource availability
```go
package main 

import (
    "context"
    "fmt"
    "os"
	"github.com/ziflex/waitfor/pkg/runner"
)

func main() {
    err := runner.Test(
        context.Background(),
        []string{"postgres://locahost:5432/mydb?user=user&password=test"},
        runner.WithAttempts(5),
    )
    
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```


### Test resource availability and run a program
```go
package main 

import (
    "context"
    "fmt"
    "os"
	"github.com/ziflex/waitfor/pkg/runner"
)

func main() {
    program := runner.Program{
        Executable: "myapp",
        Args:       []string{"--database", "postgres://locahost:5432/mydb?user=user&password=test"},
        Resources:  []string{"postgres://locahost:5432/mydb?user=user&password=test"},
    }
    
    out, err := runner.Run(
        context.Background(),
        program,
        runner.WithAttempts(5),
    )
    
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    fmt.Println(string(out))
}
```

## CLI
CLI is a simple wrapper around this library.

### Basic usage
```bash
    waitfor -r postgres://locahost:5432/mydb?user=user&password=test -r http://myservice:8080 npm start
```

### Options
```bash
NAME:
   waitfor - Tests and waits on the availability of a remote resource

USAGE:
   waitfor [global options] command [command options] [arguments...]

DESCRIPTION:
   Tests and waits on the availability of a remote resource before executing a command with exponential backoff

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --resource value, -r value  -r http://localhost:8080 [$WAITFOR_RESOURCE]
   --attempts value, -a value  amount of attempts (default: 5) [$WAITFOR_ATTEMPTS]
   --interval value            interval between attempts (sec) (default: 5) [$WAITFOR_INTERVAL]
   --max-interval value        maximum interval between attempts (sec) (default: 60) [$WAITFOR_MAX_INTERVAL]
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)

```
