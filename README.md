# godotenv

Go dot env lib. Some of other go dot env library doesn't
support [go embed feature](https://github.com/golang/go/blob/37588ffcb221c12c12882b591a16243ae2799fd1/src/embed/internal/embedtest/embed_test.go)
, so you can NOT embed the static file into go build executable binaries. This lib is a standard repo which support this
feature, once you integrate with this repo, you can just simply run go build as usually we do, you don't need to care
that you need to copy the static files to somewhere manually.

# Usage

```commandline
go get github.com/driftprogramming/godotenv@v1.0.0
```

#### Example

see `godotenv_test.go` in current repo Or see this:

```go
.

├── envs.go // example here                    
├── envs               # local, dev, test, live env files.
│   ├──.env           # common env variables, which will be loaded in all other envs
│   ├──.env.dev       # dev env variables
│   ├──.env.live      # live env variables
│   ├──.env.local     # local env variables
│   └──.env.test      # test env variables
└── ...
```

```go
// envs/.env.local
version=local.1.0
otherkey=abcd
```

```go
// envs/.env.live
version=live.1.0
```

```go
// envs.go
package root

import (
	"embed"

	"github.com/driftprogramming/godotenv"
)

//go:embed envs/*
var envs embed.FS

func SetupExample(env string) {
	_ = godotenv.Load(envs, "envs/.env.live")
	version1 := godotenv.Get("version") // version1 will be live.1.0

	// the below method will not overwrite if env key already existed.
	_ = godotenv.LoadWithoutOverwrite(envs, "envs/.env.local")
	version2 := godotenv.Get("version") // version2 will still be live.1.0

	_ = godotenv.Load(envs, "envs/.env.local")
	version3 := godotenv.Get("version") // version3 will be local.1.0
}

```


