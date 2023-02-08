package backendapi

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	HelperSetup(m)

	os.Exit(m.Run())
}
