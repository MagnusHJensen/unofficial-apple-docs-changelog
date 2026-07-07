# StatusContentCacheStatusContentCacheStatusObject

The basic set of AssetCache status items

**Platforms:** macOS 27.0 (Beta)

## Properties

### activated

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device has enabled the Content Cache. Enabling the Content Caching doesn’t guarantee service. See the `Active` key for the readiness of the Content Cache to serve requests.

### active

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the Content Cache is ready to serve requests.

### cache-status

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LOWSPACE`, `OK`

The level of cache pressure. `LOWSPACE` means cache pressure is high.

### port

- **Type:** `integer`
- **Required:** No

The IP port number the Content Cache listens to for requests from clients, peers, and children.

### private-addresses

- **Type:** `[string]`
- **Required:** No

An array of the Content Cache’s local IP addresses.

### public-address

- **Type:** `string`
- **Required:** No

The public IP address of the Content Cache.

### registration-error

- **Type:** `string`
- **Required:** No

If present, the reason the Content Cache failed to register itself with Apple.

### registration-response-code

- **Type:** `integer`
- **Required:** No

If present, the HTTP response code that the Content Cache received when it failed to register itself with Apple.

### registration-started

- **Type:** `string`
- **Required:** No

The RFC 3339 timestamp for when the Content Cache began registering itself with Apple. This value is only available during registration attempts.

### registration-status

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `-1`, `0`, `1`

The status of the Content Cache’s registration with Apple, which is one of the following values:

- `-1`: The registration failed.
- `0`: The registration is pending.
- `1`: The registration succeeded.

### report-error

- **Type:** `string`
- **Required:** No

When present, indicates why the Content Cache failed to send the metrics.

### report-response-code

- **Type:** `integer`
- **Required:** No

When present, contains the HTTP response code that the Content Cache received when it failed to send the metrics report to the target URL.

### sending-reports

- **Type:** `boolean`
- **Required:** No

When present, a value of `true` indicates that the cache is sending metrics reports to the URL specified in the `ManagementStatusTarget` key in the installed [ContentCaching](/documentation/devicemanagement/contentcaching) configuration.

### server-guid

- **Type:** `string`
- **Required:** Yes

The unique identifier of the Content Cache.

### startup-status

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `FAILED`, `MIGRATING_DATA`, `OK`, `PENDING`

The status of the Content Cache’s registration with Apple.

### tetherator-status

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `-1`, `0`, `1`

The status of tethered caching, which is the Content Cache with a shared internet connection, which is one of the following values:

- ‘-1’ : Unknown
- ‘0’ : Disabled
- ‘1’ : Enabled

### version

- **Type:** `string`
- **Required:** Yes

The version number of the Content Cache software.

