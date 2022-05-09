# Submarine Position

I'm going to try a different strategy this time through.
Create a map or data structure that just takes each of the possible
directions, and sums the magnitude of that direction through all of
the instructions.

Order doesn't matter here. So I'll just have forward x, down, y, up z.
Then calculate the total position. I'll basically just have to iterate
the full array of directions once.

This should also work fine for part two.

## Scoring

Forward - horizontal + n
Down - Depth + n
Up - Depth - n

1. Do all forward first (all positive)
2. Do all down next (all positive)
3. Subtract up from down

Final result is horizontal * depth

## Part Two

down - aim + n
up - aim - n
forward:
  - horizontal + n
  - depth = aim * n

So now down and up only affect aim. Forward will update horizontal
and depth. The final value is again, horizontal * depth.

### Update to Part Two

In this part *order does matter*. So I am unable to parse this the
same way. I can still use the struct methods that I wrote, but unfortunately
I cannot use the directions map that I created in the first step.

I'm going to have to parse the directions line by line instead. This is due
to the fact that the multiplication factor is going to be different at different
points in the instructions.

This shouldn't be a difficult change.
