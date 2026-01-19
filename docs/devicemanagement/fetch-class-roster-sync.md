# Sync the List of Classes

Get updates about the list of classes the server manages.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This sync service uses a cursor that is returned by the full class-roster service. It returns a list of all modifications (additions or deletions) made since the cursor date, for up to 7 days.

This service may return the same class more than once. You can identify duplicates by matching their `unique_identifier` values.

