package utils

import (
	"path/filepath"
	"runtime"
)

// Root returns the absolute path to the project root directory
// if it cannot determine the filepath of the caller it panics
func Root() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Couldn't find file!")
	}
	return filepath.Join(file, "../..")
}
