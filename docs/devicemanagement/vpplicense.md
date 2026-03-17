# VppLicense

A license for a product in the purchase program.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### adamId

- **Type:** `int64`
- **Required:** No

The unique identifier for a product in the iTunes Store.

### adamIdStr

- **Type:** `string`
- **Required:** No

The unique identifier for a product in the iTunes Store.

### clientUserIdStr

- **Type:** `string`
- **Required:** No

The client user ID for the user to whom this license is assigned.

### isIrrevocable

- **Type:** `boolean`
- **Required:** No

If true, once a license is assigned for specified adamId, the license may not be revoked or disassociated.

### itsIdHash

- **Type:** `string`
- **Required:** No

The hash of the iTunes Store ID for the user to whom this license is assigned. The itsIdHash field is included only when the user is associated with an iTunes Store account.

### licenseIdStr

- **Type:** `string`
- **Required:** No

The identifier of the assigned license.

### pricingParam

- **Type:** `string`
- **Required:** No

The quality of a product in the iTunes Store. Possible values are:

- `STDQ`: Standard quality
- `PLUS`: High quality

### productTypeId

- **Type:** `int32`
- **Required:** No

The type of a product. Possible values are:

- `7` = macOS software
- `8` = iOS or macOS app from the App Store
- `10` = Book

### productTypeName

- **Type:** `string`
- **Required:** No

The name of the product type.

### serialNumber

- **Type:** `string`
- **Required:** No

The device serial number to which this license is assigned.

### status

- **Type:** `string`
- **Required:** No

The current state of the license. Possible values are:

- `Associated`
- `Available`
- `Refunded`

### userId

- **Type:** `int64`
- **Required:** No

The unique identifier assigned by the VPP for the user to whom this license is assigned.

### userIdStr

- **Type:** `string`
- **Required:** No

The string representation of the unique identifier assigned by the VPP for the user to whom this license is assigned.

