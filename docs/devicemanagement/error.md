# Error

Information about an error that occurred while processing a request.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### code

- **Type:** `string`
- **Required:** Yes

The specific code for the underlying cause of the error.

### detail

- **Type:** `string`
- **Required:** No

More detailed information about the cause of the error, intended to help identify possible resolutions.

### id

- **Type:** `string`
- **Required:** Yes

The identifier of the error, mapping to the occurrence.

### source

- **Type:** `Error.Source`
- **Required:** No

An object containing a reference to the source of the error.

### status

- **Type:** `string`
- **Required:** Yes

The HTTP status code the error maps to.

### title

- **Type:** `string`
- **Required:** Yes

A developer-friendly title for the error.

## Topics

### Related Objects

- [Error.Source](/documentation/devicemanagement/error/source-data.dictionary) - An object that represents the source of an error.

