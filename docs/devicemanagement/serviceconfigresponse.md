# ServiceConfigResponse

Service configuration, including request limits, available URLs, supported notification types, and error code reference information.

## Properties

### errorCodes

- **Type:** `[ResponseErrorCode]`
- **Required:** No

### limits

- **Type:** `ServiceConfigResponse.Limits`
- **Required:** No

### notificationTypes

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `ASSET_COUNT`, `ASSET_MANAGEMENT`, `USER_MANAGEMENT`, `USER_ASSOCIATED`, `SUBSCRIPTION_MANAGEMENT`, `SUBSCRIPTION_COUNT`, `UNLIMITED_ASSET`

### urls

- **Type:** `ServiceConfigResponse.Urls`
- **Required:** No

## Topics

### Dictionaries

- [ServiceConfigResponse.Limits](/documentation/devicemanagement/serviceconfigresponse/limits-data.dictionary) - Request limits for the managed location. Each entry maps a limit name to its current integer value.
- [ServiceConfigResponse.Urls](/documentation/devicemanagement/serviceconfigresponse/urls-data.dictionary) - Service URLs for the managed location. Each entry maps a URL name to its corresponding endpoint.

