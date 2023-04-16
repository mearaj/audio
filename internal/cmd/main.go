package main

import (
	"fmt"
	"github.com/mearaj/audio/internal"
	_ "github.com/mearaj/audio/internal"
)

func main() {
	devices := internal.QueryDevicesInfoCache()
	for _, d := range devices.GetAllDevicesInfo() {
		fmt.Println(d.Name)
		fmt.Println(string(d.ID))
	}
}
