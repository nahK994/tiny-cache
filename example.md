# Here’s an updated **Summary of Inputs and Outputs** with the addition of `SET count 10` and examples for `INCR` and `DECR` on an uninitialized variable (`age`).

# ### Summary of Inputs and Outputs

# | Command            | Input                                                                                  | Output                                                      |
# |--------------------|-----------------------------------------------------------------------------------------|-------------------------------------------------------------|
# | **`SET name John`** | `*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$4\r\nJohn\r\n`                                        | `+OK\r\n`                                                   |
# | **`GET name`**      | `*2\r\n$3\r\nGET\r\n$4\r\nname\r\n`                                                      | `$4\r\nJohn\r\n` or `$-1\r\n` (if not exists)               |
# | **`SET count 10`**  | `*3\r\n$3\r\nSET\r\n$5\r\ncount\r\n$2\r\n10\r\n`                                         | `+OK\r\n`                                                   |
# | **`INCR count`**    | `*2\r\n$4\r\nINCR\r\n$5\r\ncount\r\n`                                                    | `:11\r\n` (increments `count` from `10` to `11`)            |
# | **`DECR count`**    | `*2\r\n$4\r\nDECR\r\n$5\r\ncount\r\n`                                                    | `:10\r\n` (decrements `count` back to `10`)                 |
# | **`INCR age`**      | `*2\r\n$4\r\nINCR\r\n$3\r\nage\r\n`                                                      | `:1\r\n` (initializes `age` to `1` and increments)          |
# | **`DECR age`**      | `*2\r\n$4\r\nDECR\r\n$3\r\nage\r\n`                                                      | `:0\r\n` (decrements `age` from `1` to `0`)                 |
# | **`EXISTS name`**   | `*2\r\n$6\r\nEXISTS\r\n$4\r\nname\r\n`                                                   | `:1\r\n` (exists) or `:0\r\n` (not exists)                  |
# | **`EXISTS age`**    | `*2\r\n$6\r\nEXISTS\r\n$3\r\nage\r\n`                                                    | `:1\r\n` (exists after `INCR`/`DECR`)                       |

# ### Breakdown of New Additions:

# 1. **`SET count 10`**:
#    - **Input**: `*3\r\n$3\r\nSET\r\n$5\r\ncount\r\n$2\r\n10\r\n`
#    - **Output**: `+OK\r\n`
#      - This sets the value of `"count"` to `10`.

# 2. **`INCR age`** (Uninitialized Variable):
#    - **Input**: `*2\r\n$4\r\nINCR\r\n$3\r\nage\r\n`
#    - **Output**: `:1\r\n`
#      - If `"age"` is not initialized, `INCR` will create the key `"age"` with an initial value of `1`.

# 3. **`DECR age`** (After Uninitialized Increment):
#    - **Input**: `*2\r\n$4\r\nDECR\r\n$3\r\nage\r\n`
#    - **Output**: `:0\r\n`
#      - After `INCR` sets `"age"` to `1`, `DECR` will decrement `"age"` to `0`.

# ### How It Works:
# - **`SET`**: The command sets a value for a key, and the response is `+OK\r\n` to indicate success.
# - **`INCR` and `DECR` for Initialized Keys**: For numeric keys like `"count"`, `INCR` increments the value and returns an integer reply, and `DECR` decrements it.
# - **`INCR` and `DECR` for Uninitialized Keys**: If a key doesn’t exist, `INCR` initializes it to `1`, and `DECR` initializes it to `-1` or decreases it accordingly.
# - **`EXISTS`**: This checks if a key exists. After `INCR`/`DECR`, the previously uninitialized key will exist, returning `:1`.

# These commands and their RESP outputs are fundamental to building an in-memory cache system. They ensure basic functionality for storing, retrieving, and manipulating key-value pairs.