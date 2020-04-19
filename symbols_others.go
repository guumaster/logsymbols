// +build !windows

package logsymbols

import (
	"os"
)

func init() {
	osBaseSymbols = normal

	AutodetectTTY(os.Stdout)
}
