# ManifestDeclaration

A dictionary that describes a declaration.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### Identifier

- **Type:** `string`
- **Required:** Yes

The declaration’s identifier.

### ServerToken

- **Type:** `string`
- **Required:** Yes

The `ServerToken` value of the declaration.

The client uses this to determine if the actual payload is different from the one on the client. Servers must compute the token over the entire declaration content to ensure the value always changes whenever there’s any change to the content.

