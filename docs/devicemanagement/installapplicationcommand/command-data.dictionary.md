# InstallApplicationCommand.Command

The command to install a third-party app on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### Attributes

- **Type:** `InstallApplicationCommand.Command.Attributes`
- **Required:** No

A dictionary that contains the initial attributes of the app, if you choose to provide it.

Available: iOS 7+ | iPadOS 7+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### ChangeManagementState

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Managed`

The change management state. This value doesn’t work with the user enrollment feature introduced in iOS 13, or any type of account driven enrollment. The only possible value is:

- `Managed`: Take management of the app if the user installed it already and `InstallAsManaged` is `true`.

Available: iOS 9+ | iPadOS 9+ | macOS 11+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### Configuration

- **Type:** `InstallApplicationCommand.Command.Configuration`
- **Required:** No

A dictionary that contains the initial configuration of the app, if you choose to provide it.

Available: iOS 7+ | iPadOS 7+ | macOS 11+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### Identifier

- **Type:** `string`
- **Required:** No

The app’s bundle identifier.

> 

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### InstallAsManaged

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, install the app as a managed app. Otherwise, the system installs the app as unmanaged. If you reinstall a manged app and omit this value or set it to `false`, the app becomes unmanaged.

For manifest-based installs, if `true`, the system only considers apps installed in `/Applications` as managed. In macOS 11 through 13, the system requires that the `pkg` only contains a single signed app.

Available: macOS 11+

### iOSApp

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the app is an iOS app that can run on a Mac with Apple silicon in macOS 11 and later.

Available: macOS 11+

### iTunesStoreID

- **Type:** `integer`
- **Required:** No

The app’s iTunes Store identifier.

### ManagementFlags

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `4`, `5`

A bitwise OR of the management flags. The possible values are:

- `1`: If `InstallAsManaged` is `true`, remove the app upon removal of the MDM profile.
- `4`: Prevent backup of app data.
- `5`: Both `1` and `4`.

Available: iOS 5+ | iPadOS 5+ | macOS 11+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### ManifestURL

- **Type:** `string`
- **Required:** No

The URL of the app manifest, which needs to begin with `https:`. The manifest is returned as a property list that uses the [ManifestURL](/documentation/devicemanagement/manifesturl) format.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### Options

- **Type:** `InstallApplicationCommand.Command.Options`
- **Required:** No

A dictionary that contains the app installation options.

Available: iOS 7+ | iPadOS 7+ | macOS 10.9+ | tvOS 10.2+ | visionOS 1.1+ | watchOS 10+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallApplication`

The request type for this command.

## Topics

### Objects

- [InstallApplicationCommand.Command.Attributes](/documentation/devicemanagement/installapplicationcommand/command-data.dictionary/attributes-data.dictionary) - A dictionary that contains the initial attributes of the app.
- [InstallApplicationCommand.Command.Configuration](/documentation/devicemanagement/installapplicationcommand/command-data.dictionary/configuration-data.dictionary) - A dictionary that contains the configuration to install an enterprise app.
- [InstallApplicationCommand.Command.Options](/documentation/devicemanagement/installapplicationcommand/command-data.dictionary/options-data.dictionary) - A dictionary that contains the app installation options.

