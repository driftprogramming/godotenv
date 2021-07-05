# godotenv

Go dot env lib. Some of other go dot env library doesn't support [go embed feature](https://github.com/golang/go/blob/37588ffcb221c12c12882b591a16243ae2799fd1/src/embed/internal/embedtest/embed_test.go), so you can NOT embed the static file
into go build executable binaries. This lib is a standard repo which support this feature, once you integrate with this
repo, you can just simply run go build as usually we do, you don't need to care that you need to copy the static files
to somewhere manually.

# Usage

```commandline
go get github.com/driftprogramming/godotenv@v1.0.0
```

#### Example
see `godotenv_test.go` in current repo Or see this:
```go
package root

import (
	"embed"

	"github.com/driftprogramming/godotenv"
)

//go:embed envs/*
var envs embed.FS

func SetupExample(env string) {
	_ = godotenv.Load(envs, "envs/.env.live")
	_ = godotenv.LoadWithoutOverwrite(envs, "envs/.env.local")
}

```


