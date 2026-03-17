# ReturnToServiceResponse

The return-to-service response details.

**Platforms:** iOS 26.0, iPadOS 26.0, visionOS 26.0

## Properties

### PreserveDataPlan

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device preserves the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 26.4 and later.

### ReturnToService

- **Type:** `ReturnToServiceResponse.ReturnToService`
- **Required:** Yes

A dictionary containing the configuration for return to service.

## Topics

### Objects

- [ReturnToServiceResponse.ReturnToService](/documentation/devicemanagement/returntoserviceresponse/returntoservice-data.dictionary) - A dictionary containing the configuration for return to service.

