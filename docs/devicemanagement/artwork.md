# Artwork

An object that represents artwork.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### assetToken

- **Type:** `string`
- **Required:** No



### bgColor

- **Type:** `string`
- **Required:** No

The average background color for the image to use as a placeholder.

### gradient

- **Type:** `Artwork.Gradient`
- **Required:** No

Gradient information. If absent, use existing gradient behavior. If present, but the body is an empty object {}, then gradient is disabled. This overrides any setting in the editorial item.

### hasP3

- **Type:** `boolean`
- **Required:** No



### height

- **Type:** `number`
- **Required:** Yes

The largest height available for the source image.

### pictureFileType

- **Type:** `string`
- **Required:** No



### supportsLayeredImage

- **Type:** `boolean`
- **Required:** No



### textColor1

- **Type:** `string`
- **Required:** No

A primary text color appropriate to use with the artwork.

### textColor2

- **Type:** `string`
- **Required:** No

A secondary text color appropriate to use with the artwork.

### textColor3

- **Type:** `string`
- **Required:** No

An auxiliary text color appropriate to use with the artwork.

### textColor4

- **Type:** `string`
- **Required:** No

An auxiliary text color appropriate to use with the artwork.

### url

- **Type:** `string`
- **Required:** Yes

The URL of the image. May be templated to allow customizing the height {h}, width {w}, and crop code {c}.

### width

- **Type:** `number`
- **Required:** Yes

The largest width available for the source image.

## Topics

### Related Objects

- [Artwork.Gradient](/documentation/devicemanagement/artwork/gradient-data.dictionary) - An object that represents the properties of a color gradient for an artwork object.

