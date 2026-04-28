# SystemMigration.CustomBehaviorItem.PathsItem

The custom behavior path dictionary.

**Platforms:** macOS 10.12.4, Device Assignment Services , VPP License Management 

## Properties

### SourcePath

- **Type:** `string`
- **Required:** Yes

The path to the migrating file or directory on the source system.

### SourcePathInUserHome

- **Type:** `boolean`
- **Required:** Yes

If `true`, the source path is located within a user home directory.

### TargetPath

- **Type:** `string`
- **Required:** Yes

The path to the destination file or directory on the target system.

### TargetPathInUserHome

- **Type:** `boolean`
- **Required:** Yes

If `true`, the target path is located within a user home directory.

