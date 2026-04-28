# InstallMediaCommand.Command

The command to install a book on a device.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### Author

- **Type:** `string`
- **Required:** No

The name of the book’s author. This value is available in iOS 8 and later.

### iTunesStoreID

- **Type:** `integer`
- **Required:** No

The book’s iTunes Store identifier.

### Kind

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `pdf`, `epub`, `ibooks`

The kind of the media, which can be one of the following values:

- `pdf`: A PDF file
- `epub`: An EPUB file in `gzip` format.
- `ibooks`: An iBooks Author file in `gzip` format.

If you omit this value, its value is the file extension in the URL. This value is available in iOS 8 and later.

### MediaType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Book`

The media type, which can only be `Book`.

### MediaURL

- **Type:** `string`
- **Required:** No

The URL to retrieve the book. This value is available in iOS 8 and later.

### PersistentID

- **Type:** `string`
- **Required:** No

The book’s persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`. This value is available in iOS 8 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `InstallMedia`

The request type for this command.

### Title

- **Type:** `string`
- **Required:** No

The book’s title. This value is available in iOS 8 and later.

### Version

- **Type:** `string`
- **Required:** No

The book’s version number. This value is available in iOS 8 and later.

