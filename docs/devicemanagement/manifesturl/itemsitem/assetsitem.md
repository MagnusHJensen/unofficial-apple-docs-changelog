# ManifestURL.ItemsItem.AssetsItem

An array of dictionaries that describe an item to install.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### kind

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `asset-pack-manifest`, `display-image`, `full-size-image`, `software-package`

The kind of manifest item to install. Use `software-package` for apps and macOS packages.

### md5

- **Type:** `string`
- **Required:** No

The MD5 hash value the device uses when verifying the hash of the manifest item data. When this key is present, the device ignores the `md5-size` and `md5s` keys.

### md5-size

- **Type:** `integer`
- **Required:** No

The data ** size the device uses when verifying the hash of the manifest item data. Required when the `md5s` key is present.

### md5s

- **Type:** `[string]`
- **Required:** No

An array of strings representing a set of MD5 hash values. The device uses these values to verify the integrity of the downloaded manifest item data. Required when the `md5-size` key is present.

### sha256

- **Type:** `string`
- **Required:** No

The SHA-256 hash value the device uses when verifying the hash of the manifest item data. When this key is present, the device ignores the `sha256-size` and `sha256` keys.

### sha256-size

- **Type:** `integer`
- **Required:** No

The data ** size the device uses when verifying the hash of the manifest item data. Required when the `sha256s` key is present.

### sha256s

- **Type:** `[string]`
- **Required:** No

An array of strings representing a set of SHA-256 hash values. The device uses these values to verify the integrity of the downloaded manifest item data. Required when the `sha256-size` key is present.

### url

- **Type:** `string`
- **Required:** Yes

The URL that hosts the manifest item data. The URL needs to start with `https://`.

