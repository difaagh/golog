Simple & lightweight logging file for golang


### Installation
```
    go get github.com/difaagh/golog
```

#### Usage

```go
package main

import (
  "github.com/difaagh/golog"
)

func main() {
    // create struct for golog settings
    settings := golog.Config {
        FileName: "mylog", // mylog.log
        FolderName: "logs"
    }
    golog.Setup(&settings)

    // golog.Info
    // golog.Warn
    // golog.Error
    // golog.Fatal
    // golog.Debug
}
```
