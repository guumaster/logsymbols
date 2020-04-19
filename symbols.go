package logsymbols

import (
	"os"

	"github.com/mattn/go-isatty"
	"gopkg.in/gookit/color.v1"
)

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

func Colorize(s Symbols) Symbols {
	return Symbols{
		Info:    symbol(color.Notice.Render(s.Info)),
		Success: symbol(color.Success.Render(s.Success)),
		Ok:      symbol(color.Success.Render(s.Ok)),
		Warning: symbol(color.Warn.Render(s.Warning)),
		Warn:    symbol(color.Warn.Render(s.Warn)),
		Error:   symbol(color.Danger.Render(s.Error)),
	}
}

func ForceColors() {
	colorOn = true

	setGlobals(Colorize(osBaseSymbols))
}

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

func NormalSymbols() Symbols {
	if colorOn {
		return Colorize(normal)
	}

	return normal
}

func FallbackSymbols() Symbols {
	if colorOn {
		return Colorize(fallback)
	}

	return fallback
}
