# InstallProvisioningProfileCommand.Command

The command to install a provisioning profile on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### ProvisioningProfile

- **Type:** `data`
- **Required:** Yes

The provisioning profile.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallProvisioningProfile`

The request type for this command.

