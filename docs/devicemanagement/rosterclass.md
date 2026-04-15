# RosterClass

A class’s properties and their values.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### class_display_name

- **Type:** `string`
- **Required:** No

The Class Nickname in Apple School Manager. This name appears in the Schoolwork app and is modifiable by teachers within Schoolwork or by admins within Apple School Manager. Use this field for Classroom and Shared iPad class configurations.

Available in X-Protocol Version 5 and later.

### class_icon_identifier

- **Type:** `string`
- **Required:** No

The class icon selected by teachers when they create or update the class via Schoolwork.

Available in X-Protocol Version 5 and later.

### class_number

- **Type:** `string`
- **Required:** No

The class number. Maximum string length is 256 UTF-8 characters.

Available in X-Protocol Version 4 and later.

### course

- **Type:** `BaseRosterCourse`
- **Required:** No

The base course for this course.

### instructor_unique_identifiers

- **Type:** `[string]`
- **Required:** No

Unique identification for instructors. Each string in the array has a maximum length of 256 UTF-8 characters. Value can be null.

### location

- **Type:** `BaseRosterLocation`
- **Required:** No

The base location for this course.

### name

- **Type:** `string`
- **Required:** No

The class name as displayed in Apple School Manager. Maximum length is 1024 UTF-8 characters.

### op_date

- **Type:** `string`
- **Required:** No

The most recent update timestamp for this class.

Available in X-Protocol Version 3 and later.

### room

- **Type:** `string`
- **Required:** No

The room where class is held. Maximum length is 512 UTF-8 characters.

### source

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `LDAP`, `SIS`, `CSV`, `MANUAL`, `SYSTEM`, `ENROLLMENT`, `DEP`, `FORMULA_GENERATED`, `SFTP`

The data source where class was created. Maximum length is 64 UTF-8 characters.

### source_system_identifier

- **Type:** `string`
- **Required:** No

The identifier configured by the organization for its classes. Maximum length is 256 UTF-8 characters. See Note.

### status

- **Type:** `string`
- **Required:** No

The status of the class. Possible values: `Active and Inactive`.

Available in X-Protocol Version 3 and later.

### student_unique_identifiers

- **Type:** `[string]`
- **Required:** No

The unique identifiers for students. Each string in the array has a maximum length of 256 UTF-8 characters. Value can be null.

### unique_identifier

- **Type:** `string`
- **Required:** No

The global unique identifier for the class. Maximum length is 256 UTF-8 characters.

