# Tetris Optimizer

This project aims to solve Tetris puzzles by finding the optimal arrangement of tetrominos within a grid.

## Installation
1. Clone the repository:
   ```bash
   git clone https://learn.zone01kisumu.ke/git/johnotieno0/tetris-optimizer.git
   ```
## Usage
1. Run the following command to start the application:
    ```bash
    go run main.go path/to/tetromino_file.txt
    ```
2. Example of a text File: sample.txt

```
#...
#...
#...
#...

....
....
..##
..##
```

If it isn't possible to form a complete square, the program should leave spaces between the tetrominoes. For example:

```
ABB.
ABB.
A...
A...
```


Another example of a sample.txt
```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
```

```bash
 go run . sample.txt | cat -e
```

The output should look like this:
```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```



## Contributing

Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the MIT License.

## Author
[John Paul Nyunja](https://learn.zone01kisumu.ke/git/johnotieno0)

Apprenctice Software Developer at Zone01Kisumu