# Nintaco Go API

###About
This API provides programmatic control of [the Nintaco NES/Famicom emulator](https://nintaco.com/) at a very granular level. It supports direct access to CPU/PPU/APU Memory and registers, to controller input and to save states. And it offers a multitude of listeners that enable programs to tap into emulation events.

Go programs control the emulator remotely using an internal socket connection for interprocess communication.

###Concurrency
The API is _not_ safe for concurrent use. After invoking `API.Run`, the only goroutine that can safely invoke API methods is the one that executes the listeners. While a listener is running, the emulator is effectively frozen; listeners need to return in a timely manner to avoid slowing down emulation. Programs can start additional goroutines to perform parallel computations; however, the results of those computations should be exposed to and used from the listeners to act on the API.
