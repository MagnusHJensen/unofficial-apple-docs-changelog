# StatusContentCacheInfoContentCacheInfoObject

A dictionary that contains info about the usage of the Content Cache on the device

**Platforms:** macOS 27.0 (Beta)

## Properties

### cache-details

- **Type:** `[string]`
- **Required:** No

The amount of disk space that various categories of cached content use. Apple defines these categories and they’re subject to change.

### cache-free

- **Type:** `integer`
- **Required:** No

The amount of disk space in bytes, available to the Content Cache.

### cache-limit

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space in bytes, available to the Content Cache. A value of `0` indicates an unlimited amount. This value corresponds to `CacheLimit` in the installed [ContentCaching](/documentation/devicemanagement/contentcaching) configuration.

### cache-status

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LOWSPACE`, `OK`

The level of cache pressure. `LOWSPACE` means cache pressure is high.

### cache-used

- **Type:** `integer`
- **Required:** No

The amount of disk space in bytes, cached content uses. The Content Cache allocates space in its cache for entire files even when it stores only part of those files in its cache.

### max-cache-pressure-last-hour

- **Type:** `number`
- **Required:** No

A floating-point number between `0.0` and `1.0` that represents how often the cache needed more disk space over the last hour of operation. A lower value is better.

### personal-cache-free

- **Type:** `integer`
- **Required:** No

The amount of disk space in bytes, available to the Content Cache for personal iCloud content.

### personal-cache-limit

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space in bytes, available to the Content Cache for personal iCloud content. A value of `0` indicates an unlimited amount.

### personal-cache-used

- **Type:** `integer`
- **Required:** No

The amount of disk space, in bytes, available to the Content Cache for personal iCloud content.

