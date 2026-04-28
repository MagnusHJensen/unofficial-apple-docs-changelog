# UserListCommand.Command

The command to get a list of users with active accounts on a device.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `UserList`

The request type for this command.

