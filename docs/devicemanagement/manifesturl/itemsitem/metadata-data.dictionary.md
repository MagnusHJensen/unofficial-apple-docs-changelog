# ManifestURL.ItemsItem.Metadata

The metadata for an application or package manifest item.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### bundle-identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app or package manifest item.

### bundle-version

- **Type:** `string`
- **Required:** No

The bundle version of the app or package manifest item.

### kind

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `software`

The type of manifest item. For an app or package, this needs to be `software`.

### subtitle

- **Type:** `string`
- **Required:** No

The name of the app or package developer.

### title

- **Type:** `string`
- **Required:** Yes

The title of the app or package being installed.

