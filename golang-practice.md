# Golang Practice

Implement a command-line application `fops` written in **Golang** that satisfies following requirements:

1. Specific `fops` features (see below).
2. Push your code to a GitHub public repository with all commit history.
3. A `README.md` describes the application design, 3rd-party libraries you used, how to build/run your project, what features not yet implemented and known issues.
4. Integrate with GitHub Actions, Travis CI or CircleCI for
   - Unit tests (_bonus1_)
   - Build (_bonus2_)
   - Publish release (_bonus3_)

**Notes**
If you’re not familiar with Golang, do the Go Tour first. Feel free to use any go libraries you can find. (eg. Awesome Go)

## Features

1. Prepare input file

```bash
$ cat <<EOF > myfile.txt
how do
you
turn this
on
EOF
```

2. Print the line count of file, support both short and long options

```bash
$ fops linecount -f myfile.tx
4
$ fops linecount --file myfile.txt
4
```

3. Print the checksum of file, support multiple algorithms: md5, sha1 and sha256

```bash
$ fops checksum -f myfile.txt --md5
a8c5d553ed101646036a811772ffbdd8

$ fops checksum -f myfile.txt --sha1
a656582ca3143a5f48718f4a15e7df018d286521

$ fops checksum -f myfile.txt --sha256
495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d
```

4. Handle non-existent, invalid input file

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

5. Show version and help

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

6. Show subcommand help

```bash
$ fops help linecount
Print line count of file

Usage:
  fops linecount [flags]

Flags:
  -f, --file   the input file
```
