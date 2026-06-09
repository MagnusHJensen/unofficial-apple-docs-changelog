# InstalledApplicationListCommand.Command

The command to get a list of the installed apps on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### Identifiers

- **Type:** `[string]`
- **Required:** No

An array of app identifiers. Provide this value to limit the response to only include these apps.

> 

Available: iOS 7+ | iPadOS 7+ | macOS 10.15+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### Items

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `AdHocCodeSigned`, `AppStoreVendable`, `BetaApp`, `BundleSize`, `DeviceBasedVPP`, `DistributorIdentifier`, `DynamicSize`, `ExternalVersionIdentifier`, `HasUpdateAvailable`, `Identifier`, `Installing`, `IsAppClip`, `IsValidated`, `Name`, `ShortVersion`, `Version`

An array of strings that represent keys in [InstalledApplicationListResponse.InstalledApplicationListItem](/documentation/devicemanagement/installedapplicationlistresponse/installedapplicationlistitem). If present, the response only contains the keys listed here, except `Identifier` is always included. If not present, the response contains all keys. Starting in iOS 26, macOS 26, tvOS 26, watchOS 26, and visionOS 26, if this key isn’t present, the response omits values that are expensive to calculate.

> 

Available: iOS 14+ | iPadOS 14+ | tvOS 14+ | visionOS 1.1+ | watchOS 10+

### ManagedAppsOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, only get a list of managed apps, excluding ones that Declarative Device Management is managing.

> 

Available: iOS 7+ | iPadOS 7+ | macOS 10.15+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstalledApplicationList`

The request type for this command.

