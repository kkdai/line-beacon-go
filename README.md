# LINE Simple Beacon in Go

LINE Simple Beacon is a specification for beacon hardware for LINE Bot developers.

[Spec in English from Line](https://github.com/line/line-simple-beacon/blob/master/README.en.md)

## Install in MacOSX (WIP)

`go install github.com/kkdai/line-beacon-go/tools/macosx`

## How to use it

```go

import lineBeacon "github.com/kkdai/line-beacon-go"

hwID := []byte{0x01, 0x02, ...}
msgData := []byte{0x01, 0x02, ...}

totalFrameData := lineBeacon.CreateLineSimpleBeaconAdvertisingPDU(hwID, msgData)
fmt.Println("Your Line Beacon frame raw data:", totalFrameData)
```

## Tools

WIP

## Inspire by Line Simple Beacon

[Repo](https://github.com/line/line-simple-beacon)

## Disclaimer

This is WIP libray and I still don't get realone line-beacon on hand, but it fullfil the testing case in Line Simple Beacon