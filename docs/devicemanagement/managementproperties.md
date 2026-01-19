# ManagementProperties

The declaration to configure the properties on the device.

**Platforms:** iOS 16.0, iPadOS 16.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.management.properties` as the declaration type.

### Management declaration example

```json
{
    "Type": "com.apple.management.properties",
    "Identifier": "187C9F47-297C-4811-80C4-2E19D3C11963",
    "ServerToken": "526CE2FB-DB79-409A-825D-8C5DC5EE873E",
    "Payload": {
        "is-part-time": false,
        "groups": [
            "teacher",
            "grade 1",
            "grade 2",
            "it-admin"
        ]
    }
}
```

