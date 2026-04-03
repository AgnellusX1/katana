package common

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/projectdiscovery/ratelimit"
)

func BenchmarkGlobalRateLimit(b *testing.B) {
	for _, hosts := range []int{1, 5, 20} {
		b.Run(fmt.Sprintf("hosts=%d", hosts), func(b *testing.B) {
			limiter := ratelimit.New(context.Background(), 150, time.Second)
			defer limiter.Stop()

			b.ResetTimer()
			var wg sync.WaitGroup
			for i := range b.N {
				wg.Add(1)
				go func() {
					defer wg.Done()
					_ = i // host selection irrelevant for global limiter
					limiter.Take()
				}()
			}
			wg.Wait()
		})
	}
}

func BenchmarkPerHostRateLimit(b *testing.B) {
	for _, hosts := range []int{1, 5, 20} {
		b.Run(fmt.Sprintf("hosts=%d", hosts), func(b *testing.B) {
			limiter := ratelimit.NewAutoLimiter(
				context.Background(),
				ratelimit.WithMaxCount(150),
				ratelimit.WithDuration(time.Second),
			)
			defer limiter.Stop()

			hostnames := make([]string, hosts)
			for i := range hosts {
				hostnames[i] = fmt.Sprintf("host-%d.example.com", i)
			}

			b.ResetTimer()
			var wg sync.WaitGroup
			for i := range b.N {
				wg.Add(1)
				go func() {
					defer wg.Done()
					host := hostnames[i%hosts]
					_ = limiter.Take(host)
				}()
			}
			wg.Wait()
		})
	}
}
