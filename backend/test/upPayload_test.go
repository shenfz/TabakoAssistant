package test

import (
	"github.com/shenfz/TabakoAssistant/backend/pkg/spark"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

/**
 * @Author shenfz
 * @Date 2024/3/4 16:06
 * @Email 1328919715@qq.com
 * @Description:
 **/

type Demo struct {
	locker *sync.RWMutex
	cache  map[string][]spark.Text
}

// Benchmark_001-16         7850859               162.1 ns/op
func Benchmark_001(b *testing.B) {
	de := Demo{
		locker: &sync.RWMutex{},
		cache:  make(map[string][]spark.Text),
	}

	for i := 0; i < b.N; i++ {
		de.Up01("key", append([]spark.Text{}, spark.Text{
			Role:    "some",
			Content: "some",
		}))
	}
}

// Benchmark_000-16         8587020               138.7 ns/op
func Benchmark_000(b *testing.B) {
	de := Demo{
		locker: &sync.RWMutex{},
		cache:  make(map[string][]spark.Text),
	}
	for i := 0; i < b.N; i++ {
		de.Up("key", append([]spark.Text{}, spark.Text{
			Role:    "some",
			Content: "some",
		}))
	}
}

func Test_UP(t *testing.T) {
	de := Demo{
		locker: &sync.RWMutex{},
		cache:  make(map[string][]spark.Text),
	}

	go func() {
		c := time.NewTicker(1 * time.Second)
		key := strconv.Itoa(int(time.Now().Unix()))
		defer c.Stop()
		for {
			select {
			case <-c.C:
				var upObj []spark.Text
				for i := 0; i < rand.Intn(10); i++ {
					upObj = append(upObj, spark.Text{
						Role:    "some001",
						Content: "some",
					})
				}
				dstObj := de.Up01(key, upObj)
				t.Logf("Some001: length = %d ,ranIn : %d", len(dstObj), len(upObj))
			default:
			}
		}
	}()
	time.Sleep(2 * time.Minute)
	t.Logf("Finished")
}

func (d *Demo) Up(key string, texts []spark.Text) []spark.Text {
	d.locker.Lock()
	defer d.locker.Unlock()
	dst := append(d.cache[key], texts...)
	d.cache[key] = dst
	return dst
}

func (d *Demo) Up01(key string, texts []spark.Text) []spark.Text {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.cache[key] = append(d.cache[key], texts...)
	return d.cache[key]
}
