package cache_elimination

import (
	"testing"
	"time"
)

func TestLruCache(t *testing.T) {
	lc := NewLruCache(2)
	lc.Set("1", "1", 60*time.Second)
	lc.Set("2", "2", 60*time.Second)
	exist, value := lc.Get("1")
	if !exist {
		t.Errorf("get 1 failed")
	}
	t.Log(value)
	lc.Set("3", "3", 60*time.Second)
	exist, value = lc.Get("2")
	if !exist {
		t.Errorf("get 2 failed")
	}
	lc.Set("4", "4", 60*time.Second)
	exist, value = lc.Get("1")
	if !exist {
		t.Errorf("get 1 failed")
	}
	exist, value = lc.Get("3")
	if !exist {
		t.Errorf("get 3 failed")
	}
	exist, value = lc.Get("4")
	if !exist {
		t.Errorf("get 1 failed")
	}
}
