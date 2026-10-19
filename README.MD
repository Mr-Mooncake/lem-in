# Lem-In

## Description

Lem-in is a Go program that simulates an ant colony. The goal is to move all ants from the start room to the end room in the fewest possible turns.

The program reads a file containing rooms, tunnels, and ants, finds efficient paths, and displays the ants' movements.

## Authors

- ahmsalem
- asami
- htamam

## Usage: how to run



From the root project directory, :

```bash
go run . test0.txt
```

## Implementation details: algorithm

The program uses a graph to represent the ant colony. Each room is a node, and each tunnel is an edge connecting two rooms.

The algorithm works as follows:

1. Parse the input: Read the number of ants, rooms, and tunnels, and validate the data.

2. Build the graph: Store each room and its connected neighbors.

3. Find paths: Use a minimum-cost flow approach to find paths from the start room to the end room while avoiding shared intermediate rooms.

4. Extract paths: Convert the selected routes into sequences of room names.

5. Distribute ants: Assign ants to paths based on path length and the number of ants already assigned.

6. Optimize turns: Compare different path combinations and keep the one that allows all ants to finish in the fewest turns.

7. Display movements: Print the ants' movements turn by turn, ensuring that intermediate rooms are not occupied by multiple ants simultaneously.