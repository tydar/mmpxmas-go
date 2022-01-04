# ModMyPi Programmable Christmas Tree examples in Go


This small program contains examples similar to the examples provided by [ModMyPi](https://github.com/modmypi/Programmable-Xmas-Tree) for interacting with the GPIO-controlled LED Christmas Tree. Tested with a Raspberry Pi Zero 2 W running 32-bit Raspbian Lite. Using the [go-rpio](https://github.com/stianeikeland/go-rpio) library to interface with the GPIO system.

## Usage


Build cross-compiled from an x86 PC:
```shell-session
$ make build
```

Use `scp` or another tool to move the binary to your Raspberry Pi. Then:

Run a similar pattern to the "ants" example from the original ModMyPi repo.

```shell-session
$ ./mmpxmas-go ants
```

Turn off all LEDs:

```shell-session
$ ./mmpxmas-go clear
```

Similar to "alt" example:

```shell-session
$ ./mmpxmas-go alt
```

Even and odd indicies:

```shell-session
$ ./mmpxmas-go mod2
```

Light one row at a time:

```shell-session
$ ./mmpxmas-go rows
```
