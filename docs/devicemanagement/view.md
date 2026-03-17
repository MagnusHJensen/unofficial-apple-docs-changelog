# View

A view for the resource.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### attributes

- **Type:** `View.Attributes`
- **Required:** No

The attribute metadata for the view.

### data

- **Type:** `[Resource]`
- **Required:** Yes

A paginated collection of resources in the view.

### href

- **Type:** `string`
- **Required:** No

A relative location to fetch the view, if it’s directly fetchable.

### meta

- **Type:** `View.Meta`
- **Required:** No

Contextual data about the view.

### next

- **Type:** `string`
- **Required:** No

A relative cursor to fetch the next paginated collection of resources in the view if more exist.

## Topics

### Objects

- [View.Attributes](/documentation/devicemanagement/view/attributes-data.dictionary) - The attribute metadata for the view.
- [View.Meta](/documentation/devicemanagement/view/meta-data.dictionary) - Contextual data about the view.

