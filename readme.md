# OS library

This module contains utilities to read and write files to multiple sources using URI schemes as:

| URI scheme | destination                     |
|---|---------------------------------|
| `://` | File System folder | 
| `s3://` or `s3s://` | Simple Storage Service endpoint |
| `http://` or` https://` | Http Service endpoint |

Examples of use can be found in the [read](readfile_test.go) and [write](writefile_test.go) tests.