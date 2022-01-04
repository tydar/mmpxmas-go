# ModMyPi Programmable Christmas Tree examples in Go


This small program contains examples similar to the examples provided by [ModMyPi](https://github.com/modmypi/Programmable-Xmas-Tree) for interacting with the GPIO-controlled LED Christmas Tree. Tested with a Raspberry Pi Zero 2 W.

## Usage


Build cross-compiled from an x86 PC:
```shell-session
$ make build
```

Use `scp` or another tool to move the binary to your Raspberry Pi. Then:

Run a similar pattern to the "ants" example from the original ModMyPi repo.

```shell-session
$ ./mmpxmas ants
```

Turn off all LEDs:

```shell-session
$ ./mmpxmas clear
```

