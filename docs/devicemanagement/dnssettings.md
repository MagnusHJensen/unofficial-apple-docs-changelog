# DNSSettings

The payload that configures encrypted DNS settings.

**Platforms:** iOS 14.0, iPadOS 14.0, macOS 11.0, visionOS 1.0

## Properties

### DNSSettings

- **Type:** `DNSSettings.DNSSettings`
- **Required:** Yes

A dictionary that defines a configuration for an encrypted DNS server.

### OnDemandRules

- **Type:** `[DNSSettings.OnDemandRulesElement]`
- **Required:** No

An array of rules that define the DNS settings. If not set, the system always applies the DNS settings. These rules are identical to the `OnDemandRules` array in VPN payloads.

### ProhibitDisablement

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prohibits users from disabling DNS settings. This key is only available on supervised devices.

## Discussion

Specify `com.apple.dnsSettings.managed` as the payload type.

When installed from an MDM, the setting only applies to managed Wi-Fi networks.

When installed manually, this setting also applies to cellular networks.

### Profile availability

## Topics

### Objects

- [DNSSettings.DNSSettings](/documentation/devicemanagement/dnssettings/dnssettings-data.dictionary) - A dictionary that defines a configuration for an encrypted DNS server.
- [DNSSettings.OnDemandRulesElement](/documentation/devicemanagement/dnssettings/ondemandruleselement) - A list of domain strings that determine which DNS queries use the DNS server.

