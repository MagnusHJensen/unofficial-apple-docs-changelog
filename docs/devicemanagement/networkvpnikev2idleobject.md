# NetworkVPNIKEV2IdleObject

Specifies details about how the system handles idle VPN connections.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### DeadPeerDetectionRate

- **Type:** `string`
- **Required:** No
- **Default:** `Medium`
- **Allowed Values:** `None`, `Low`, `Medium`, `High`

One of the following:

- `None`: No keepalive.
- `Low`: Send keepalive every 30 minutes.
- `Medium`: Send keepalive every 10 minutes.
- `High`: Send keepalive every 1 minute.

### Disconnect

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disconnects after an on-demand connection idles.

### Timer

- **Type:** `integer`
- **Required:** No

The length of time to wait, in seconds, before disconnecting an on-demand connection.

