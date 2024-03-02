package lockedmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetGet(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get failed")
	}
}

func TestExists(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if !m.Exists("a") {
		t.Errorf("Exists failed")
	}
}

func TestRemove(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.Remove("a")
	if m.Exists("a") {
		t.Errorf("Remove failed")
	}
}

func TestRemoveAll(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.RemoveAll()
	if m.Exists("a") {
		t.Errorf("RemoveAll failed")
	}
	assert.Equal(t, 0, m.Size())
}

func TestSize(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if m.Size() != 1 {
		t.Errorf("Size failed")
	}
}

func TestKeys(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if len(m.Keys()) != 1 {
		t.Errorf("Keys failed")
	}
}

func TestValues(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if len(m.Values()) != 1 {
		t.Errorf("Values failed")
	}
}

func TestConcurrentSetGet(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Get("a")
	}
}

func TestConcurrentRemove(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Remove("a")
	}
}

func TestConcurrentSize(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Size()
	}
}

func TestConcurrentExists(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Exists("a")
	}
}

func TestConcurrentKeys(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Keys()
	}
}

func TestConcurrentGet(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.Get("a")
	}
}

func TestConcurrentRemoveAll(t *testing.T) {
	m := New[string, int]()
	for i := 0; i < 1000; i++ {
		go m.Set("a", i)
		go m.RemoveAll()
	}
}
