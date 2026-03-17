# AccountDetail

The response that contains the details for an account.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### admin_id

- **Type:** `string`
- **Required:** No

The Apple Account of the person who generated the currently in-use tokens.

### facilitator_id

- **Type:** `string`
- **Required:** No

The legacy equivalent to the `admin_id` key. This key is deprecated and may not be returned in future responses.

### org_address

- **Type:** `string`
- **Required:** No

The organization address.

### org_email

- **Type:** `string`
- **Required:** No

The organization email address.

### org_id

- **Type:** `string`
- **Required:** No

The customer ID. This key is available only in protocol version 3 and later.

### org_id_hash

- **Type:** `string`
- **Required:** No

The SHA hash of an organization identifier. This helps Mobile Device Management (MDM) server match the hash with the `organizationIdHash` key in the [Client Configuration](/documentation/devicemanagement/client-configuration) API. This key is available only in protocol version 3 and later.

### org_name

- **Type:** `string`
- **Required:** No

The organization name.

### org_phone

- **Type:** `string`
- **Required:** No

The organization phone.

### org_type

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `edu`, `org`, `tvprovider`

The type of organization. Possible values are `edu` or `org`. This key is available only in protocol version 3 and later.

### org_version

- **Type:** `string`
- **Required:** No

Possible values: `v1` or `v2`. `v1` is for Apple Deployment Programs (like Device Enrollment Program or Volume Purchase Program) organizations and `v2` is for Apple School Manager (ASM) organizations. Currently `v2` is applicable only to educational organizations. This key is available only in protocol version 3 and later.

### server_name

- **Type:** `string`
- **Required:** No

The name of the MDM server.

### server_uuid

- **Type:** `string`
- **Required:** No

The system-generated server identifier.

### urls

- **Type:** `[Url]`
- **Required:** No

The list of URLs available in the MDM service. This key is valid in X-Server-Protocol-Version 3 and later.

