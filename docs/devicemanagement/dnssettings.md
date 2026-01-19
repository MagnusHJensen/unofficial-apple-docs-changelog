# DNSSettings

The payload that configures encrypted DNS settings.

**Platforms:** iOS 14.0, iPadOS 14.0, macOS 11.0, visionOS 1.0

## Discussion

Specify `com.apple.dnsSettings.managed` as the payload type.

When installed from an MDM, the setting only applies to managed Wi-Fi networks.

When installed manually, this setting also applies to cellular networks.

### Profile availability

## Topics

### Objects

- [DNSSettings.DNSSettings](/documentation/devicemanagement/dnssettings/dnssettings-data.dictionary) - A dictionary that defines a configuration for an encrypted DNS server.
- [DNSSettings.OnDemandRulesElement](/documentation/devicemanagement/dnssettings/ondemandruleselement) - A list of domain strings that determine which DNS queries use the DNS server.

