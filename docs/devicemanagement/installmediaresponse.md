# InstallMediaResponse

A response from the device after it processes the command to install a book on a device.

**Platforms:** iOS 8.0, iPadOS 8.0, macOS 10.9

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

### ErrorChain

- **Type:** `[InstallMediaResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

### iTunesStoreID

- **Type:** `integer`
- **Required:** No

The book’s iTunes Store identifier, if present in the command.

### MediaType

- **Type:** `string`
- **Required:** No

The media type, which can only be `Book`.

### MediaURL

- **Type:** `string`
- **Required:** No

The URL to retrieve the book, if present in the command. This value is available in iOS 8 and later.

### PersistentID

- **Type:** `string`
- **Required:** No

The book’s persistent identifier, if present in the command. This value is available in iOS 8 and later.

### RejectionReason

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `CouldNotVerifyITunesStoreID`, `PurchaseNotFound`, `AppStoreDisabled`, `WrongMediaType`, `DownloadInvalid`, `EnterpriseBooksNotSupportedInMultiUser`

The reason, if installation fails, which is one of the following values:

- `CouldNotVerifyITunesStoreID`: The `iTunesStoreID` is invalid.
- `PurchaseNotFound`: The Volume Purchase Program (VPP) license isn’t in the user’s history.
- `AppStoreDisabled`: App Store isn’t available on the device.
- `WrongMediaType`: The media type is invalid. The only valid type is `Book`.
- `DownloadInvalid`: The URL doesn’t lead to a valid book.
- `EnterpriseBooksNotSupportedInMultiUser`: Multiuser mode doesn’t support enterprise books.

### State

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Queued`, `PromptingForLogin`, `Updating`, `Installing`, `Managed`, `ManagedButUninstalled`, `Installed`, `Uninstalled`, `Failed`, `Unknown`

The installation state of this book. The `Failed` and `Unknown` states are transient and the device only reports them once. Books from the Book Store report their state as `Installed` instead of `Managed`.

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

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

### UserLongName

- **Type:** `string`
- **Required:** Yes

## Topics

### Objects

- [InstallMediaResponse.ErrorChainItem](/documentation/devicemanagement/installmediaresponse/errorchainitem) - A dictionary that describes an error chain item.

