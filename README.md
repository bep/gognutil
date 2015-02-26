# gognutil

This isn't necessarily very useful, just a placeholder project for Go implementations of some of the [GNU Core Utils](https://www.gnu.org/software/coreutils/).

The functionality should be very similar, the implementation very different (I haven't looked too close). It's not tuned and can surely be improved. PRs are welcome!

As you see below, this project just started -- and will probably never be complete.

For a similar project with lots of these util implemented in Go, see [go-coreutils](https://github.com/aisola/go-coreutils).

```
Available Commands: 
  seq         Print a sequence of numbers
```

## Building and running

If you have `make` installed, just run `make`. The executale will be built to `bin`. Here it sits alongside scripts for running the single commands. They are named as theyr GNU counterparts, so if you place `bin` in the path before `/usr/bin`, these will be used instead.

You can of course also just run it as a CLI with `bin/gognutil`.
