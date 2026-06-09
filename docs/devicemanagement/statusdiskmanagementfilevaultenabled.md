# StatusDiskManagementFileVaultEnabled

The status item that reports whether FileVault is enabled.

**Platforms:** macOS 14.0

## Properties

### diskmanagement.filevault.enabled

- **Type:** `boolean`
- **Required:** Yes

A Boolean value that specifies the File Vault enabled status on the device.

## Discussion

### Status item availability

### Status item example

```json
{
    "diskmanagement": {
        "filevault": {
            "enabled": true
        }
    }
}
```

