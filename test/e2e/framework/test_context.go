package framework

import (
	"flag"
	"fmt"
	"os"
)

type TestContextType struct {
	FRPClientPath string
	FRPServerPath string
	LogLevel      string
	Debug         bool
}

var TestContext TestContextType

// RegisterCommonFlags registers flags common to all e2e test suites.
// The flag set can be flag.CommandLine (if desired) or a custom
// flag set that then gets passed to viperconfig.ViperizeFlags.
//
// The other Register*Flags methods below can be used to add more
// test-specific flags. However, those settings then get added
// regardless whether the test is actually in the test suite.
func RegisterCommonFlags(flags *flag.FlagSet) {
	flags.StringVar(&TestContext.FRPClientPath, "cccc-path", "../../bin/cccc", "The cccc client binary to use.")
	flags.StringVar(&TestContext.FRPServerPath, "cccs-path", "../../bin/cccs", "The cccs server binary to use.")
	flags.StringVar(&TestContext.LogLevel, "log-level", "debug", "Log level.")
	flags.BoolVar(&TestContext.Debug, "debug", false, "Enable debug mode to print detail info.")
}

func ValidateTestContext(t *TestContextType) error {
	if t.FRPClientPath == "" || t.FRPServerPath == "" {
		return fmt.Errorf("cccc and cccs binary path can't be empty")
	}
	if _, err := os.Stat(t.FRPClientPath); err != nil {
		return fmt.Errorf("load cccc-path error: %v", err)
	}
	if _, err := os.Stat(t.FRPServerPath); err != nil {
		return fmt.Errorf("load cccs-path error: %v", err)
	}
	return nil
}
