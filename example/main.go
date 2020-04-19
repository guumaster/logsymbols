// + build example
package main

import (
	"fmt"

	"github.com/guumaster/logSymbols"
)

func main() {
	fmt.Println("\nAutodetect: ")
	fmt.Printf("%s info\n", logSymbols.Info)
	fmt.Printf("%s success\n", logSymbols.Ok)
	fmt.Printf("%s warning\n", logSymbols.Warning)
	fmt.Printf("%s error\n", logSymbols.Error)

	// Output:
	//
	// Autodetect:
	// ℹ info
	// ✔ success
	// ⚠ warning
	// ✖ error

	logSymbols.ForceNoColors()
	fmt.Println("\nwithout-color")
	fmt.Println("\tNormal ", logSymbols.NormalSymbols())
	fmt.Println("\tFallback", logSymbols.FallbackSymbols())

	// Output:
	//
	// without-color
	//        Normal  {ℹ ✔ ⚠ ✖}
	//        Fallback {i √ ‼ ×}

	logSymbols.ForceColors()
	fmt.Println("\nwith-color")
	fmt.Println("\tNormal ", logSymbols.NormalSymbols())
	fmt.Println("\tFallback ", logSymbols.FallbackSymbols())

	// Output:
	//
	// with-color
	//        Normal  {ℹ ✔ ✔ ⚠ ⚠ ✖}
	//        Fallback  {i √ √ ‼ ‼ ×}
}
