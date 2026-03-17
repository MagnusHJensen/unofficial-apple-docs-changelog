# DeviceInformationCommand.Command

The command to get detailed information about a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### DeviceAttestationNonce

- **Type:** `data`
- **Required:** No

A freshness code that appears in the resulting attestation. This value can contain up to 32 bytes of data. If specified, queries need to contain `DevicePropertiesAttestation`.

The MDM server uses this value to prove that an attestation was recently generated. The system caches the most recently generated attestation on the device. If omitted or if the value matches the cached attestation, the system returns the cached attestation. To request a new attestation, provide a new freshness code. Requests for new attestations are rate limited. If it is fewer than 7 days since the system generated an attestation, the device returns the cached attestation rather than generating a new one.

Available in iOS 16 and later, macOS 14 and later, tvOS 16 and later, and watchOS 10 and later. The hardware requirements for attestation are described below.

### Queries

- **Type:** `[DeviceInformationCommand.Command.Queries]`
- **Required:** Yes

An array of query dictionaries to get information about a device.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeviceInformation`

The request type for this command.

## Topics

### Objects

- [DeviceInformationCommand.Command.Queries](/documentation/devicemanagement/deviceinformationcommand/command-data.dictionary/queries-data.dictionary) - An array of query dictionaries to get information about a device.

