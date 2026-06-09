# UnlimitedResponseAsset

An asset with an unlimited license that the organization owns.

## Properties

### adamId

- **Type:** `string`
- **Required:** No

The unique identifier for the product in the store.

### assignedCount

- **Type:** `int32`
- **Required:** No

The assigned amount of the asset.

### deviceAssignable

- **Type:** `boolean`
- **Required:** No

The flag denoting whether the asset is device-assignable.

### pricingParam

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `STDQ`, `PLUS`

The quality of the product in the store.

### productType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `App`, `Book`

The asset product type.

### revocable

- **Type:** `boolean`
- **Required:** No

The flag denoting whether the asset is revocable.

### supportedPlatforms

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `iOS`, `macOS`, `tvOS`, `watchOS`, `visionOS`

The platforms that the asset supports.

## Overview

The server returns this object in the `unlimitedAssets` array of [GetAssetsResponse](/documentation/devicemanagement/getassetsresponse) when the `unlimited` query parameter is set to `true`. This object omits the `availableCount`, `totalCount`, and `retiredCount` fields, which don’t apply to unlimited licenses.

