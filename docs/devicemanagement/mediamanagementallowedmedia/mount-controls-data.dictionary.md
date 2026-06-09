# MediaManagementAllowedMedia.Mount-controls

A dictionary of volumes to control volume mounting.

**Platforms:** macOS 10.7

## Properties

### all-media

- **Type:** `string`
- **Required:** No

Unused; set to an empty string.

Deprecated: macOS 11+

### bd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### blankbd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### blankcd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### blankdvd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### cd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### disk-image

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### dvd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### dvdram

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### harddisk-external

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A string or an array of media action strings. The hard disk-external category includes internally installed SD cards and USB flash drives.

This key is the default for media types that don’t fall into other categories.

Deprecated: macOS 11+

### harddisk-internal

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

### networkdisk

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

Deprecated: macOS 11+

