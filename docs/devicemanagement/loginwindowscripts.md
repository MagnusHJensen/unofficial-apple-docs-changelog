# LoginWindowScripts

The payload that configures scripts to run at login and logout.

**Platforms:** macOS 10.7

## Discussion

Specify `com.apple.mcxloginscripts` as the payload type.

The MCX login and logout managed-scripts payload contains information about executable scripts that can run at user login and logout. To use this payload, set `EnableMCXLoginScripts` to `true` in `/var/root/Library/Preferences/com.apple.loginwindow.plist`; otherwise, the system ignores this payload.

`Loginwindow` uses the `LoginHook` and `LogoutHook` string keys in `/var/root/Library/Preferences/com.apple.loginwindow.plist` to indicate a path to the executable script files, which run during user login and logout. The system passes the current user name as an argument to the file.

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
            <key>loginscripts</key>
            <array>
                <dict>
                    <key>filedata</key>
                    <data>ExampleD</data>
                    <key>filename</key>
                    <string>ping.sh</string>
                </dict>
            </array>
            <key>skipLoginHook</key>
            <true/>
            <key>skipLogoutHook</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myloginwindowscriptspayload</string>
            <key>PayloadType</key>
            <string>com.apple.mcxloginscripts</string>
            <key>PayloadUUID</key>
            <string>2bcd5563-f44d-4f74-b706-050a628c0caf</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Login Window Scripts</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>5cbf617d-16a5-4564-ba2c-728fd7f7d732</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [LoginWindowScripts.ScriptsItems](/documentation/devicemanagement/loginwindowscripts/scriptsitems) - A dictionary of login scripts.

