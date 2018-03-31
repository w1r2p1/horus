package config

import "strings"
import "os"

// TODO - it could be `pkg` instead of `src` so, add that check to `GetProjectPath`.
const projectPath = "/src/github.com/DanielOchoa/horus"

func GetProjectPath() string {
	return projectPath
}

// Get $GOPATH for use in opening files for example. Since we may be setting the path to use `:` for multiple
// lookups, we need to split that up. It returns the first path without the string 'bin' in it.
func GetGoPath() string {
	if strings.Contains(os.Getenv("GOPATH"), ":") {
		paths := strings.Split(os.Getenv("GOPATH"), ":")
		for _, path := range paths {
			if strings.Contains(path, "bin") {
				continue
			}
			return path
		}
		// all paths contain bin?!?! whatever.
		return paths[0]
	}
	return os.Getenv("GOPATH")
}
