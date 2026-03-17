# InstallApplicationCommand.Command.Attributes

A dictionary that contains the initial attributes of the app.

**Platforms:** iOS 7.0, iPadOS 7.0, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### AssociatedDomains

- **Type:** `[string]`
- **Required:** No

An array that contains the associated domains to add to this app. Available in iOS 13 and later.

### AssociatedDomainsEnableDirectDownloads

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, perform claimed site association verification directly at the domain instead of on Apple’s servers. Only set this to `true` for domains that can’t access the internet. Available in iOS 14 and later.

### CellularSliceUUID

- **Type:** `string`
- **Required:** No

The data network name (DNN) or app category. For DNN, the value is `DNN:name`, where `name` is the carrier-provided DNN name. For app category, the value is `AppCategory:category`, where `category` is a carrier-provided string like “Enterprise1”.

Available in iOS 17 and later.

### ContentFilterUUID

- **Type:** `string`
- **Required:** No

The content filter UUID for this app. Available in iOS 16 and later.

### DNSProxyUUID

- **Type:** `string`
- **Required:** No

The DNS proxy UUID for this app. Available in iOS 16 and later.

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

The relay UUID for this app. Available in iOS 17 and later.

### Removable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, this app isn’t removable while it’s a managed app. Available in iOS 14 and later, and tvOS 14 and later.

### TapToPayScreenLock

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, Tap to Pay on iPhone requires users to use Face ID or a passcode to unlock their device after every transaction that requires a customer’s card PIN. If `false`, the user can configure this setting on their device.

Available in iOS 16.4 and later.

### VPNUUID

- **Type:** `string`
- **Required:** No

A per-app VPN unique identifier for this app. Available in iOS 7 and later.

