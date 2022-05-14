package limit

import (
	"github/leave8080/go_package/pkg/log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	c := &Config{
		Rate:       2,  // 每秒允许通过 2 个
		BucketSize: 50, // 桶容量 50 个
	}
	var (
		allowNum = int32(1)
		wg       sync.WaitGroup
	)
	rl := NewLimiter(c)

	f := func() {
		allow := rl.LimiterGroup.Get("test_1").Allow()
		if allow {
			atomic.AddInt32(&allowNum, 1)
		}
		wg.Done()
	}

	start := time.Now()
	// 总共请求个数：10000 / 10 = 1000
	// 最大允许通过量：size + 10 * Rate = 50 + 10 * 2 = 70
	end := start.Add(10 * time.Second)
	for time.Now().Before(end) {
		wg.Add(1)
		go f()
		time.Sleep(10 * time.Millisecond) // Rate = 1000 / 10 = 100
	}
	wg.Wait()

	elapsed := time.Since(start)
	ideal := int64(c.BucketSize) + (int64(c.Rate) * int64(elapsed) / int64(time.Second))

	want := int32(ideal + 1)
	log.Warning("allowNum:% ,want:%d", allowNum, want)
	if allowNum > want {
		t.Errorf("count = %d, want %d (ideal %d)", allowNum, want, ideal)
	}

	want2 := int32(ideal - 1)
	log.Warning("allowNum:% ,want2:%d", allowNum, want2)
	if allowNum < want2 {
		t.Errorf("count = %d, want2 %d (ideal %d)", allowNum, want2, ideal)
	}
}
