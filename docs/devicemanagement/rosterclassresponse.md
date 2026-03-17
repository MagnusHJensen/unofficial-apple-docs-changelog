# RosterClassResponse

The response that contains a list of classes.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### classes

- **Type:** `[RosterClass]`
- **Required:** No

An array of classes, sorted in lexical order by a class `source_system_identifier`. The organization must provide this identifier to Apple.

### cursor

- **Type:** `string`
- **Required:** No

A hex string that should be used for the next request to paginate. This field data type has a maximum length of 512 UTF-8 characters.

### more_to_follow

- **Type:** `boolean`
- **Required:** No

Indicates whether the request’s limit and cursor values resulted in only a partial list of classes. If `true`, the MDM server should then make another request (starting from the newly returned cursor) to obtain additional records.

