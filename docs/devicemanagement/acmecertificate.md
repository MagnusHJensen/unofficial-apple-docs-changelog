# ACMECertificate

The payload that configures Automated Certificate Management Environment (ACME) settings.

**Platforms:** iOS 16.0, iPadOS 16.0, macOS 13.1, tvOS 16.0, visionOS 1.0, watchOS 9.0

## Discussion

Specify `com.apple.security.acme` as the payload type.

Use this payload to specify how the device requests a client certificate from an Automated Certificate Management Environment (ACME) server. Other payloads can reference the resulting client identity by the payload’s `PayloadUUID`.

First the device generates an asymmetric key pair based upon the `KeyType`, `KeySize`, and `HardwareBound` fields. Then the device communicates with the ACME server. It requests a new order using the `ClientIdentifier` as the `permanent-identifier`. The ACME server responds with a challenge type of `device-attest-01`. If `Attest` is `true` the device requests an attestation of the key and device properties. Then it replies to the challenge with a WebAuthn attestation statement, and this contains the attestation if the device obtained one. The device submits a certificate signing request matching the key and containing the `ClientIdentifier`, `Subject`, `SubjectAltName`, `UsageFlags`, and `ExtendedKeyUsage` fields. The ACME server issues a certificate, and the device stores the resulting identity.

For details on the content of the attestation provided to the ACME server, see the documentation of the `DevicePropertiesAttestation` key in the [DeviceInformationResponse.QueryResponses](/documentation/devicemanagement/deviceinformationresponse/queryresponses-data.dictionary)response. In the attestation certificate the value of the freshness code OID is the SHA-256 hash of the `token` from the `device-attest-01` challenge.

### ACME attestation hardware support

The following table indicates which System on Chips (SoCs) support ACME attestation. If the Attest key is false or ignored, the ACME server does not receive an attestation.

### Profile availability

### Example Profile

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>ClientIdentifier</key>
            <string>this is an identifier</string>
            <key>ExtendedKeyUsage</key>
            <array>
                <string>1.3.6.1.5.5.7.3.2</string>
            </array>
            <key>HardwareBound</key>
            <true/>
            <key>KeySize</key>
            <integer>384</integer>
            <key>KeyType</key>
            <string>ECSECPrimeRandom</string>
            <key>UsageFlags</key>
            <integer>5</integer>
            <key>PayloadIdentifier</key>
            <string>com.example.myacmepayload</string>
            <key>PayloadType</key>
            <string>com.apple.security.acme</string>
            <key>PayloadUUID</key>
            <string>cbdc6238-feec-4171-878d-34e576bbb813</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
            <key>Subject</key>
            <array>
                <array>
                    <array>
                        <string>C</string>
                        <string>US</string>
                    </array>
                </array>
                <array>
                    <array>
                        <string>O</string>
                        <string>Example Inc.</string>
                    </array>
                </array>
                <array>
                    <array>
                        <string>1.2.840.113635.100.6.99999.99999</string>
                        <string>test custom OID value</string>
                    </array>
                </array>
            </array>
            <key>SubjectAltName</key>
            <dict>
                <key>dNSName</key>
                <string>site.example.com</string>
            </dict>
            <key>DirectoryURL</key>
            <string>https://acme.example.com/acme</string>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>ACME</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>ce876f81-abf0-46f9-9e68-9b3a7ede8097</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [ACMECertificate.SubjectAltName](/documentation/devicemanagement/acmecertificate/subjectaltname-data.dictionary) - The subject’s alternative name details.

