# Regex-Engine

## Stage 1/6: Single character strings

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

## Stage 2/6: Matching two equal length strings
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
## Stage 3/6: Working with strings of different length

Repeatedly invoke a function and check if there is a match of characters. If there isn't,
another section of the string should be passed.

### Example
```
Input: ‘tion|Section’     Output: false
Input: ‘tion|ection’      Output: false
Input: ‘tion|ction’       Output: false
Input: ‘tion|tion’        Output: true
```
## Objectives

The improved regex engine should do the following:

    * A new function is created as an entry point;
    * It should repeatedly invoke the function that compares two equal length patterns;
    * If that function returns true, the new function should also return true;
    * If that function returns false, the input string should be passed to the matching
    function with an incremented starting position, and the regex should be passed unmodified;
    * The process goes on until the entire input string has been consumed.

## Stage 4/6: Implementing the operators ^ and $

Your task is to add some metacharacters to the already existing regex engine.

At this stage, you should add the following special cases:

   * `^` can occur at the beginning of the regex, and it means that the 
    following regex should be matched only at the beginning of the input string.
   * `$` can occur at the end of the regex, and it means that the preceding regex
    should be matched only at the end of the input string.

### Example
```
Input:    '^app|apple'           Output: true
Input:     'le$|apple'           Output: true
Input:      '^a|apple'           Output: true
Input:      '.$|apple'           Output: true
Input:  'apple$|tasty apple'     Output: true
Input:  '^apple|apple pie'       Output: true
Input: '^apple$|apple'           Output: true
Input: '^apple$|tasty apple'     Output: false
Input: '^apple$|apple pie'       Output: false
Input:    'app$|apple'           Output: false
Input:     '^le|apple'           Output: falsa
```

## Stage 5/6: Controlling repetition 
Add the following metacharacters to your engine:

   * `?` matches the preceding character zero times or once;
   * `*` matches the preceding character zero or more times;
   * `+` matches the preceding character once or more times.

## Objectives

In the case of the operator ?, there are two possible scenarios:

   * The preceding character occurs zero times, so basically it is skipped. This means that only the part of the regex, if present, after the metacharacter ? is passed to the recursive function along with the unmodified input string.
   * The preceding character occurs once. This means that if the character preceding ? matches the first character of the input string, the part of the regex after ? is passed to the recursive function along with the part of the input string without the character that is already matched.

In the case of the operator *, there are the following scenarios:

   * The preceding character occurs zero times (just like with ?). The condition from the previous case can be reused.
   * The preceding character occurs one or more times. Like in the case of ?, the character preceding * should match the first character of the input string. Since we don’t know how many times it is going to be repeated, the regex should be passed to the recursive function without any modification, and the first character of the input string should be chopped off. In this case, since the metacharacter * is not removed, the second case is repeated until the preceding character can be matched. After that, the first case applies and the function comes out of the recursion.

Finally, here is what can happen with the operator +:

   * The preceding character occurs once. This case is the same as the second case with the operator ?.
   * The preceding character occurs more than once. This case is basically the same as the second case with the operator *.
   * If there are character(s) after the operator, you need to pass to the recursive function a modified regex and a modified string.

### Example
```
Input: 'colou?r|color'       Output: true
Input: 'colou?r|colour'      Output: true
Input: 'colou?r|colouur'     Output: false
Input: 'colou*r|color'       Output: true
Input: 'colou*r|colour'      Output: true
Input: 'colou*r|colouur'     Output: true
Input:  'col.*r|color'       Output: true
Input:  'col.*r|colour'      Output: true
Input:  'col.*r|colr'        Output: true
Input:  'col.*r|collar'      Output: true
Input: 'col.*r$|colors'      Output: false
```
## Stage 6/6: Escaping

The character that follows a backward slash should be interpreted as a 
literal. This way, a simple string comparison is enough to determine a
match. You should also take care of skipping the escape sequence before
continuing the recursion. This can be done using the same logic that we
used to handle the repetition operators ?, *, and +.

### Example
```
Input:      '\.$|end.'              Output: true
Input:     '3\+3|3+3=6'             Output: true
Input:       '\?|Is this working?'  Output: true
Input:       '\\|\'                 Output: true
Input: 'colou\?r|color'             Output: false
Input: 'colou\?r|colour'            Output: false
```
