# yeswecode

This repository is a collection of challenges I wrote for a coding challenge I co-organized at my company. To solve a challenge, you need to find the solution (~flag). While all these challenges are original, some of them were inspired by existing challenges from platforms such as [Advent of Code](https://adventofcode.com/), [CodinGame](https://www.codingame.com/) or [LeetCode](https://leetcode.com/).

Each challenge consists of:

- A description
- A level of difficulty (ranging from 1 for very easy to 5 for very hard)
- A flag (the expected answer)
- Associated files (e.g., example or input files)

To prevent accidental viewing, I have base64 encoded the flags.

To decode a flag:

```bash
base64 -d flag.b64 > flag.txt
```

To decode all flags:

```bash
find . -name "flag.b64" -exec sh -c 'base64 -d "$0" > "${0%.b64}.txt"' {} \;
```

Feel free to ask question or share the challenges!
