package godotenv

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed envs/*
var envFS embed.FS

func Test_In_UnParallel(t *testing.T) {
	t.Parallel()

	err := testLoadShouldGetEmbeddedLocalEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "local.1.0", Get("version"))

	err = testLoadShouldGetEmbeddedLiveeEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "live.1.0", Get("version"))

	err = testLoadShouldGetEmbeddedCommonEnvFileSuccessfully()
	assert.Nil(t, err)
	assert.Equal(t, "abcd", Get("shared_key"))

	testLoadShouldGetEmbeddedCommonAndLiveEnvFileSuccessfully()
	assert.Equal(t, "abcd", Get("shared_key"))
	assert.Equal(t, "live.1.0", Get("version"))

	testLoadTestEnvWithoutOverwriteShouldNotBeOverwrite()
	assert.Equal(t, "live.1.0", Get("version"))
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
