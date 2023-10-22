## *Problem:* String Manipulation - Word Reverser

*Objective:* Write a Go program that takes a string as input and returns the string with each word reversed, but the order of the words should remain the same.

*Example:*

Input: "Hello World"
Output: "olleH dlroW"

*Requirements:*

The program should be able to handle multiple spaces between words.
The program should handle punctuation correctly, i.e., punctuation should stay at the end of the word.
The program should be case-sensitive.

*Steps:*

Take a string as input from the user.
Split the string into words.
For each word, reverse the characters.
Join the reversed words to form the output string.
Print the output string.

*Bonus:*

Ensure the program can handle other forms of whitespace, like tabs.
Provide an option for the user to input a filename, and the program should read the content of the file, process it, and save the result in a new file.

*Hints:*

You can use the strings package in Go to split and join strings.
Use a loop to reverse the characters of each word.
Consider using the bufio package for reading and writing to files.
After you've attempted the problem, I can provide a sample solution if you'd like!