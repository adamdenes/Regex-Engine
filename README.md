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

### Example
```
Input: 'a|a'     Output: true
Input: '.|a'     Output: true
Input:  '|a'     Output: true
Input:  '|'      Output: true
Input: 'a|'      Output: false
```

## Matching two equal length strings
A single character is not a lot, so let's extend our new regex engine to handle
regex-string pairs of equal length. Don't forget about supporting the wildcard 
symbol! It is still not the most realistic way to use a regex engine, but we are
slowly getting there.

## Objectives

Create a new function that calls the function from the first stage on every character
of the regex-string pair and returns true only if there is a match for every character.
In other words, for a complete match, either every character pair should be the same,
or the regex should contain a wild card. There are different ways to achieve this, but
the most elegant is probably recursion.

Recall that recursion is when a function calls itself from its own code. It can be
used to break down a problem into smaller steps, thus simplifying the code.

This is exactly what you are going to do in this stage! First, invoke the function
from the first stage on the first characters of the regex-string pair. If there is a match,
pass the remainder of the string recursively to the same function, but this time without
the first characters of the regex and the input string, thus "consuming" them step by step.

Some terminating conditions should be added to stop the function from entering infinite
recursion:

   - If the regex has been entirely consumed, the function should return true
   since it means that all the characters in the regex-string pair are the same.
   - If the regex has not been consumed but the input string has, the function
   should return false since it means that the regex is longer than the input 
   string, which is undefined behavior at this point.
   - If the first character of the regex does not match the first character of 
   the input string, the function should return false because it guarantees that
   the two patterns are different.

If none of the above apply, the recursion should continue until the regex-string
pair is entirely consumed through slicing.

### Example
```
Input: 'apple|apple'     Output: true
Input: '.pple|apple'     Output: true
Input: 'appl.|apple'     Output: true
Input: '.....|apple'     Output: true
Input: 'peach|apple'     Output: false
```
