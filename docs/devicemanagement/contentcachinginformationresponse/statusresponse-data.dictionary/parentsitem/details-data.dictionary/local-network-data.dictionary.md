# ContentCachingInformationResponse.StatusResponse.ParentsItem.Details.Local-network

A dictionary that describes the parent content cache’s connection to its local network.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### speed

- **Type:** `integer`
- **Required:** No

The transfer speed, in megabits per second, of the parent content cache’s connection to its local network.

### wired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the parent content cache has a wired connection to its local network. If `false`, it has a wireless connection; for example, Wi-Fi.

