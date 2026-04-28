# RestrictionsResponse.GlobalRestrictions

A dictionary that contains the global restrictions in effect.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### intersection

- **Type:** `RestrictionsResponse.GlobalRestrictions.Intersection`
- **Required:** No

A dictionary of intersected profile restrictions. Intersected restrictions indicate that new restrictions can only reduce the number of strings in the set.

### restrictedBool

- **Type:** `RestrictionsResponse.GlobalRestrictions.RestrictedBool`
- **Required:** No

A dictionary of Boolean profile restrictions.

### restrictedValue

- **Type:** `RestrictionsResponse.GlobalRestrictions.RestrictedValue`
- **Required:** No

A dictionary of numeric profile restrictions.

### union

- **Type:** `RestrictionsResponse.GlobalRestrictions.Union`
- **Required:** No

A dictionary of unioned profile restrictions. Unioned restrictions indicate that new restrictions can add to the set.

## Topics

### Objects

- [RestrictionsResponse.GlobalRestrictions.Intersection](/documentation/devicemanagement/restrictionsresponse/globalrestrictions-data.dictionary/intersection-data.dictionary) - A dictionary that contains intersected restrictions.
- [RestrictionsResponse.GlobalRestrictions.RestrictedBool](/documentation/devicemanagement/restrictionsresponse/globalrestrictions-data.dictionary/restrictedbool-data.dictionary) - A dictionary that contains Boolean restrictions.
- [RestrictionsResponse.GlobalRestrictions.RestrictedValue](/documentation/devicemanagement/restrictionsresponse/globalrestrictions-data.dictionary/restrictedvalue-data.dictionary) - A dictionary that contains numeric restrictions.
- [RestrictionsResponse.GlobalRestrictions.Union](/documentation/devicemanagement/restrictionsresponse/globalrestrictions-data.dictionary/union-data.dictionary) - A dictionary that contains unioned restrictions.

