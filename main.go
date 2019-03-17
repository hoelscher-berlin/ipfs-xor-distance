package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"

	u "github.com/ipfs/go-ipfs-util"
	kb "github.com/libp2p/go-libp2p-kbucket"
	peer "github.com/libp2p/go-libp2p-peer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(`
This tool calculates the matching prefix of two IPFS peer IDs.
Usage:

  %s {id1} {id2} {bytes to include}
		`, os.Args[0])
		os.Exit(1)
	}

	if os.Args[1] == "-l" {
		handleList(os.Args[2])
	} else {
		fmt.Println("Matching prefix: ", matchingPrefix(os.Args[1], os.Args[2]))
	}
}

func handleList(path string) {
	file, err := os.Open(path)
	check(err)
	line := ""
	//avg := 0

	compareTo := os.Args[3]

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line, " ", matchingPrefix(line, compareTo))
	}
}

func matchingPrefix(a, b string) int {
	id1, err := peer.IDB58Decode(a)
	if err != nil {
		fmt.Println("converting ID 1 failed: ", err)
	}

	id2, err := peer.IDB58Decode(b)
	if err != nil {
		fmt.Println("converting ID 2 failed: ", err)
	}

	xor := u.XOR(kb.ConvertPeerID(id1), kb.ConvertPeerID(id2))

	xorInt := byteArrayToInt(xor, 4)

	leadingZeros := bits.LeadingZeros32(uint32(xorInt))
	return leadingZeros
}

func power(a, n int) int {
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}

func byteArrayToInt(byteSlice []byte, bytes int) int {
	sum := 0
	for i := 0; i < bytes; i++ {
		sum = sum + power(2, ((bytes-i-1)*8))*int(byteSlice[i])
	}

	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
