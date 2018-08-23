package main

import (
	"fmt"

	lineBeacon "github.com/kkdai/line-beacon-go"
)

func main() {
	hwID := []byte{0x01, 0x02}
	msgData := []byte{0x01, 0x02}

	totalFrameData := lineBeacon.CreateLineSimpleBeaconAdvertisingPDU(hwID, msgData)
	fmt.Println("Your Line Beacon frame raw data:", totalFrameData)

}
