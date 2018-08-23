package linebeacon

var (
	ADTYPE_FLAGS                        byte = 0x01
	ADTYPE_SERVICE_DATA                 byte = 0x16
	ADTYPE_COMPLETE_16_BIT_SERVICE_UUID byte = 0x03

	UUID16LE_FOR_LINECORP []byte = []byte{0x6f, 0xfe}
)

var (
	frameTypeData byte = 0x02
	// hwidData            []byte
	measuredTxPowerData byte = 0x7F
	// deviceMessageData   []byte
)

func getBeaconFrame(hwid, msgData []byte) []byte {
	var ret []byte
	ret = append(ret, frameTypeData)
	ret = append(ret, hwid...)
	ret = append(ret, measuredTxPowerData)
	ret = append(ret, msgData...)
	return ret
}

type LineBeacon struct {
	raw []byte //raw data
}

func (l *LineBeacon) AddField(header byte, data []byte) *LineBeacon {
	l.raw = append(l.raw, byte(len(data)+1))
	l.raw = append(l.raw, header)
	l.raw = append(l.raw, data...)
	return l
}

func CreateLineSimpleBeaconAdvertisingPDU(hwid, msgData []byte) []byte {
	l := &LineBeacon{}
	l.AddField(ADTYPE_FLAGS, []byte{0x06})
	l.AddField(ADTYPE_COMPLETE_16_BIT_SERVICE_UUID, UUID16LE_FOR_LINECORP)
	l.AddField(ADTYPE_SERVICE_DATA, append(UUID16LE_FOR_LINECORP, getBeaconFrame(hwid, msgData)...))
	return l.raw
}
