package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	lineBeacon "github.com/kkdai/line-beacon-go"
)

func main() {
	var hwid string
	var msg string
	flag.StringVar(&hwid, "h", "a1b2c3d4e5", "Hardward ID")
	flag.StringVar(&msg, "m", "00", "Message")
	flag.Parse()

	hwID, err := hex.DecodeString(hwid)
	if err != nil {
		log.Println("HW ID is not valid, err:", err)
	}

	msgData, err := hex.DecodeString(msg)
	if err != nil {
		log.Println("Msg data is not valid, err:", err)
	}

	beacon := lineBeacon.NewLineBeacon(hwID, msgData)
	fmt.Println("Your Line Beacon frame raw data:", hex.EncodeToString(beacon.OutFrame))
	beacon.Advertise()
}
