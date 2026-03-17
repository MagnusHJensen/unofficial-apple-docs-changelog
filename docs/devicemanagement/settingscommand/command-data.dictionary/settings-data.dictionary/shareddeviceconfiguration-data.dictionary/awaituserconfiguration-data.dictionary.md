# SettingsCommand.Command.Settings.SharedDeviceConfiguration.AwaitUserConfiguration

Enables the user configuration Setup Assistant workflow.

**Platforms:** iOS 17.0, iPadOS 17.0

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device stops at the Setup Assistant pane after user login. The user can’t use the device until it receives a [User Configured](/documentation/devicemanagement/user-configured-command) command.

