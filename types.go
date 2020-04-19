package logSymbols

var colorOn bool

type symbol string

var (
	Info    symbol
	Success symbol
	Ok      symbol
	Warning symbol
	Warn    symbol
	Error   symbol
)

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
