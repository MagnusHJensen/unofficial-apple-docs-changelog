# Dock.StaticItem.Tile-data

The dictionary that contains details about a Dock item.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### file-data

- **Type:** `Dock.StaticItem.Tile-data.File-data`
- **Required:** No

The data in a file. For Apple use only.

### file-type

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `0`, `1`, `3`

The type of tile:

- `0`: URL
- `1`: File
- `3`: Directory

### label

- **Type:** `string`
- **Required:** Yes

The label of the Dock item.

### url

- **Type:** `string`
- **Required:** No

The URL string.

## Topics

### Objects

- [Dock.StaticItem.Tile-data.File-data](/documentation/devicemanagement/dock/staticitem/tile-data-data.dictionary/file-data-data.dictionary) - For Apple use only.

