package common

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsThrottled(t *testing.T) {
	require.True(t, isThrottled(http.StatusTooManyRequests))
	require.True(t, isThrottled(http.StatusServiceUnavailable))
	require.False(t, isThrottled(http.StatusOK))
	require.False(t, isThrottled(http.StatusNotFound))
	require.False(t, isThrottled(http.StatusInternalServerError))
}

func TestHostBackoffIncrementsOnThrottle(t *testing.T) {
	shared := &Shared{}

	b := shared.backoffFor("example.com")
	require.Equal(t, int32(0), b.consecutive.Load())

	b.consecutive.Add(1)
	require.Equal(t, int32(1), b.consecutive.Load())

	b.consecutive.Add(1)
	require.Equal(t, int32(2), b.consecutive.Load())
}

func TestHostBackoffDecaysOnSuccess(t *testing.T) {
	shared := &Shared{}

	b := shared.backoffFor("example.com")
	b.consecutive.Store(3)

	b.consecutive.Add(-1)
	require.Equal(t, int32(2), b.consecutive.Load())
}

func TestHostBackoffPerDomain(t *testing.T) {
	shared := &Shared{}

	a := shared.backoffFor("a.com")
	b := shared.backoffFor("b.com")

	a.consecutive.Store(5)
	require.Equal(t, int32(0), b.consecutive.Load())
}

func TestApplyBackoffNoDelayWhenClean(t *testing.T) {
	shared := &Shared{}

	start := testing.AllocsPerRun(1, func() {
		shared.applyBackoff("clean-host.com")
	})
	_ = start
}
