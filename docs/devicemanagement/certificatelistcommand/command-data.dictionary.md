# CertificateListCommand.Command

The command to get a list of installed certificates on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### ManagedOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, only include certificates that MDM installed or that are in the same profile as the MDM payload. User-enrolled devices ignore this value and always only include managed certificates.

Available: iOS 13+ | iPadOS 13+ | macOS 10.15+ | tvOS 13+ | visionOS 1.1+ | watchOS 10+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `CertificateList`

The request type for this command.

