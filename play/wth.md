## My notes
### 2017-03-26
Doing some clean up. Moving files into playground links, even though some of them
don't actually work on the playground. I'm not sure if they changed something 
somewhere, but reading the raw JSON off the GH repo seems not to work on playground,
although [this sort of seems to work](https://play.golang.org/p/72egNg9ebH), at least 
it's not throwing an error.

Some the files moving out are actually a bit outdated, given how the Quiz struct
has evolved in response to the tests. That's fine - it's actually how it's supposed
to go, so I'm cool with it.

I think the next real step needs to be completing the JSON file. I'm reasonably happy
with what it's doing. I guess the next nut to crack is how to put rich text in the 
About section. Either that or come up with a way to index the questions without
relying on them being sequential ints.


### 2016-12-23
[Got it reading the json, finally](https://play.golang.org/p/1uR4FLTchA).
This is where I found out I'll have to change the casing on the 
section headers in the JSON.

After some messing around trying to get go to parse my data file, 
I finally got it going. I bit off a little more than I could chew
to start, but this was a good exercise. Go is not javascript. :-)

The `for` structure is similar to the `foreach` in C#, but different
in syntax. I like the `_` structure as way of dealing with unneeded
returns from functions. 

[I'm snapshotting this now.](https://play.golang.org/p/bpvjBNronw)

[Here's a slightly nicer format.](https://play.golang.org/p/O23pqgNQdv)


--------------