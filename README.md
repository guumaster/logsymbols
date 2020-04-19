# logSymbols

> Colored symbols for various log levels
Includes fallbacks for Windows CMD which only supports a [limited character set](https://en.wikipedia.org/wiki/Code_page_437).


### Linux/Mac
![logSymbols on Linux](docs/screenshot.png)

### Windows

![logSymbols on Windows](docs/screenshot_windows.png)


## Install

```
$ go get github.com/guumaster/logSymbols
```



## Usage

### Basic example
```go

package main

import (
  "fmt"
  "github.com/guumaster/logSymbols"
)

func main() {
  fmt.Println(logSymbols.Success, 'Finished successfully!')
  fmt.Println(logSymbols.Error, 'Something broke')

  // On good OSes:  ✔ Finished successfully!
  //                ✖ Something broke

  // On Windows:    √ Finished successfully!
  //                × Something broke
}

```

### Forcing colors

`logSymbols` will autodetect if its not in TTY mode and remove all colors. You can force color output with this example:

```go
  
  fmt.Println(logSymbols.Success, 'Finished successfully!')
  
  // Terminal Output:  ✔ Finished successfully!
  // Redirected Output:       ^[[1;32m✔^[[0m Finished successfully!
}

```



## Ported from npm version

- [sindresorhus/log-symbols](https://github.com/sindresorhus/log-symbols) - Colored symbols for various log levels


## Related

- [prabirshrestha/go-log-symbols](https://github.com/prabirshrestha/go-log-symbols)
