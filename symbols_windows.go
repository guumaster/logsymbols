// +build windows

package logsymbols

func init() {
	osBaseSymbols = fallback

	AutodetectTTY(os.Stdout)
}
