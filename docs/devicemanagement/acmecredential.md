# ACMECredential

An ACME identity that the device generates.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

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

