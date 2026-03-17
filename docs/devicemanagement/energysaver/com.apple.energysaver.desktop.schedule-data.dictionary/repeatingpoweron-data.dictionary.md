# EnergySaver.Com.apple.EnergySaver.desktop.Schedule.RepeatingPowerOn

The triggers for turning the device on.

**Platforms:** macOS 10.7

## Properties

### eventtype

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `wake`, `poweron`, `wakepoweron`, `sleep`, `shutdown`, `restart`

The type of action defined by this schedule.

### time

- **Type:** `integer`
- **Required:** No

The time, in minutes, since midnight.

### weekdays

- **Type:** `integer`
- **Required:** No

One or more days of the week in an unsigned integer bitmap:

- `1` = Mon
- `2` = Tue
- `4` = Wed
- `8` = Thu
- `16` = Fri
- `32` = Sat
- `64` = Sun

