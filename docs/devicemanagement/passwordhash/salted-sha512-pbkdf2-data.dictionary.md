# PasswordHash.SALTED-SHA512-PBKDF2

A dictionary that contains the elements to create the password hash.

**Platforms:** macOS 10.11

## Properties

### entropy

- **Type:** `data`
- **Required:** Yes

The derived key from the password hash; for example, from `CCKeyDerivationPBKDF()`.

### iterations

- **Type:** `integer`
- **Required:** Yes

The number of iterations; for example, from `CCCalibratePBKDF()` using a minimum hash time of 100 milliseconds, or if unknown, a number in the range of 20,000 to 40,000 iterations.

### salt

- **Type:** `data`
- **Required:** Yes

The 32-byte randomized data; for example, from `CCRandomCopyBytes()`.

