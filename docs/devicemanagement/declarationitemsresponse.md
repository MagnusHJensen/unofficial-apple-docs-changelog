# DeclarationItemsResponse

The set of available declarations on the server.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### Declarations

- **Type:** `DeclarationItemsResponse.ManifestDeclarationItems`
- **Required:** Yes

The set of available declarations on the server.

### DeclarationsToken

- **Type:** `string`
- **Required:** Yes

The current value of the declarations token. Clients use this to detect when declarations change so they can refetch the token.

## Topics

### Supporting Objects

- [DeclarationItemsResponse.ManifestDeclarationItems](/documentation/devicemanagement/declarationitemsresponse/manifestdeclarationitems) - The dictionary that contains the lists of declarations available on the server.
- [ManifestDeclaration](/documentation/devicemanagement/manifestdeclaration) - A dictionary that describes a declaration.

