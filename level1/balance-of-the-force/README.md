# Balance of the force

The input of this challenge is composed of matching opening and closing characters:

- Square brackets (`[` and `]`).
- Curly brackets (`{` and `}`).
- Closing parentheses (`(` and `)`).

For each type of characters (square brackets, curly brackets and parentheses), you need to keep track of a counter, that increments if you encounter an opening character, or
decrements for a closing character. Whenever the counter becomes negative, you found the unbalanced position for this type of character.

Here is an example of input:

```
(()[]{}))[]]}
```

At position **0**:
- square brackets: 0
- curly brackets: 0
- parentheses: 0

At position **2**:
- square brackets: 0
- curly brackets: 0
- parentheses: 1

At position **5**:
- square brackets: 0
- curly brackets: 1
- parentheses: 1

At position **8**, we find the unbalanced position of parentheses:
- square brackets: 0
- curly brackets: 1
- parentheses: **-1** 

At position **11**, we find the unbalanced position of square brackets.
At position **12**, we find the unbalanced position of curly brackets.

The flag of this challenge is the multiplication of the three unbalanced positions.
For the example input, it would be: **1056** (`8*11*12`).

You can find the input attached to this challenge.
