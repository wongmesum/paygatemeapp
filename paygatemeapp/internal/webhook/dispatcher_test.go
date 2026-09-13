package webhook

import (
	"net/netip"
	"testing"
)

func TestPublicIP(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{"1.1.1.1", true},
		{"8.8.8.8", true},
		{"127.0.0.1", false},
		{"10.0.0.1", false},
		{"169.254.169.254", false},
		{"::1", false},
		{"fc00::1", false},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			if got := publicIP(netip.MustParseAddr(tt.ip)); got != tt.want {
				t.Fatalf("publicIP(%s) = %v, want %v", tt.ip, got, tt.want)
			}
		})
	}
}
