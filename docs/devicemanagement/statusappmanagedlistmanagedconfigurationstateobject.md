# StatusAppManagedListManagedConfigurationStateObject

A dictionary that contains details about a declarative managed app’s managed configuration.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, visionOS 2.4, Device Assignment Services , VPP License Management 

## Properties

### state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `unknown`, `invalid`, `valid`

The managed configuration status.

- `unknown`: The managed configuration has not been read
- `invalid`: The managed configuration was read and deemed to be invalid
- `valid`: The managed configuration was read and deemed to be valid

