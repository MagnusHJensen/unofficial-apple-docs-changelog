# ManagedMediaListResponse.BooksItem

A dictionary that describes a managed book.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, Device Assignment Services , VPP License Management 

## Properties

### Author

- **Type:** `string`
- **Required:** No

The name of the book’s author.

### iTunesStoreID

- **Type:** `integer`
- **Required:** Yes

The book’s iTunes Store identifier.

### Kind

- **Type:** `string`
- **Required:** No

The kind of the media, which is one of the following values:

- `pdf`: A PDF file
- `epub`: An EPUB file in `gzip` format
- `ibooks`: An iBooks Author file in `gzip` format
- The file extension in the URL

This value is available in iOS 8 and later.

### PersistentID

- **Type:** `string`
- **Required:** No

The book’s persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`.

### State

- **Type:** `string`
- **Required:** No

The installation state of this book, which is one of the following values:

- `Queued`
- `PromptingForLogin`
- `Updating`
- `Installing`
- `Managed`
- `ManagedButUninstalled`
- `Installed`
- `Uninstalled`
- `Failed`
- `Unknown`

The `Failed` and `Unknown` states are transient and the device only reports them once. Books from the Book Store report their state as `Installed` instead of `Managed`.

### Title

- **Type:** `string`
- **Required:** No

The book’s title.

### Version

- **Type:** `string`
- **Required:** No

The book’s version number.

