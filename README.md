# wexec

`wexec` is a tool that executes files directly from the web. Instead of piping curl to bash you just use `wexec url arg1 arg2`

## Usage

wexec <url> [arguments...]

## Test this yourself

Make sure wexec is in PATH

run script

```shell
wexec https://raw.githubusercontent.com/dnkmmr69420/wexec/main/test.sh
```

try with one argument

```shell
wexec https://raw.githubusercontent.com/dnkmmr69420/wexec/main/test.sh arg1
```

Try with multiple arguments

```shell
wexec https://raw.githubusercontent.com/dnkmmr69420/wexec/main/test.sh arg1 arg2 arg3 arg4 arg5
```

Works with flags

```shell
wexec https://raw.githubusercontent.com/dnkmmr69420/wexec/main/test.sh arg1 arg2 arg3 arg4 arg5 -sdw --flag-string
```

Exit code 1 test

```shell
wexec https://raw.githubusercontent.com/dnkmmr69420/wexec/main/exit-code-1.sh
```
