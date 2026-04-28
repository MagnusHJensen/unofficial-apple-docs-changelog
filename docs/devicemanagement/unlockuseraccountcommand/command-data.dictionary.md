# UnlockUserAccountCommand.Command

The command to unlock a user account that the system locked because of too many failed password attempts.

**Platforms:** macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `UnlockUserAccount`

The request type for this command.

### UserName

- **Type:** `string`
- **Required:** Yes

The user name of the local account, which can be any local account on the system, not just a managed user account.

