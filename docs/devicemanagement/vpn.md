# VPN

The payload that configures a VPN.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0

## Discussion

Specify `com.apple.vpn.managed` as the payload type.

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
            <key>IPSec</key>
            <dict>
                <key>AuthenticationMethod</key>
                <string>SharedSecret</string>
                <key>LocalIdentifierType</key>
                <string>KeyID</string>
                <key>SharedSecret</key>
                <data>
                UVhCd2JHVXhNak1o
                </data>
            </dict>
            <key>IPv4</key>
            <dict>
                <key>OverridePrimary</key>
                <integer>0</integer>
            </dict>
            <key>PPP</key>
            <dict>
                <key>AuthName</key>
                <string>username</string>
                <key>AuthPassword</key>
                <string>password</string>
                <key>CommRemoteAddress</key>
                <string>vpn.example.com</string>
            </dict>
            <key>Proxies</key>
            <dict>
                <key>HTTPEnable</key>
                <integer>0</integer>
                <key>HTTPSEnable</key>
                <integer>0</integer>
            </dict>
            <key>UserDefinedName</key>
            <string>VPN Server</string>
            <key>VPNType</key>
            <string>L2TP</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myvpnmanagedprofile</string>
            <key>PayloadType</key>
            <string>com.apple.vpn.managed</string>
            <key>PayloadUUID</key>
            <string>74615F25-3B51-4386-A31B-ACB1F1094EF9</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>VPN</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>01E7F1C0-2DD0-4E36-82FF-EC6F29BB6C45</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [VPN.AlwaysOn](/documentation/devicemanagement/vpn/alwayson-data.dictionary) - The dictionary that contains IPSec settings.
- [VPN.DNS](/documentation/devicemanagement/vpn/dns-data.dictionary) - The dictionary to configure DNS settings for the VPN.
- [VPN.IKEv2](/documentation/devicemanagement/vpn/ikev2-data.dictionary) - The dictionary to use for an IKEv2 VPN type.
- [VPN.IPSec](/documentation/devicemanagement/vpn/ipsec-data.dictionary) - The dictionary to use for an IPSec VPN type.
- [VPN.IPv4](/documentation/devicemanagement/vpn/ipv4-data.dictionary) - The dictionary that contains IPV4 settings.
- [VPN.PPP](/documentation/devicemanagement/vpn/ppp-data.dictionary) - The dictionary that contains PPP settings.
- [VPN.Proxies](/documentation/devicemanagement/vpn/proxies-data.dictionary) - The dictionary that contains the Proxies settings.
- [VPN.TransparentProxy](/documentation/devicemanagement/vpn/transparentproxy-data.dictionary) - The dictionary to use for a transparent proxy VPN type.
- [VPN.VPN](/documentation/devicemanagement/vpn/vpn-data.dictionary) - The dictionary that contains VPN, IPSec, and IKEv2 settings.
- [VPN.VendorConfig](/documentation/devicemanagement/vpn/vendorconfig-data.dictionary) - The vendor-specific configuration dictionary.

