# ContentCachingInformationResponse.StatusResponse.ParentsItem.Alert

A dictionary that describes a parent content cache alert.

**Platforms:** macOS 10.15.4

## Properties

### addresses

- **Type:** `[string]`
- **Required:** Yes

An array of local IP addresses of parent content caches.

### className

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AssetCacheParentCycleAlert`, `AssetCacheParentDepthAlert`

The type of the alert.

### postDate

- **Type:** `date`
- **Required:** Yes

The date of the alert.

