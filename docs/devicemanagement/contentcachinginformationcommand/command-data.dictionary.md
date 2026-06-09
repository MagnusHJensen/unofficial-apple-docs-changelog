# ContentCachingInformationCommand.Command

The command to get the status of the content caches on a device.

**Platforms:** macOS 10.15.4

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ContentCachingInformation`

The request type for this command.

