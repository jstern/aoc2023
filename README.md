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
created aoc/y2023d1_test.go``
```

Note that keys **must** be prefixed `YYYY:D:` ... this is how the runner fetches the puzzle input from the advent of code site.

The stubs target sets up the common situation of wanting to run a part 1 and part 2, but you can add additional solutions by providing another correctly prefixed `key`, e.g.

```go
# y2023d1.go
package aoc

init() {
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

Note that keys **must** be prefixed `YYYY:D:` ... this is how the runner fetches the puzzle input from the advent of code site.

## Run a given solution

### Fetching puzzle input

The runner attempts to fetch input from `https://adventofcode.com/YYYY/day/D/input` given a key prefixed `YYYY:D:`. You'll need to paste a valid adventofcode.com session cookie value in `.aoc-session` for this to work.

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
