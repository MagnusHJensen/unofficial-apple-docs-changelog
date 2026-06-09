# ContentCachingService.Ranges

A range of IP addresses to cache.

**Platforms:** macOS 10.13.4

## Properties

### first

- **Type:** `string`
- **Required:** Yes

The first IP address in the range.

Deprecated: macOS 27+

### last

- **Type:** `string`
- **Required:** Yes

The last IP address in the range.

Deprecated: macOS 27+

### type

- **Type:** `string`
- **Required:** No
- **Default:** `IPv4`
- **Allowed Values:** `IPv4`, `IPv6`

The IP address type.

Deprecated: macOS 27+

