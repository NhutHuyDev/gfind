## GFIND
Searches for a string of text in a file or files, and displays lines of text that contain the specified string, like **FIND command** in Windows.
- Implement by Go

- Recursive file search into subdirectories is not implemented.

- The /off[line] flag is not supported

- Using Factory Method & Decorator 

- folder "./poems" for testing

## Syntax & Parameters
```
gfind [-v] [-c] [-n] [-i] -f <"string"> [[drive:][path]filename[...]]
```
| **Parameter**                       | **Description**                                                                                         |
|--------------------------------------|---------------------------------------------------------------------------------------------------------|
| `-v`                                 | Displays all lines that don't contain the specified `<string>`.                                           |
| `-c`                                 | Counts the lines that contain the specified `<string>` and displays the total.                           |
| `-n`                                 | Precedes each line with the file's line number.                                                          |
| `-i`                                 | Specifies that the search is not case-sensitive.                                                         |
| `-f`                                 | Required, defines the declaration for the string to search for and the file or files to search in.       |
| `<"string">`                         | Required. Specifies the group of characters (enclosed in quotation marks) that you want to search for.   |
| `[<drive>:][<path>]<filename>`       | Specifies the location and name of the file in which to search for the specified string.                 |
| `-help`                                 | Displays help at the command prompt.                                                                     |

## Reference
[1]. Học lập trình cùng Nam. (2024, April 17). Học lập trình  NET - Phần 15 - Giải bài tập, unit test, abstract factory và decorator design pattern [Video]. YouTube. https://www.youtube.com/watch?v=SJL664hxJ-4

[2]. Robinharwood. (2023, February 3). find. Microsoft Learn. https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/find
