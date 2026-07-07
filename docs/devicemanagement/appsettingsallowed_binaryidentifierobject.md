# AppSettingsAllowed_BinaryIdentifierObject

Dictionary containing one or more identifier fields to match a binary.

**Platforms:** macOS 27.0 (Beta)

## Properties

### CDHash

- **Type:** `string`
- **Required:** No

The code signature code directory hash of the binary.

### PathPrefix

- **Type:** `string`
- **Required:** No

The file system path prefix to match binaries.

### SigningID

- **Type:** `string`
- **Required:** No

The code signature signing identifier of the binary.

### SigningState

- **Type:** `string`
- **Required:** No
- **Default:** `All`
- **Allowed Values:** `All`, `TestFlight`, `DeveloperID`, `Enterprise`, `AppStore`, `Apple`

The code signing state to match binaries.

### TeamID

- **Type:** `string`
- **Required:** No

The code signature team identifier of the binary. Use the value “*APPLE*” instead of an empty string for Apple binaries with an empty team identifier.

