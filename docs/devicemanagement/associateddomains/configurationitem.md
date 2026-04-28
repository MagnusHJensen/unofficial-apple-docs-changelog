# AssociatedDomains.ConfigurationItem

A dictionary that maps apps to their associated domains.

**Platforms:** macOS 10.15, Device Assignment Services , VPP License Management 

## Properties

### ApplicationIdentifier

- **Type:** `string`
- **Required:** Yes

The app identifier to associate the domains with.

### AssociatedDomains

- **Type:** `[string]`
- **Required:** Yes

The domains to associate with the app. Each string is in the form of `service:domain`. Use fully qualified hostnames, such as `www.example.com`. See [Supporting associated domains](/documentation/Xcode/supporting-associated-domains) for more information.

### EnableDirectDownloads

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables direct download of data for this domain instead of through a CDN. Set the entitlement value for this domain to `service:domain?mode=managed`; otherwise, the system ignores this value. Available in macOS 11 and later.

