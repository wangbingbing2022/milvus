package typeutil

import (
	"log"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestUint64(t *testing.T) {
	var i int64 = -1
	var u uint64 = uint64(i)
	t.Log(i)
	t.Log(u)
}

func TestHash32_Uint64(t *testing.T) {
	var u uint64 = 0x12
	h, err := Hash32Uint64(u)
	assert.Nil(t, err)

	t.Log(h)

	b := make([]byte, unsafe.Sizeof(u))
	b[0] = 0x12
	h2, err := Hash32Bytes(b)
	assert.Nil(t, err)

	t.Log(h2)
	assert.Equal(t, h, h2)
}

func TestHash32_String(t *testing.T) {
	var u string = "ok"
	h, err := Hash32String(u)
	assert.Nil(t, err)

	t.Log(h)
	log.Println(h)

	b := []byte(u)
	h2, err := Hash32Bytes(b)
	assert.Nil(t, err)
	log.Println(h2)

	assert.Equal(t, uint32(h), h2)

}