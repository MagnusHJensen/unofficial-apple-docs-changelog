# DiskManagementSettingsRestrictionsObject

The restrictions for the disk.

**Platforms:** macOS 15.0

## Properties

### ExternalStorage

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `ReadOnly`, `Disallowed`

Specifies the mount policy for external storage:

- `Allowed`: The system can mount external storage that is read-write or read-only.
- `ReadOnly`: The system can only mount read-only external storage. Note that external storage that is read-write will not be mounted read-only.
- `Disallowed`: The system can’t mount any external storage.

### NetworkStorage

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `ReadOnly`, `Disallowed`

Specifies the mount policy for network storage:

- `Allowed`: The system can mount network storage that is read-write or read-only.
- `ReadOnly`: The system can only mount read-only network storage. Note that network storage that is read-write will not be mounted read-only.
- `Disallowed`: The system can’t mount any network storage.

