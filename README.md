# 📝 Taskli

![Go version](https://img.shields.io/github/go-mod/go-version/carlosmms1/taskli)
![Project License](https://img.shields.io/github/license/carlosmms1/taskli)

A command-line (CLI) tasks manager application written in Go

## 🚀 Features

- Add new tasks
- List added tasks
- Mark tasks as done
- JSON file storage

## 📦 Installation

```bash
git clone https://github.com/carlosmms1/taskli
cd taskli
go mod tidy
```

## 🎉 How to use

```bash
go run cmd/main.go add "My new task"
# Task created succefully!

go run cmd/main.go list
# [] 1 - My new task

go run cmd/main.go done 1
# Task marked as done successfully!
```

## 📁 Project Structure

```
taskli/
├── cmd/
│   └── main.go        # entrypoint (main package)
├── internal/
│   └── task/
│       └── task.go    # task logic
├── tasks.json         # generated JSON file
└── go.mod             # go dependencies file
```

## 🛠️ Requirements

- Go 1.24.3 or higher

## 📃 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Author

Made with 🤍 by [Carlos Silva](https://github.com/carlosmms1)
