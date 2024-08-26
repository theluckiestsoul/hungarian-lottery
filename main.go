package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"sync"
	"time"
)

type Player struct {
	numbers [5]int
}

var validPlayer = regexp.MustCompile(`^\d+ \d+ \d+ \d+ \d+$`)

// NewPlayer creates a new player from a string
func NewPlayer(s string) (Player, error) {
	if !validPlayer.MatchString(s) {
		return Player{}, fmt.Errorf("invalid input format")
	}
	var numbers [5]int
	fmt.Sscanf(s, "%d %d %d %d %d", &numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])

	// Check for distinct numbers
	numberSet := make(map[int]struct{})
	for _, number := range numbers {
		if _, exists := numberSet[number]; exists {
			return Player{}, fmt.Errorf("numbers must be distinct")
		}
		numberSet[number] = struct{}{}
	}

	return Player{numbers}, nil
}

type Pick struct {
	numbers [5]int
}

var validPick = regexp.MustCompile(`^\d+ \d+ \d+ \d+ \d+$`)

// NewPick creates a new pick from a string
func NewPick(s string) (Pick, error) {
	if !validPick.MatchString(s) {
		return Pick{}, fmt.Errorf("invalid input format")
	}
	var numbers [5]int
	fmt.Sscanf(s, "%d %d %d %d %d", &numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])

	// Check for distinct numbers
	numberSet := make(map[int]struct{})
	for _, number := range numbers {
		if _, exists := numberSet[number]; exists {
			return Pick{}, fmt.Errorf("numbers must be distinct")
		}
		numberSet[number] = struct{}{}
	}

	return Pick{numbers}, nil
}

// CountWinners counts the winners for a given lottery pick using concurrency
func (pick Pick) CountWinners(players []Player) [4]int {
	var winners [4]int
	var pickSet [91]bool // Assuming lottery numbers are between 1 and 90

	for _, number := range pick.numbers {
		if number >= 1 && number <= 90 {
			pickSet[number] = true
		}
	}

	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup
	chunkSize := (len(players) + numCPU - 1) / numCPU

	localWinners := make([][4]int, numCPU)

	for i := 0; i < numCPU; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= len(players) {
			break
		}
		if end > len(players) {
			end = len(players)
		}

		wg.Add(1)
		go func(i int, playersChunk []Player) {
			defer wg.Done()
			for _, player := range playersChunk {
				matches := 0
				for _, number := range player.numbers {
					if number >= 1 && number <= 90 && pickSet[number] {
						matches++
					}
				}
				if matches >= 2 {
					localWinners[i][matches-2]++
				}
			}
		}(i, players[start:end])
	}

	wg.Wait()

	for i := 0; i < numCPU; i++ {
		for j := 0; j < 4; j++ {
			winners[j] += localWinners[i][j]
		}
	}

	return winners
}

func main() {
	fileName := parseFlags()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := make(chan string, 100)
	players := make(chan Player, 100)
	var wg sync.WaitGroup

	// Start worker pool
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lines {
				player, err := NewPlayer(line)
				if err != nil {
					fmt.Printf("skipping invalid line: %s, error: %v\n", line, err)
					continue
				}
				players <- player
			}
		}()
	}

	// Read the file line by line and send lines to workers
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}()

	// Close players channel when all workers are done
	go func() {
		wg.Wait()
		close(players)
	}()

	// Collect players
	var allPlayers []Player
	for player := range players {
		allPlayers = append(allPlayers, player)
	}

	fmt.Println("READY")

	// Read lottery picks and print winners
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pick, err := NewPick(scanner.Text())
		if err != nil {
			fmt.Printf("invalid input format: %s\n", scanner.Text())
			continue
		}
		start := time.Now()
		winners := pick.CountWinners(allPlayers)
		fmt.Printf("%d %d %d %d\n", winners[0], winners[1], winners[2], winners[3])
		fmt.Fprintf(os.Stderr, "Elapsed time: %v ms\n", time.Since(start).Milliseconds())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
