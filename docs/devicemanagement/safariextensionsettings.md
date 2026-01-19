# SafariExtensionSettings

The declaration to configure Safari Extensions.

**Platforms:** iOS 18.0, iPadOS 18.0, macOS 15.0, visionOS 26.0

## Discussion

Specify `com.apple.configuration.safari.extensions.settings` as the declaration type.

Safari supports the following values for `AllowedDomains` and `DeniedDomains`:

- A specific domain such as “example.com” or “www.example.com”.
- A wildcard domain that uses a single “*” character as a prefix for the domain, such as “*example.com”. This matches both the exact domain “example.com”, and any sub-domains such as “www.example.com”. It won’t match other domains with a similar string suffix such as “myexample.com”.
- A global wildcard specified as a single “*” character that matches any domain.

Safari determines whether a domain is allowed or denied using the following precedence rules:

1. A specific domain takes precedence over the global wildcard or a wildcard domain.
2. A wildcard domain takes precedence over the global wildcard.

If the same value appears in both `AllowedDomains` and `DeniedDomains`, Safari denies use of a matching domain.

The user can configure any domains not matched by the values in `AllowedDomains` or `DeniedDomains`.

### Examples

Give an extension access to only “example.com” and its sub-domains, and deny access to everywhere else.

```json
"AllowedDomains": ["*example.com"],
"DeniedDomains": ["*"]
```

Give an extension access to “example.com” and its sub-domains, without deny anywhere else. The user can make their own choice for other domains.

```json
"AllowedDomains": ["*example.com"]
```

Give an extension access to “example.com” and its sub-domains, but deny access to “private.example.com” or anywhere else.

```json
"AllowedDomains": ["*example.com"],
"DeniedDomains": ["private.example.com", "*"]
```

Give an extension access to “public.example.com”, but deny access to “example.com” or any other of its sub-domains. The user can make their own choice for other domains.

```json
"AllowedDomains": ["public.example.com"],
"DeniedDomains": ["*example.com"]
```

### Configuration availability

### Configuration examples

## Topics

### Objects

- [SafariExtensionSettingsManagedExtensionsObject](/documentation/devicemanagement/safariextensionsettingsmanagedextensionsobject) - The dictionary that defines the managed extension.

