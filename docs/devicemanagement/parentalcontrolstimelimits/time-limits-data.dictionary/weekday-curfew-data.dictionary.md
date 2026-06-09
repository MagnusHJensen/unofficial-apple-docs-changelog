# ParentalControlsTimeLimits.Time-limits.Weekday-curfew

The weekday curfew dictionary.

**Platforms:** macOS 10.7

## Properties

### enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, enable these settings.

### end

- **Type:** `string`
- **Required:** No

The curfew end time, in the format `%d:%d:%d`.

### rangeType

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `0`, `1`

The type of day range, which has the following possible values:

- `0`: Weekday
- `1`: Weekend

### secondsPerDay

- **Type:** `integer`
- **Required:** No

The allowance for that day, in seconds.

### start

- **Type:** `string`
- **Required:** No

The curfew start time, in the format ‘%d:%d:%d’.

