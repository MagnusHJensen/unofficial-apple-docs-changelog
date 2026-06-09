# NSExtensionMappingsCommand.Command

The command to get a list of the installed extensions for a user on a device.

**Platforms:** macOS 10.13

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `NSExtensionMappings`

The request type for this command.

