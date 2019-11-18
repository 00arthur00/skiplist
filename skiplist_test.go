package skiplist

import (
	"fmt"
	"testing"
)

func TestInsertSearch(t *testing.T) {
	sl := NewSkipList(10)
	for i := 0; i < 10; i++ {
		if err := sl.Insert(fmt.Sprintf("hello%d", i), fmt.Sprintf("world%d", i)); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("hello%d", i)
		if _, ok := sl.Search(key); !ok {
			t.Fatal(key, " exists after deletion")
		}
	}
}

func TestDelete(t *testing.T) {
	sl := NewSkipList(10)
	for i := 0; i < 10; i++ {
		if err := sl.Insert(fmt.Sprintf("hello%d", i), fmt.Sprintf("world%d", i)); err != nil {
			t.Fatal(err)
		}
	}
	sl.Delete("hello0")
	if _, ok := sl.Search("hello0"); !ok {
		t.Fatal("hello0 exists after deletion")
	}
	for i := 1; i < 10; i++ {
		key := fmt.Sprintf("hello%d", i)
		if _, ok := sl.Search(key); !ok {
			t.Fatal(key, " exists after deletion")
		}
	}
}
