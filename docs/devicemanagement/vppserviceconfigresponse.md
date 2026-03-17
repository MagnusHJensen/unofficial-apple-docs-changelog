# VppServiceConfigResponse

The response with the service configuration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### associateLicenseSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the `Associate License` endpoint. Note the `Associate License` endpoint has been deprecated. Use [Manage Licenses](/documentation/devicemanagement/manage-licenses) instead.

### clientConfigSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Client Configuration](/documentation/devicemanagement/client-configuration) endpoint.

### contentMetadataLookupUrl

- **Type:** `string`
- **Required:** No

The URL that returns metadata about a product in the iTunes Store.

See [Getting App and Book Information (Legacy)](/documentation/devicemanagement/getting-app-and-book-information-legacy), for more information.

### disassociateLicenseSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the `Disassociate License` endpoint. Note the `Disassociate License`  endpoint has been deprecated. Use [Manage Licenses](/documentation/devicemanagement/manage-licenses) instead.

### editUserSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Edit a User](/documentation/devicemanagement/edit-a-user) endpoint.

### errorCodes

- **Type:** `VppErrorCode`
- **Required:** No

List of possible error numbers and their human-readable explanations.

### errorMessage

- **Type:** `string`
- **Required:** No

The human-readable explanation of the error.

### errorNumber

- **Type:** `int32`
- **Required:** No

The numeric code of the error.

### getLicensesSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Get Licenses](/documentation/devicemanagement/get-licenses) endpoint.

### getUserSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Get a User](/documentation/devicemanagement/get-a-user) endpoint.

### getUsersSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Get Users](/documentation/devicemanagement/get-users-5boi1) endpoint.

### getVPPAssetsSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Get Assets](/documentation/devicemanagement/get-assets-44p83) endpoint.

### invitationEmailUrl

- **Type:** `string`
- **Required:** No

The URL template for inviting users to an organization.

### manageVPPLicensesByAdamIdSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Manage Licenses](/documentation/devicemanagement/manage-licenses) endpoint.

### maxBatchAssociateLicenseCount

- **Type:** `int32`
- **Required:** No

The maximum number of entries allowed in the arrays for associating licenses with [Manage Licenses](/documentation/devicemanagement/manage-licenses). The MDM server should check this value every 5 minutes, because it could change without notice.

### maxBatchDisassociateLicenseCount

- **Type:** `int32`
- **Required:** No

The maximum number of entries allowed in the arrays for disassociating licenses from [Manage Licenses](/documentation/devicemanagement/manage-licenses). The MDM server should check this value every 5 minutes, because it could change without notice.

### registerUserSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Register a User](/documentation/devicemanagement/register-a-user) endpoint.

### retireUserSrvUrl

- **Type:** `string`
- **Required:** No

The URL for the [Retire a User](/documentation/devicemanagement/retire-a-user) endpoint.

### status

- **Type:** `int32`
- **Required:** No

The status code for the response. Possible values are:

`0` = Success. `-1` = Failure.

### vppWebsiteUrl

- **Type:** `string`
- **Required:** No

The URL for the VPP website.

