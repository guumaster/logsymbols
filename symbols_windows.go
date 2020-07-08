// +build windows

package logsymbols

import (
	"os"
)

func init() {
	osBaseSymbols = fallback

	AutodetectTTY(os.Stdout)
}
