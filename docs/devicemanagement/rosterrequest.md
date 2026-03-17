# RosterRequest

The request for a list of classes.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### cursor

- **Type:** `string`
- **Required:** No

A hex string that represents the starting position for a request. This is used for pagination. On the initial request, this should be omitted.

### limit

- **Type:** `int32`
- **Required:** No
- **Default:** `1000`

The maximum number of entries to return.

