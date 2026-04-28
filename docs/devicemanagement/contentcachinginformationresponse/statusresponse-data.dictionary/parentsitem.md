# ContentCachingInformationResponse.StatusResponse.ParentsItem

A dictionary that describes a parent content cache.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### address

- **Type:** `string`
- **Required:** Yes

The local IP address of the parent content cache.

### alert

- **Type:** `ContentCachingInformationResponse.StatusResponse.ParentsItem.Alert`
- **Required:** No

A dictionary that describes an alert related to the parent content cache.

### details

- **Type:** `ContentCachingInformationResponse.StatusResponse.ParentsItem.Details`
- **Required:** Yes

A dictionary that contains additional details about the parent content cache.

### guid

- **Type:** `string`
- **Required:** Yes

The unique identifier of the parent content cache.

### healthy

- **Type:** `boolean`
- **Required:** Yes

If `true,` the parent content cache is able to respond to requests from this content cache.

### port

- **Type:** `integer`
- **Required:** Yes

The IP port number the parent content cache listens to for requests.

### version

- **Type:** `string`
- **Required:** Yes

The version number of the parent content cache software.

## Topics

### Objects

- [ContentCachingInformationResponse.StatusResponse.ParentsItem.Alert](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/parentsitem/alert-data.dictionary) - A dictionary that describes a parent content cache alert.
- [ContentCachingInformationResponse.StatusResponse.ParentsItem.Details](/documentation/devicemanagement/contentcachinginformationresponse/statusresponse-data.dictionary/parentsitem/details-data.dictionary) - A dictionary that contains additional details about the parent content cache.

