# ManagedApplicationFeedbackCommand.Command

The command to get app feedback from a managed app on the device.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 11.0, tvOS 10.2, visionOS 1.1

## Properties

### DeleteFeedback

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, delete the app’s feedback dictionary after the server reads it. Apps that are managed by Declarative Device Management are ignored.

### Identifiers

- **Type:** `[string]`
- **Required:** Yes

The bundle identifiers of the managed apps.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ManagedApplicationFeedback`

The request type for this command.

