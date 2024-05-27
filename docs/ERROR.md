error codes are unique

each code represent scope of occurrence based on its numbers.
error codes contains 5 number `xyzab`. each part of numbers explain in below:

### x:
first part of error represent Internal section errors

| Code | Section    |
|------|------------|
| 1    | controller |
| 2    | models     |
| 9    | pkg        |

### y:
second part of error represent one level inside Internal section. for example
in models section it represent each model or in controllers represent each api scope.

#### controller
| Code | Section  |
|------|----------|
| 1    | auth     |
| 2    | panel    |
| 3    | redirect |

#### models
| Code | Section  |
|------|----------|
| 1    | link app |
| 2    | user app |

#### pkg
| Code | Section  |
|------|----------|
| 1    | db pkg   |

### z:
The `z` number represent inner section of each scope

#### Apps
| Code | Section    |
|------|------------|
| 1    | repository |
| 2    | service    |
| 3    | workers    |

#### APIs
- each business scope of api

#### ab
The last part is `ab` represent error code of each section.