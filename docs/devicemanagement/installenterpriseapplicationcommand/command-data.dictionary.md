# InstallEnterpriseApplicationCommand.Command

The command to install an enterprise app on a device.

**Platforms:** macOS 10.13.6, Device Assignment Services , VPP License Management 

## Properties

### ChangeManagementState

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Managed`

The change management state. This value doesn’t work with the user enrollments. The only possible value is:

Available in macOS 11 and later.

### Configuration

- **Type:** `InstallEnterpriseApplicationCommand.Command.Configuration`
- **Required:** No

A dictionary that contains the initial configuration of the app, if you choose to provide it. Available in macOS 11 and later.

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

### ManagementFlags

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`

The management flags. The possible values are:

Available in macOS 11 and later.

### Manifest

- **Type:** `InstallEnterpriseApplicationCommand.Command.Manifest`
- **Required:** No

A dictionary that specifies where to download the app. This value uses the [ManifestURL](/documentation/devicemanagement/manifesturl) format.

### ManifestURL

- **Type:** `string`
- **Required:** No

The URL of the app manifest, which needs to begin with `https:`. The manifest is returned as a property list that uses the [ManifestURL](/documentation/devicemanagement/manifesturl) format.

### ManifestURLPinningCerts

- **Type:** `[data]`
- **Required:** No

An array of DER-encoded certificates to pin the connection when fetching the `ManifestURL`.

### PinningRevocationCheckRequired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, certificate revocation checks require a positive response when using certificate pinning with `ManifestURLPinningCerts`.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallEnterpriseApplication`

The request type for this command.

## Topics

### Objects

- [InstallEnterpriseApplicationCommand.Command.Configuration](/documentation/devicemanagement/installenterpriseapplicationcommand/command-data.dictionary/configuration-data.dictionary) - A dictionary that contains the configuration to install an enterprise app.
- [InstallEnterpriseApplicationCommand.Command.Manifest](/documentation/devicemanagement/installenterpriseapplicationcommand/command-data.dictionary/manifest-data.dictionary) - A dictionary that contains a manifest.

