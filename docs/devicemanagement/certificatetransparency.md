# CertificateTransparency

The payload that configures certificate transparency enforcement.

**Platforms:** iOS 12.1.1, iPadOS 12.1.1, Mac Catalyst 12.1.1, macOS 10.14.2, tvOS 12.1.1, visionOS 1.0, watchOS 5.1.1, Device Assignment Services , VPP License Management 

## Properties

### DisabledForCerts

- **Type:** `[CertificateTransparency.SubjectPublicKeyInfoHashDict]`
- **Required:** No

An array of certificates for which certificate transparency is disabled. One of the following conditions needs to be met to disable certificate transparency enforcement when this policy is set:

- The hash is of the server certificate’s `subjectPublicKeyInfo`.
- The hash is of a `subjectPublicKeyInfo` that appears in a CA certificate in the certificate chain; the CA certificate is constrained through the X.509v3 `nameConstraints` extension. One or more `directoryName` `nameConstraints` are present in the `permittedSubtrees`, and the `directoryName` contains an `organizationName` attribute.
- The hash is of a `subjectPublicKeyInfo` that appears in a CA certificate in the certificate chain. The CA certificate has one or more `organizationName` attributes in the certificate `Subject`, and the server’s certificate contains the same number of `organizationName` attributes, in the same order, and with byte-for-byte identical values.

### DisabledForDomains

- **Type:** `[string]`
- **Required:** No

An array of strings that represent the domains to exclude from certificate transparency enforcement. The system supports using a leading period (`.`) to signify subdomains. However, the system doesn’t support wildcards. If you include a leading period, the domain can’t be a top-level domain, such as `.com` and `.co.uk`.

## Discussion

Specify `com.apple.security.certificatetransparency` as the payload type.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>DisabledForCerts</key>
            <array>
                <dict>
                    <key>Algorithm</key>
                    <string>sha256</string>
                    <key>Hash</key>
                    <data>AAolBg==</data>
                </dict>
            </array>
            <key>DisabledForDomains</key>
            <array>
                <string>example.com</string>
            </array>
            <key>PayloadDescription</key>
            <string>Configures Certificate Transparency</string>
            <key>PayloadDisplayName</key>
            <string>Domains</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mycerttransparencypayload</string>
            <key>PayloadType</key>
            <string>com.apple.security.certificatetransparency</string>
            <key>PayloadUUID</key>
            <string>0ae54b4a-cbf5-4323-8524-262a3cc4b733</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Certificate Transparancy</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>a54d018e-864e-4ec9-9638-85fc50410ae3</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [CertificateTransparency.SubjectPublicKeyInfoHashDict](/documentation/devicemanagement/certificatetransparency/subjectpublickeyinfohashdict) - A dictionary of hashed public keys.

