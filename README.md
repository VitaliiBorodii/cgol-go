## Implementation of Convay's Game of Life in go


### Example of usage
`go run main.go --cycles 20`

will run `20` cycles of the game life. 

the program will print it's state after each cycle with `1s` delay

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
| patern | enum | `''` | pattern to use as a seed`*`|
| overflow | bool | `false` | allow to overflow the grid boundaries and make it cyclical|
| infinite | bool | `false` | make the game of life infinite (ignores the `cycles` flag)|
| speed | uint | 1 | animation speed (1 frame per second by default)|

`*`allowed values for a `pattern` argument: `glider`, `lwss`, `bee-hive`

in case you're running the game with `--infinite` flag - click `Ctrl+C` to stop it

`overflow` mode will allow cells to reach out of the grid size and appear on the opposite side, which creates a semblance of the infinite grid

`speed = 4` will result in 4x animation speed (`250ms` delay between life cycles)

-----
By default the game of life will start with a [`glider`](https://conwaylife.com/wiki/Glider) pattern in the middle of `25x25` grid
