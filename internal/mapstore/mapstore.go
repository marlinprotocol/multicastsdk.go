package mapstore

import (
	"sync"
	"unsafe"
)

var (
	mutex sync.Mutex
	store = map[unsafe.Pointer]interface{}{}
)

func Save(key unsafe.Pointer, v interface{}) bool {
	if key == nil || v == nil {
		return false
	}

	mutex.Lock()
	store[key] = v
	mutex.Unlock()

	return true
}

func Restore(key unsafe.Pointer) (v interface{}) {
	if key == nil {
		return nil
	}

	mutex.Lock()
	v = store[key]
	mutex.Unlock()
	return
}

func Unref(key unsafe.Pointer) bool {
	if key == nil {
		return false
	}

	mutex.Lock()
	delete(store, key)
	mutex.Unlock()

	return true
}
