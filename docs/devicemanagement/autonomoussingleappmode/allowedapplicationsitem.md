# AutonomousSingleAppMode.AllowedApplicationsItem

A dictionary that specifies an app that can be granted access to the Accessibilty APIs.

**Platforms:** macOS 10.13.4, Device Assignment Services , VPP License Management 

## Properties

### BundleIdentifier

- **Type:** `string`
- **Required:** Yes

The unique bundle identifier. If two dictionaries contain the same `BundleIdentifier` value but a different `TeamIdentifier` value, an error occurs and the profile won’t be installed.

### TeamIdentifier

- **Type:** `string`
- **Required:** Yes

The developer’s team identifier that the system used when it signed the app.

