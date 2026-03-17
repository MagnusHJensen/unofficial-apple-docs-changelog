# GetVppAssetRequest

The request for an asset.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### includeLicenseCounts

- **Type:** `boolean`
- **Required:** No

If `true`, returns the total number of licenses, the number of assigned licenses, and the number of unassigned licenses in the response for each asset.

### pricingParam

- **Type:** `string`
- **Required:** No

The quality of a product in the iTunes Store. If a pricing parameter is specified, only records with that parameter are included in the results. Possible values are:

- `STDQ`: Standard quality
- `PLUS`: High quality

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authentication](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authentication).

