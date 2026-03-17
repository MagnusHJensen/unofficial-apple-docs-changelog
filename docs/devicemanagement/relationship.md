# Relationship

A to-one or to-many relationship from one resource object to others.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### data

- **Type:** `[Resource]`
- **Required:** Yes

A paginated collection of resources in the relationship.

### href

- **Type:** `string`
- **Required:** No

A relative location to fetch the relationship, if it may be fetched directly.

### meta

- **Type:** `Relationship.Meta`
- **Required:** No

Contextual data about the relationship.

### next

- **Type:** `string`
- **Required:** No

A relative cursor to fetch the next paginated collection of resources in the relationship if more exist.

## Topics

### Related Objects

- [Relationship.Meta](/documentation/devicemanagement/relationship/meta-data.dictionary)

