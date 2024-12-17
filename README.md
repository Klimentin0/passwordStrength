# Password Strength

As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the previous lesson to iterate over the characters of a string.

## Assignment

Implement the _isValidPassword_ function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.
