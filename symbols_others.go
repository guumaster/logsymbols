// +build !windows

package logSymbols

import (
	"os"
)

func init() {
	osBaseSymbols = normal
	AutodetectTTY(os.Stdout)
}
