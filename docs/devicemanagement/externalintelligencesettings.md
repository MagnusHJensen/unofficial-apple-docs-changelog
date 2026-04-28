# ExternalIntelligenceSettings

The declaration to configure External Intelligence Integrations settings.

**Platforms:** iOS 26.4, iPadOS 26.4, Mac Catalyst 26.4, macOS 26.4, visionOS 26.4, Device Assignment Services , VPP License Management 

## Properties

### AllowedWorkspaceIDs

- **Type:** `[string]`
- **Required:** No

An array of strings, but currently restricted to a single element. If present, Apple Intelligence allows use of only the given external integration workspace ID, and requires a sign-in to make requests. The user is required to sign in to integrations that support signing in. Multiple values combine using an intersect operation.

### AllowSignIn

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables sign-in for external intelligence integrations.

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables external intelligence integrations.

## Discussion

Specify `com.apple.configuration.external-intelligence.settings` as the declaration type.

### Configuration availability

### Configuration examples

This configuration restricts external intelligence integrations.

```json
{
    "Type": "com.apple.configuration.external-intelligence.settings",
    "Identifier": "A1B2C3D4-E5F6-4A5B-9C8D-7E6F5A4B3C2D",
    "ServerToken": "F1E2D3C4-B5A6-4D5E-8F9A-0B1C2D3E4F5A",
    "Payload": {
        "Enabled": false,
        "AllowSignIn": false,
        "AllowedWorkspaceIDs": [
            "ABCDEFGH"
        ]
    }
}
```

