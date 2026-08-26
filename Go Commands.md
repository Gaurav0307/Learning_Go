# Important Go Commands

## Initialize Go Project

```bash
 go mod init <Package-Name>
```

### Example:

```bash
 go mod init phonebook
```

## Run Go Program

```bash
 go run .
```

Or

```bash
 go run main.go
```

## Create Build of Go Project

```bash
 go build -o <File-Name>
```

### Example:

Windows

```bash
 go build -o phonebook.exe
```

Linux

```bash
 go build -o phonebook
```

### Or

```bash
 $ GOOS=<OS> GOARCH=<Architecture> go build
```

### Example:

Windows

```bash
 $ GOOS=windows GOARCH=amd64 go build
```

Linux

```bash
 $ GOOS=linux GOARCH=amd64 go build
```

macOS

```bash
 $ GOOS=darwin GOARCH=amd64 go build
```
