# InstallApplicationCommand.Command

The command to install a third-party app on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Attributes

- **Type:** `InstallApplicationCommand.Command.Attributes`
- **Required:** No

A dictionary that contains the initial attributes of the app, if you choose to provide it. Available in iOS 7 and later, and tvOS 10.2 and later.

### ChangeManagementState

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Managed`

The change management state. This value doesn’t work with the user enrollment feature introduced in iOS 13, or any type of account driven enrollment. Available in iOS 9 and later, macOS 11 and later, and tvOS 10.2 and later. The only possible value is:

### Configuration

- **Type:** `InstallApplicationCommand.Command.Configuration`
- **Required:** No

A dictionary that contains the initial configuration of the app, if you choose to provide it. Available in iOS 7 and later, macOS 11 and later, and tvOS 10.2 and later.

### Identifier

- **Type:** `string`
- **Required:** No

The app’s bundle identifier.

> 

### InstallAsManaged

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, install the app as a managed app. Otherwise, the system installs the app as unmanaged. If you reinstall a manged app and omit this value or set it to `false`, the app becomes unmanaged.

For manifest-based installs, if `true`, the system only considers apps installed in `/Applications` as managed. In macOS 11 through 13, the system requires that the `pkg` only contains a single signed app.

Available in macOS 11 and later.

### iOSApp

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the app is an iOS app that can run on a Mac with Apple silicon in macOS 11 and later.

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

Available in iOS 5 and later, macOS 11 and later, and tvOS 10.2 and later.

### ManifestURL

- **Type:** `string`
- **Required:** No

The URL of the app manifest, which needs to begin with `https:`. The manifest is returned as a property list that uses the [ManifestURL](/documentation/devicemanagement/manifesturl) format.

### Options

- **Type:** `InstallApplicationCommand.Command.Options`
- **Required:** No

A dictionary that contains the app installation options.

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

