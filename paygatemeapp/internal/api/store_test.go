package api

import "testing"

func TestValidWebhookURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{"public https", "https://merchant.example/webhooks/payment", true},
		{"http rejected", "http://merchant.example/webhook", false},
		{"localhost rejected", "https://localhost/webhook", false},
		{"loopback rejected", "https://127.0.0.1/webhook", false},
		{"private IPv4 rejected", "https://10.0.0.8/webhook", false},
		{"link local rejected", "https://169.254.169.254/latest/meta-data", false},
		{"credentials rejected", "https://user:pass@merchant.example/webhook", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validWebhookURL(tt.url); got != tt.want {
				t.Fatalf("validWebhookURL(%q) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}
