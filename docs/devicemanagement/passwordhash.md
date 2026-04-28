# PasswordHash

A dictionary that contains the password hash for the account.

**Platforms:** macOS 10.11, Device Assignment Services , VPP License Management 

## Properties

### SALTED-SHA512-PBKDF2

- **Type:** `PasswordHash.SALTED-SHA512-PBKDF2`
- **Required:** Yes

A dictionary that contains the `entropy`, `iterations`, and `salt` elements to create the password hash using the CommonCrypto libraries, or equivalent. Convert this dictionary to binary data before setting it as the value for the password hash.

## Topics

### Objects

- [PasswordHash.SALTED-SHA512-PBKDF2](/documentation/devicemanagement/passwordhash/salted-sha512-pbkdf2-data.dictionary) - A dictionary that contains the elements to create the password hash.

