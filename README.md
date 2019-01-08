# MargeTexFile
this script can marge splited tex files by [subfiles](https://ctan.org/pkg/subfiles)

When you use subfile, you couldn't use `\ref{hoge}` to refer `\label{hoge}` written by another file.

### This script will resolve this problem.

# Getting Started
download this script and just execute

```
go run MargeTex.go -m ]path/to/master.tex] -o [path/to/output.tex]
```
if `master.tex` call subfiles like `\subfile{path/to/subfile.tex}` whether or not including `\ref` which is not refer correctly, a large file `output.tex`will be created based on `master.tex` and expanded `subfile.tex` expanded with resolving `\ref`, `\label`.

# Options
- `-m` : master Tex file path (default `./master.tex`)
- `-o` : output Tex file path (default `./output.tex`)
- `-h` : help
