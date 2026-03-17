# EnergySaver

The payload that configures Energy Saver settings.

**Platforms:** macOS 10.7

## Properties

### com.apple.EnergySaver.desktop.ACPower

- **Type:** `EnergySaver.Com.apple.EnergySaver.desktop.ACPower`
- **Required:** No

The settings for a desktop computer.

### com.apple.EnergySaver.desktop.Schedule

- **Type:** `EnergySaver.Com.apple.EnergySaver.desktop.Schedule`
- **Required:** No

The schedule for turning a computer on and off.

### com.apple.EnergySaver.portable.ACPower

- **Type:** `EnergySaver.Com.apple.EnergySaver.portable.ACPower`
- **Required:** No

The settings for a laptop computer using AC power.

### com.apple.EnergySaver.portable.BatteryPower

- **Type:** `EnergySaver.Com.apple.EnergySaver.portable.BatteryPower`
- **Required:** No

The settings for a laptop computer using battery power.

### SleepDisabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, disables sleep.

## Discussion

Specify `com.apple.MCX` as the payload type.

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
            <key>com.example.EnergySaver.desktop.ACPower</key>
            <dict>
                <key>Automatic Restart On Power Loss</key>
                <integer>1</integer>
                <key>Disk Sleep Timer</key>
                <integer>180</integer>
                <key>Display Sleep Timer</key>
                <integer>60</integer>
                <key>Sleep On Power Button</key>
                <integer>0</integer>
                <key>System Sleep Timer</key>
                <integer>120</integer>
                <key>Wake On LAN</key>
                <integer>1</integer>
            </dict>
            <key>com.apple.EnergySaver.desktop.ACPower-ProfileNumber</key>
            <integer>-1</integer>
            <key>com.apple.EnergySaver.desktop.Schedule</key>
            <dict>
                <key>RepeatingPowerOff</key>
                <dict>
                    <key>eventtype</key>
                    <string>sleep</string>
                    <key>time</key>
                    <integer>0</integer>
                    <key>weekdays</key>
                    <integer>31</integer>
                </dict>
                <key>RepeatingPowerOn</key>
                <dict>
                    <key>eventtype</key>
                    <string>wakepoweron</string>
                    <key>time</key>
                    <integer>0</integer>
                    <key>weekdays</key>
                    <integer>31</integer>
                </dict>
            </dict>
            <key>com.apple.EnergySaver.portable.ACPower</key>
            <dict>
                <key>Automatic Restart On Power Loss</key>
                <integer>1</integer>
                <key>Disk Sleep Timer</key>
                <integer>180</integer>
                <key>Display Sleep Timer</key>
                <integer>5</integer>
                <key>System Sleep Timer</key>
                <integer>10</integer>
                <key>Wake On LAN</key>
                <integer>1</integer>
            </dict>
            <key>com.apple.EnergySaver.portable.ACPower-ProfileNumber</key>
            <integer>-1</integer>
            <key>com.apple.EnergySaver.portable.BatteryPower</key>
            <dict>
                <key>Automatic Restart On Power Loss</key>
                <integer>1</integer>
                <key>Disk Sleep Timer</key>
                <integer>180</integer>
                <key>Display Sleep Timer</key>
                <integer>5</integer>
                <key>System Sleep Timer</key>
                <integer>10</integer>
                <key>Wake On LAN</key>
                <integer>1</integer>
            </dict>
            <key>com.apple.EnergySaver.portable.BatteryPower-ProfileNumber</key>
            <integer>-1</integer>
            <key>PayloadIdentifier</key>
            <string>com.example.myenergysaverpayload</string>
            <key>PayloadType</key>
            <string>com.apple.MCX</string>
            <key>PayloadUUID</key>
            <string>e66c94e5-f059-404e-9188-ff152cb80c64</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>EnergySaver</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>2627f09c-a6ea-4a2d-94b3-f6df28b3c6af</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [EnergySaver.Com.apple.EnergySaver.desktop.ACPower](/documentation/devicemanagement/energysaver/com.apple.energysaver.desktop.acpower-data.dictionary) - The desktop AC power Energy Saver settings.
- [EnergySaver.Com.apple.EnergySaver.desktop.Schedule](/documentation/devicemanagement/energysaver/com.apple.energysaver.desktop.schedule-data.dictionary) - The schedule for turning the device on or off.
- [EnergySaver.Com.apple.EnergySaver.portable.ACPower](/documentation/devicemanagement/energysaver/com.apple.energysaver.portable.acpower-data.dictionary) - The laptop AC power Energy Saver settings.
- [EnergySaver.Com.apple.EnergySaver.portable.BatteryPower](/documentation/devicemanagement/energysaver/com.apple.energysaver.portable.batterypower-data.dictionary) - The laptop battery power Energy Saver settings.

