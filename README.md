# TinyCache - A Lightweight Cache

**TinyCache** is an efficient, lightweight caching system designed to provide easy and fast data storage and retrieval. This project is currently in progress and open for contributions!

## Roadmap

### Milestone 1: Basic Cache System (In Progress)
- [x] **Cache server and client**: Set up a cache server that handles requests and a client that interacts with the server to store and retrieve data.
- [x] **RESP**: Build a simple communication protocol to send and receive data.
- [x] **In-memory cache storage**: Implement some commands(like PING, SET, GET etc.) to create a basic caching system.
- [x] **Thread-safe operations**: Ensure that the cache works correctly in programs with multiple tasks running at the same time.
- [ ] **TTL (Time-to-Live)**: Add a feature that automatically removes cached data after a set time.
- [ ] **Unit tests**: Write tests to make sure all parts of the cache work as expected.

### Milestone 2: Feature Expansion (Planned)
- [ ] **LRU (Least Recently Used) eviction policy**: Automatically remove the least used items when the cache gets full.
- [ ] **Pluggable storage backends**: Allow users to store cache data in different places (e.g., in a file or a database like Redis).
- [ ] **Persistent storage**: Ensure that cached data remains available even after the program restarts.

### Milestone 3: Advanced Features (Planned)
- [ ] **LFU (Least Frequently Used) eviction policy**: Add an LFU eviction policy to remove the least frequently used items when the cache is full, providing an alternative to LRU.


##  📦 Project Structure Overview

This project is organized into several directories to maintain a clean and modular structure. Below is a breakdown of each folder and its purpose:

### `cmd`
Contains the entry points for running different components of the project.

- **server**: The main server application.
- **client**: The client application to communicate with the server.
- **playground**: A playground for testing and experimenting with features.

### `connection`
Handles server and client connections.

- **server**: Manages server-side connections.
- **client**: Manages client-side connections.

### `pkg`
Contains core logic and utilities used across the project.

- **utils**: General utility functions used across the project.
- **resp**: Handles RESP serializing and deserializing logic.
- **cache**: Contains caching logic.
- **errors**: Contains custom error definations.
- **handlers**: Contains the request handlers.

### `tests`
Includes unit tests for different components.

- **resp**: Tests for RESP command handling.
- **cache**: Tests for caching functionality.

---

This structure ensures that the project is organized, modular, and easy to maintain. Each directory has a specific purpose, promoting clean separation of concerns.


## 🚀 How to Run the Project 🛠️

To start the project, use the `run.sh` script. This script provides several options for running different components of the project. 

### Steps:

1. Run the following command to start the script:
    ```bash
    bash run.sh
    ```
2. You will be presented with the following options:
   ```bash
   1) Start server
   2) Start client
   3) Start playground
   4) Kill running server and client
   5) Run test
   Type: 
    ```
3. Type 1 to start the server, or type 2 to start the client.

4. To stop any running process, select the Kill process option by typing 4.

5. To run the tests, type 5 and press enter.


## 🤝 Contributing

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated. Please follow our [contribution guideline](./CONTRIBUTING.md).

## 📬 Contact

If you have any questions, feel free to reach out via [GitHub Issues](https://github.com/nahK994/TinyCache/issues) or email at nkskl6@gmail.com.



## 📝 License

This project is open-source and available under the [MIT License](./LICENSE).

