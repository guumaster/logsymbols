package logsymbols

import (
	"os"

	"github.com/mattn/go-isatty"
	"gopkg.in/gookit/color.v1"
)

// AutodetectTTY detects if it's on a TTY output
func AutodetectTTY(f *os.File) {
	if f == nil {
		f = os.Stdout
	}

	if isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd()) {
		ForceColors()
		return
	}

	ForceNoColors()
}

// Colorize adds colors to all symbols
func Colorize(s Symbols) Symbols {
	return Symbols{
		Info:    Symbol(color.Notice.Render(s.Info)),
		Success: Symbol(color.Success.Render(s.Success)),
		Ok:      Symbol(color.Success.Render(s.Ok)),
		Warning: Symbol(color.Warn.Render(s.Warning)),
		Warn:    Symbol(color.Warn.Render(s.Warn)),
		Error:   Symbol(color.Danger.Render(s.Error)),
	}
}

// ForceColors force adding color to all symbols
func ForceColors() {
	colorOn = true

	setGlobals(Colorize(osBaseSymbols))
}

// ForceNoColors force removing color to all symbols
func ForceNoColors() {
	colorOn = false

	setGlobals(osBaseSymbols)
}

func setGlobals(s Symbols) {
	Success = s.Success
	Ok = s.Ok
	Info = s.Info
	Warning = s.Warning
	Warn = s.Warn
	Error = s.Error
}

// CurrentSymbols returns a set with the current OS symbols (colored if color enabled)
func CurrentSymbols() Symbols {
	if colorOn {
		return Colorize(osBaseSymbols)
	}

	return osBaseSymbols
}

// BaseSymbols returns a set with the current OS symbols
func BaseSymbols() Symbols {
	return osBaseSymbols
}

// NormalSymbols returns a set with the normal symbols
func NormalSymbols() Symbols {
	if colorOn {
		return Colorize(normal)
	}

	return normal
}

// FallbackSymbols returns a set with the fallback symbols
func FallbackSymbols() Symbols {
	if colorOn {
		return Colorize(fallback)
	}

	return fallback
}
