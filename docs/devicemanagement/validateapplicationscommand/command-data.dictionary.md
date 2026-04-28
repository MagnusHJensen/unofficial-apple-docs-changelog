# ValidateApplicationsCommand.Command

The command to force validation of developer and universal provisioning profiles for enterprise apps.

**Platforms:** iOS 9.2, iPadOS 9.2, Mac Catalyst 9.2, tvOS 10.2, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### Identifiers

- **Type:** `[string]`
- **Required:** No

The bundle identifiers of the enterprise apps to include for validation of associated provisioning profiles, if you choose to provide them. Otherwise, validation occurs for the provisioning profiles for the installed managed apps.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ValidateApplications`

The request type for this command.

