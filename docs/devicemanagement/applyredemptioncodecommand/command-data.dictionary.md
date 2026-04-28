# ApplyRedemptionCodeCommand.Command

The command to complete the installation of an app using a redemption code.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, Device Assignment Services , VPP License Management 

## Properties

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app.

### RedemptionCode

- **Type:** `string`
- **Required:** Yes

The redemption code that applies to the app pending installation.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ApplyRedemptionCode`

The request type for this command.

