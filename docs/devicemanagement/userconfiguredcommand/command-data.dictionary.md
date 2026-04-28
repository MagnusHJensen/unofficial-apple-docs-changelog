# UserConfiguredCommand.Command

The command to inform the device that it can continue past Setup Assistant and finish login.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `UserConfigured`

The request type for this command.

