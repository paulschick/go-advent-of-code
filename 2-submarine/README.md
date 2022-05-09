# Submarine Position

I'm going to try a different strategy this time through.
Create a map or data structure that just takes each of the possible
directions, and sums the magnitude of that direction through all of
the instructions.

Order doesn't matter here. So I'll just have forward x, down, y, up z.
Then calculate the total position. I'll basically just have to iterate
the full array of directions once.

This should also work fine for part two.
