# Books.Attributes

The attributes for a books resource.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### artistName

- **Type:** `string`
- **Required:** Yes

The name of the artist for this content.

### artwork

- **Type:** `Artwork`
- **Required:** Yes

The artwork for this content.

### genreNames

- **Type:** `[string]`
- **Required:** Yes

A list of genre names associated with this content.

### isbn

- **Type:** `string`
- **Required:** No

The ISBN of this book.

### name

- **Type:** `string`
- **Required:** Yes

The (potentially) censored name of the content.

### offers

- **Type:** `[Books.Attributes.Offers]`
- **Required:** Yes

A map of offer and asset information for the associated content.

### seriesInfo

- **Type:** `Books.Attributes.SeriesInfo`
- **Required:** No

Info about the series this book is a part of.

### taxExclusivePrices

- **Type:** `[Books.Attributes.TaxExclusivePrices]`
- **Required:** No

**** Tax-exclusive prices for this salable.

### taxRate

- **Type:** `number`
- **Required:** No

**** Tax rate for this salable for the current account.

### url

- **Type:** `string`
- **Required:** Yes

A canonical URL to the content that may be used for sharing or linking to the content externally.

### userRating

- **Type:** `Books.Attributes.UserRating`
- **Required:** Yes

User rating information for the content.

## Topics

### Related Objects

- [Books.Attributes.Offers](/documentation/devicemanagement/books/attributes-data.dictionary/offers-data.dictionary)
- [Books.Attributes.SeriesInfo](/documentation/devicemanagement/books/attributes-data.dictionary/seriesinfo-data.dictionary)
- [Books.Attributes.TaxExclusivePrices](/documentation/devicemanagement/books/attributes-data.dictionary/taxexclusiveprices-data.dictionary)
- [Books.Attributes.UserRating](/documentation/devicemanagement/books/attributes-data.dictionary/userrating-data.dictionary)

