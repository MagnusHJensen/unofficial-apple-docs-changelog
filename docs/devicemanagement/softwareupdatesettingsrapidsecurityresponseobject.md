# SoftwareUpdateSettingsRapidSecurityResponseObject

The object that configures Background Security Improvement settings.

**Platforms:** iOS 18.0, iPadOS 18.0, macOS 15.0

## Properties

### Enable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `false`, Background Security Improvements aren’t offered for user installation. The system can still install Background Security Improvements with `com.apple.configuration.softwareupdate.enforcement.specific` configurations.

If set to `true`, the system offers Background Security Improvements to the user.

### EnableRollback

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If set to `false`, the system doesn’t offer Background Security Improvement rollbacks to the user.

If set to `true`, the system offers Background Security Improvement rollbacks to the user.

