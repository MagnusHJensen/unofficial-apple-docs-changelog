# SCEP

The payload that configures Simple Certificate Enrollment Protocol (SCEP) settings.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 3.0

## Discussion

Specify `com.apple.security.scep` as the payload type.

A SCEP payload automates the request of a client certificate from a SCEP server, as described in [Over-the-Air Profile Delivery and Configuration](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/iPhoneOTAConfiguration/Introduction/Introduction.html).

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
            <key>PayloadContent</key>
            <dict>
                <key>Challenge</key>
                <string>Example</string>
                <key>Key Type</key>
                <string>RSA</string>
                <key>Key Usage</key>
                <integer>5</integer>
                <key>Keysize</key>
                <integer>0</integer>
                <key>Name</key>
                <string>example.org</string>
                <key>Subject</key>
                <array>
                    <array>
                        <array>
                            <string>C</string>
                            <string>US</string>
                        </array>
                        <array>
                            <string>O</string>
                            <string>Example Inc.</string>
                        </array>
                        <array>
                            <string>CN</string>
                            <string>foo</string>
                        </array>
                        <array>
                            <string>1.2.5.3</string>
                            <string>bar</string>
                        </array>
                    </array>
                </array>
                <key>SubjectAltName</key>
                <dict>
                    <key>dNSName</key>
                    <string>example.com</string>
                    <key>ntPrincipalName</key>
                    <string>hostname.example.com</string>
                </dict>
                <key>URL</key>
                <string>https://hostname.example.com/</string>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.myscepcertpayload</string>
            <key>PayloadType</key>
            <string>com.apple.security.scep</string>
            <key>PayloadUUID</key>
            <string>c0264fd7-1d89-4385-8806-759fbe78a622</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>SCEP Certificate</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>bc0328a9-c199-4572-9e5e-ed59a73454fa</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [SCEP.PayloadContent](/documentation/devicemanagement/scep/payloadcontent-data.dictionary) - The SCEP dictionary.

