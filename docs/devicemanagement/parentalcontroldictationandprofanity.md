# ParentalControlDictationAndProfanity

The payload that configures parental control for dictation and profanity.

**Platforms:** macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### Ironwood Allowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables dictation. Use `allowDictation` in Restrictions instead.

### Profanity Allowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, suppresses profanity. Use `forceAssistantProfanityFilter` in Restrictions instead.

## Discussion

Specify `com.apple.ironwood.support` as the payload type.

### Profile availability

