# AppLock.App.Options

The dictionary of options to set for the app.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, tvOS 10.2

## Properties

### DisableAutoLock

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device doesn’t automatically go to sleep after an idle period.

### DisableDeviceRotation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables device rotation sensing.

Available: iOS 7+ | iPadOS 7+

### DisableRingerSwitch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the ringer switch. When disabled, the ringer behavior depends on what position the switch was in when it was first disabled.

Available: iOS 7+ | iPadOS 7+

### DisableSleepWakeButton

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the sleep/wake button.

Available: iOS 7+ | iPadOS 7+

### DisableTouch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the touch screen. In tvOS, it disables the touch surface on the Apple TV Remote.

### DisableVolumeButtons

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the volume buttons.

Available: iOS 7+ | iPadOS 7+

### EnableAssistiveTouch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables AssistiveTouch.

Available: iOS 7+ | iPadOS 7+

### EnableInvertColors

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Invert Colors.

### EnableMonoAudio

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Mono Audio.

Available: iOS 7+ | iPadOS 7+

### EnableSpeakSelection

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Speak Selection.

Available: iOS 7+ | iPadOS 7+

### EnableVoiceControl

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Voice Control.

Available: iOS 13+ | iPadOS 13+

### EnableVoiceOver

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables VoiceOver.

### EnableZoom

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Zoom.

