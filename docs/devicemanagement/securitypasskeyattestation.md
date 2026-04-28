# SecurityPasskeyAttestation

The declaration to configure the device to allow WebAuthn enterprise attestation for certain passkeys.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### AttestationIdentityAssetReference

- **Type:** `string`
- **Required:** Yes

The identifier of an asset declaration that contains the identity to install and use for passkey attestation.

### AttestationIdentityKeyIsExtractable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the private key for the attestation identity is extractable in the keychain.

### RelyingParties

- **Type:** `[string]`
- **Required:** Yes

An array of the relying parties to allow enterprise attestation.

## Discussion

Specify `com.apple.configuration.security.passkey.attestation` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.security.passkey.attestation",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "AttestationIdentityAssetReference": "AD0A8CB5-64EE-4CC9-8CB6-22DCBE6ED38A",
        "RelyingParties": [
            "example.com"
        ]
    }
}
```

