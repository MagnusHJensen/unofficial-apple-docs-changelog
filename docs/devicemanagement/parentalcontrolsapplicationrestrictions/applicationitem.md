# ParentalControlsApplicationRestrictions.ApplicationItem

A dictionary defining an app for parental control.

**Platforms:** macOS 10.15

## Properties

### appID

- **Type:** `data`
- **Required:** Yes

The identifier of the app. Obtain this value from the Security framework using [SecCodeCopyDesignatedRequirement(_:_:_:)](/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)).

Deprecated: macOS 27+

### bundleID

- **Type:** `string`
- **Required:** Yes

The bundle ID of the app.

Deprecated: macOS 27+

### detachedSignature

- **Type:** `data`
- **Required:** No

The signature for an unsigned binary.

Deprecated: macOS 27+

### disabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this app isn’t added to the allow list.

Deprecated: macOS 27+

### displayName

- **Type:** `string`
- **Required:** No

The name used for display purposes.

Deprecated: macOS 27+

### subApps

- **Type:** `[ParentalControlsApplicationRestrictions.ApplicationItem]`
- **Required:** No

An array of nested helper applications.

Deprecated: macOS 27+

