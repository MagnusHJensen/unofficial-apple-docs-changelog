# EducationConfiguration.UsersItem

A user in the organization.

**Platforms:** iOS 9.3, iPadOS 9.3, macOS 10.14

## Properties

### AppleID

- **Type:** `string`
- **Required:** No

The Managed Apple Account for this user.

Not required to configure Classroom, but if set the system uses it.

Required to configure the Shared iPad login screen.

### FamilyName

- **Type:** `string`
- **Required:** No

The family name of the user.

### FullScreenImageURL

- **Type:** `string`
- **Required:** No

Deprecated in iOS 9.3.1 and later. The URL pointing to an image of the user. The system uses the  `ResourcePayloadCertificateUUID` identity certificate or the MDM client identity to perform authentication when fetching the specified resource.

### GivenName

- **Type:** `string`
- **Required:** No

The given name of the user.

### Identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier for a user in the organization.

### ImageURL

- **Type:** `string`
- **Required:** No

A string that contains a URL pointing to an image of the user. The system displays this image in the iOS login screen and in the Classroom app. The recommended resolution is 256 x 256 pixels (512 x 512 pixels on a 2x device). The recommended formats are JPEG, PNG, and TIFF. The system uses the `ResourcePayloadCertificateUUID` identity certificate or the MDM client identity to perform authentication when fetching the image.

### Name

- **Type:** `string`
- **Required:** Yes

The name of the user.

### PasscodeType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `complex`, `four`, `six`

The type of passcode UI to show when the user is at the Login Window.

### PhoneticFamilyName

- **Type:** `string`
- **Required:** No

The user’s phonetic family name. The system uses this name to sort users in the Classroom app and the Shared iPad login screen.

### PhoneticGivenName

- **Type:** `string`
- **Required:** No

The user’s phonetic given name. The system uses this name to sort users in the Classroom app and the Shared iPad Login Screen.

