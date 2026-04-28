# ActiveNSExtensionsResponse.ExtensionsItem

A dictionary that contains information about an extension.

**Platforms:** macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### ContainerDisplayName

- **Type:** `string`
- **Required:** No

The display name of the container.

### ContainerIdentifier

- **Type:** `string`
- **Required:** No

The identifier of the container.

### DisplayName

- **Type:** `string`
- **Required:** Yes

The extension’s display name.

### ExtensionPoint

- **Type:** `string`
- **Required:** Yes

The [NSExtensionPointIdentifier](/documentation/BundleResources/Information-Property-List/NSExtension/NSExtensionPointIdentifier) for the extension.

### Identifier

- **Type:** `string`
- **Required:** Yes

The identifier of the extension.

### Path

- **Type:** `string`
- **Required:** Yes

The path to the extension.

### UserElection

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Default`, `Use`, `Ignore`

The user-selected state of the extension, which a user sets in the Extensions preference pane in System Preferences.

### Version

- **Type:** `string`
- **Required:** Yes

The version of the extension.

