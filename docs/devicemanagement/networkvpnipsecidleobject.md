# NetworkVPNIPSecIdleObject

Specifies details about how the system handles idle VPN connections.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, visionOS 27.0

## Properties

### Disconnect

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disconnects after an on-demand connection idles.

### Timer

- **Type:** `integer`
- **Required:** No

The length of time to wait, in seconds, before disconnecting an on-demand connection.

