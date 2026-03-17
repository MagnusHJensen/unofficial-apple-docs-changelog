# RosterLocation

A location’s properties and their values.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### name

- **Type:** `string`
- **Required:** No

The location name. Maximum length 1024 UTF-8 characters.

### op_date

- **Type:** `string`
- **Required:** No

The time stamp, in iSO 8601 format, when the location was added, updated, or deleted.

### source

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LDAP`, `SIS`, `CSV`, `MANUAL`, `SYSTEM`, `ENROLLMENT`, `DEP`, `FORMULA_GENERATED`, `SFTP`

The data source where the location was created.

### source_system_identifier

- **Type:** `string`
- **Required:** No

The identifier configured by organization for the location. Maximum length 256 UTF-8 characters.

### status

- **Type:** `string`
- **Required:** No

The status of the location.

### unique_identifier

- **Type:** `string`
- **Required:** No

The unique identifier for the location. Maximum length 256 UTF-8 characters.

