# NetworkVPNAlwaysOnSecurityAssociationParametersObject

These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### DiffieHellmanGroup

- **Type:** `integer`
- **Required:** No
- **Default:** `14`
- **Allowed Values:** `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `31`, `32`

The Diffie-Hellman group.

For `AlwaysOn` VPN, the minimum allowed value is `14`.

### EncryptionAlgorithm

- **Type:** `string`
- **Required:** No
- **Default:** `AES-256`
- **Allowed Values:** `AES-128`, `AES-256`, `AES-128-GCM`, `AES-256-GCM`, `ChaCha20Poly1305`

The encryption algorithm.

In tvOS the default value is `AES-256-GCM`.

### IntegrityAlgorithm

- **Type:** `string`
- **Required:** No
- **Default:** `SHA2-256`
- **Allowed Values:** `SHA2-256`, `SHA2-384`, `SHA2-512`

The integrity algorithm.

### LifeTimeInMinutes

- **Type:** `integer`
- **Required:** No
- **Default:** `1440`

The SA lifetime (rekey interval) in minutes.

### PostQuantumKeyExchangeMethods

- **Type:** `[integer]`
- **Required:** No
- **Allowed Values:** `0`, `36`, `37`

An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.

