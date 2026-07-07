# NetworkVPNIKEV2NetworkRoutingObject

Specifies details about how the VPN routes different types of network traffic.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### EnforceRoutes

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, all the VPN’s non-default routes take precedence over any locally defined routes.

If `IncludeAllNetworks` is `true`, the system ignores the value of `EnforceRoutes`.

### ExcludeAPNs

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true` and `IncludeAllNetworks` is `true`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel.

### ExcludeCellularServices

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true` and `IncludeAllNetworks` is `true`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.

### ExcludeDeviceCommunication

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `true` and `IncludeAllNetworks` is set to `true`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.

### ExcludeLocalNetworks

- **Type:** `boolean`
- **Required:** No

If `true` and `IncludeAllNetworks` is `true`, routes all local network traffic outside the VPN.

### IncludeAllNetworks

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs` and `ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:

- Traffic necessary for connecting and maintaining the device’s network connection, such as DHCP.
- Traffic necessary for connecting to captive networks.
- Certain cellular services traffic that’s not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.

