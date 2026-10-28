# mcls

![Go](https://img.shields.io/badge/go-1.21%2B-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-Apache%202.0-green)
![Stars](https://img.shields.io/github/stars/MohssineX/mcls?style=social)

**mcls** is a simple command-line tool written in **Go** that clears the screen and the scrollback buffer

## Features

- Clears the screen and the scrollback buffer
- Single binary with no external dependencies

> [!IMPORTANT]
> Make sure your system or terminal supports ANSI escape sequences by default.
> Otherwise, `mcls` may not work correctly.

## Supported Platforms

- Windows
- Linux
- macOS
- BSD (FreeBSD, OpenBSD, NetBSD)

Tested on : Linux, Windows 

## Requirements

- Go 1.21 or later (only needed to build or install from source)
- A terminal with ANSI escape sequence support

## Installation

### Using go install

```bash
go install github.com/MohssineX/mcls/src/mcls@latest
```

Then run :

```bash
mcls
```

> [!NOTE]
> If the `mcls` command is not found, add your Go `bin` folder to your `PATH`.

### Build from source

```bash
git clone https://github.com/MohssineX/mcls.git
cd mcls
```

```bash
go build ./src/mcls
```

Then run it from the current folder :

```bash
./mcls
```

## License

Licensed under the [Apache License 2.0](LICENSE).

## Author

**Mohssine :**
[https://github.com/MohssineX](https://github.com/MohssineX)

---

## Special Thanks 🐱

Special thanks to **Mimi**, the legendary, great, and gentle cat.

---

### If you like it, give it a star :)
