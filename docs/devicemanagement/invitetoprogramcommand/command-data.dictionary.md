# InviteToProgramCommand.Command

The command to invite a user to join the Volume Purchase Program (VPP).

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### InvitationURL

- **Type:** `string`
- **Required:** Yes

The Volume Purchase Program (VPP) invitation URL.

### ProgramID

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `com.apple.cloudvpp`

The program’s identifier, which can only be `com.apple.cloudvpp`.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InviteToProgram`

The request type for this command.

