# ManagedMenuExtras

The payload that configures menu extras.

**Platforms:** macOS 10.7

## Properties

### AirPort.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the AirPort menu extra.

### Battery.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Battery menu extra.

### Bluetooth.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Bluetooth menu extra.

### Clock.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Clock menu extra.

### CPU.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the CPU menu extra.

### delaySeconds

- **Type:** `number`
- **Required:** No
- **Default:** `2.5`

The number of seconds to delay after login before adding or removing menu extras. If the delay is too short, the menu extras don’t appear, or disappear from the menu bar.

### Displays.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Displays menu extra.

### Eject.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Eject menu extra.

### Fax.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Fax menu extra.

### HomeSync.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the HomeSync menu extra.

### iChat.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the iChat menu extra.

### Ink.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Ink menu extra.

### IrDA.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the IrDA menu extra.

### maxWaitSeconds

- **Type:** `number`
- **Required:** No
- **Default:** `20`

The maximum wait, in seconds, for all menu extras to be added or removed.

### PCCard.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the PCCard menu extra.

### PPP.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the PPP menu extra.

### PPPoE.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the PPPoE menu extra.

### RemoteDesktop.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Remote Desktop menu extra.

### Script Menu.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Script menu extra.

### Spaces.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Spaces menu extra.

### Sync.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Sync menu extra.

### TextInput.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Text Input menu extra.

### TimeMachine.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the TimeMachine menu extra.

### UniversalAccess.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Universal Access menu extra.

### User.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the User menu extra.

### Volume.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the Volume menu extra.

### VPN.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the VPN menu extra.

### WWAN.menu

- **Type:** `boolean`
- **Required:** No

If `true`, enables the WWAN menu extra.

## Discussion

Specify `com.apple.mcxMenuExtras` as the payload type.

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
            <key>Battery.menu</key>
            <false/>
            <key>delaySeconds</key>
            <integer>30</integer>
            <key>maxWaitSeconds</key>
            <integer>60</integer>
            <key>PayloadIdentifier</key>
            <string>com.example.mymanagedmenuextraspayload</string>
            <key>PayloadType</key>
            <string>com.apple.mcxMenuExtras</string>
            <key>PayloadUUID</key>
            <string>93bd5b68-0141-4055-aaaf-a6cebc1cfeeb</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Menu Extras</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>dc2618ce-736c-4af7-b652-f9cdf3eb9ce4</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

