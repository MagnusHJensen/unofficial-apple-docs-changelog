# ContentCachingInformationResponse.StatusResponse.AlertsItem

A dictionary that describes an alert from the content cache.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### cacheLimit

- **Type:** `integer`
- **Required:** No

The limit, in bytes, for the content cache at the time of the alert. This value only applies to `AssetCacheLowSpaceAlert` and `AssetCacheNoSpaceAlert` types.

### className

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AssetCacheLowSpaceAlert`, `AssetCacheNoSpaceAlert`, `AssetCacheRegistrationRejectedAlert`, `AssetCacheRegistrationUnavailableAlert`, `AssetCacheResourceMissingAlert`

The type of the alert.

### pathPreventingAccess

- **Type:** `string`
- **Required:** No

The subpath of the resource that was missing or inaccessible at the time of the alert. This value only applies to the `AssetCacheResourceMissingAlert` type.

### postDate

- **Type:** `date`
- **Required:** Yes

The date of the alert.

### reservedVolumeSpace

- **Type:** `integer`
- **Required:** No

The space, in bytes, that the system reserves at the time of the alert. This value only applies to the `AssetCacheLowSpaceAlert` and `AssetCacheNoSpaceAlert` types.

### resource

- **Type:** `string`
- **Required:** No

The resource that was missing or inaccessible at the time of the alert. This value only applies to the `AssetCacheResourceMissingAlert` type.

