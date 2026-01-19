# GetVppLicensesRequest

The request for licenses.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

The `batchToken` and `sinceModifiedToken` encode whether `adamId` and `pricingParam` were originally passed; therefore, if the `batchToken` or `sinceModifiedToken` is present on the request, the `adamId` and `pricingParam` fields (if passed) are ignored.

