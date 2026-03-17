# ACMECredential

An ACME identity that the device generates.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Properties

### Attest

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device provides attestations that describe the device and the generated key to the ACME server. The server can use the attestations as strong evidence that the key is bound to the device, and that the device has properties listed in the attestation. The server can use that as part of a trust score to decide whether to issue the requested certificate. When `Attest` is `true`, set `HardwareBound` to `true`. On macOS, set this key, if present, to `false`. The hardware requirements for attestation are described below.

### ClientIdentifier

- **Type:** `string`
- **Required:** Yes

The server can use this as a one-time code to prevent issuing multiple certificates. It also indicates to the ACME server that the device has access to a valid client identifier that the enterprise infrastructure issued. This can help the ACME server determine whether to trust the device, however this is a relatively weak indication because of the risk that an attacker may intercept and duplicate the client identifier.

### DirectoryURL

- **Type:** `string`
- **Required:** Yes

Specifies the directory URL of the ACME server. Use the `https` scheme for the URL.

### ExtendedKeyUsage

- **Type:** `[string]`
- **Required:** No

The device requests this extended key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.

The value is an array of strings. Each string is an OID in dotted notation. For example, `["1.3.6.1.5.5.7.3.2", "1.3.6.1.5.5.7.3.4"]` indicates client authentication and email protection.

### HardwareBound

- **Type:** `boolean`
- **Required:** Yes

If `false`, the private key isn’t bound to the device.

If `true`, the private key is bound to the device. The Secure Enclave generates the key pair, and the private key is cryptographically entangled with a system key. This protects the private key from being exported.

If `true`, `KeyType` needs to be `ECSECPrimeRandom` and `KeySize` needs to be `256` or `384`.

On macOS, this is a required key. Set the value to `false`.

### KeySize

- **Type:** `integer`
- **Required:** Yes

The valid values for `KeySize` depend on the values of `KeyType` and `HardwareBound`. See those keys for specific requirements.

### KeyType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RSA`, `ECSECPrimeRandom`

Specifies the type of key pair to generate.

`RSA` specifies an RSA key pair. If you set this value to `RSA`, set `KeySize` in the range `[1024..4096]` inclusive and a multiple of `8`, and set `HardwareBound` to `false`.

`ECSECPrimeRandom` specifies a key pair on the P-256, P-384 or P-521 curves as defined in FIPS Pub 186-4, and `KeySize` determines the specific curve. If you set this value to `ECSECPrimeRandom`, set `KeySize` to `256`, `384`, or `521`. The system only supports `256` and `384` for hardware bound keys.

> 

### Subject

- **Type:** `[[[string]]]`
- **Required:** Yes

The device requests this subject for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.

The representation of an X.500 name is an array of OID and value. For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` corresponds to:

`[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], [ [ "CN", "foo"] ], [ [ "1.2.5.3", "bar" ] ] ]`

You can represent OIDs as dotted numbers or use shortcuts for country (`C`), locality (`L`), state (`ST`), organization (`O`), organizational unit (`OU`), and common name (`CN`).

### SubjectAltName

- **Type:** `ACMECredentialSubjectAltNameObject`
- **Required:** No

Specifies the subject’s alternative name that the device requests for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.

### UsageFlags

- **Type:** `integer`
- **Required:** No

The device requests this key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.

The value is a bit field. Bit `0x01` indicates digital signature, and bit `0x04` indicates key encipherment.

## Discussion

This schema specifies how the device requests a client certificate from an Automated Certificate Management Environment (ACME) server. Use this to create a JSON document that the device downloads when resolving an asset.

When the device resolves the asset, first it generates an asymmetric key pair based upon the `KeyType`, `KeySize`, and `HardwareBound` fields. Then the device communicates with the ACME server. It requests a new order using the `ClientIdentifier` as the `permanent-identifier`. The ACME server responds with a challenge type of `device-attest-01`. If `Attest` is `true` the device requests an attestation of the key and device properties. Then it replies to the challenge with a WebAuthn attestation statement, and this contains the attestation if the device obtained one. The device submits a certificate signing request matching the key and containing the `ClientIdentifier`, `Subject`, `SubjectAltName`, `UsageFlags`, and `ExtendedKeyUsage` fields. The ACME server issues a certificate, and the device stores the resulting identity.

For details on the content of the attestation provided to the ACME server, see the documentation of the `DevicePropertiesAttestation` key in the [DeviceInformationResponse.QueryResponses](/documentation/devicemanagement/deviceinformationresponse/queryresponses-data.dictionary) response. In the attestation certificate the value of the freshness code OID is the SHA-256 hash of the `token` from the `device-attest-01` challenge.

### ACME attestation hardware support

The following table indicates which System on Chips (SoCs) support ACME attestation. If the Attest key is ignored, the ACME server does not receive an attestation.

### Credential example

```json
{
    "DirectoryURL": "https://example.com/acme/dir",
    "ClientIdentifier": "This is a ClientIdentifier",
    "KeySize": 384,
    "KeyType": "ECSECPrimeRandom",
    "HardwareBound": true,
    "Subject": [[["C", "US"]], [["O", "Example Inc."]], [["1.2.840.113635.100.6.99999.99999", "test custom OID value"]]],
    "SubjectAltName": {
        "dNSName": "site.example.com",
        "rfc822Name": "rfc822",
        "uniformResourceIdentifier": "https://www.example.com/",
        "ntPrincipalName": "Principal Example"
    },
    "UsageFlags": 5,
    "ExtendedKeyUsage": [
        "1.3.6.1.5.5.7.3.2"
    ],
    "Attest": true
}
```

## Topics

### Objects

- [ACMECredentialSubjectAltNameObject](/documentation/devicemanagement/acmecredentialsubjectaltnameobject) - Specifies the subject’s alternative name that the device requests for the certificate that the ACME server issues.

