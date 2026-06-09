# DeclarativeManagementRequest

The declarative management request details.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### Data

- **Type:** `data`
- **Required:** No

A Base64-encoded JSON object using the [SynchronizationTokens](/documentation/devicemanagement/synchronizationtokens) schema.

### Endpoint

- **Type:** `string`
- **Required:** Yes

The type of operation the declaration is requesting. This key needs to be one of these values:

- `tokens`: For fetching synchronization tokens from the server
- `declaration-items`: For fetching the declaration manifest from the server
- `status`: For sending a status report to the server
- `declaration/…/…`: For fetching a specific declaration from the server. Include the declaration type and identifier separated by slash characters (`/`).

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available: iOS 15+ | iPadOS 15+ | macOS 13+ | visionOS 1.1+

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.

Available: macOS 13+

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeclarativeManagement`

The message type, which requires a value of `DeclarativeManagement`.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UserID

- **Type:** `string`
- **Required:** No

For macOS, this value is the ID of the user.

For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn’t occur.

Available: iOS 15+ | iPadOS 15+ | macOS 13+

### UserLongName

- **Type:** `string`
- **Required:** Yes

The full name of the user.

Available: macOS 13+

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

Available: iOS 15+ | iPadOS 15+ | macOS 13+

## Topics

### Declaration Endpoints

- [declaration/activation/{identifier}](/documentation/devicemanagement/declaration-activation-_identifier_) - The endpoint for fetching an activation declaration.
- [declaration/asset/{identifier}](/documentation/devicemanagement/declaration-asset-_identifier_) - The endpoint for fetching an asset declaration.
- [declaration/configuration/{identifier}](/documentation/devicemanagement/declaration-configuration-_identifier_) - The endpoint for fetching a configuration declaration.
- [declaration/management/{identifier}](/documentation/devicemanagement/declaration-management-_identifier_) - The endpoint for fetching a management declaration.

### Declaration Response

- [DeclarationResponse](/documentation/devicemanagement/declarationresponse)

