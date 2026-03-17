# ReturnToServiceResponse.ReturnToService

A dictionary containing the configuration for return to service.

**Platforms:** iOS 26.0, iPadOS 26.0, visionOS 26.0

## Properties

### BootstrapToken

- **Type:** `data`
- **Required:** No

The system uses the bootstrap token for return to service with app preservation. Required when Automated Device Enrollment enables return to service for the device.

If the bootstrap token isn’t present, the device performs a full erasure and a regular return to service, and can’t preserve any data for app preservation.

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device automatically erases itself and then performs reenrollment.

### MDMProfileData

- **Type:** `data`
- **Required:** No

The MDM profile that installs after erasure when using return to service. If provided, the device uses this profile directly instead of fetching it from the server. This key is required if the device’s Automated Device Enrollment profile contains the `configuration-web-url` key.

The device always downloads the Automated Device Enrollment profile even when this key is present, so the supervision identity, MDM removability, and other settings still apply. However, the device doesn’t use the specified URL in the Automated Device Enrollment profile to fetch the MDM profile.

### WiFiProfileData

- **Type:** `data`
- **Required:** No

The Wi-Fi profile that installs after erasure when using return to service. This is required when the device doesn’t have Ethernet access.

