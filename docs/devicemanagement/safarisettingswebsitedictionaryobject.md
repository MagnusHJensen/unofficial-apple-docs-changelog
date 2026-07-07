# SafariSettingsWebsiteDictionaryObject

The dictionary that defines the website privacy permission defaults. Each key represents a website.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

## Properties

### Camera

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `None`, `Allow`

Controls whether a website privacy permission default is set.

- `None`: Safari sets no website privacy permission default for use of the camera.
- `Allow`: Safari sets the website privacy permission default to allow use of the camera.

### Microphone

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `None`, `Allow`

Controls whether a website privacy permission default is set.

- `None`: Safari sets no website privacy permission default for use of the microphone.
- `Allow`: Safari sets the website privacy permission default to allow use of the microphone.

### OrganizationJustification

- **Type:** `string`
- **Required:** Yes

Text that clearly explains to the Safari user the reason why the organization requires these website privacy permission defaults. Safari includes this text in the permission consent prompt it displays when it first displays the website.

