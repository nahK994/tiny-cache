## TinyCache - A Lightweight Cache

**TinyCache** is an efficient, lightweight caching system designed to provide easy and fast data storage and retrieval. This project is currently in progress and open for contributions!

### Roadmap

#### Milestone 1: Basic Cache System (In Progress)
- [x] **Cache server and client**: Set up a cache server that handles requests and a client that interacts with the server to store and retrieve data.
- [ ] **RESP**: Build a simple communication protocol to send and receive data.
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

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated.

#### How to Contribute

1. **Fork the repository** to your GitHub account.
2. **Clone your forked repository** to your local machine:

   ```bash
   git clone https://github.com/nahK994/TinyCache.git
   ```
   
4. **Create a new branch** for your feature or fix:
    ```bash
    git checkout -b feature/YourFeature
    ```
5. **Make your changes** and commit them with descriptive messages:
    ```bash
    git commit -m "Add new feature"
    ```
6. **Push your changes** to your forked repository:
    ```bash
    git push origin feature/YourFeature
    ```
7. **Open a pull request** to merge your changes into the main repository.




### 📝 License

This project is open-source and available under the MIT License.


#### Feel free to reach out if you have any questions or need help!
