// Package shopee contains the Shopee merchant adapter.
package shopee

// DeviceRiskBlob is intentionally empty.
//
// Earlier upstream snapshots embedded a device-risk report captured from a
// real browser. Replaying a shared fingerprint links unrelated installations,
// may expose capture-device metadata, and can violate provider anti-fraud
// controls. An empty value asks the provider to issue a fresh risk token using
// its normal JSON flow. Keep any provider-authorized device report in a secret
// store and pass it through ProviderConfig.DeviceReport at runtime; never
// commit it to source control.
const DeviceRiskBlob = ""
