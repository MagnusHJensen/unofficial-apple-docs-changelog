# NetworkDNSProxy

The declaration to configure DNS proxy settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AppBundleIdentifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app containing the DNS proxy network extension.

### DNSProxyUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this DNS proxy configuration. The proxy processes DNS lookups traffic for managed apps with the same `DNSProxyUUID` in their app attributes. This key is required for user enrollment.

Available: iOS 27+ | iPadOS 27+ | visionOS 27+

### ProviderComposedIdentifier

- **Type:** `string`
- **Required:** No

The identifier of the provider to use for this configuration. Useful for apps that contain more than one DNS proxy extension.

In iOS and visionOS, the identifier is a bundle ID, for example, “com.example.app”.

In macOS, the identifier is a composed identifier. The format of the composed identifier is either “Bundle-ID” or “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the provider. “Designated-Requirement” is the designated requirement string from the code signature of the provider. For example, “com.example.app” for the bundle ID format, or “com.example.app {anchor apple generic}” for the designated requirement format.

### ProviderConfiguration

- **Type:** `NetworkDNSProxyProviderConfigurationObject`
- **Required:** No

The dictionary of vendor-specific configuration items.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the DNS proxy configuration that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.dns-proxy` as the declaration type.

### Configuration availability

### Configuration example

This configuration sets up a DNS proxy using a Network Extension app.

```json
{
    "Type": "com.apple.configuration.network.dns-proxy",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "DNS Proxy",
        "AppBundleIdentifier": "com.example.mydnsproxyapp",
        "ProviderComposedIdentifier": "com.example.mydnsproxyapp.mydnsproxyprovider",
        "ProviderConfiguration": {
            "resolver": {
                "ipaddress": "9.9.9.9"
            }
        }
    }
}
```

## Topics

### Objects

- [NetworkDNSProxyProviderConfigurationObject](/documentation/devicemanagement/networkdnsproxyproviderconfigurationobject) - The dictionary of vendor-specific configuration items.

