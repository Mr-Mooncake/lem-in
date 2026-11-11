# Lem-In

A Go program that simulates an ant farm. It reads a colony of rooms and tunnels from a file and moves every ant from `##start` to `##end` in as few turns as possible.

## Usage

```bash
go run . test1.txt
```

The input file holds the number of ants, then the rooms (`name x y`), then the tunnels (`room1-room2`):

```
3
##start
1 23 3
2 16 7
##end
0 9 5
1-2
2-0
```

The program prints the input, then one line per turn. `L1-2` means ant 1 moves to room 2:

```
L1-3 L3-2
L1-4 L2-3 L3-5
L1-0 L2-4 L3-6
L2-0 L3-0
```

## How it works

1. Read and validate the file
2. Build a graph of rooms and tunnels
3. Find paths that do not share rooms (min-cost flow)
4. Spread the ants across the paths to finish in the fewest turns
5. Print the moves turn by turn, with at most one ant per room

## Testing

```bash
go test ./...
```

## Authors

- ahmsalem
- asami
- htamam
