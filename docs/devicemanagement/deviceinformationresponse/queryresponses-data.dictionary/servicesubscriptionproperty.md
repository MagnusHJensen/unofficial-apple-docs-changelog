# DeviceInformationResponse.QueryResponses.ServiceSubscriptionProperty

The response dictionary that contains information about the active service subscription.

**Platforms:** iOS 12.0, iPadOS 12.0, Mac Catalyst 12.0, Device Assignment Services , VPP License Management 

## Properties

### CarrierSettingsVersion

- **Type:** `string`
- **Required:** No

The version of the carrier settings. This value is available in iOS 12 and later.

### CurrentCarrierNetwork

- **Type:** `string`
- **Required:** No

The name of the current carrier network. This value is available in iOS 12 and later.

### CurrentMCC

- **Type:** `string`
- **Required:** No

The current mobile country code (MCC). This value is available in iOS 12 and later.

### CurrentMNC

- **Type:** `string`
- **Required:** No

The current mobile network code (MNC). This value is available in iOS 12 and later.

### EID

- **Type:** `string`
- **Required:** No

The eSIM identifier. This value is available in iOS 14 and later.

### ICCID

- **Type:** `string`
- **Required:** No

The integrated circuit card identifier (ICCID) value. This value is available in iOS 12 and later.

### IMEI

- **Type:** `string`
- **Required:** No

The device International Mobile Equipment Identity (IMEI) number. This value is available in iOS 12 and later.

### IsDataPreferred

- **Type:** `boolean`
- **Required:** No

If `true`, this subscription is the preference for data. This value is available in iOS 12 and later.

### IsRoaming

- **Type:** `boolean`
- **Required:** No

If `true`, the phone is roaming. This value is available in iOS 12 and later.

### IsVoicePreferred

- **Type:** `boolean`
- **Required:** No

If `true`, this subscription is the preference for voice. This value is available in iOS 12 and later.

### Label

- **Type:** `string`
- **Required:** No

The label of this subscription. This value is available in iOS 12 and later.

### LabelID

- **Type:** `string`
- **Required:** No

The unique identifier for this subscription. This value is available in iOS 12 and later.

### MEID

- **Type:** `string`
- **Required:** No

The device Mobile Equipment Identifier (MEID) number. This query is available in iOS 12 and later.

### PhoneNumber

- **Type:** `string`
- **Required:** No

The raw phone number without punctuation and including country code. This value is available in iOS 12 and later.

### Slot

- **Type:** `string`
- **Required:** No

The description of the slot that contains the SIM representing this subscription. This value is available in iOS 12 and later.

### SubscriberCarrierNetwork

- **Type:** `string`
- **Required:** No

The name of the home carrier network. This value is available in iOS 16 and later.

