# ContentCachingInformationResponse.StatusResponse.PeersItem

A dictionary that describes a peer content cache.

**Platforms:** macOS 10.15.4

## Properties

### address

- **Type:** `string`
- **Required:** Yes

The local IP address of the peer content cache.

### alert

- **Type:** `ContentCachingInformationResponse.StatusResponse.PeersItem.Alert`
- **Required:** No

A dictionary that describes an alert related to the peer content cache.

### details

- **Type:** `ContentCachingInformationResponse.StatusResponse.PeersItem.Details`
- **Required:** Yes

A dictionary that contains additional details about the peer content cache.

### friendly

- **Type:** `boolean`
- **Required:** Yes

If `true`, the peer content cache is able to respond to requests from the content cache.

### guid

- **Type:** `string`
- **Required:** Yes

The unique identifier of the peer content cache.

### healthy

- **Type:** `boolean`
- **Required:** Yes

If `true`, the peer content cache is able to respond to requests from the content cache.

### port

- **Type:** `integer`
- **Required:** Yes

The IP port number the peer content cache listens to for requests.

### version

- **Type:** `string`
- **Required:** Yes

The version number of the peer content cache software.

## Topics

### Objects

- [ContentCachingInformationResponse.StatusResponse.PeersItem.Alert](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/peersitem/alert-data.dictionary) - A dictionary that describes a peer content cache alert.
- [ContentCachingInformationResponse.StatusResponse.PeersItem.Details](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/peersitem/details-data.dictionary) - A dictionary that contains additional details about the peer content cache.

