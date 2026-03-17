# DeviceInformationResponse.QueryResponses.OSUpdateSettings

The response dictionary that contains operating system update settings.

**Platforms:** macOS 10.11

## Properties

### AutoCheckEnabled

- **Type:** `boolean`
- **Required:** No

The preference to automatically check for app updates. This value is available in macOS 10.11 and later.

### AutomaticAppInstallationEnabled

- **Type:** `boolean`
- **Required:** No

The preference to automatically install app updates. This value is available in macOS 10.11 and later.

### AutomaticOSInstallationEnabled

- **Type:** `boolean`
- **Required:** No

The preference to automatically install operating system updates. This value is available in macOS 10.11 and later.

### AutomaticSecurityUpdatesEnabled

- **Type:** `boolean`
- **Required:** No

The preference to automatically install system data files and security updates. This value is available in macOS 10.11 and later.

### BackgroundDownloadEnabled

- **Type:** `boolean`
- **Required:** No

The preference to download app updates in the background. This value is available in macOS 10.11 and later.

### CatalogURL

- **Type:** `string`
- **Required:** No

The URL to the software update catalog the client is using. This value is available in macOS 10.11 and later.

### IsDefaultCatalog

- **Type:** `boolean`
- **Required:** No

If `true`, `CatalogURL` is the default catalog. This value is available in macOS 10.11 and later.

### PerformPeriodicCheck

- **Type:** `boolean`
- **Required:** No

If `true`, start a new scan. This value is available in macOS 10.11 and later.

### PreviousScanDate

- **Type:** `date`
- **Required:** No

The date of the last software update scan. This value is available in macOS 10.11 and later.

