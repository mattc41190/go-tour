package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	hosts := map[string]IPAddr{
		// Remember we do not need to specify type here (https://tour.golang.org/moretypes/21)
		"loopback": {127, 0, 0, 1},
		"google":   {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// Example w/o `fmt` package
// func (ipAddr IPAddr) String() string {
// 	var result string
// 	for i, group := range ipAddr {
// 		sGroup := strconv.Itoa(int(uint8(group)))
// 		if i == 3 {
// 			result += sGroup
// 			continue
// 		}
// 		result += sGroup + "."
// 	}
// 	return result
// }
