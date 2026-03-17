# StatusAppManagedListAppObject

A dictionary that describes a declarative managed app.

**Platforms:** iOS 17.2, iPadOS 17.2, macOS 26.0, visionOS 2.4

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.

### config-state

- **Type:** `StatusAppManagedListManagedConfigurationObject`
- **Required:** No

The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the declaration that controls the app.

### external-version-id

- **Type:** `integer`
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

### reasons

- **Type:** `[StatusAppManagedListStatusReasonObject]`
- **Required:** No

An array that contains additional details about the app state, including errors.

### short-version

- **Type:** `string`
- **Required:** No

The short version of the app.

### state

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `optional`, `queued`, `not-present`, `prompting-for-consent`, `prompting-for-login`, `prompting-for-management`, `downloading`, `installing`, `managed`, `managed-but-uninstalled`, `failed`

The status of the app, which has the following possible values:

- `optional`: The app is optional and the user has to trigger its installation.
- `queued`: The system has started installation of the app.
- `not-present`: Management of the app occurs after it is installed.
- `prompting-for-consent`: The system is displaying a prompt to the user to proceed with app installation.
- `prompting-for-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
- `prompting-for-management`: The system is displaying a prompt to the user to allow changing the installed app to a managed app.
- `downloading`: The system is downloading the app.
- `installing`: The system is installing the app.
- `managed`: The app is installed and managed.
- `managed-but-uninstalled`: The app is required, but the system hasn’t installed it. The app becomes managed if the system installs it again. If the user removes an optional app, its state is `optional`, not `managed-but-uninstalled`.
- `failed`: The app install failed.

### update-state

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `available`, `prompting-for-update`, `prompting-for-update-login`, `updating`, `failed`

The update status of the app, which has the following possible values:

- `available`: An update is available for the app.
- `prompting-for-update`: The system is displaying a prompt to the user to proceed with app installation.
- `prompting-for-update-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
- `updating`: The app is updating.
- `failed`: The app update failed.

> 

### version

- **Type:** `string`
- **Required:** No

The version of the app.

## Topics

### Objects

- [StatusAppManagedListManagedConfigurationObject](/documentation/devicemanagement/statusappmanagedlistmanagedconfigurationobject) - A dictionary that contains details about a declarative managed app’s managed configuration.
- [StatusAppManagedListStatusReasonObject](/documentation/devicemanagement/statusappmanagedliststatusreasonobject) - A dictionary that contains details about a declarative managed app’s state.

