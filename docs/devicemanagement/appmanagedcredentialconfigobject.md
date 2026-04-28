# AppManagedCredentialConfigObject

A dictionary of values associated with a credential config.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, visionOS 2.4, Device Assignment Services , VPP License Management 

## Properties

### AssetReference

- **Type:** `string`
- **Required:** Yes

Specifies the identifier of an asset declaration containing a username and password. The [ManagedApp](/documentation/ManagedApp) framework makes the password available to the app or extension. The [ManagedApp](/documentation/ManagedApp) framework ignores the username.

### Identifier

- **Type:** `string`
- **Required:** Yes

The app or extension uses this identifier to fetch the corresponding password using the [ManagedApp](/documentation/ManagedApp) framework. App developers define the values for these identifiers.

