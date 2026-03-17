# WiFi.EAPClientConfiguration

A dictionary that configures an enterprise network.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 3.2

## Properties

### AcceptEAPTypes

- **Type:** `[integer]`
- **Required:** Yes
- **Allowed Values:** `13`, `17`, `18`, `21`, `23`, `25`, `43`

The EAP types that the system accepts. Allowed values:

- `13`: EAP-TLS
- `17`: LEAP
- `18`: EAP-SIM
- `21`: EAP-TTLS
- `23`: EAP-AKA
- `25`: PEAPv0/v1
- `43`: EAP-FAST

For EAP-TLS authentication without a network payload, install the necessary identity certificates and have your users select EAP-TLS mode in the 802.1X credentials dialog that appears when they connect to the network. For other EAP types, a network payload is necessary and must specify the correct settings for the network.

### EAPFASTProvisionPAC

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, allows PAC provisioning.

This value is only applicable if ‘EAPFASTUsePAC’ is ‘true’. This value must be ‘true’ for EAP-FAST PAC usage to succeed because there’s no other way to provision a PAC.

### EAPFASTProvisionPACAnonymously

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, provisions the device anonymously. Note that there are known machine-in-the-middle attacks for anonymous provisioning.

### EAPFASTUsePAC

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, the device uses an existing PAC if it’s present. Otherwise, the server must present its identity using a certificate.

### EAPSIMNumberOfRANDs

- **Type:** `integer`
- **Required:** No
- **Default:** `3`
- **Allowed Values:** `2`, `3`

The minimum number of RAND values to accept from the server. For use with EAP-SIM only.

Available in iOS 8 and later, macOS 10.7 and later, tvOS 9 and later, visionOS 1 and later, and watchOS 3.2 and later.

### OneTimeUserPassword

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, the user receives a prompt for a password each time they connect to the network.

Available in iOS 8 and later, macOS 10.8 and later, tvOS 9 and later, visionOS 1 and later, and watchOS 3.2 and later.

### OuterIdentity

- **Type:** `string`
- **Required:** No

A name that hides the user’s true name. The user’s actual name appears only inside the encrypted tunnel. For example, you might set this to anonymous or anon, or anon@mycompany.net. It can increase security because an attacker can’t see the authenticating user’s name in the clear. This key is only relevant to TTLS, PEAP, and EAP-FAST. This field is required if ‘TLSMinimumVersion’ is ‘1.3’.

### PayloadCertificateAnchorUUID

- **Type:** `[string]`
- **Required:** No

An array of the UUID of each certificate payload in the same profile to trust for authentication. Use this key to prevent the device from asking the user whether to trust the listed certificates. Dynamic trust (the certificate dialogue) is in a disabled state if you specify this property without also enabling ‘TLSAllowTrustExceptions’.

### SystemModeCredentialsSource

- **Type:** `string`
- **Required:** No

Set this string to ‘ActiveDirectory’ to use the AD computer name and password credentials. If using this property, you can’t use ‘SystemModeUseOpenDirectoryCredentials’.

### SystemModeUseOpenDirectoryCredentials

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, the system mode connection tries to use the Open Directory credentials. If using this property, you can’t use ‘SystemModeCredentialsSource’.

### TLSCertificateIsRequired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, allows for two-factor authentication for EAP-TTLS, PEAP, or EAP-FAST. If ‘false’, allows for zero-factor authentication for EAP-TLS. If you don’t specify a value, the default is ‘true’ for EAP-TLS, and ‘false’ for other EAP types.

Available in iOS 7 and later, macOS 10.7 and later, tvOS 9 and later, visionOS 1 and later, and watchOS 3.2 and later.

### TLSMaximumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.2`
- **Allowed Values:** `1.0`, `1.1`, `1.2`, `1.3`

The maximum TLS version for EAP authentication.

Available in iOS 11 and later, macOS 10.13 and later, tvOS 11 and later, visionOS 1 and later, and watchOS 3.2 and later.

### TLSMinimumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.0`
- **Allowed Values:** `1.0`, `1.1`, `1.2`, `1.3`

The minimum TLS version for EAP authentication.

Available in iOS 11 and later, macOS 10.13 and later, tvOS 11 and later, visionOS 1 and later, and watchOS 3.2 and later.

### TLSTrustedCertificates

- **Type:** `[string]`
- **Required:** No

An array of trusted certificates. Each entry in the array must contain certificate data that represents an anchor certificate used for verifying the server certificate.

### TLSTrustedServerNames

- **Type:** `[string]`
- **Required:** No

The list of accepted server certificate common names. If a server presents a certificate that isn’t in this list, the system doesn’t trust it. If you specify this property, the system disables dynamic trust (the certificate dialog) unless you also specify ‘TLSAllowTrustExceptions’ with the value ‘true’. If necessary, use wildcards to specify the name, such as ‘wpa.*.example.com’.

### TTLSInnerAuthentication

- **Type:** `string`
- **Required:** No
- **Default:** `MSCHAPv2`
- **Allowed Values:** `PAP`, `EAP`, `CHAP`, `MSCHAP`, `MSCHAPv2`

The inner authentication that the TTLS module uses.

### UserName

- **Type:** `string`
- **Required:** No

The user name for the account. If you don’t specify a value, the system prompts the user during login.

### UserPassword

- **Type:** `string`
- **Required:** No

The user’s password. If you don’t specify a value, the system prompts the user during login.

