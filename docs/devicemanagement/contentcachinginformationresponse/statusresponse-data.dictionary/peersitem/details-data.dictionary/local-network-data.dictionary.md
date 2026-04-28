# ContentCachingInformationResponse.StatusResponse.PeersItem.Details.Local-network

The network details about the peer cache.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### speed

- **Type:** `integer`
- **Required:** No

The transfer speed, in megabits per second, of the peer content cache’s connection to its local network.

### wired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the peer content cache has a wired connection to its local network. If `false`, it has a wireless connection; for example, Wi-Fi.

