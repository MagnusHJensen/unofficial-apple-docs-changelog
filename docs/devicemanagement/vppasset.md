# VppAsset

A particular asset in the purchase program.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### adamIdStr

- **Type:** `string`
- **Required:** No

The unique identifier for a product in the iTunes Store.

### assignedCount

- **Type:** `int32`
- **Required:** No

The total number of licenses assigned to a user or a device for the specified adamId and pricingParam.

### availableCount

- **Type:** `int32`
- **Required:** No

The total number of licenses available to be assigned for the specified adamId and pricingParam.

### deviceAssignable

- **Type:** `boolean`
- **Required:** No

If `true`, a license for the specified adamId may be assigned to a device.

### isIrrevocable

- **Type:** `boolean`
- **Required:** No

If `true`, once a license is assigned for specified adamId, the license may not be revoked or disassociated.

### pricingParam

- **Type:** `string`
- **Required:** No

The quality of a product in the iTunes Store. Possible values are:

STDQ: Standard quality

PLUS: High quality

### productTypeId

- **Type:** `int32`
- **Required:** No

The type of a product. Possible values are:

7 = macOS software

8 = iOS or macOS app from the App Store

10 = Book

### productTypeName

- **Type:** `string`
- **Required:** No

The name of the product type.

### retiredCount

- **Type:** `int32`
- **Required:** No

The total number of licenses that have been retired for the specified adamId and pricingParam.

### totalCount

- **Type:** `int32`
- **Required:** No

The total number of licenses managed by your organization for the specified adamId and pricingParam.

