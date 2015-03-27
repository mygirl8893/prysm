package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Memory struct {
	store []byte
}

func NewMemory() *Memory {
	return &Memory{nil}
}

func (m *Memory) Set(offset, size uint64, value []byte) {
	// If the length of the store is 0 this is a complete failure
	// memory size is set prior to calling this method so enough size
	// should always be available.
	if len(m.store) == 0 {
		panic("INVALID memory: store empty")
	}

	value = common.RightPadBytes(value, int(size))

	totSize := offset + size
	lenSize := int64(len(m.store) - 1)
	if totSize > lenSize {
		// Calculate the diff between the sizes
		diff := totSize - lenSize
		if diff > 0 {
			// Create a new empty slice and append it
			newSlice := make([]byte, diff-1)
			// Resize slice
			m.store = append(m.store, newSlice...)
		}
	}
	copy(m.store[offset:offset+size], value)
}

func (m *Memory) Resize(size uint64) {
	if uint64(m.Len()) < size {
		m.store = append(m.store, make([]byte, size-uint64(m.Len()))...)
	}
}

func (self *Memory) Get(offset, size int64) (cpy []byte) {
	if size == 0 {
		return nil
	}

	if len(self.store) > int(offset) {
		cpy = make([]byte, size)
		copy(cpy, self.store[offset:offset+size])

		return
	}

	return
}

func (m *Memory) Len() int {
	return len(m.store)
}

func (m *Memory) Data() []byte {
	return m.store
}

func (m *Memory) Print() {
	fmt.Printf("### mem %d bytes ###\n", len(m.store))
	if len(m.store) > 0 {
		addr := 0
		for i := 0; i+32 <= len(m.store); i += 32 {
			fmt.Printf("%03d: % x\n", addr, m.store[i:i+32])
			addr++
		}
	} else {
		fmt.Println("-- empty --")
	}
	fmt.Println("####################")
}