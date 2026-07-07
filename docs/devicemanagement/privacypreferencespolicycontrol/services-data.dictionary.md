# PrivacyPreferencesPolicyControl.Services

The privacy policy control services dictionary that controls access on a per app basis.

**Platforms:** macOS 10.14

## Properties

### Accessibility

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for the app via the Accessibility subsystem. This profile deprecated its ability to grant access as of macOS 26.2, and removes that ability in macOS 27.0.

Deprecated: use the `Privacy` key in the declarative management `com.apple.configuration.app-settings` configuration.

Deprecated: macOS 27+

### AddressBook

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for contact information managed by the Contacts.app.

### AppleEvents

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for the app sending restricted AppleEvents to another process.

### BluetoothAlways

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for the app to access Bluetooth devices.

Deprecated: use the `Privacy` key in the declarative management `com.apple.configuration.app-settings` configuration.

Available: macOS 11+
Deprecated: macOS 27+

### Calendar

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for calendar information managed by the Calendar.app.

### Camera

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

A system camera. A profile can’t grant access to the camera; it can only deny it.

Deprecated: macOS 27+

### FileProviderPresence

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows a File Provider application to know when the user is using files managed by the File Provider.

Available: macOS 10.15+

### ListenEvent

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to use CoreGraphics and HID APIs to listen to (receive) CGEvents and HID events from all processes. A profile can’t grant access to these events; it can only deny it.

Available: macOS 10.15+

### MediaLibrary

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access Apple Music, music and video activity, and the media library.

Available: macOS 10.15+

### Microphone

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

A system microphone. A profile can’t grant access to the microphone; it can only deny it.

Deprecated: macOS 27+

### Photos

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

The pictures managed by the Photos app in `~/Pictures/.photoslibrary`.

### PostEvent

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for the application to use CoreGraphics APIs to send CGEvents to the system event stream.

### Reminders

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for reminders information managed by the Reminders app.

### ScreenCapture

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to capture (read) the contents of the system display. A profile can’t grant access to the contents; it can only deny it.

Available: macOS 10.15+

### SpeechRecognition

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to use the system Speech Recognition facility and to send speech data to Apple.

Deprecated: use the `Privacy` key in the declarative management `com.apple.configuration.app-settings` configuration.

Available: macOS 10.15+
Deprecated: macOS 27+

### SystemPolicyAllFiles

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application access to all protected files, including system administration files.

### SystemPolicyAppBundles

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to update or delete other apps.

Available: macOS 13+

### SystemPolicyAppData

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Specifies the policies for the app to access the data of other apps.

Available: macOS 14+

### SystemPolicyDesktopFolder

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access files in the user’s Desktop folder.

Available: macOS 10.15+

### SystemPolicyDocumentsFolder

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access files in the user’s Documents folder.

Available: macOS 10.15+

### SystemPolicyDownloadsFolder

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access files in the user’s Downloads folder.

Available: macOS 10.15+

### SystemPolicyNetworkVolumes

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access files on network volumes.

Available: macOS 10.15+

### SystemPolicyRemovableVolumes

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application to access files on removable volumes.

Available: macOS 10.15+

### SystemPolicySysAdminFiles

- **Type:** `[PrivacyPreferencesPolicyControl.Services.Identity]`
- **Required:** No

Allows the application access to some files used in system administration.

## Topics

### Objects

- [PrivacyPreferencesPolicyControl.Services.Identity](/documentation/devicemanagement/privacypreferencespolicycontrol/services-data.dictionary/identity) - A dictionary listing apps and the privacy policy to apply to them.

