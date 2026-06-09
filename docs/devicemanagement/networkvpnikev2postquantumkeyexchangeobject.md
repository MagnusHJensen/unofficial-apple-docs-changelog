# NetworkVPNIKEV2PostQuantumKeyExchangeObject

Post Quantum Key Exchange settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AllowFallback

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If set to `false`, the VPN doesn’t establish a connection if the server does not support or doesn’t allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` is not present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.

### PPK

- **Type:** `string`
- **Required:** No

The Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPKIdentifier` must also be present.

### PPKIdentifier

- **Type:** `string`
- **Required:** No

The identifier for the Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPK` must also be present.

### PPKMandatory

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `true`, the VPN doesn’t establish a connection if the server doesn’t support RFC 8784 or doesn’t accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` are not present.

