# StatusContentCachePeersPeerItemObject

A peer Content Cache.

**Platforms:** macOS 27.0 (Beta)

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the peer entry and only this key and the `identifier` key are present in the status item object.

### address

- **Type:** `string`
- **Required:** Yes

The local IPv4 address of the peer Content Cache.

### friendly

- **Type:** `boolean`
- **Required:** Yes

If `true`, the peer Content Cache is willing to respond to requests from this Content Cache.

### healthy

- **Type:** `boolean`
- **Required:** Yes

If `true`, the peer Content Cache is able to respond to requests from this Content Cache.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier of the peer Content Cache.

### port

- **Type:** `integer`
- **Required:** Yes

The IP port number the peer Content Cache listens to for requests.

### version

- **Type:** `string`
- **Required:** Yes

The version number of the peer Content Cache software.

