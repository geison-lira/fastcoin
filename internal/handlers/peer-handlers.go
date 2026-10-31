package handlers

import (
	"bufio"
	"os"
	"strings"
)

func PeersLoad(peers string) ([]string, error) {
	file, err := os.Open(peers)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var parsedPeers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			parsedPeers = append(parsedPeers, line)
		}
	}
	return parsedPeers, scanner.Err()
}
