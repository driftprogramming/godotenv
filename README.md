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
see `godotenv_test.go` in current repo:
```go
package godotenv

import (
	"embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed envs/*
var envFS embed.FS

func Test_In_UnParallel(t *testing.T) {
	t.Parallel()

	err := testLoadShouldGetEmbeddedLocalEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "local.1.0", os.Getenv("version"))

	err = testLoadShouldGetEmbeddedLiveeEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "live.1.0", os.Getenv("version"))

	err = testLoadShouldGetEmbeddedCommonEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "abcd", os.Getenv("shared_key"))

	testLoadShouldGetEmbeddedCommonAndLiveEnvFileSuccessfully()
	assert.Equal(t, "abcd", os.Getenv("shared_key"))
	assert.Equal(t, "live.1.0", os.Getenv("version"))

	testLoadTestEnvWithoutOverwriteShouldNotBeOverwrite()
	assert.Equal(t, "live.1.0", os.Getenv("version"))
}

func testLoadShouldGetEmbeddedLocalEnvFileSuccessfully() error {
	return Load(envFS, "envs/.env.local")
}

func testLoadShouldGetEmbeddedLiveeEnvFileSuccessfully() error {
	return Load(envFS, "envs/.env.live")
}

func testLoadShouldGetEmbeddedCommonEnvFileSuccessfully() error {
	return Load(envFS, "envs/.env")
}

func testLoadShouldGetEmbeddedCommonAndLiveEnvFileSuccessfully() {
	_ = Load(envFS, "envs/.env")
	_ = Load(envFS, "envs/.env.live")
}

func testLoadTestEnvWithoutOverwriteShouldNotBeOverwrite() {
	_ = LoadWithoutOverwrite(envFS, "envs/.env.test")
}

```


