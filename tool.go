package CoapPubsub

import (
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Parse interface (which is []uint8) to string
func ParseUint8ToString(in interface{}) string {
	val, ok := in.([]uint8)
	if ok {
		return string(val)
	} else {
		return ""
	}
}

//GetLocalRandomInt gets random number locally
func GetLocalRandomInt() uint16 {
	rand.Seed(time.Now().UnixNano())
	return uint16(rand.Intn(1000))
}

//GetIPv4Int16 gets local IPV4 address to uint16 by <<8
func GetIPv4Int16() uint16 {

	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		log.Println("No network:", err)
		return 0
	}

	for _, i := range ifaces {
		if strings.Contains(i.Name, "en0") {
			addrs, err := i.Addrs()
			// handle err
			if err != nil {
				log.Println("No IP:", err)
				return 0
			}

			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				// process IP address
				if ip[0] == 0 {
					//target XX.XX.XX.XX ipv4
					var myIP uint16
					myIP = uint16(ip[12])<<8 + uint16(ip[13])<<7 + uint16(ip[14])<<6 + uint16(ip[13])<<6
					return myIP
				}
			}
		}
	}

	return 0
}
