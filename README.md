# Golang-Practice


File Operations (Fops) CLI, a command written in Golang, can be used to count lines and verify files.
## How to build/run this project


## Commands

### Features

This CLI has the following commands:

- `linecount` - Calculate the number of lines in the document

- `checksum` - Calculate the checksum of the file using md5, sha1 or sha256

- `version` - Get the version of the command

- `help` - Show help


### Command 

The command follows the following format:
```bash
fops <command> <file> <flags>
```

### Command examples

Print the line count of file, support both short and long options:
```bash
fops linecount -f myfile.tx
fops linecount --file myfile.txt
```

Print the checksum of file, support multiple algorithms: md5, sha1 and sha256:
```bash
fops checksum -f myfile.txt --md5
fops checksum --file myfile.txt --sha1
fops checksum -f myfile.txt --sha256
```

Handle non-existent, invalid input file:
```bash
$ fops linecount -f non-exist-file.tx
error: No such file 'non-exist-file.txt'

$ fops linecount -f /tmp
error: Expected file got directory '/tmp'

$ fops linecount -f fops
error: Cannot do linecount for binary file 'fops'

$ fops checksum -f fops --sha256
f07bb6a888308db77fda750aa3739b7c643b07675c5c6a2d6de6c9e69de05ceb
```

Show version and help:
```bash
$ fops version
fops v0.0.1

$ fops help
File Ops

Usage:
  fops [flags]
  fops [command]

Available Commands:
  linecount    Print line count of file
  checksum     Print checksum of file
  version      Show the version info
  help         Help about commands

Flags:
  -h, --help   help for fops
```

Show subcommand help
```bash
$ fops help linecount
Print line count of file

Usage:
  fops linecount [flags]

Flags:
  -f, --file   the input file
```



## List of features to be accomplished


### linecount
- [x] verify file
- [x] read file
- [x] count lines in file

### file verification
- [x] non-existent
- [x] is a dir
- [x] is binary

### checksum
- [x] implement 3 algorithm with different flags

### version
- [x] get from build

### help
- [x] subcommands

### Integrate withCI
- [x] Unit tests
- [ ] Integrate with CI
- [ ] Build
- [ ] Release

## 3rd-party libraries I used
- [cobra](https://github.com/spf13/cobra) v1.7.0
- [mimetype](https://github.com/gabriel-vasile/mimetype) v1.4.2
- [testify](https://github.com/stretchr/testify) v1.8.3


