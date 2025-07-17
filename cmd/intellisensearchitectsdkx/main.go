// cmd/intellisensearchitectsdkx/main.go
package main

import (
"flag"
"log"
"os"

"intellisensearchitectsdkx/internal/intellisensearchitectsdkx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisensearchitectsdkx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
