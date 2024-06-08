//go:build go1.20

package lockedmap

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

// LockedMap的value是指针，并发修改value的属性会有问题（因为Get()返回的是value的指针）。
var unsafeLockedMap = New[string, *TestStruct]()

func TestDataRaceError(t *testing.T) {
	unsafeLockedMap.Set("a", &TestStruct{Name: "Alice"})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		val, ok := unsafeLockedMap.Get("a")
		assert.True(t, ok)
		val.Name = "Alice" + time.Now().Format(time.RFC3339Nano) //write
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		val, ok := unsafeLockedMap.Get("a")
		assert.True(t, ok)
		_ = val.Name //read
	}()

	wg.Wait()
}
