# LDAP

The payload that configures a Lightweight Directory Access Protocol (LDAP) account.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, visionOS 1.1

## Discussion

Specify `com.apple.ldap.account` as the payload type.

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
            <key>LDAPAccountDescription</key>
            <string>Company LDAP Account</string>
            <key>LDAPAccountHostName</key>
            <string>com.apple.ldap.account</string>
            <key>LDAPAccountUseSSL</key>
            <true/>
            <key>LDAPAccountUserName</key>
            <string>JuanChavez4</string>
            <key>LDAPSearchSettings</key>
            <array>
                <dict>
                    <key>LDAPSearchSettingDescription</key>
                    <string>My Search</string>
                    <key>LDAPSearchSettingScope</key>
                    <string>LDAPSearchSettingScopeSubtree</string>
                    <key>LDAPSearchSettingSearchBase</key>
                    <string>o=My Company,ou=My Department</string>
                </dict>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.myldappayload</string>
            <key>PayloadType</key>
            <string>com.apple.ldap.account</string>
            <key>PayloadUUID</key>
            <string>7f846724-1bf7-4501-b8cd-ce7026e95280</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>LDAP</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>c5208028-7e96-4669-8d83-4fbbeb48845a</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [LDAP.LDAPSearchSettingsItem](/documentation/devicemanagement/ldap/ldapsearchsettingsitem) - An array of search settings dictionaries.

