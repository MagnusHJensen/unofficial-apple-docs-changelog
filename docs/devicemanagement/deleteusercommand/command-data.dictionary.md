# DeleteUserCommand.Command

The command to delete a user’s account from a device.

**Platforms:** iOS 9.3, iPadOS 9.3, macOS 10.13

## Properties

### DeleteAllUsers

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system attempts to delete all users from the device. If `ForceDeletion` is `false`, the system generates an error instead and doesn’t delete users who have data that’s pending sync. This value is available in iOS 14 and later.

### ForceDeletion

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system deletes the account even if the user has data that’s pending sync to the cloud. This value is available on iOS 9.3 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeleteUser`

The request type for this command.

### UserName

- **Type:** `string`
- **Required:** No

The user name of the account to delete. This key is required when the value for `DeleteAllUsers` is absent or `false`.

