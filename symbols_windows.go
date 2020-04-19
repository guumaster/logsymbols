// +build windows

package logSymbols

func init() {
	osBaseSymbols = fallback
	AutodetectTTY(os.Stdout)
}
