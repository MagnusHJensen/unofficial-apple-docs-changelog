# ContentCachingInformationResponse.StatusResponse.PeersItem.Alert

A dictionary that describes a peer content cache alert.

**Platforms:** macOS 10.15.4

## Properties

### addresses

- **Type:** `[string]`
- **Required:** No

An array of local IP addresses of peer content caches.

### className

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AssetCachePeerCycleAlert`, `AssetCacheUnfriendlyPeerAlert`

The type of the alert.

### peerAddress

- **Type:** `string`
- **Required:** No

The local IP address of a peer content cache.

### postDate

- **Type:** `date`
- **Required:** Yes

The date of the alert.

