# Artwork.Gradient

An object that represents the properties of a color gradient for an artwork object.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### color

- **Type:** `string`
- **Required:** No

An ‘rrggbb’ field that defines the color of the gradient. The textColor* fields are not used to define the gradient color when this field is present.

### y1

- **Type:** `number`
- **Required:** No

A height between 0 and 1 that indicates where the gradient starts, measured from the bottom. If y1 is not present, it defaults to 0.

### y2

- **Type:** `number`
- **Required:** No

A height between 0 and 1 that indicates where the gradient stops, measured from the bottom of the image.

