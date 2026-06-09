# NetworkDNSSettings

The declaration to configure encrypted DNS settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### DNSSettings

- **Type:** `NetworkDNSSettingsDNSSettingsObject`
- **Required:** Yes

A dictionary that defines a configuration for an encrypted DNS server.

### OnDemandRules

- **Type:** `[NetworkDNSSettingsOnDemandRulesElementObject]`
- **Required:** No

An array of rules that define the DNS settings. If not set, the system always applies the DNS settings. These rules are identical to the `OnDemandRules` array in VPN payloads.

### ProhibitDisablement

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prohibits users from disabling DNS settings.

Allowed enrollments: supervised

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the DNS settings that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.dns-settings` as the declaration type.

The following rules determine which networks the settings apply to:

- For supervised enrollments, the settings apply to all networks.
- For device enrollments, the settings are limited to only managed networks.
- For local installs, the settings apply to all networks.

### Configuration availability

### Configuration example

This configuration sets up encrypted DNS using DNS-over-HTTPS.

```json
{
    "Type": "com.apple.configuration.network.dns-settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "DNS Settings",
        "DNSSettings": {
            "DNSProtocol": "HTTPS",
            "ServerURL": "https://dns.example.com/dns-query",
            "ServerAddresses": [
                "12.12.12.12"
            ],
            "AllowFailover": false,
            "SupplementalMatchDomains": [
                "example.com"
            ]
        }
    }
}
```

## Topics

### Objects

- [NetworkDNSSettingsDNSSettingsObject](/documentation/devicemanagement/networkdnssettingsdnssettingsobject) - A dictionary that defines a configuration for an encrypted DNS server.
- [NetworkDNSSettingsOnDemandRulesElementObject](/documentation/devicemanagement/networkdnssettingsondemandruleselementobject) - An array of rules that define the DNS settings. If not set, the system always applies the DNS settings. These rules are identical to the `OnDemandRules` array in VPN payloads.

