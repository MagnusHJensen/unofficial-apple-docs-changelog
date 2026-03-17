# ResponseAsset

The asset that the organization owns.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### adamId

- **Type:** `string`
- **Required:** No

The unique identifier for the product in the store.

### assignedCount

- **Type:** `int32`
- **Required:** No

The assigned amount of the asset.

### availableCount

- **Type:** `int32`
- **Required:** No

The available amount of the asset.

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

### retiredCount

- **Type:** `int32`
- **Required:** No

The retired amount of the asset.

### revocable

- **Type:** `boolean`
- **Required:** No

The flag denoting whether the asset is revocable.

### totalCount

- **Type:** `int32`
- **Required:** No

The total amount of the asset.

### supportedPlatforms

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `iOS`, `macOS`, `tvOS`, `watchOS`

The platforms that the asset supports.

