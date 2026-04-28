# LogOutUserCommand.Command

The command to force the current user to log out of a device.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `LogOutUser`

The request type for this command.

