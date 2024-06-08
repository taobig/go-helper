package lockedmap

import (
	"sync"
	"testing"
	"time"
)

// LockedMap的value是struct，并发修改value的属性不会有问题（因为Get()返回的是value的拷贝）。
var safeLockedMap = New[string, TestStruct]()

// 虽然LockedMap的value是指针，但是返回的结构体是只读的，所以不会存在修改value的属性的问题。
var safeLockedMap2 = New[string, *ReadonlyTestStruct]()

type TestStruct struct {
	Name string
	Age  int64
}

type ReadonlyTestStruct struct {
	name string
}

func (r ReadonlyTestStruct) Name() string {
	return r.name
}

func TestTestStructDataRaceOK(t *testing.T) {
	safeLockedMap.Set("a", TestStruct{Name: "Alice", Age: 20})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap.Set("a", TestStruct{Name: "Alice" + time.Now().Format(time.RFC3339Nano)})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_, _ = safeLockedMap.Get("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap.Remove("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap.RemoveAll()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap.Size()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap.Exists("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap.Keys()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap.Values()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap.Range(func(key string, value TestStruct) bool {
			return true
		})
	}()

	wg.Wait()
}

func TestReadonlyTestStructDataRaceOK(t *testing.T) {
	safeLockedMap2.Set("a", &ReadonlyTestStruct{name: "Alice"})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap2.Set("a", &ReadonlyTestStruct{name: "Alice" + time.Now().Format(time.RFC3339Nano)})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_, _ = safeLockedMap2.Get("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap2.Remove("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap2.RemoveAll()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap2.Size()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap2.Exists("a")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap2.Keys()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = safeLockedMap2.Values()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		safeLockedMap2.Range(func(key string, value *ReadonlyTestStruct) bool {
			return true
		})
	}()

	wg.Wait()
}
