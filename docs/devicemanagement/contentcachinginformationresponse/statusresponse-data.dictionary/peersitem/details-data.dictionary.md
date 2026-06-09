# ContentCachingInformationResponse.StatusResponse.PeersItem.Details

A dictionary that contains additional details about the peer content cache.

**Platforms:** macOS 10.15.4

## Properties

### ac-power

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache power source is AC; otherwise, an internal battery provides its power.

### cache-size

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space, in bytes, available to the peer content cache.

### capabilities

- **Type:** `ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Capabilities`
- **Required:** No

A dictionary that describes the capabilities of the peer content cache.

### is-portable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache computer is portable; for example, a laptop.

### local-network

- **Type:** `ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Local-network`
- **Required:** No

A dictionary that describes the peer content cache’s connection to its local network.

## Topics

### Objects

- [ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Capabilities](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/peersitem/details-data.dictionary/capabilities-data.dictionary) - The capabilities of the peer content cache.
- [ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Local-network](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/peersitem/details-data.dictionary/local-network-data.dictionary) - The network details about the peer cache.

