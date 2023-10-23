# Problem: String Compression
*Objective:*
Write a Go program that implements a basic string compression algorithm using the counts of repeated characters. If the compressed string is not smaller than the original string, return the original string.

*Example:*

Input: "aabcccccaaa"
Output: "a2b1c5a3"

Input: "abcde"
Output: "abcde" // Compressed string is "a1b1c1d1e1", which is longer than the original string.

*Requirements:*

The input string consists only of uppercase and lowercase letters (a-z).
If two similar letters with different cases are adjacent, they should be treated as distinct characters.
The program should handle empty strings correctly. For an empty string input, it should return an empty string.

*Hints:*

You can traverse the string and keep a count of the current character. When the current character changes, append the previous character and its count to the compressed string.
Use a strings.Builder for efficient string concatenation.

*Bonus:*

Implement a decompression function that takes the compressed string as input and returns the original string.
Provide a mechanism for the user to choose between compression and decompression.
After attempting the problem, let me know if you'd like a sample solution or any further hints!