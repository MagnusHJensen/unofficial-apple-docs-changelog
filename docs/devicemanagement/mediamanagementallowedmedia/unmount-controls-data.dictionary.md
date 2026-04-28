# MediaManagementAllowedMedia.Unmount-controls

A dictionary to control volume unmounting.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### all-media

- **Type:** `string`
- **Required:** No

Unused; set to an empty string.

### bd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### blankbd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### blankcd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### blankdvd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### cd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### disk-image

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### dvd

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### dvdram

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### harddisk-external

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A string or an array of media action strings. Internally installed SD cards and USB flash drives are included in the hard disk-external category.

This key is the default for media types that don’t fall into other categories.

### harddisk-internal

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

### networkdisk

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `authenticate`, `read-only`, `deny`, `eject`

A media action string or an array of media action strings.

