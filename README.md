## TinyCache - A Lightweight Cache

**TinyCache** is an efficient, lightweight caching system designed to provide easy and fast data storage and retrieval. This project is currently in progress and open for contributions!

### Roadmap

#### Milestone 1: Basic Cache System (In Progress)
- [x] **Cache server and client**: Set up a cache server that handles requests and a client that interacts with the server to store and retrieve data.
- [x] **RESP**: Build a simple communication protocol to send and receive data.
- [ ] **In-memory cache storage**: Create a basic system that stores data in memory for fast access.
- [ ] **TTL (Time-to-Live)**: Add a feature that automatically removes cached data after a set time.
- [ ] **Thread-safe operations**: Ensure that the cache works correctly in programs with multiple tasks running at the same time.
- [ ] **Unit tests**: Write tests to make sure all parts of the cache work as expected.

#### Milestone 2: Feature Expansion (Planned)
- [ ] **LRU (Least Recently Used) eviction policy**: Automatically remove the least used items when the cache gets full.
- [ ] **Pluggable storage backends**: Allow users to store cache data in different places (e.g., in a file or a database like Redis).
- [ ] **Persistent storage**: Ensure that cached data remains available even after the program restarts.

#### Milestone 3: Advanced Features (Planned)
- [ ] **LFU (Least Frequently Used) eviction policy**: Add an LFU eviction policy to remove the least frequently used items when the cache is full, providing an alternative to LRU.


### 🤝 Contributing

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated. Please follow our [contribution guideline](./CONTRIBUTING.md).


### 📝 License

This project is open-source and available under the MIT License.


#### Feel free to reach out if you have any questions or need help!
