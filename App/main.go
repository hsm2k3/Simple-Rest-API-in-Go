package main

// just a comment regarding imports in golang
// github.com/hsm2k3/Simple-Rest-API-in-Go functions as a virtual path rather than an absolute path or path relative to the current directory.
// In other words it is a path relative to the go mod init {path} that can be found in the go.mod file.
import (
	"github.com/hsm2k3/Simple-Rest-API-in-Go/Routes"
)

func main() {
	Routes.HandleRequests()
}
