# InstallProfileCommand.Command

The command to install a configuration profile on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Payload

- **Type:** `data`
- **Required:** Yes

The profile to install, which you can encrypt using any identity certificate installed on the device. You can also sign the profile.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallProfile`

The request type for this command.

