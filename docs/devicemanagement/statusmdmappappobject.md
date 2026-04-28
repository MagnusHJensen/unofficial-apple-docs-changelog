# StatusMDMAppAppObject

A status report that contains details about an MDM-installed app.

**Platforms:** iOS 16.0, iPadOS 16.0, Mac Catalyst 16.0, tvOS 16.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object. The device reports an MDM-installed app as removed if management of the app has been transferred to Declarative Device Management.

### external-version-id

- **Type:** `string`
- **Required:** No

The app’s external version identifier. You can also retrieve this value from the App Store. For more information, see [Apps and Books for Organizations](/documentation/devicemanagement/apps-and-books-for-organizations).

If the current external version identifier of an app on the App Store doesn’t match the external version identifier reported by the device, there may be an app update available for the device.

### identifier

- **Type:** `string`
- **Required:** Yes

The app’s bundle id, which is unique.

### name

- **Type:** `string`
- **Required:** No

The name of the app.

### short-version

- **Type:** `string`
- **Required:** No

The short version of the app.

### state

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `queued`, `needs-redemption`, `redeeming`, `prompting`, `prompting-for-login`, `validating-purchase`, `prompting-for-update`, `prompting-for-update-login`, `prompting-for-management`, `validating-update`, `updating`, `installing`, `managed`, `managed-but-uninstalled`, `unknown`, `user-installed-app`, `user-rejected`, `update-rejected`, `management-rejected`, `failed`

The status of the app that [ManagedApplicationListCommand](/documentation/devicemanagement/managedapplicationlistcommand) reports.

### version

- **Type:** `string`
- **Required:** No

The version of the app.

