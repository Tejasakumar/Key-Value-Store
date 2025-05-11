





<div align="center">
  <h1>🗄️ KVS</h1>
  <p><strong>A Simple In-Memory Key-Value Store</strong></p>
  <p>
    <a href="#features">Features</a> •
    <a href="#getting-started">Getting Started</a> •
    <a href="#command-reference">Commands</a> •
    <a href="#example-session">Examples</a> •
    <a href="#implementation-details">Implementation</a>
  </p>
</div>

---

## ✨ Features

- **⚡ Lightning Fast**: In-memory storage for rapid read and write operations
- **⏱️ TTL Support**: Set expiration times for keys with automatic cleanup
- **🗂️ Multiple Databases**: Create and switch between different database instances
- **🖥️ Simple CLI**: Intuitive interactive command interface

## 🚀 Getting Started

### Requirements

- Go 1.18 or higher

### Installation & Running

```bash
# Clone the repository
git clone https://github.com/Tejasakumar/Key-Value-Store.git
cd KVS

# Run the interactive CLI
go run Main.go
```

## 📝 Command Reference

### Basic Operations

| Command | Description | Example |
|:--------|:------------|:--------|
| `put key value` | Store a key-value pair | `put username john` |
| `get key` | Retrieve a value by key | `get username` |
| `delete key` | Remove a key-value pair | `delete username` |


### TTL Operations

| Command | Description | Example |
|:--------|:------------|:--------|
| `setttl key seconds` | Set a time-to-live for a key | `setttl username 3600` |
| `getttl key` | Get remaining TTL for a key | `getttl username` |
| `rmttl key` | Remove TTL from a key | `rmttl username` |
| `upttl key seconds` | Update TTL for a key | `upttl username 7200` |
| `list` | List all keys in expiration order in the database | `list` |

### Database Management

| Command | Description | Example |
|:--------|:------------|:--------|
| `use dbname` | Switch to a specific database | `use users` |
| `showdbs` | List all available databases | `showdbs` |
| `curdb` | Show the current active database | `curdb` |
| `dropdb dbname` | Delete a database | `dropdb temp` |
| `exit` | Exit the program | `exit` |
| `keys` | List all keys in the database | `keys` |

## 💻 Example Session

```
Interactive Input Processor
Type 'exit' to quit
> use users
Creating new database: users
> curdb
Current database: users
> put name John
> put email john@example.com
> get name
 Value for 'name': John
> setttl email 3600
> getttl email
 TTL for 'email': 3600 seconds
> list
Keys in database:
 - name
 - email
> delete name
 Deleted key: name
> list
Keys in database:
 - email
> showdbs
Available databases:
 - default
 - users
> dropdb users
Database 'users' deleted
> use default
Switched to database: default
> exit
Goodbye!
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

<div align="center">
  <p>Made by <a href="https://github.com/yourusername">Tejas A Kumar</a></p>
</div>
