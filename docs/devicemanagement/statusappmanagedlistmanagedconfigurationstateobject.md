# StatusAppManagedListManagedConfigurationStateObject

The status of any app managed configuration. This key is only present when the managed app has a managed configuration.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, macOS 27.0 (Beta), visionOS 2.4

## Properties

### state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `unknown`, `invalid`, `valid`

The managed configuration status.

- `unknown`: The managed configuration has not been read
- `invalid`: The managed configuration was read and deemed to be invalid
- `valid`: The managed configuration was read and deemed to be valid

