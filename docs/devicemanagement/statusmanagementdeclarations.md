# StatusManagementDeclarations

The status item that reports the device’s processed declarations.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### management.declarations

- **Type:** `StatusManagementDeclarationsDeclarationsObject`
- **Required:** Yes

A collection of the client’s processed declarations.

## Discussion

The name of the declaration status item is `management.declarations`.

### Status item availability

### Status item example

```json
{
    "management": {
        "declarations": {
            "activations": [
                {
                    "identifier": "com.example.activation.main",
                    "server-token": "A1B2C3D4-E5F6-7890-ABCD-EF1234567890",
                    "active": true,
                    "valid": "valid"
                }
            ],
            "configurations": [
                {
                    "identifier": "com.example.config.passcode",
                    "server-token": "B2C3D4E5-F6A7-8901-BCDE-F01234567891",
                    "active": true,
                    "valid": "valid"
                }
            ],
            "assets": [],
            "management": []
        }
    }
}
```

## Topics

### Objects

- [StatusManagementDeclarationsDeclarationsObject](/documentation/devicemanagement/statusmanagementdeclarationsdeclarationsobject) - A collection of the client’s processed declarations.

