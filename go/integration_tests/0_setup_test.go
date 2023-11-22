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
	server := server.SetupServer()
	testHTTP = httptest.NewServer(server)

	code := m.Run()

	testHTTP.Close()
	os.Exit(code)
}
