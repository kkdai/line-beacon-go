package linebeacon

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

const (
	advTypeAllUUID16     = 0x03 // Complete List of 16-bit Service Class UUIDs
	advTypeServiceData16 = 0x16 // Service Data - 16-bit UUID
)

const (
	advFlagGeneralDiscoverable = 0x02
	advFlagLEOnly              = 0x04
)

type LineBeaconGATT struct {
	//Output frame data for line beacon
	OutFrame []byte

	//Tx power is the received power at 0 meters, in dBm, and the value ranges from -100 dBm to +20 dBm to a resolution of 1 dBm.
	//
	//Note to developers:
	// The best way to determine the precise value to put into this field is to measure the actual output of
	// your beacon from 1 meter away and then add 41dBm to that. 41dBm is the signal loss that occurs over 1 meter.
	PowerLevel int
}

func NewLineBeacon(hwid []byte, data []byte) *LineBeaconGATT {
	return &LineBeaconGATT{
		OutFrame: CreateLineSimpleBeaconAdvertisingPDU(hwid, data),
	}
}

// Start to Advertise your beacon, it is block API.
func (eb *LineBeaconGATT) Advertise() {
	a := &gatt.AdvPacket{}
	a.AppendFlags(advFlagGeneralDiscoverable | advFlagLEOnly)
	a.AppendField(advTypeAllUUID16, UUID16LE_FOR_LINECORP)
	a.AppendField(advTypeServiceData16, eb.OutFrame)

	fmt.Println(a.Len(), a)

	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			d.Advertise(a)
		default:
			log.Println(s)
		}
	}

	d.Init(onStateChanged)
	select {}
}
