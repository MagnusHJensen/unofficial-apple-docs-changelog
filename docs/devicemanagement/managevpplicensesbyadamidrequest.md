# ManageVppLicensesByAdamIdRequest

The request to manage licenses.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

The maximum number of entries allowed in the `associate*` and `disassociate*` arrays are indicated by the `maxBatchAssociateLicenseCount` or `maxBatchDisassociateLicenseCount` fields available from the [Service Configuration](/documentation/devicemanagement/service-configuration) response.

Requests that exceed the current limits are rejected with the error code 9602 “Invalid Argument” and no work is done. If you receive this error code, query the [Service Configuration](/documentation/devicemanagement/service-configuration) endpoint to retrieve new `maxBatchAssociateLicenseCount` and `maxBatchDisassociateLicenseCount` values, correct the last request that was rejected, and resend the request.

