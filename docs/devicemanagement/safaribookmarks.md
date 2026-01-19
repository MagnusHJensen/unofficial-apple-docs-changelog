# SafariBookmarks

The declaration to configure managed bookmarks in Safari.

**Platforms:** iOS 26.0, iPadOS 26.0, macOS 26.0, visionOS 26.0

## Discussion

Specify `com.apple.configuration.safari.bookmarks` as the declaration type.

### Configuration availability

### Configuration example

This configuration applies a set of managed bookmarks: two bookmarks and one bookmarks folder, which contains two more bookmarks.

```json
{
    "Type": "com.apple.configuration.safari.bookmarks",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "ManagedBookmarks": [
            {
                "GroupIdentifier": "Group1",
                "Title": "Company Bookmarks",
                "Bookmarks": [
                    {
                        "Title": "Public Site",
                        "URL": "https://www.example.com"
                    },
                    {
                        "Title": "Store",
                        "URL": "https://store.example.com"
                    },
                    {
                        "Title": "Internal Developers",
                        "Folder": [
                            {
                                "Title": "Developers",
                                "URL": "https://developer.example.com"
                            },
                            {
                                "Title": "Repositories",
                                "URL": "https://repo.example.com"
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
```

## Topics

### Objects

- [SafariBookmarksBookmarkGroupObject](/documentation/devicemanagement/safaribookmarksbookmarkgroupobject) - A group of managed bookmarks.

