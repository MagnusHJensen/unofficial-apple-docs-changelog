# DNSProxy

The payload that configures DNS proxies.

**Platforms:** iOS 11.0, iPadOS 11.0, macOS 10.15, visionOS 1.1

## Properties

### AppBundleIdentifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app containing the DNS proxy network extension.

### DNSProxyUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this DNS proxy configuration. The proxy processes DNS lookups traffic for managed apps with the same `DNSProxyUUID` in their app attributes. This key is required for user enrollment.

### ProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

The bundle identifier of the DNS proxy network extension to use. Declaring the bundle identifier is useful for apps that contain more than one DNS proxy extension.

### ProviderConfiguration

- **Type:** `DNSProxy.ProviderConfiguration`
- **Required:** No

The dictionary of vendor-specific configuration items.

## Discussion

Specify `com.apple.dnsProxy.managed` as the payload type.

Beginning with iOS 15, this profile is unsupervised and needs to be installed through MDM.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>AppBundleIdentifier</key>
            <string>com.example.mydnsproxyapp</string>
            <key>ProviderBundleIdentifier</key>
            <string>com.example.mydnsproxyapp.mydnsproxyprovider</string>
            <key>ProviderConfiguration</key>
            <dict>
                <key>resolver</key>
                <dict>
                    <key>ipaddress</key>
                    <string>9.9.9.9</string>
                </dict>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.mydnsproxypayload</string>
            <key>PayloadType</key>
            <string>com.apple.dnsProxy.managed</string>
            <key>PayloadUUID</key>
            <string>D6B3F3E4-A72E-49F1-AE2E-742A3A11BE1D</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>DNS Proxy</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>6a60bbe0-242c-493d-b338-c5885107f2af</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [DNSProxy.ProviderConfiguration](/documentation/devicemanagement/dnsproxy/providerconfiguration-data.dictionary) - The dictionary of vendor-specific configuration items.

