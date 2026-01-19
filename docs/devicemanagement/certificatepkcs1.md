# CertificatePKCS1

The payload that configures a PKCS #1-formatted certificate.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 3.0

## Discussion

Specify `com.apple.security.pkcs1` as the payload type.

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
            <key>PayloadContent</key>
            <data>Exam</data>
            <key>PayloadDisplayName</key>
            <string>CertificatePKCS1</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mycertpkcs1payload</string>
            <key>PayloadType</key>
            <string>com.apple.security.pkcs1</string>
            <key>PayloadUUID</key>
            <string>72d2c549-2a97-4032-b818-d8ebf7cb88f2</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>CertificatePKCS1</string>
    <key>PayloadIdentifier</key>
    <string>com.example.profile</string>
    <key>PayloadType</key>
    <string>com.apple.security.pkcs1</string>
    <key>PayloadUUID</key>
    <string>d7d678c5-87ea-457d-82b9-25db21cd7868</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

