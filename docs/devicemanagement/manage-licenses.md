# Manage Licenses

Associate and disassociate licenses with users and devices.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This endpoint operates on a single asset (specified by the `{adamIdStr, pricingParam}` tuple) for multiple associations and disassociations in a single request.

Licenses are disassociated from all users specified by the `disassociateClientUserIdStrs` array, the devices specified by the `disassociateSerialNumbers` array, or the licenses specified by the `disassociateLicenseIdStrs` array (which must only specify licenses assigned to the specified asset). At most one of these `disassociate*` arrays may be specified per request.

Then licenses are associated either with the users specified by the `associateClientUserIdStrs` array or the devices specified by the `associateSerialNumbers` array. Device assignment doesn’t trigger notifcation.

At most, one `associate*` and one `disassociate*` array is allowed per request. Specifying more than one of either `associate*` or `disassociate*` arrays result in undefined behavior.

### Example Request and Response with a Serial Number

### Example Request and Response with a Client User ID String

## Topics

### Request and Response

- [ManageVppLicensesByAdamIdRequest](/documentation/devicemanagement/managevpplicensesbyadamidrequest) - The request to manage licenses.
- [ManageVppLicensesByAdamIdResponse](/documentation/devicemanagement/managevpplicensesbyadamidresponse) - The response from managing licenses.

