# ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Capabilities

The capabilities of the peer content cache.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### im

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of imports and uploads.

### ns

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of handling namespaces, which is an aspect of personal caching.

### pc

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of caching personal iCloud content.

### query-parameters

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of handling query parameters in URLs.

### sc

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of caching shared non-iCloud content.

### ur

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache is capable of prioritizing imports and uploads.

