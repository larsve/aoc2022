package main

import (
	"bufio"
	"fmt"
	"io"
)

func part1Ofs(r io.Reader) int {
	b := bufio.NewReader(r)
	var sopm [4]byte
	ofs, err := b.Read(sopm[:])
	if err != nil {
		if err == io.EOF {
			return -1
		}
		fmt.Printf("Error %v while reading start-of-packet-marker\n", err)
		return -1
	}
	if ofs != 4 {
		return -1
	}

	for {
		if sopm[0] != sopm[1] && sopm[0] != sopm[2] && sopm[0] != sopm[3] &&
			sopm[1] != sopm[2] && sopm[1] != sopm[3] &&
			sopm[2] != sopm[3] {
			return ofs
		}
		d, err := b.ReadByte()
		if err != nil {
			if err == io.EOF {
				return -1
			}
			fmt.Printf("Error %v while reading data\n", err)
			return -1
		}
		sopm[ofs%4] = d
		ofs++
	}
}

func part2Ofs(r io.Reader) int {
	b := bufio.NewReader(r)
	var somm [14]byte
	ofs, err := b.Read(somm[:])
	if err != nil {
		if err == io.EOF {
			return -1
		}
		fmt.Printf("Error %v while reading start-of-message-marker\n", err)
		return -1
	}
	if ofs != 14 {
		return -1
	}

	buf := make(map[byte]int, 14)
	for _, d := range somm {
		buf[d]++
	}

	for {
		if len(buf) == 14 {
			return ofs
		}
		d, err := b.ReadByte()
		if err != nil {
			if err == io.EOF {
				return -1
			}
			fmt.Printf("Error %v while reading data\n", err)
			return -1
		}
		pos := ofs % 14
		ofs++
		if somm[pos] == d {
			continue
		}
		buf[somm[pos]]--
		if buf[somm[pos]] == 0 {
			delete(buf, somm[pos])
		}

		somm[pos] = d
		buf[d]++
	}
}
