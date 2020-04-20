package logsymbols

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/gookit/color.v1"
)

func ExampleForceNoColors() {
	ForceNoColors()
	fmt.Println("\nno color")
	fmt.Printf("%s info\n", Info)
	fmt.Printf("%s success\n", Success)
	fmt.Printf("%s warning\n", Warn)
	fmt.Printf("%s error\n", Error)
	// Output: no color
	// ℹ info
	// ✔ success
	// ⚠ warning
	// ✖ error
}

func ExampleFallbackSymbols() {
	f := FallbackSymbols()

	fmt.Println("\nfallback")
	fmt.Printf("%s info\n", f.Info)
	fmt.Printf("%s success\n", f.Success)
	fmt.Printf("%s warning\n", f.Warning)
	fmt.Printf("%s error\n", f.Error)
	// Output: fallback
	// i info
	// √ success
	// ‼ warning
	// × error
}

func TestNormalSymbols(t *testing.T) {
	sets := map[string]Symbols{
		"normal":   NormalSymbols(),
		"fallback": FallbackSymbols(),
		"current":  CurrentSymbols(),
	}

	for n, set := range sets {
		s := set
		name := n

		t.Run("no color "+name, func(t *testing.T) {
			ForceNoColors()

			var expected Symbols
			switch name {
			case "normal":
				expected = NormalSymbols()
			case "current":
				expected = CurrentSymbols()
			case "fallback":
				expected = FallbackSymbols()
			}

			assert.Equal(t, s.Info, expected.Info)
			assert.Equal(t, s.Success, expected.Success)
			assert.Equal(t, s.Ok, expected.Ok)
			assert.Equal(t, s.Warning, expected.Warning)
			assert.Equal(t, s.Error, expected.Error)
		})

		t.Run("color "+name, func(t *testing.T) {
			ForceColors()

			var expected Symbols
			switch name {
			case "normal":
				expected = NormalSymbols()
			case "current":
				expected = CurrentSymbols()
			case "fallback":
				expected = FallbackSymbols()
			}

			assert.Contains(t, expected.Info, s.Info)
			assert.Contains(t, expected.Success, s.Success)
			assert.Contains(t, expected.Ok, s.Ok)
			assert.Contains(t, expected.Warning, s.Warning)
			assert.Contains(t, expected.Error, s.Error)
		})
	}
}

func TestFallbackSymbols(t *testing.T) {
	t.Run("no color", func(t *testing.T) {
		ForceNoColors()
		s := FallbackSymbols()

		assert.Equal(t, s.Info, fallback.Info)
		assert.Equal(t, s.Success, fallback.Success)
		assert.Equal(t, s.Ok, fallback.Ok)
		assert.Equal(t, s.Warning, fallback.Warning)
		assert.Equal(t, s.Error, fallback.Error)
	})

	t.Run("color", func(t *testing.T) {
		ForceColors()
		s := FallbackSymbols()

		assert.Equal(t, s.Info, Symbol(color.Notice.Render(fallback.Info)))
		assert.Equal(t, s.Success, Symbol(color.Success.Render(fallback.Success)))
		assert.Equal(t, s.Ok, Symbol(color.Success.Render(fallback.Ok)))
		assert.Equal(t, s.Warning, Symbol(color.Warn.Render(fallback.Warning)))
		assert.Equal(t, s.Error, Symbol(color.Danger.Render(fallback.Error)))
	})
}

func TestAutodetectTTY(t *testing.T) {
	t.Run("autodetect", func(t *testing.T) {
		AutodetectTTY(os.Stdout)

		// During tests Stdout is not a TTY
		assert.False(t, colorOn)
	})

	t.Run("try tty1", func(t *testing.T) {
		f, err := os.OpenFile("/dev/pts/1", os.O_RDONLY, 0444)
		if err != nil {
			t.SkipNow() // this environment don't have TTY or PTS
		}

		AutodetectTTY(f)
		assert.True(t, colorOn)
	})

	t.Run("nil", func(t *testing.T) {
		AutodetectTTY(nil)
		assert.False(t, colorOn)
	})
}

func TestSymbols_String(t *testing.T) {
	s := osBaseSymbols.String()
	assert.Equal(t, "Info: ℹ Success: ✔ Warning: ⚠ Error: ✖", s)
}
