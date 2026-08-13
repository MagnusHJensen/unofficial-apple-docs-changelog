# ServiceConfigResponse

The service configuration for the Asset Management API.

## Properties

### errorCodes

- **Type:** `[ResponseErrorCode]`
- **Required:** No

The set of possible error numbers and their human-readable explanations.

### limits

- **Type:** `ServiceConfigResponse.Limits`
- **Required:** No

The set of current request limits.

### notificationTypes

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `ASSET_COUNT`, `ASSET_MANAGEMENT`, `USER_MANAGEMENT`, `USER_ASSOCIATED`, `SUBSCRIPTION_MANAGEMENT`, `SUBSCRIPTION_COUNT`, `UNLIMITED_ASSET`

The set of supported notification types.

### urls

- **Type:** `ServiceConfigResponse.Urls`
- **Required:** No

The set of current service URLs.

## Discussion

The values in `limits` and `urls` are dynamic and can change without notice. Sync them every 5 minutes rather than hard-coding them into your device management service.

## Topics

### Objects and Data Types

- [ServiceConfigResponse.Limits](/documentation/devicemanagement/serviceconfigresponse/limits-data.dictionary) - The set of current request limits.
- [ServiceConfigResponse.Urls](/documentation/devicemanagement/serviceconfigresponse/urls-data.dictionary) - The set of current service URLs.
- [ResponseErrorCode](/documentation/devicemanagement/responseerrorcode) - An error code.

