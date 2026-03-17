# LOMSetupRequestCommand.Command

The command to get information from a device to set up lights-out management (LOM).

**Platforms:** macOS 11.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `LOMSetupRequest`

The request type for this command.

