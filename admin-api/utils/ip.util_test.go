package utils

import "testing"

func TestIpAddress(t *testing.T) {
	address := IpAddress("103.151.173.196")
	t.Log(address)
}
