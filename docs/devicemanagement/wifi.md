# WiFi

The payload that configures Wi-Fi settings.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 3.2

## Properties

### AllowJoinBeforeFirstUnlock

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device makes this network available for joining before the device is unlocked for the first time following a reboot, on a device configured for return to service. The device places any network credentials into Class D storage within the keychain and stores information about the network on disk in Class D.

There are several restrictions on the use of this flag:

- This property is only available in the return to service mode.
- You can designate only one network as available before first unlock.
- `EAPClientConfiguration` must not be present.
- If `IsHotspot` is present, it must be set to `false`.
- `QoSMarkingPolicy` must not be present.
- If `ProxyType` is present, it must be set to `None`.

The device fails to install the profile payload if any of these conditions aren’t met.

Available: visionOS 26+

### AutoJoin

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the device joins the network automatically.

If `false`, the user must tap the network name to join it.

Available: iOS 5+ | iPadOS 5+ | macOS 10.7+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### CaptiveBypass

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system bypasses Captive Network detection when the device connects to the network.

Available: iOS 10+ | iPadOS 10+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### DisableAssociationMACRandomization

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true,` disables MAC address randomization for a Wi-Fi network while associated with that network. This feature also shows a privacy warning in Settings indicating that the network has reduced privacy protections.

If `false`, then the system enables MAC address randomization on iOS, watchOS, and visionOS.

This value is only locked when MDM installs the profile. If the profile is manually installed, the system sets the value but the user can change it.

Available: iOS 14+ | iPadOS 14+ | macOS 15+ | visionOS 1+ | watchOS 7+

### DisplayedOperatorName

- **Type:** `string`
- **Required:** No

The operator name to display when connected to this network. Used only with Wi-Fi Hotspot 2.0 access points.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### DomainName

- **Type:** `string`
- **Required:** No

The primary domain of the tunnel.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### EAPClientConfiguration

- **Type:** `WiFi.EAPClientConfiguration`
- **Required:** No

The enterprise network configuration.

### EnableIPv6

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables IPv6 on this interface.

### EncryptionType

- **Type:** `string`
- **Required:** No
- **Default:** `Any`
- **Allowed Values:** `WEP`, `WPA`, `WPA2`, `WPA3`, `Any`, `None`

The encryption type for the network.

If set to anything except `None`, the payload may contain the following three keys: `Password`, `PayloadCertificateUUID`, or `EAPClientConfiguration`.

As of iOS 16, tvOS 16, watchOS 9, and macOS 13:

- `WPA` allows joining WPA or WPA2 networks
- `WPA2` allows joining WPA2 or WPA3 networks
- `WPA3` allows joining WPA3 networks only
- `Any` allows joining WPA, WPA2, WPA3, and WEP networks

Prior to iOS 16, tvOS 16, and watchOS 9, specifying `WPA`, `WPA2`, and `WPA3` were equivalent and would allow joining any WPA network.

Prior to macOS 13, the encryption type, if specified explicitly, needed to match the encryption type of the network exactly.

### HESSID

- **Type:** `string`
- **Required:** No

The HESSID used for Wi-Fi Hotspot 2.0 negotiation.

Available: iOS 7+ | iPadOS 7+ | macOS 10.7+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### HIDDEN_NETWORK

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, defines this network as hidden.

### IsHotspot

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device treats the network as a hotspot.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### MCCAndMNCs

- **Type:** `[string]`
- **Required:** No

An array of Mobile Country Code/Mobile Network Code (MCC/MNC) pairs used for Wi-Fi Hotspot 2.0 negotiation. Each string must contain exactly six digits.

Available: iOS 7+ | iPadOS 7+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### NAIRealmNames

- **Type:** `[string]`
- **Required:** No

An array of Network Access Identifier Realm names used for Wi-Fi Hotspot 2.0 negotiation.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### Password

- **Type:** `string`
- **Required:** No

The password for the access point.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use for the client credential.

### ProxyPACFallbackAllowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows connecting directly to the destination if the PAC file is unreachable.

### ProxyPACURL

- **Type:** `string`
- **Required:** No

The URL of the PAC file that defines the proxy configuration.

### ProxyPassword

- **Type:** `string`
- **Required:** No

The password used to authenticate to the proxy server.

### ProxyServer

- **Type:** `string`
- **Required:** No

The proxy server’s network address.

### ProxyServerPort

- **Type:** `integer`
- **Required:** No

The proxy server’s port number.

### ProxyType

- **Type:** `string`
- **Required:** No
- **Default:** `None`
- **Allowed Values:** `None`, `Manual`, `Auto`

The proxy type, if any, to use. If you choose the manual proxy type, you need the proxy server address, including its port and optionally a user name and password into the proxy server. If you choose the auto proxy type, you can enter a proxy autoconfiguration (PAC) URL.

### ProxyUsername

- **Type:** `string`
- **Required:** No

The user name used to authenticate to the proxy server.

### QoSMarkingPolicy

- **Type:** `WiFi.QoSMarkingPolicy`
- **Required:** No

A dictionary that contains the list of apps that the system allows to benefit from L2 and L3 marking. When this dictionary isn’t present, the system allows all apps to use L2 and L3 marking when the Wi-Fi network supports Cisco QoS fast lane.

Available: iOS 10+ | iPadOS 10+ | macOS 10.13+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### RoamingConsortiumOIs

- **Type:** `[string]`
- **Required:** No

An array of Roaming Consortium Organization Identifiers used for Wi-Fi Hotspot 2.0 negotiation.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### ServiceProviderRoamingEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows connection to roaming service providers.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### SetupModes

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `System`, `Loginwindow`

An array of strings that contain the type of connection mode to attach.

Available: macOS 10.7+

### SSID_STR

- **Type:** `string`
- **Required:** No

The SSID of the Wi-Fi network to use. In iOS 7.0 and later, the SSID is optional if a value exists for `DomainName` value.

Available: iOS 7+ | iPadOS 7+ | macOS 10.7+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### TLSCertificateRequired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows for two-factor authentication for EAP-TTLS, PEAP, or EAP-FAST. If `false`, allows for zero-factor authentication for EAP-TLS.

## Discussion

Specify `com.apple.wifi.managed` as the payload type.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>New item</key>
    <dict>
        <key>PayloadContent</key>
        <array>
            <dict>
                <key>AutoJoin</key>
                <true/>
                <key>CaptiveBypass</key>
                <false/>
                <key>DisableAssociationMACRandomization</key>
                <false/>
                <key>EncryptionType</key>
                <string>WPA</string>
                <key>HIDDEN_NETWORK</key>
                <true/>
                <key>IsHotspot</key>
                <false/>
                <key>Password</key>
                <string>Password123</string>
                <key>PayloadDisplayName</key>
                <string>Wi-Fi</string>
                <key>ProxyType</key>
                <string>None</string>
                <key>SSID_STR</key>
                <string>Example</string>
                <key>PayloadIdentifier</key>
                <string>com.example.mywifipayload</string>
                <key>PayloadType</key>
                <string>com.apple.wifi.managed</string>
                <key>PayloadUUID</key>
                <string>94c487e0-d6f8-41e3-b66d-a89994e6919b</string>
                <key>PayloadVersion</key>
                <integer>1</integer>
            </dict>
        </array>
        <key>PayloadDisplayName</key>
        <string>Wi-Fi</string>
        <key>PayloadIdentifier</key>
        <string>com.example.myprofile</string>
        <key>PayloadType</key>
        <string>Configuration</string>
        <key>PayloadUUID</key>
        <string>71e9b0f7-02f8-4aea-b365-b381d872909a</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
    </dict>
</dict>
</plist>
```

## Topics

### Objects

- [WiFi.EAPClientConfiguration](/documentation/devicemanagement/wifi/eapclientconfiguration-data.dictionary) - A dictionary that configures an enterprise network.
- [WiFi.QoSMarkingPolicy](/documentation/devicemanagement/wifi/qosmarkingpolicy-data.dictionary) - A dictionary that defines the quality-of-service settings.

