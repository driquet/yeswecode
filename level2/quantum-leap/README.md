# Quantum AI Assistant
You are working on a computer science project and you would like to improve it, because you are convinced it is not optimized. You seek assistance from an advanced Quantum AI Assistant, which is supposed to help you identify areas where you could optimize your code. Unfortunately, the AI answer is not clear, and based on `qubits` (or quantum bits). A `qubit` is a two-state system (enabled or disabled), roaming in the answer given by the AI. According to the AI, observing the qubits will reveal the answer to you.

Here is an example input:

```
(6, 6) (-3, -3)
(0, 4) (1, -2)
(4, 0) (0, 0)
(4, 5) (-2, -2)
(-4, 3) (3, -1)
(2, 3) (1, -1)
(-6, -2) (3, 2)
(-5, 6) (3, -2)
(6, 8) (-2, -3)
(6, 4) (-1, -1)
(-4, 9) (2, -3)
(8, -1) (-3, 2)
(0, -3) (2, 3)
(4, 2) (-2, 1)
(-2, 6) (2, -1)
(0, 2) (2, 1)
```

Each line represents a qubit, following this format: `(pos_x,pos_y) (velocity_x,velocity_y)`. At each quantum tick, the velocity of each qubit is added to its position. So, a qubit with velocity `(1, -2)` is moving to the right, but is moving upward twice as quickly. If the initial position of the qubit was `(3,9)`, after `3` ticks, its position would be `(6,3)`.

Over time, the example input would look like this:

`tick 0`
```
......#........
#..............
..............#
......#...#....
...............
......#...#....
..#.....#......
......#.....#..
..........#....
.#..#.......#..
...............
............#..
..#............
```

`tick 1`
```
#..#.#.#.
........#
..#.#.#..
.....##.#
.#.......
...#...#.
.#.......
```

`tick 2`
```
#.#.#
#.#.#
###.#
#.#.#
#.#.#
```

After 2 ticks, the message appears: `HI`.

Discover the answer sent by the AI, which may require many more quantum ticks to appear.
