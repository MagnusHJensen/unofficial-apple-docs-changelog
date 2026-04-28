# ActiveDirectoryCertificate

The payload that configures Active Directory Certificate settings.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### AllowAllAppsAccess

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, gives apps access to the private key. Available in macOS 10.10 and later.

### CertificateAcquisitionMechanism

- **Type:** `string`
- **Required:** No

This value is most commonly `RPC`; if using web enrollment, use `HTTP`. Available in macOS 10.8 and later.

### CertificateAuthority

- **Type:** `string`
- **Required:** No

The name of the certificate authority (CA), which is determined from the common name (CN) of the Active Directory entry. Available in macOS 10.8 and later. Valid values:

- CN=
- CN=`Certification Authorities`
- CN=`Public Key Services`
- CN=`Services`
- CN=`Configuration`
- CN=

### CertificateRenewalTimeInterval

- **Type:** `integer`
- **Required:** No

The number of days in advance of certificate expiration that the notification center notifies the user.

### CertServer

- **Type:** `string`
- **Required:** Yes

The fully qualified host name of the CA.

### CertTemplate

- **Type:** `string`
- **Required:** Yes

The certificate template for your environment. The default user certificate value is `User`. The default computer certificate value is `Machine`.

### Description

- **Type:** `string`
- **Required:** No

A user-friendly description of the certification identity.

### EnableAutoRenewal

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the certificate obtained with this payload attempts auto-renewal. Auto-renewal can only be used with device Active Directory certificate payloads. Available in macOS 10.13.4 and later.

### KeyIsExtractable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system allows exporting the private key. Available in macOS 10.10 and later.

### Keysize

- **Type:** `integer`
- **Required:** No
- **Default:** `2048`

The RSA key size for the certificate signing request (CSR). Available in macOS 10.11 and later.

### PromptForCredentials

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prompts the user for credentials when is installs the profile. This key applies only to user certificates with the Manual Download profile delivery method. Omit this key for computer certificates. Available in macOS 10.8 and later.

## Discussion

Specify `com.apple.ADCertificate.managed` as the payload type.

To get a certificate from a Microsoft CA, follow the instructions at [Request a certificate from a Microsoft Certificate Authority](https://support.apple.com/en-us/HT204602).

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
            <key>CertServer</key>
            <string>server.example.com</string>
            <key>CertTemplate</key>
            <string>MachineUser</string>
            <key>CertificateAcquisitionMechanism</key>
            <string>RPC</string>
            <key>CertificateAuthority</key>
            <string>Example</string>
            <key>Description</key>
            <string>Active Directory Certificate</string>
            <key>PromptForCredentials</key>
            <false/>
            <key>PayloadIdentifier</key>
            <string>com.example.myADcertpayload</string>
            <key>PayloadType</key>
            <string>com.apple.myadcertificate.managed</string>
            <key>PayloadUUID</key>
            <string>59729e65-4c09-4fa1-b367-7a38cfd1b190</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Active Directory Certificate</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>com.apple.ADCertificate.managed</string>
    <key>PayloadUUID</key>
    <string>55a22a34-02b7-49d8-8116-ea95c3545261</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

