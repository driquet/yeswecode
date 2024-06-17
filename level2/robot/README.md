# Robot explorer
You are given instructions for a robot that moves on a 2D grid.
Here are the instructions:

- `<`: Move the robot one step to the left.
- `>`: Move the robot one step to the right.
- `^`: Move the robot one step upwards.
- `v`: Move the robot one step downwards.


Additionally, the robot may leave a trail behind its path. This feature is disabled by default, and is toggled with the `!` character.

Here is an example of input:

```
^!^^>>>!v!vv<<<
```

The trail left by the robot draws the following (it begins in the bottom-left corner):

```
###.
#..#
#..#
.###
```

Note: In the output, the trail left by the robot is represented using the `#` character, while empty cells are represented by `.`.

Draw the trail left by the robot and find the flag.
