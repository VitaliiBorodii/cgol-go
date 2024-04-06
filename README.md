## Implementation of Convay's Game of Life in go


### Example of usage
`go run main.go --cycles 20`

will run `20` cycles of the game life. 

the program will print it's state after each cycle with `100ms` delay

example output:
```
cycle: 20 | alive cells: 5
____________________________________________________
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                   *              |
|                                     *            |
|                                 * * *            |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
|                                                  |
----------------------------------------------------
```

#### command line arguments:
| argument | type | default value | usage |
| - | - | - | - |
| cycles | int | 0 | number of life cycles to run |
| xsize | int | 25 | width of the grid |
| ysize | int | 25 | height of the grid |
| patern | enum | `'glider'` | pattern to use as a seed for the 'life'`*`|

`*`allowed values for a `pattern` argument: `glider`, `lwss`, `bee-hive`

-----
By default the game of life will start with a [`glider`](https://conwaylife.com/wiki/Glider) pattern in the middle of `25x25` grid
