# DeviceInformationResponse.QueryResponses.ServiceSubscriptionProperty

The response dictionary that contains information about the active service subscription.

**Platforms:** iOS 12.0, iPadOS 12.0, Mac Catalyst 12.0

## Properties

### CarrierSettingsVersion

- **Type:** `string`
- **Required:** No

The version of the carrier settings.

### CurrentCarrierNetwork

- **Type:** `string`
- **Required:** No

The name of the current carrier network.

### CurrentMCC

- **Type:** `string`
- **Required:** No

The current mobile country code (MCC).

### CurrentMNC

- **Type:** `string`
- **Required:** No

The current mobile network code (MNC).

### EID

- **Type:** `string`
- **Required:** No

The eSIM identifier.

Available: iOS 14+ | iPadOS 14+

### ICCID

- **Type:** `string`
- **Required:** No

The integrated circuit card identifier (ICCID) value.

### IMEI

- **Type:** `string`
- **Required:** No

The device International Mobile Equipment Identity (IMEI) number.

### IsDataPreferred

- **Type:** `boolean`
- **Required:** No

If `true`, this subscription is the preference for data.

### IsRoaming

- **Type:** `boolean`
- **Required:** No

If `true`, the phone is roaming.

### IsVoicePreferred

- **Type:** `boolean`
- **Required:** No

If `true`, this subscription is the preference for voice.

### Label

- **Type:** `string`
- **Required:** No

The label of this subscription.

### LabelID

- **Type:** `string`
- **Required:** No

The unique identifier for this subscription.

### MEID

- **Type:** `string`
- **Required:** No

The device Mobile Equipment Identifier (MEID) number.

### PhoneNumber

- **Type:** `string`
- **Required:** No

The raw phone number without punctuation and including country code.

### Slot

- **Type:** `string`
- **Required:** No

The description of the slot that contains the SIM representing this subscription.

### SubscriberCarrierNetwork

- **Type:** `string`
- **Required:** No

The name of the home carrier network.

Available: iOS 16+ | iPadOS 16+

