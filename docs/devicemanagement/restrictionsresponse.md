# RestrictionsResponse

A response from the device after it processes the command to get a list of restrictions on the device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### CommandUUID

- **Type:** `string`
- **Required:** No

The unique identifier of the command for this response.

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.

Available in macOS 10.15 and later.

### ErrorChain

- **Type:** `[RestrictionsResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

### GlobalRestrictions

- **Type:** `RestrictionsResponse.GlobalRestrictions`
- **Required:** Yes

A dictionary that contains the global restrictions in effect. This value is available in iOS 4 and later, and tvOS 6.1 and later.

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device isn’t on-console.

### ProfileRestrictions

- **Type:** `RestrictionsResponse.ProfileRestrictions`
- **Required:** Yes

A dictionary that contains dictionaries of restrictions from each profile. This value is only available when `ProfileRestrictions` is `true` in the command. The keys are the identifiers of the profiles. This value is available in iOS 4 and later, and tvOS 6.1 and later.

### Status

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Acknowledged`, `Error`, `CommandFormatError`, `Idle`, `NotNow`

The status of the response, which is one of the following values:

- `Acknowledged`: The device processed the command successfully.
- `Error`: An error occurred. See the `ErrorChain` for more details.
- `CommandFormatError`: A protocol error occurred, which can result from a malformed command.
- `Idle`: The device is idle; there’s no status.
- `NotNow`: The device received the command, but can’t run it.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UserID

- **Type:** `string`
- **Required:** No

For macOS, this value is the ID of the user.

For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn’t occur.

### UserLongName

- **Type:** `string`
- **Required:** Yes

The full name of the user.

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

## Topics

### Objects

- [RestrictionsResponse.ErrorChainItem](/documentation/devicemanagement/restrictionsresponse/errorchainitem) - A dictionary that describes an error chain item.
- [RestrictionsResponse.GlobalRestrictions](/documentation/devicemanagement/restrictionsresponse/globalrestrictions-data.dictionary) - A dictionary that contains the global restrictions in effect.
- [RestrictionsResponse.ProfileRestrictions](/documentation/devicemanagement/restrictionsresponse/profilerestrictions-data.dictionary) - A dictionary that contains restrictions from each profile.

