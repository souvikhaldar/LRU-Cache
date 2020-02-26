package goLRU

import (
	"testing"
)

func TestSet(t *testing.T) {
	c := NewLRUCache(2)
	if err := c.Set("a", 1); err != nil {
		t.Fatal("Can't set to cache: %s", err)
	}
}
