# ContentCachingInformationResponse.StatusResponse.ParentsItem.Details

A dictionary that contains additional details about the parent content cache.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### ac-power

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the parent content cache power source is AC; otherwise, an internal battery provides its power.

### cache-size

- **Type:** `integer`
- **Required:** No

The maximum amount of disk space, in bytes, available to the parent content cache.

### capabilities

- **Type:** `ContentCachingInformationResponse.StatusResponse.ParentsItem.Details.Capabilities`
- **Required:** No

A dictionary that describes the capabilities of the parent content cache.

### is-portable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the parent content cache computer is portable; for example, a laptop.

### local-network

- **Type:** `ContentCachingInformationResponse.StatusResponse.ParentsItem.Details.Local-network`
- **Required:** No

A dictionary that describes the parent content cache’s connection to its local network.

## Topics

### Objects

- [ContentCachingInformationResponse.StatusResponse.ParentsItem.Details.Capabilities](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/parentsitem/details-data.dictionary/capabilities-data.dictionary) - A dictionary that describes the capabilities of the parent content cache.
- [ContentCachingInformationResponse.StatusResponse.ParentsItem.Details.Local-network](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/parentsitem/details-data.dictionary/local-network-data.dictionary) - A dictionary that describes the parent content cache’s connection to its local network.

