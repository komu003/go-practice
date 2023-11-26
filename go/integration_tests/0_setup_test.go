package integration_tests

import (
	"app/pkg/server"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testHTTP *httptest.Server
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	config.InitConfig()
	server := server.SetupServer()
	testHTTP = httptest.NewServer(server)

	code := m.Run()

	testHTTP.Close()
	os.Exit(code)
}
