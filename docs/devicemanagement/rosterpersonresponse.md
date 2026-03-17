# RosterPersonResponse

The response that contains the people in the organization.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### cursor

- **Type:** `string`
- **Required:** No

A hex string that should be used for the next request to paginate. This field data type has a maximum length of 512 UTF-8 characters.

### more_to_follow

- **Type:** `boolean`
- **Required:** No

Indicates whether the request’s limit and cursor values resulted in only a partial list of classes. If `true`, the Mobile Device Management server should then make another request (starting from the newly returned cursor) to obtain additional records.

### persons

- **Type:** `[RosterPerson]`
- **Required:** No

Provides information about persons, both teachers and students, sorted in lexical order by a person `source_system_identifier`. The organization must provide this identifier to Apple.

