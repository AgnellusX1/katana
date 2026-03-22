package browser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLauncherDefaults(t *testing.T) {
	t.Run("empty strategy defaults to heuristic", func(t *testing.T) {
		l, err := NewLauncher(LauncherOptions{
			MaxBrowsers: 1,
		})
		require.NoError(t, err)
		require.Equal(t, "heuristic", l.opts.PageLoadStrategy)
	})

	t.Run("explicit strategy is preserved", func(t *testing.T) {
		for _, strategy := range []string{"none", "load", "domcontentloaded", "networkidle", "heuristic"} {
			l, err := NewLauncher(LauncherOptions{
				MaxBrowsers:      1,
				PageLoadStrategy: strategy,
			})
			require.NoError(t, err)
			require.Equal(t, strategy, l.opts.PageLoadStrategy)
		}
	})

	t.Run("zero DOMWaitTime defaults to 5", func(t *testing.T) {
		l, err := NewLauncher(LauncherOptions{
			MaxBrowsers: 1,
		})
		require.NoError(t, err)
		require.Equal(t, 5, l.opts.DOMWaitTime)
	})

	t.Run("negative DOMWaitTime defaults to 5", func(t *testing.T) {
		l, err := NewLauncher(LauncherOptions{
			MaxBrowsers: 1,
			DOMWaitTime: -1,
		})
		require.NoError(t, err)
		require.Equal(t, 5, l.opts.DOMWaitTime)
	})

	t.Run("positive DOMWaitTime is preserved", func(t *testing.T) {
		l, err := NewLauncher(LauncherOptions{
			MaxBrowsers: 1,
			DOMWaitTime: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, l.opts.DOMWaitTime)
	})

	t.Run("ChromeWSUrl is passed through", func(t *testing.T) {
		l, err := NewLauncher(LauncherOptions{
			MaxBrowsers: 1,
			ChromeWSUrl: "ws://localhost:9222/devtools/browser/abc",
		})
		require.NoError(t, err)
		require.Equal(t, "ws://localhost:9222/devtools/browser/abc", l.opts.ChromeWSUrl)
	})
}
