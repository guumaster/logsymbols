package logsymbols

import (
	"fmt"
)

var colorOn bool

type symbol string

var (
	// Info represents the information symbol
	Info symbol
	// Success represents the success symbol
	Success symbol
	// Ok alias of Success
	Ok symbol
	// Warning represents the warning symbol
	Warning symbol
	// Warn alias of Warning
	Warn symbol
	// Error represents the error symbol
	Error symbol
)

// Symbols struct contains all symbols
type Symbols struct {
	Info    symbol
	Success symbol
	Ok      symbol
	Warning symbol
	Warn    symbol
	Error   symbol
}

var osBaseSymbols Symbols

var normal = Symbols{
	Info:    symbol("ℹ"),
	Success: symbol("✔"),
	Ok:      symbol("✔"),
	Warning: symbol("⚠"),
	Warn:    symbol("⚠"),
	Error:   symbol("✖"),
}

var fallback = Symbols{
	Info:    symbol("i"),
	Success: symbol("√"),
	Ok:      symbol("√"),
	Warning: symbol("‼"),
	Warn:    symbol("‼"),
	Error:   symbol("×"),
}

// String returns a printable representation of Symbols struct
func (s Symbols) String() string {
	return fmt.Sprintf("Info: %s Success: %s Warning: %s Error: %s", s.Info, s.Success, s.Warning, s.Error)
}
