# MachineInfo

A device’s information in response to a MDM enrollment profile request.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### IMEI

- **Type:** `string`
- **Required:** No

The device’s IMEI (if available).

Available: iOS 7+ | iPadOS 7+ | watchOS 10+

### LANGUAGE

- **Type:** `string`
- **Required:** No

The user’s currently-selected language, for example, `en`.

### MANDATORY_SOFTWARE_UPDATE_REQUIRED

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, indicates that the device requires a mandatory software update during Setup Assistant. The MDM server can return a 403 with a [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) error to force the device to update to a specific version instead of the device choosing a version.

Available: macOS 26.1+

### MDM_CAN_REQUEST_PSSO_CONFIG

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, indicates that the server can trigger the device to do a required Platform SSO authentication before enrolling.

Available: macOS 26+

### MDM_CAN_REQUEST_SOFTWARE_UPDATE

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, indicates that the server can trigger the device to do a required software update.

Available: iOS 17+ | iPadOS 17+ | macOS 14+

### MEID

- **Type:** `string`
- **Required:** No

The device’s MEID (if available).

Available: iOS 7+ | iPadOS 7+ | watchOS 10+

### OS_VERSION

- **Type:** `string`
- **Required:** Yes

The OS version installed on the device, for example, 17.0.

Available: iOS 17+ | iPadOS 17+ | macOS 14+ | tvOS 17+ | visionOS 1.1+ | watchOS 10+

### PAIRING_TOKEN

- **Type:** `data`
- **Required:** No

The pairing token to validate when a watch is enrolling.

Available: watchOS 10+

### PRODUCT

- **Type:** `string`
- **Required:** Yes

The device’s product type, for example, `iPhone5,1`.

### SERIAL

- **Type:** `string`
- **Required:** Yes

The device’s serial number.

### SOFTWARE_UPDATE_DEVICE_ID

- **Type:** `string`
- **Required:** No

The device model identifier used to lookup available OS updates through https://gdmf.apple.com/v2/pmv.

Available: iOS 17.4+ | iPadOS 17.4+ | macOS 14.4+ | visionOS 1.1+ | watchOS 10+

### SUPPLEMENTAL_BUILD_VERSION

- **Type:** `string`
- **Required:** No

The device’s operating system supplemental build version (if available).

Available: iOS 17+ | iPadOS 17+ | macOS 14+ | tvOS 17+ | visionOS 1.1+ | watchOS 10+

### SUPPLEMENTAL_OS_VERSION_EXTRA

- **Type:** `string`
- **Required:** No

The device’s operating system supplemental OS version extra (if available).

Available: iOS 17+ | iPadOS 17+ | macOS 14+ | tvOS 17+ | visionOS 1.1+ | watchOS 10+

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID.

### VERSION

- **Type:** `string`
- **Required:** Yes

The build version installed on the device, for example, `7A182`.

## Discussion

This dictionary is CMS-signed with the device identity certificate. The system includes the device’s certificate and all necessary intermediate certificates. The certificate chain should validate against the Apple Root CA.

