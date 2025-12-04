I'm optimizing prematurely.  I expected the input to grow in size, to a level that wasn't maintainable as `[][]byte`, but it didn't.  The logic to maintain the prev/current/next lookups wasn't straight forward to me so I should have just not done it to start with.

I also expected the distance we could look to grow from "8 adjacent" to "distance 3" or something, and again, optimized too early