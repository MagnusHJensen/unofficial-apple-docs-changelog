# TriggerEnhancedLogCollectionCommand.Command

The command to trigger enhanced log collection on the device.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0

## Properties

### AppleCareToken

- **Type:** `string`
- **Required:** Yes

The AppleCare token the device uses for authorizing the enhanced log collection session.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `TriggerEnhancedLogCollection`

The request type for this command.

