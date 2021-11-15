# kmap [![CI/CD](https://github.com/noah-friedman/kmap/actions/workflows/CICD.yml/badge.svg)](https://github.com/noah-friedman/kmap/actions/workflows/CICD.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/noah-friedman/kmap.svg)](https://pkg.go.dev/github.com/noah-friedman/kmap)

A program for generating [k-maps](https://en.wikipedia.org/wiki/Karnaugh_map) based on user input.

### EXAMPLES
```shell
> kmap
  
  What is the size of the k-map? (3): 
  3
  What are the arguments to the k-map?: 
  0,1,6,7
  What are the don`t care conditions of of the k-map?:
  2, 4
  
                y
    -----------------
    | 1 | 1 | 0 | X |
    -----------------
  x | X | 0 | 1 | 1 |
    -----------------
            z
```
```shell
> kmap -s 2 -a '1, 2' -dc ''
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
  > go install github.com/noah-friedman/kmap/bin/kmap@latest
  ```

- Without Go 1.17+ installed:
  
  See [Releases](https://github.com/noah-friedman/kmap/releases) for pre-compiled binaries.

### USAGE
After running the `kmap` program the user is prompted for three inputs:
- The size of the k-map (a.k.a. the number of variables in the k-map)
  
  Valid: [2-4]
- The arguments of the k-map

  Valid: [0-2<sup>size</sup>)
- The don't care conditions of the k-map
  
  Valid: [0-2<sup>size</sup>)

**NOTE**: The arguments and don't care conditions must not overlap.

When inputting the arguments of the k-map the user may provide a series of valid numbers seperated by any characters as long as each separation contains the *same* characters. Otherwise, an error will occur.

Instead of inputting parameters directly, the user may provide arguments to the `kmap` program:
- `-s` or `-size` for the size parameter
- `-a` or `-args` for the arguments to the k-map
- `-dc` or `-dont-care` for the don't care conditions of the k-map

### EXIT CODES
- `0`: Everything went fine, the program executed successfully
- `1`: Some kind of bug or error in the way I wrote the program.
- `2`: Some kind of bug in a library called by the program (possibly due to a misunderstanding of the library on my part).