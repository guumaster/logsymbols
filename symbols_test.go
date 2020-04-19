package logSymbols

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
	t.Run("no color", func(t *testing.T) {
		ForceNoColors()
		s := NormalSymbols()

		assert.Equal(t, s.Info, symbol("ℹ"))
		assert.Equal(t, s.Success, symbol("✔"))
		assert.Equal(t, s.Ok, symbol("✔"))
		assert.Equal(t, s.Warning, symbol("⚠"))
		assert.Equal(t, s.Error, symbol("✖"))
	})

	t.Run("color", func(t *testing.T) {
		ForceColors()
		s := NormalSymbols()

		assert.Equal(t, s.Info, symbol(color.Notice.Render(normal.Info)))
		assert.Equal(t, s.Success, symbol(color.Success.Render(normal.Success)))
		assert.Equal(t, s.Ok, symbol(color.Success.Render(normal.Ok)))
		assert.Equal(t, s.Warning, symbol(color.Warn.Render(normal.Warning)))
		assert.Equal(t, s.Error, symbol(color.Danger.Render(normal.Error)))
	})
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

		assert.Equal(t, s.Info, symbol(color.Notice.Render(fallback.Info)))
		assert.Equal(t, s.Success, symbol(color.Success.Render(fallback.Success)))
		assert.Equal(t, s.Ok, symbol(color.Success.Render(fallback.Ok)))
		assert.Equal(t, s.Warning, symbol(color.Warn.Render(fallback.Warning)))
		assert.Equal(t, s.Error, symbol(color.Danger.Render(fallback.Error)))
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
