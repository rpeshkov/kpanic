# kpanic

Command-line utility that simplifies viewing of MacOS kernel panics in your terminal window. Panics are taken from default location: /Library/Logs/DiagnosticReports/

## Installation

### Homebrew

```
brew tap rpeshkov/kpanic
brew install rpeshkov/kpanic/kpanic
```

### Manual

- Go to [releases](releases) page and download latest tar.gz archive.
- Unpack it
- Move binary to /usr/local/bin/kpanic

## Commands

`kpanic ls`

Display list of available kernel panics.

`kpanic show <name>`

Display contents of panic file in your terminal window. Name parameters refers to the panic name you may obtain from `kpanic ls` command.

`kpanic last`

Display content of last occurred panic in your terminal window.

## Shell completion

You may find Bash and ZSH completion files in [shell](shell) folder.

## TODO

- [ ] atos integration
- [ ] otool integration
- [ ] configuration

## [License](LICENSE)

The MIT License (MIT)

Copyright (c) 2018 Roman Peshkov
