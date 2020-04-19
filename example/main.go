// + build example
package main

import (
	"fmt"

	"github.com/guumaster/logsymbols"
)

func main() {
	fmt.Println("\nAutodetect: ")
	fmt.Printf("%s info\n", logsymbols.Info)
	fmt.Printf("%s success\n", logsymbols.Ok)
	fmt.Printf("%s warning\n", logsymbols.Warning)
	fmt.Printf("%s error\n", logsymbols.Error)

	// Output:
	//
	// Autodetect:
	// ℹ info
	// ✔ success
	// ⚠ warning
	// ✖ error

	logsymbols.ForceNoColors()
	fmt.Println("\nwithout-color")
	fmt.Println("\tNormal ", logsymbols.NormalSymbols())
	fmt.Println("\tFallback", logsymbols.FallbackSymbols())

	// Output:
	//
	// without-color
	//        Normal  {ℹ ✔ ⚠ ✖}
	//        Fallback {i √ ‼ ×}

	logsymbols.ForceColors()
	fmt.Println("\nwith-color")
	fmt.Println("\tNormal ", logsymbols.NormalSymbols())
	fmt.Println("\tFallback ", logsymbols.FallbackSymbols())

	// Output:
	//
	// with-color
	//        Normal  {ℹ ✔ ✔ ⚠ ⚠ ✖}
	//        Fallback  {i √ √ ‼ ‼ ×}
}
