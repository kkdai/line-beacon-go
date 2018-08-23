package linebeacon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrameData(t *testing.T) {
	// new Buffer('02 01 06 03 03 6f fe 0b 16 6f fe 02 fd 5e a0 ad 1e 7f 00'.replace(/ /g, ''), 'hex'),
	// simplebacon.createLineSimpleBeaconAdvertisingPDU('fd5ea0ad1e', '00'));

	expect := []byte{0x02, 0x01, 0x06, 0x03, 0x03, 0x6f, 0xfe, 0x0b, 0x16, 0x6f, 0xfe, 0x02, 0xfd, 0x5e, 0xa0, 0xad, 0x1e, 0x7f, 0x00}
	real := CreateLineSimpleBeaconAdvertisingPDU([]byte{0xfd, 0x5e, 0xa0, 0xad, 0x1e}, []byte{0x00})
	t.Log(expect)
	t.Log(real)
	assert.Equal(t, expect, real)
}
