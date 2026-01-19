# Get Licenses

Get the set of licenses managed by your organization.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This call returns an exhaustive list of individual licenses managed by your organization. For performance reasons, it’s recommended that you use a combination of [Get Assets](/documentation/devicemanagement/get-assets-44p83)  and [Get Assignments](/documentation/devicemanagement/get-assignments-158kc) as described in the initial import process in [Managing Apps and Books Through Web Services](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy) instead of this API.

Additionally, it’s not recommended to use the Apple license IDs returned in the [GetVppLicensesResponse](/documentation/devicemanagement/getvpplicensesresponse) in MDM implementations because it creates an unnecessary dependence on internal data models, which can change.

### Example Request and Response

## Topics

### Request and Response

- [GetVppLicensesRequest](/documentation/devicemanagement/getvpplicensesrequest) - The request for licenses.
- [GetVppLicensesResponse](/documentation/devicemanagement/getvpplicensesresponse) - The response with the licenses.

