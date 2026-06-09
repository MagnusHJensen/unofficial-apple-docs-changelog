# WiFi.QoSMarkingPolicy

A dictionary that defines the quality-of-service settings.

**Platforms:** iOS 10.0, iPadOS 10.0, Mac Catalyst 10.0, macOS 10.13, tvOS 9.0, visionOS 1.0, watchOS 3.2

## Properties

### QoSMarkingAllowListAppIdentifiers

- **Type:** `[string]`
- **Required:** No

An array of app bundle identifiers that defines the allow list for L2 and L3 marking for traffic that goes to the Wi-Fi network. If the array isn’t present, but the `QoSMarkingPolicy` key is present — even empty — no apps can use L2 and L3 marking.

Available: iOS 14.5+ | iPadOS 14.5+ | macOS 14+ | tvOS 9+ | visionOS 1+ | watchOS 3.2+

### QoSMarkingAppleAudioVideoCalls

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, adds audio and video traffic of built-in audio or video services, such as FaceTime and Wi-Fi Calling, to the allow list for L2 and L3 marking for traffic that goes to the Wi-Fi network.

### QoSMarkingEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, disables L3 marking and only uses L2 marking for traffic that goes to the Wi-Fi network.

If `false`, the system behaves as if Wi-Fi doesn’t have an association with a Cisco QoS fast lane network.

### QoSMarkingWhitelistedAppIdentifiers

- **Type:** `[string]`
- **Required:** No

Use `QoSMarkingAllowListAppIdentifiers` instead.

Available: iOS 10+ | iPadOS 10+ | macOS 10.13+ | tvOS 9+ | watchOS 3.2+
Deprecated: iOS 14.5+ | iPadOS 14.5+ | macOS 14+

