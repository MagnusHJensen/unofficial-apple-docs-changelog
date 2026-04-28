# VPN.IKEv2.IKESecurityAssociationParameters

The dictionary that contains security association parameters.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### DiffieHellmanGroup

- **Type:** `integer`
- **Required:** No
- **Default:** `14`
- **Allowed Values:** `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `31`, `32`

The Diffie-Hellman group.

For `AlwaysOn` VPN in iOS 14.2 and later, the minimum allowed value is `14`.

`1`, `2`, and `5` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.

### EncryptionAlgorithm

- **Type:** `string`
- **Required:** No
- **Default:** `AES-256`
- **Allowed Values:** `DES`, `3DES`, `AES-128`, `AES-256`, `AES-128-GCM`, `AES-256-GCM`, `ChaCha20Poly1305`

The encryption algorithm.

In watchOS and tvOS, the default value is `AES-256-GCM`. `DES` and `3DES` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.

### IntegrityAlgorithm

- **Type:** `string`
- **Required:** No
- **Default:** `SHA2-256`
- **Allowed Values:** `SHA1-96`, `SHA1-160`, `SHA2-256`, `SHA2-384`, `SHA2-512`

The integrity algorithm.

`SHA1-96` and `SHA1-160` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.

### LifeTimeInMinutes

- **Type:** `integer`
- **Required:** No
- **Default:** `1440`

The SA lifetime (rekey interval) in minutes.

### PostQuantumKeyExchangeMethods

- **Type:** `[integer]`
- **Required:** No
- **Allowed Values:** `0`, `36`, `37`

An array of strings representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.

