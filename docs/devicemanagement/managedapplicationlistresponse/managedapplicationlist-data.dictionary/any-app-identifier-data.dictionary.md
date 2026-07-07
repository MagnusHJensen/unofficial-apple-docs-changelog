# ManagedApplicationListResponse.ManagedApplicationList.ANY app identifier

The bundle identifier of the managed app.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### ExternalVersionIdentifier

- **Type:** `integer`
- **Required:** Yes

The app’s external version identifier. You can also retrieve this value from the App Store. For more information, see [Apps and books metadata for organizations](/documentation/devicemanagement/apps-and-books-metadata-for-organizations).

If the current external version identifier of an app on the App Store doesn’t match the external version identifier reported by the device, there may be an app update available for the device.

> 

Available: iOS 10.3+ | iPadOS 10.3+ | macOS 11.3+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### HasConfiguration

- **Type:** `boolean`
- **Required:** Yes

If ‘true’, the app has a server-provided managed configuration.

Available: iOS 7+ | iPadOS 7+ | macOS 11+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### HasFeedback

- **Type:** `boolean`
- **Required:** Yes

If ‘true’, the app has feedback for the server. On macOS 11.3 and later, this value is available if the device management server sent the request on the user channel.

Available: iOS 7+ | iPadOS 7+ | macOS 11.3+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### IsValidated

- **Type:** `boolean`
- **Required:** Yes

If ‘true’, the app is valid and can run on the device. If the app is enterprise-distributed and unvalidated, it won’t be able to run until validation has occurred.

Available: iOS 9.2+ | iPadOS 9.2+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### ManagementFlags

- **Type:** `integer`
- **Required:** Yes

The bitwise OR of the following management flags:

- ‘1’: Remove app upon removal of MDM profile.
- ‘4’: Prevent backup of app data.

### Status

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Queued`, `NeedsRedemption`, `Redeeming`, `Prompting`, `PromptingForLogin`, `ValidatingPurchase`, `PromptingForUpdate`, `PromptingForUpdateLogin`, `PromptingForManagement`, `ValidatingUpdate`, `Updating`, `Installing`, `Managed`, `ManagedButUninstalled`, `Unknown`, `UserInstalledApp`, `UserRejected`, `UpdateRejected`, `ManagementRejected`, `Failed`

The status of the managed app, which is one of the following values:

- `Queued`: The app is scheduled for installation.
- `NeedsRedemption`: The app needs a redemption code to complete installation.
- `Redeeming`: The device is redeeming the redemption code for the app.
- `Prompting`: The app installation is prompting the user.
- `PromptingForLogin’ - The app installation is prompting the user for App Store credentials.
- `ValidatingPurchase`: Validation of the app purchase is occurring.
- `PromptingForUpdate`: An app update is prompting the user.
- `PromptingForUpdateLogin`: An app update is prompting the user for App Store credentials.
- `PromptingForManagement`: Changing the app to a managed app is prompting the user.
- `ValidatingUpdate`: Validation of an app update is occurring.
- `Updating`: The app is updating.
- `Installing`: The app is installing.
- `Managed`: The installed app is a managed app.
- `ManagedButUninstalled`: The app is a managed app and the user removed it. Reinstalling the app reinstates it as a managed app.
- `Unknown`: The app state is unknown.

The following statuses are transient and report only once:

- `UserInstalledApp`: The user installed the app before managed app installation could occur.
- `UserRejected`: The user rejected the offer to install the app.
- `UpdateRejected`: The user rejected the offer to update the app.
- `ManagementRejected`:The user rejected management of an installed app.
- `Failed`: The app installation failed.

### UnusedRedemptionCode

- **Type:** `string`
- **Required:** Yes

If the user already purchased a paid app, this code is available for use by another user. This code reports only once.

Available: iOS 5+ | iPadOS 5+ | visionOS 1.1+ | watchOS 10+

