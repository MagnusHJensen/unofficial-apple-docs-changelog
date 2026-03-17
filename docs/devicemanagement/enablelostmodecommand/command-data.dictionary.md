# EnableLostModeCommand.Command

The command to enable Lost Mode on a device, which provides a message and phone number on the Lock Screen.

**Platforms:** iOS 9.3, iPadOS 9.3

## Properties

### Footnote

- **Type:** `string`
- **Required:** No

If present, the device displays this text at the bottom of the Lock Screen.

### Message

- **Type:** `string`
- **Required:** No

If present, the device displays this text on the Lock Screen. You must provide this value if you don’t provide a value for `PhoneNumber`.

### PhoneNumber

- **Type:** `string`
- **Required:** No

If present, the device displays this phone number on the Lock Screen. You must provide this value if you don’t provide a value for `Message`.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `EnableLostMode`

The request type for this command.

