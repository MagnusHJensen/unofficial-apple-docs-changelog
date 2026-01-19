# Sync the List of Courses

Get updates about the list of courses the server manages.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This sync service uses a cursor returned by the full course-roster service. It returns a list of all modifications (additions or deletions) made since the cursor date, for up to 7 days.

This service may return the same course more than once. You can identify duplicates by matching their `unique_identifier` values.

