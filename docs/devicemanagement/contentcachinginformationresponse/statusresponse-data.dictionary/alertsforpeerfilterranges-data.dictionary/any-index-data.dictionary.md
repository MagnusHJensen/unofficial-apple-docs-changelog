# ContentCachingInformationResponse.StatusResponse.AlertsForPeerFilterRanges.ANY index

A dictionary that describes the alerts for the peer filter ranges. The key name is the index into the `PeerFilterRanges` array in the installed `com.apple.AssetCache.managed` payload.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### addresses

- **Type:** `[string]`
- **Required:** Yes

An array of local IP addresses of peer content caches that rejected requests from the content cache.

### className

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AssetCacheUnfriendlyPeersInFilterRangeAlert`

The type of the alert.

### peerFilterRangeIndex

- **Type:** `integer`
- **Required:** Yes

The index into the `PeerFilterRanges` in the installed ContentCaching payload.

### postDate

- **Type:** `date`
- **Required:** Yes

The date of the alert.

