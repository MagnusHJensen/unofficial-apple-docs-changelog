# ContentCachingInformationResponse.StatusResponse

A dictionary that contains the status of content caching on a device.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### Activated

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device has enabled content caching. Enabling content caching doesn’t guarantee service. See the `Active` key for the readiness of content caching to serve requests.

### Active

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, content caching is ready to serve requests.

### ActualCacheUsed

- **Type:** `integer`
- **Required:** No

The actual amount of disk space, in bytes, that cached content uses. See related values `CacheUsed` and `PersonalCacheUsed`.

### Alerts

- **Type:** `[ContentCachingInformationResponse.StatusResponse.AlertsItem]`
- **Required:** No

An array that contains the error conditions the content cache detected that aren’t related to peer filter ranges, parent content caches, or peer content caches.

See `AlertsForPeerFilterRanges` for errors related to peer filter ranges.

See `Parents` and `Peers` for errors related to parent and peer content caches.

To display these alerts on the device, set `DisplayAlerts` to `true` in the installed [ContentCaching](/documentation/devicemanagement/contentcaching) profile.

### AlertsForPeerFilterRanges

- **Type:** `ContentCachingInformationResponse.StatusResponse.AlertsForPeerFilterRanges`
- **Required:** No

The error conditions the content cache detected in the `PeerFilterRanges` in the installed `com.apple.AssetCache.managed` payload.

To display these alerts on the device, set `DisplayAlerts` to `true` in the installed [ContentCaching](/documentation/devicemanagement/contentcaching) profile.

### CacheDetails

- **Type:** `ContentCachingInformationResponse.StatusResponse.CacheDetails`
- **Required:** No

The amount of disk space that various categories of cached content use. Apple defines these categories and they’re subject to change.

### CacheFree

- **Type:** `integer`
- **Required:** No

The amount of disk space, in bytes, available to the content cache.

### CacheLimit

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space, in bytes, available to the content cache. A value of `0` indicates an unlimited amount. This value corresponds to `CacheLimit` in the installed [ContentCaching](/documentation/devicemanagement/contentcaching) profile.

### CacheStatus

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LOWSPACE`, `OK`

The level of cache pressure. `LowSpace` means cache pressure is high.

### CacheUsed

- **Type:** `integer`
- **Required:** No

The amount of disk space, in bytes, cached content uses. Content caching allocates space in its cache for entire files even when it stores only part of those files in its cache.

### DataMigrationCompleted

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the content cache finished moving from one volume to another.

### DataMigrationError

- **Type:** `ContentCachingInformationResponse.StatusResponse.DataMigrationError`
- **Required:** No

The error that occurred while the content cache moved from one volume to another.

### DataMigrationProgress

- **Type:** `number`
- **Required:** No

A floating-point number between `0.0` and `1.0` that indicates the percentage of progress in moving the content cache from one volume to another. A value of `1.0` indicates that the content cache has fully migrated.

### MaxCachePressureLast1Hour

- **Type:** `number`
- **Required:** No

A floating-point number between `0.0` and `1.0` that represents how often the cache needed more disk space over the last hour of operation. A lower value is better.

### Parents

- **Type:** `[ContentCachingInformationResponse.StatusResponse.ParentsItem]`
- **Required:** No

An array of dictionaries that describes parent content caches.

### Peers

- **Type:** `[ContentCachingInformationResponse.StatusResponse.PeersItem]`
- **Required:** No

An array of dictionaries that describes peer content caches.

### PersonalCacheFree

- **Type:** `integer`
- **Required:** No

The amount of disk space, in bytes, available to the content cache for personal iCloud content.

### PersonalCacheLimit

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space, in bytes, available to the content cache for personal iCloud content. A value of `0` indicates an unlimited amount.

### PersonalCacheUsed

- **Type:** `integer`
- **Required:** No

The amount of disk space, in bytes, available to the content cache for personal iCloud content.

### Port

- **Type:** `integer`
- **Required:** No

The IP port number the content cache listens to for requests from clients, peers, and children.

### PrivateAddresses

- **Type:** `[string]`
- **Required:** No

An array of the content cache’s local IP addresses.

### PublicAddress

- **Type:** `string`
- **Required:** No

The public IP address of the content cache.

### RegistrationError

- **Type:** `string`
- **Required:** No

If present, the reason the content cache failed to register itself with Apple.

### RegistrationResponseCode

- **Type:** `integer`
- **Required:** No

If present, the HTTP response code the content cache received when it failed to register itself with Apple.

### RegistrationStarted

- **Type:** `date`
- **Required:** No

The date when the content cache began registering itself with Apple. This value is only available during registration attempts.

### RegistrationStatus

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `-1`, `0`, `1`

The status of the content cache’s registration with Apple, which is one of the following values:

- `-1:` Failed
- `0:` Pending
- `1:` Succeeded

### RestrictedMedia

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, a restriction prevents caching of certain content types.

### ServerGUID

- **Type:** `string`
- **Required:** No

The unique identifier of the content cache.

### StartupStatus

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `FAILED`, `MIGRATING_DATA`, `OK`, `PENDING`

The status of the content cache’s registration with Apple.

### TetheratorStatus

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `-1`, `0`, `1`

The status of tethered caching, which is content caching with a shared internet connection, which is one of the following values:

- `-1:` Unknown
- `0:` Disabled
- `1:` Enabled

### TotalBytesAreSince

- **Type:** `date`
- **Required:** No

The start date to use when collecting data for the other `TotalBytes` values.

### TotalBytesDropped

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache downloaded, but couldn’t add to its cache, since the `TotalBytesAreSince` date.

### TotalBytesImported

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache received since the `TotalBytesAreSince` date.

### TotalBytesReturnedToChildren

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache served to its child content cache since the `TotalBytesAreSince` date.

### TotalBytesReturnedToClients

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache served to client iOS, macOS, and tvOS devices since the `TotalBytesAreSince` date.

### TotalBytesReturnedToPeers

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache served to peer content caches since the `TotalBytesAreSince` date.

### TotalBytesStoredFromOrigin

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache saved from the internet since the `TotalBytesAreSince` date.

### TotalBytesStoredFromParents

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache saved from parent content caches since the `TotalBytesAreSince` date.

### TotalBytesStoredFromPeers

- **Type:** `integer`
- **Required:** No

The amount of data, in bytes, that the content cache saved from peer content caches since the `TotalBytesAreSince` date.

## Topics

### Objects

- [ContentCachingInformationResponse.StatusResponse.AlertsForPeerFilterRanges](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/alertsforpeerfilterranges-data.dictionary) - A dictionary that contains alerts for peer filter ranges.
- [ContentCachingInformationResponse.StatusResponse.AlertsItem](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/alertsitem) - A dictionary that describes an alert from the content cache.
- [ContentCachingInformationResponse.StatusResponse.CacheDetails](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/cachedetails-data.dictionary) - A dictionary that describes disk space the content cache uses.
- [ContentCachingInformationResponse.StatusResponse.DataMigrationError](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/datamigrationerror-data.dictionary) - A dictionary that describes a data migration error.
- [ContentCachingInformationResponse.StatusResponse.ParentsItem](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/parentsitem) - A dictionary that describes a parent content cache.
- [ContentCachingInformationResponse.StatusResponse.PeersItem](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/peersitem) - A dictionary that describes a peer content cache.

