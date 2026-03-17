# VppClientConfigRequest

The request for the client configuration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### clientContext

- **Type:** `string`
- **Required:** No

Any JSON string under 256 bytes. The server stores the value of this field, and this value is returned in all responses. To clear the field’s value, provide an empty string as the input value (””).

See [Protecting Your VPP Account](/documentation/devicemanagement/protecting-your-vpp-account) for more information.

### notificationToken

- **Type:** `string`
- **Required:** No

The token to use when sending notifications through `notificationURL`.

### sToken

- **Type:** `string`
- **Required:** Yes

The authentication token. For more information, see [Authentication](/documentation/devicemanagement/managing-apps-and-books-through-web-services-legacy#Authentication).

