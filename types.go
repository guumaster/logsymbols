package logsymbols

import (
	"fmt"
)

var colorOn bool

type Symbol string

var (
	// Info represents the information symbol
	Info Symbol
	// Success represents the success symbol
	Success Symbol
	// Ok alias of Success
	Ok Symbol
	// Warning represents the warning symbol
	Warning Symbol
	// Warn alias of Warning
	Warn Symbol
	// Error represents the error symbol
	Error Symbol
)

// Symbols struct contains all symbols
type Symbols struct {
	Info    Symbol
	Success Symbol
	Ok      Symbol
	Warning Symbol
	Warn    Symbol
	Error   Symbol
}

var osBaseSymbols Symbols

var normal = Symbols{
	Info:    Symbol("ℹ"),
	Success: Symbol("✔"),
	Ok:      Symbol("✔"),
	Warning: Symbol("⚠"),
	Warn:    Symbol("⚠"),
	Error:   Symbol("✖"),
}

var fallback = Symbols{
	Info:    Symbol("i"),
	Success: Symbol("√"),
	Ok:      Symbol("√"),
	Warning: Symbol("‼"),
	Warn:    Symbol("‼"),
	Error:   Symbol("×"),
}

// String returns a printable representation of Symbols struct
func (s Symbols) String() string {
	return fmt.Sprintf("Info: %s Success: %s Warning: %s Error: %s", s.Info, s.Success, s.Warning, s.Error)
}
