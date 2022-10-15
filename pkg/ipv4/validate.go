package ipv4

import "net"

func ValidateIPv4(ipv4 string) bool {
	parsed := net.ParseIP(ipv4)
	return parsed != nil
}
