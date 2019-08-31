# Nintaco Go API

### About

This API provides programmatic control of [the Nintaco NES/Famicom emulator](https://nintaco.com/) at a very granular level. It supports direct access to CPU/PPU/APU Memory and registers, to controller input and to save states. And it offers a multitude of listeners that enable programs to tap into emulation events.

### Initialization

Go programs control the emulator externally using an internal socket connection for interprocess communication. At the beginning of a program, specify the host and port consistent with the values entered into the Start Program Server window using:

```
nintaco.InitRemoteAPI("localhost", 9999)
```

After that, the singleton API instance can be obtained via:

```
api := nintaco.GetAPI()
```

### Listeners

The API starts out in a disabled state and while disabled only the `Add`/`Remove` listener methods work. After adding listeners, a program calls `Run` to activate the API. It signals that everything is setup and the program is ready to receive events. `Run` never returns. Instead, it enters an infinite loop that maintains the connection to the emulator. 

Listeners are cached and they rarely need to be removed. In the event that the API is temporarily disabled, listeners do not need to be re-added. They are automatically removed on program shutdown. And most of the API methods that modify internal states do not have the side effect of triggering listeners. For example, a program that receives events when a region of CPU memory is updated can modify the same region from the event listener without creating infinite recursion.

The easiest way for a program to do something once-per-frame is within a `FrameListener`, which is called back immediately after a full frame was rendered, but just before the frame is displayed to the user. `ScanlineListener` works in a similar way, but it is invoked after a specified scanline was rendered. `ScanlineCycleListener` takes that one step further and responds to a specified dot. A program can manipulate controller input from a `ControllersListener`, which is called back immediately after the controllers were probed for data, but just before the probed data is exposed to the machine. `AccessPointListener` is triggered by a specified CPU Memory read or write, or instruction execution point and `SpriteZeroListener` is triggered by sprite zero hits. Finally, `ActivateListener`, `DeactivateListener`, `StatusListener` and `StopListener` respond to API enabled events, API disabled events, status message events and Stop button events, respectively.

### Concurrency

The API is _not_ safe for concurrent use. After invoking `API.Run`, the only goroutine that can safely invoke API methods is the one that executes the listeners. While a listener is running, the emulator is effectively frozen; listeners need to return in a timely manner to avoid slowing down emulation. Programs can start additional goroutines to perform parallel computations; however, the results of those computations should be exposed to and used from the listeners to act on the API.

### Details

This API is a translation of [the Nintaco Java API](https://nintaco.com/api.html). Refer to [the Javadoc](https://nintaco.com/javadoc/index.html) for a detailed description of [the API methods](https://nintaco.com/javadoc/nintaco/api/API.html), the listeners and the constants.

### Examples

