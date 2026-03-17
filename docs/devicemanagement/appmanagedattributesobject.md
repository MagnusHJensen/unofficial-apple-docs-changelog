# AppManagedAttributesObject

A dictionary of values associated with an app.

**Platforms:** iOS 17.2, iPadOS 17.2, visionOS 2.4

## Properties

### AssociatedDomains

- **Type:** `[string]`
- **Required:** No

An array of domain names to associate with the app.

### AssociatedDomainsEnableDirectDownloads

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables direct downloads for the `AssociatedDomains`.

### CellularSliceUUID

- **Type:** `string`
- **Required:** No

The cellular slice identifier, which can be the data network name (DNN) or app category. For DNN, encode the value as “DNN:name”, where “name” is the carrier-provided DNN name. For app category, encode the value as “AppCategory:category”, where “category” is a carrier-provided string such as “Enterprise1”.

### ContentFilterUUID

- **Type:** `string`
- **Required:** No

The UUID of the content filter to associate with the app.

### DNSProxyUUID

- **Type:** `string`
- **Required:** No

The UUID of the DNS proxy to associate with the app.

### Hideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system prevents the user from hiding the app. It doesn’t affect the user’s ability to leave it in the App Library, while removing it from the Home Screen.

### Lockable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.

### RelayUUID

- **Type:** `string`
- **Required:** No

The UUID of the relay to associate with the app.

### TapToPayScreenLock

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device automatically locks after every transaction that requires a customer’s card PIN. If `false`, the user can choose the behavior.

### VPNUUID

- **Type:** `string`
- **Required:** No

The UUID of the VPN to associate with the app.

