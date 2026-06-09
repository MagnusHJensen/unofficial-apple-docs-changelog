# InstallMediaCommand.Command

The command to install a book on a device.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, macOS 10.9

## Properties

### Author

- **Type:** `string`
- **Required:** No

The name of the book’s author.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

### iTunesStoreID

- **Type:** `integer`
- **Required:** No

The book’s iTunes Store identifier.

Deprecated: macOS 11+
Removed: macOS 11+

### Kind

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `pdf`, `epub`, `ibooks`

The kind of the media, which can be one of the following values:

- `pdf`: A PDF file
- `epub`: An EPUB file in `gzip` format.
- `ibooks`: An iBooks Author file in `gzip` format.

If you omit this value, its value is the file extension in the URL.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

### MediaType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Book`

The media type, which can only be `Book`.

Deprecated: macOS 11+
Removed: macOS 11+

### MediaURL

- **Type:** `string`
- **Required:** No

The URL to retrieve the book.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

### PersistentID

- **Type:** `string`
- **Required:** No

The book’s persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

Deprecated: macOS 11+
Removed: macOS 11+

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallMedia`

The request type for this command.

Deprecated: macOS 11+
Removed: macOS 11+

### Title

- **Type:** `string`
- **Required:** No

The book’s title.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

### Version

- **Type:** `string`
- **Required:** No

The book’s version number.

Available: iOS 8+ | iPadOS 8+
Deprecated: macOS 11+
Removed: macOS 11+

