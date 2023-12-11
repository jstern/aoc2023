# aoc2023

Solutions to advent of code 2023 (plus some 2022 stuff to make sure the harness works).

## Add boilerplate for a new solution

Use the `stubs` makefile target and provide a `key` in `YYYY:D` format, e.g.

```
make stubs key=2023:1
```

```
$ make stubs key=2023:1
created aoc/y2023d1.go
created aoc/y2023d1_test.go
```

Note that keys **must** be prefixed `YYYY:D:` ... this is how the runner fetches the puzzle input from the advent of code site.

The stubs target sets up the common situation of wanting to run a part 1 and part 2, but you can add additional solutions by providing another correctly prefixed `key`, e.g.

```go
// y2023d1.go
package aoc

func init() {
        registerSolution("2023:1:1", y2023d1part1)
        registerSolution("2023:1:2", y2023d1part2)
        // additional solution
        registerSolution("2023:1:1brute", y2023d1part1brute)

}

// ...

func y2023d1part1brute(input string) string {
        // do something brutal
        return "brute force answer"
}
```

## Run a given solution

### Fetching puzzle input

The runner attempts to fetch input from `https://adventofcode.com/YYYY/day/D/input` given a key prefixed `YYYY:D:`. You'll need to paste a valid adventofcode.com session cookie value in a file named `.aoc-session` in this directory.

### Running a solution

Use the `run` make target and provide the key for the solution you want to run.

```
make run key=2022:1:1
```

```
$ make run key=2022:1:1
---
Answer in 688ns
---
wrong
```

### Listing solutions

If you forget which keys have at least the stub of a solution, you can ask:

```
$ make list

Available solutions (aka keys):
  * 2022:1:1
  * 2022:1:2
  * 2023:1:1
  * 2023:1:2
```

### Being patient

By default, the runner is not so patient ... it'll wait 5 minutes for an answer and then give up. Set the `AOC_TIMEOUT` environment variable to something higher than 300 to wait more than 300 seconds ... or lower, to be even more impatient!

### Being confident

**Note: this was working for a while, but has stopped working; seems like posts made by my code are receiving 302 responses without bodies now?**

To automatically post a solution and see if the answer is correct, use `make submit`:

```
$ make submit key=2020:15:1

Answer in 882ns
---
wrong
===
Submitting...
===
That's not the right answer.  If you're stuck, make sure you're using the full input data; there are also some general tips on the about page, or you can ask for hints on the subreddit.  Because you have guessed incorrectly 4 times on this puzzle, please wait 5 minutes before trying again. [Return to Day 15]
```

### Being suspicious

If you're worried about your logic and want to use print statements to see what's going on, but you don't want to see that junk automatically every time, you can set `AOC_VERBOSE=1` in the environment and wrap your output code in a check:

```go
if VerboseEnabled() {
    fmt.Println(
        `some really long but probably cool looking string
        that totally shows you where you've messed up`
    )
}
```

### Being aware

The `run` and `submit` targets will append information about the run to `.aoc/log.txt` so you can see which answers you've already come up with, and (for `submit`) what the site said about them.
