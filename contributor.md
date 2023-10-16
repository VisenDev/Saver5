# Explanation of source files
- `main.go` holds the program startup code, as well as important data structures like `SerialConfig`
- `upload.go` holds the ui for the upload menu
- `config.go` holds the ui for the config menu
- `download.go` holds the ui for the download menu
- `help.go` holds the ui for the help menu
- `model.go` holds the system mode and upload/download backend

# Naming Conventions
- Public functions should be `PascalCase`
- Private functions should be `camelCase`
- Public struct fields should be `PascalCase`
- Private struct fields should be `camelCase`
- Variables should be `snake_case`
