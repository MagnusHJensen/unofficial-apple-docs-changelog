# DeclarativeManagementCommand.Command

The command to enable your server to support declarative management or trigger a declarative management synchronization operation on the device.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### Data

- **Type:** `data`
- **Required:** No

The base64-encoded declarative management JSON request using a [TokensResponse](/documentation/devicemanagement/tokensresponse).

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeclarativeManagement`

The request type for this command.

