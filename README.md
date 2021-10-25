# kmap [![CI/CD](https://github.com/noah-friedman/kmap/actions/workflows/CICD.yml/badge.svg)](https://github.com/noah-friedman/kmap/actions/workflows/CICD.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/noah-friedman/kmap.svg)](https://pkg.go.dev/github.com/noah-friedman/kmap)

A program for generating [k-maps](https://en.wikipedia.org/wiki/Karnaugh_map) based on user input.

### EXAMPLES
```shell
> kmap
  
  Enter size: 3
  Enter arguments: 0,1,6,7
  
                y
    -----------------
    | 1 | 1 | 0 | 0 |
    -----------------
  x | 0 | 0 | 1 | 1 |
    -----------------
            z
```
```shell
> kmap -s 2 -a '1, 2'
          y
    ---------
    | 0 | 1 |
    ---------
  x | 1 | 0 |
    ---------
```

### INSTALLATION
- With Go 1.17+ installed:
  ```shell
  > go install github.com/noah-friedman/kmap/bin/kmap
  ```

- Without Go 1.17+ installed:
  
  See [Releases](https://github.com/noah-friedman/kmap/releases) for pre-compiled binaries.

### USAGE
After running the `kmap` program the user is prompted for two inputs:
- The size of the k-map (a.k.a. the number of variables in the k-map)
  
  Valid: [2-4]
- The arguments of the k-map

  Valid: [0-2<sup>size</sup>)

When inputting the arguments of the k-map the user may provide a series of valid numbers seperated by any characters as long as each separation contains the *same* characters. Otherwise, an error will occur.

Instead of inputting parameters directly, the user may provide arguments to the `kmap` program:
- `-s` or `--size` for the size parameter
- `-a` or `--args` for the arguments to the k-map
