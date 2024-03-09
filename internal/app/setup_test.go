package app

import (
	"fmt"
	"main/testsutils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Tests setup")
	testsutils.SetupDbForTests()
	code := m.Run()
	os.Exit(code)
}
