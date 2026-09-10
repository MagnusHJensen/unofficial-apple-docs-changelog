# NetworkVPNAlwaysOnTunnelConfigurationElement_IKEV2_IdleObject

Specifies details about how the system handles idle VPN connections.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, visionOS 27.0

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

