# TinyCache - A Lightweight Cache

**TinyCache** is an efficient, lightweight caching system designed to provide easy and fast data storage and retrieval. This project is currently in progress and open for contributions!

## Roadmap

### Milestone 1: Basic Cache System (Completed ✅)
- [x] **Cache server and client**: Set up a cache server that handles requests and a client that interacts with the server to store and retrieve data.
- [x] **RESP**: Build a simple communication protocol to send and receive data.
- [x] **In-memory cache storage**: Implement basic commands to create a basic caching system. Commands are-

    ***PING, GET, SET, EXISTS, FLUSHALL, DEL, INCR, DECR, LPUSH, RPUSH, LPOP, RPOP, LRANGE, EXPIRE, TTL, PERSIST***
- [x] **Thread-safe operations**: Ensure that the cache works correctly in programs with multiple tasks running at the same time.
- [x] **Unit tests**: Write tests to make sure all parts of the cache work as expected.

### Milestone 2: Feature Expansion (Completed ✅)
- [x] **TTL (Time-to-Live)**: Add a feature that automatically removes cached data after a set time.
- [x] **LFU (Least Frequently Used) eviction policy**: Add an LFU eviction policy to remove the least frequently used items when the cache is full, providing an alternative to LRU.

### Milestone 3: Advanced Features (Planned 🛠️)
- [ ] **Pluggable storage backends**: Allow users to store cache data in different places (e.g., in a file or a database like Redis).
- [ ] **Persistent storage**: Ensure that cached data remains available even after the program restarts.


## 🚀 How to Run the Project 🛠️

### Prerequisites

To run TinyCache, make sure you have the following installed on your machine:

- **Go 1.22** or higher. You can download Go from [here](https://golang.org/dl/).

Verify your Go version using:
```bash
go version
```

### Running the Project

1. Clone the repository or download it manually:
    ```bash
    git clone https://github.com/nahK994/TinyCache.git
    cd TinyCache
    ```
2. Use the `run.sh` script to manage the server and client. This script provides options to start the server, client, playground, and more.
3. Run the following command to start the script:
    ```bash
    bash run.sh
    ```

4. You will be presented with the following options:
    ```bash
    1) Start server
    2) Start client
    3) Start playground
    4) Kill running server and client
    5) Run test
    6) Run build
    Type: 
    ```

5. Type 1 to start the server, or 2 to start the client.
6. To stop any running process, select the Kill process option by typing 4.
7. To run the tests, type 5 and press enter.
8. Type 6 to build binaries for client and server.


## 📥 Installation
If you'd like to install TinyCache on Linux without cloning the repository, use the following command to install both the server and client:
```bash
curl -fsSL https://raw.githubusercontent.com/nahK994/TinyCache/master/install.sh | bash
```
This will download the binaries, install them to the appropriate locations, and set up the server as a systemd service.


## 🧹 Uninstallation
To uninstall from linex TinyCache, simply run the following command:
```bash
curl -fsSL https://raw.githubusercontent.com/nahK994/TinyCache/master/uninstall.sh | bash
```
This will stop the service, remove the binaries, and clean up all installed files.


##  📦 Project Structure Overview

This project is organized into several directories to maintain a clean and modular structure. Below is a breakdown of each folder and its purpose:

```shell
├── cmd/                         # Entry points for different binaries
│   ├── client/
│   │   └── main.go              # Main file to run the CLI-based cache client
│   ├── playground/
│   │   └── main.go              # For experimentaion stuffs
│   └── server/
│       └── main.go              # Main file to launch the TinyCache server

├── connection/                  # Logic for handling low-level connections
│   ├── client/
│   │   └── client.go            # Code for client-side connection handling
│   └── server/
│       └── server.go            # Code for accepting and managing client connections on the server

├── CONTRIBUTING.md              # Guide for contributors (e.g., how to fork, open PRs, coding style)
├── go.mod                       # Go module file
├── go.sum                       # Checksum of module dependencies
├── install.sh                   # Script to install or set up TinyCache locally
├── LICENSE                      # License file (e.g., MIT, Apache)
├── run.sh                       # Helper script to build and run the project
├── uninstall.sh                 # Script to cleanly uninstall TinyCache

├── pkg/                         # All core functionality and reusable code lives here
│   ├── cache/                   # Main cache logic and internal data structures
│   │   ├── helpers.go           # Utility functions for cache operations
│   │   ├── models.go            # Data models
│   │   └── store.go             # Core implementation of the cache logic
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── errors/
│   │   ├── constants.go         # Custom error constants
│   │   └── errors.go            # Custom error types and functions
│   ├── handlers/
│   │   └── handlers.go          # High-level request handlers
│   ├── resp/
│   │   ├── constants.go         # RESP protocol-specific constants
│   │   ├── deserializer.go      # Parse raw RESP input
│   │   └── serializer.go        # Encode responses into RESP format
│   ├── shared/
│   │   └── helpers.go           # Shared helper utilities across the project
│   ├── utils/
│   │   └── constants.go         # Miscellaneous constants
│   └── validators/
│       ├── client_validators.go # Validate client-side input/flags/config
│       └── server_validators.go # Validate server-side configuration/commands

├── tests/                       # Organized unit & integration tests
│   ├── cache/
│   │   └── cache_test.go        # Tests for core cache logic
│   ├── handlers/
│   │   └── handlers_test.go     # Tests for command handlers
│   ├── resp/
│   │   ├── deserializer_test.go # Tests for RESP deserialization
│   │   ├── serializer_test.go   # Tests for RESP serialization
│   │   └── test_cases.go        # Common RESP test cases
│   ├── shared/
│   │   └── helpers_test.go      # Tests for shared helper functions
│   └── validators/
│       └── client_validators_test.go # Tests for client validation logic

├── README.md                    # Project overview, how to install/use, contribution guide
```

## 🤝 Contributing

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated. Please follow our [contribution guideline](./CONTRIBUTING.md).

## 📬 Contact

If you have any questions, feel free to reach out via [GitHub Issues](https://github.com/nahK994/TinyCache/issues) or email at nkskl6@gmail.com.



## 📝 License

This project is open-source and available under the [MIT License](./LICENSE).

