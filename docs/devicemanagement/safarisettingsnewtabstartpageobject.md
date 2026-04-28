# SafariSettingsNewTabStartPageObject

Sets the start page for new tabs in Safari.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 26.0, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### ExtensionIdentifier

- **Type:** `string`
- **Required:** No

The composed identifier of the extension that provides the start page. The required format is “Identifier (TeamIdentifier)”, for example “com.example.app (ABCD1234)”. Required when setting `PageType` to `Extension`.

### HomepageURL

- **Type:** `string`
- **Required:** No

The URL of the homepage which needs to start with `https://` or `http://`. Required when setting `PageType` to `Home`.

### PageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Start`, `Home`, `Extension`

Sets the start page type in Safari:

- `Start` - Safari uses the default start page. Safari disables the Homepage.
- `Home` - Safari uses the page specified by `HomepageURL`, and Safari also sets that as the Homepage.
- `Extension` - Safari uses the page specified by the Safari extension whose identifier is `ExtensionIdentifier`. Safari disables the Homepage.

