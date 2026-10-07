# HardPass

A minimal command-line password strength analyzer written in Go, styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss).

```
  HardPass  password strength analyzer
  ──────────────────────────────────────
  length      [✓]  14 chars (min 12)
  uppercase   [✓]
  lowercase   [✓]
  numbers     [✗]
  symbols     [✓]
  ──────────────────────────────────────
  strength    ████░  medium
```

## Features

- Checks for uppercase, lowercase, numbers and symbols
- Checks for a minimum length of 12 characters
- Unicode-aware (counts characters, not bytes)
- Colored verdict: **weak**, **medium** or **strong**

## Installation

Requires [Go](https://go.dev/dl/) 1.21 or newer.

```bash
go install github.com/carlosfdsn/hardpass@latest
```

Or clone and build it yourself:

```bash
git clone https://github.com/carlosfdsn/hardpass.git
cd hardpass
go build -o hardpass .
```

## Usage

```bash
hardpass <password>
```

Example:

```bash
hardpass 'MyP@ssw0rd2024'
```

Wrap the password in single quotes so your shell does not interpret special characters like `!`, `$` or `&`.

## How the score works

Each criterion met adds 1 point (max 5):

| Criterion        | Points |
| ---------------- | ------ |
| Uppercase letter | 1      |
| Lowercase letter | 1      |
| Number           | 1      |
| Symbol           | 1      |
| 12+ characters   | 1      |

| Score | Verdict |
| ----- | ------- |
| 0-2   | weak    |
| 3-4   | medium  |
| 5     | strong  |

## Security note

Passwords passed as command-line arguments may be saved in your shell history and are visible to other processes. Do not analyze real, important passwords this way. Use test passwords or clear your history afterwards.

This is a learning project. The score only counts character types and length, so it does not detect common passwords, sequences (`1234`, `abcd`) or leaked passwords.

## Roadmap

- [ ] Common password list check
- [ ] Sequence and repetition detection
- [ ] Leetspeak normalization (`p@ssw0rd`)
- [ ] Entropy estimation
- [ ] Read password from stdin

## Built with

- [Go](https://go.dev/)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## License

This project is licensed under the [MIT License](LICENSE).