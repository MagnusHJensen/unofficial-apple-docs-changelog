# SettingsCommand.Command.Settings.MaximumResidentUsers

A dictionary that contains settings for maximum resident users.

**Platforms:** iOS 9.3, iPadOS 9.3

## Properties

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `MaximumResidentUsers`

A string that identifies this setting.

### MaximumResidentUsers

- **Type:** `integer`
- **Required:** Yes

The maximum number of users that can use the device. If this value is greater than the value for the maximum possible number of users that the device supports, the MDM server uses that value instead.

This setting requires that the device is in the `AwaitingConfiguration` phase before it receives the [DeviceConfiguredCommand](/documentation/devicemanagement/deviceconfiguredcommand) message.

When a device reaches the maximum number of resident users and a new user tries to sign in, the MDM server removes a synchronized user to make space for the new user. If there are no synchronized users, the new user sign-in fails. A synchronized user is a user that has completed syncing their data.

