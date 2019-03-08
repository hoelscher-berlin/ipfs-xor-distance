package main

import (
	"fmt"
	b "math/bits"
	"os"
	"strconv"

	u "gx/ipfs/QmNohiVssaPw3KVLZik59DBVGTSm2dGvYT9eoXt5DQ36Yz/go-ipfs-util"

	kb "github.com/libp2p/go-libp2p-kbucket"
	peer "github.com/libp2p/go-libp2p-peer"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf(`
This tool calculates the XOR distance between two peer IDs.
Usage:

  %s {id1} {id2} {bytes to include}
		`, os.Args[0])
		os.Exit(1)
	}

	id1, err := peer.IDB58Decode(os.Args[1])
	if err != nil {
		fmt.Println("converting ID 1 failed: ", err)
	}

	id2, err := peer.IDB58Decode(os.Args[2])
	if err != nil {
		fmt.Println("converting ID 2 failed: ", err)
	}

	nrOfBytes, _ := strconv.Atoi(os.Args[3])

	xor := u.XOR(kb.ConvertPeerID(id1), kb.ConvertPeerID(id2))

	printByte(xor, nrOfBytes)

	xorInt := byteArrayToInt(xor, 4)

	leadingZeros := b.LeadingZeros32(uint32(xorInt))
	fmt.Println("Matching prefix:", leadingZeros)
}

func printByte(byteSlice []byte, bytes int) {
	for i := 0; i < bytes; i++ {
		fmt.Printf("%08b", byteSlice[i])
	}
	fmt.Print("\n")
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