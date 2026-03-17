# SettingsCommand.Command.Settings.OrganizationInfo.OrganizationInfo

A dictionary that contains information about the organization operating the MDM server.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 10.9, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### OrganizationAddress

- **Type:** `string`
- **Required:** No

The organization’s address. Use the LF character (`&#10`) to insert line breaks.

### OrganizationEmail

- **Type:** `string`
- **Required:** No

The organization’s support email address.

### OrganizationMagic

- **Type:** `string`
- **Required:** No

A unique identifier for the various services a single organization manages.

### OrganizationName

- **Type:** `string`
- **Required:** Yes

A string that describes the organization operating the MDM server for display to the user during certain operations, such as purchasing or installing apps.

### OrganizationPhone

- **Type:** `string`
- **Required:** No

The organization’s phone number.

### OrganizationShortName

- **Type:** `string`
- **Required:** No

A shorter version of `OrganizationName`, preferably a single word or abbreviation, suitable for display to the user in places where a very short name is necessary.

