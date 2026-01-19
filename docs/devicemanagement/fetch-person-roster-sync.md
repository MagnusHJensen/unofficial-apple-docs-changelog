# Sync the List of People

Get updates about the list of people the server manages.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This sync service uses a cursor returned by the full person-roster service. It returns a list of all modifications (additions or deletions) made since the cursor date, for up to 7 days.

This service may return the same person more than once. You can identify duplicates by matching their `unique_identifier` values.

