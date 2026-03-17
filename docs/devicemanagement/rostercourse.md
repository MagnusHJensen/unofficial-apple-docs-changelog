# RosterCourse

A course’s properties and their values.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### course_number

- **Type:** `string`
- **Required:** No

The number of the course.

### name

- **Type:** `string`
- **Required:** No

The course name. Maximum length 1024 UTF-8 characters.

### source

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LDAP`, `SIS`, `CSV`, `MANUAL`, `SYSTEM`, `ENROLLMENT`, `DEP`, `FORMULA_GENERATED`, `SFTP`

The data source where the class was created.

### source_system_identifier

- **Type:** `string`
- **Required:** No

The identifier configured by organization for the course. Maximum length is 256 UTF-8 characters. Value can be null.

### unique_identifier

- **Type:** `string`
- **Required:** No

The unique identifier for the course. Maximum length 256 UTF-8 characters.

### op_date

- **Type:** `string`
- **Required:** No

The time stamp, in iSO 8601 format, when the course was added, updated, or deleted.

### status

- **Type:** `string`
- **Required:** No

The status for the course.

