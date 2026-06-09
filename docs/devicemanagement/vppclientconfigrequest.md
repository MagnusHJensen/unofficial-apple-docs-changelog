# VppClientConfigRequest

The request for the client configuration.

**Platforms:** VPP License Management 1.0

## Properties

### clientContext

- **Type:** `string`
- **Required:** No

Any JSON string under 256 bytes. The server stores the value of this field, and this value returns in all responses. To clear the field’s value, provide an empty string as the input value (””).

### notificationToken

- **Type:** `string`
- **Required:** No

The token to use when sending notifications through `notificationURL`.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authenticate with the web service](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authenticate-with-the-web-service).

