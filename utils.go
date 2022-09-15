package golden_image

import (
	"os"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)

const goldenDir = "__baseimages__"

func findRoot() string {
	// casually stolen from https://github.com/gkampitakis/go-snaps/blob/5079c95df0a15a1f15b72fce0de0ca2a420dda1d/snaps/utils.go#L127
	var ok bool
	var pc uintptr
	var file, prevFile, funcName string

	for i := 0; ; i++ {
		prevFile = file
		pc, file, _, ok = runtime.Caller(i)
		if !ok {
			return ""
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}

		funcName = f.Name()
		if funcName == "testing.tRunner" {
			break
		}

		// special case handling test runners
		// tested with testify/suite, packagestest and testcase
		segments := strings.Split(funcName, ".")
		for _, segment := range segments {
			if !isTest(segment, "Test") {
				continue
			}

			// packagestest is same as tRunner where we step one caller further
			// so we need to return the prevFile in testcase and testify/suite we return the current file
			// e.g. funcName golang.org/x/tools/go/packages/packagestest.TestAll.func1
			if strings.Contains(funcName, "packagestest") {
				// return only the Function Name
				// e.g. "go-snaps-testing-suite/src/issues.(*ExampleTestSuite).TestExampleSnapshot"
				// will return TestExampleSnapshot
				return prevFile
			}

			return file
		}
	}

	return prevFile
}

// Stolen from the `go test` tool
//
// isTest tells whether name looks like a test
// It is a Test (say) if there is a character after Test that is not a lower-case letter
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	r, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(r)
}

func getEnvBool(variable string, fallback bool) bool {
	e, exists := os.LookupEnv(variable)
	if !exists {
		return fallback
	}

	return e == "true"
}
