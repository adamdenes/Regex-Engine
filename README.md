# Regex-Engine

## Single character strings

In this first stage, create a function that can compare a single character regex
to a single character input string.

The reason we are only considering single characters, for now, is the presence of
the wild card indicated by a period (.). Its role in a regular expression is to
match any character in the target string. So, if the regex that we pass to our function as an argument is a wild-card, the function always returns true no matter what it is being compared to.

When working on this stage, keep in mind the following special rules of the regex syntax:

    - An empty regex should always return the bool value true.
    - An empty input string should always return the bool value false, except if the regex is also empty.
    - An empty regex against an empty string always returns true.

These rules seem random at first, but later on, they will make sense to you.
Objectives

In this stage, your program should:

    - Accept two characters, a regex and an input;
    - Compare the regex to the input and return a boolean indicating if there's a match;
    - Support . as a wild card character that matches any input;
    - Follow the special syntax rules outlined above.

## Example
```
Input: 'a|a'     Output: true
Input: '.|a'     Output: true
Input:  '|a'     Output: true
Input:  '|'      Output: true
Input: 'a|'      Output: false
```
