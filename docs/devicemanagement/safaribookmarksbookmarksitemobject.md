# SafariBookmarksBookmarksItemObject

A bookmark that specifies a title, and either a URL for the bookmark, or a nested folder of bookmarks.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 26.0, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### Folder

- **Type:** `[SafariBookmarksBookmarksItemObject]`
- **Required:** No

An array of bookmarks for each bookmark in the folder. Folders can include bookmark items and bookmark folders.

Only one of `URL` or `Folder` must be present.

### Title

- **Type:** `string`
- **Required:** Yes

The title of the bookmark shown in Safari.

### URL

- **Type:** `string`
- **Required:** No

The URL for the bookmark item.

Only one of `URL` or `Folder` must be present.

