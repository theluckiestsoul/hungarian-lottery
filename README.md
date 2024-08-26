# hungarian-lottery

## Description
This problem is related to the Hungarian lottery. 

## How to run
This program is written using Go(1.22.3). To run the program, you need to have Go installed on your machine. You can download Go from [here](https://golang.org/dl/). After installing Go, you can run the program by running the following command in the terminal:
```bash
make build
```
This command will build the program and create an executable file named `hungarian-lottery` in the `bin` folder. You can run this executable file by running the following command:
```bash
bin/hungarian-lottery <input_file>
```
The `<input_file>` is the path to the input file. The input file should contain the players' numbers. The numbers should be separated by a space. You can find sample input files in the `testfiles` folder.

Once you run the program, it will process the file ans store the players' numbers in memory. Then it will start listening for the winning numbers. You can provide the winning numbers after you see the `READY` message in the terminal. The winning numbers should be provided in the same format as the players' numbers. The program will then calculate the winners and print the result in the terminal.

## Assumptions
- If file contains invalid numbers, the program will ignore those numbers and continue processing the file.

## Asymptotic analysis

- Reading and parsing the input file:
    - Reading Lines: O(n)
    - Parsing Each Line: O(n)
- Counting Winners:
    - Marking Pick Numbers: O(1)
    - Dividing work among goroutines: 
        - The number of goroutines is proportional to the number of CPU cores `p`. This setup is: O(p)
    - Processing Each Player:
        - Each player is processed to count matches. This involves iterating over 5 numbers and checking against the pickset, which is: O(1)
        - Since there are n players, the total time complexity is: O(n)
    - Aggregating Results:
        - Aggregating the results from `p` goroutines: O(p)

- Overall Time Complexity:
    - Reading and parsing the input file: O(n)
    - Counting Winners: O(n)
Hence, the overall time complexity of the program is O(n) where n is the number of players.

## How the code could be improved further
Improving performance depends on various factors. Here are some ways to improve the performance of the program:
- Currently, we use an array to store the pick numbers. We can use a bitset which can be more cache friendly and faster for comparison operations.
- We can use SIMD instructions to speed up the comparison operations.
- We can use Go's profiling tools to identify bottlenecks in the program and optimize them.
- We need to find the optimal number of goroutines to use for processing the players. This can be done by profiling the program and finding the optimal number of goroutines that gives the best performance.
- If we have large number of players, we can use a distributed system to process the players in parallel.
- Instead of reading the file line by line, we can read the file in chunks and process the chunks in parallel.
- Without loading the entire file into memory, we can use a streaming approach to process the file in chunks.
- We can store the players in a database and use SQL queries to find the winners. This can be faster than processing the players in memory.
- These are some of the ways to improve the performance of the program. The best approach depends on the specific requirements and constraints of the problem.

Note: The program is tested on a machine with 8 CPU cores and 16GB RAM. The performance may vary depending on the machine configuration and the number of players.