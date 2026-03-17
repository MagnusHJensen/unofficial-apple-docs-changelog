# RefreshCellularPlansCommand.Command

The command to query a carrier URL for active eSIM cellular-plan profiles on a device.

**Platforms:** iOS 13.0, iPadOS 13.0

## Properties

### eSIMServerURL

- **Type:** `string`
- **Required:** Yes

The carrier’s eSIM server URL to query. Obtain this URL from each carrier separately.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RefreshCellularPlans`

The request type for this command.

