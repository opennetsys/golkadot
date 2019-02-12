package dirutil

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// UserHomeDir ...
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

// NormalizePath ...
func NormalizePath(path string) string {
	// expand tilde
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(UserHomeDir(), path[2:])
	}

	return path
}

// CreateDirIfNotExist ...
func CreateDirIfNotExist(path string) error {
	path = NormalizePath(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0757)
	}

	return nil
}
