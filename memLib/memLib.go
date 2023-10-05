// MemLib
// library that allocates memory outside the go gc
//
// Author: prr azulsoftware
// Date:   5 Oct 2023
// copyright (c) 2023 prr, azul software
//
//
package MemLib

import (
	"fmt"
	"golang.org/x/sys/unix"
	)

const BLOCKSIZE = 4096

type memObj struct {
	Size uint64 // memory in blocks
	Unit uint64
	Free uint64
	Start *[]byte
	Ctl *[]byte
}

func InitMemLib(blocks uint64) (mem *memObj, err error){

	memsize := blocks*BLOCKSIZE

	data , err := unix.Mmap(
		-1,                                  // required by MAP_ANONYMOUS
		0,                                   // offset from file descriptor start, required by MAP_ANONYMOUS
		int(memsize),                                // how much memory
		unix.PROT_READ|unix.PROT_WRITE,      // protection on memory
		unix.MAP_ANONYMOUS|unix.MAP_PRIVATE, // private so other processes don't see the changes, anonymous so that nothing gets synced to the file
	)

	if err != nil {return nil, fmt.Errorf("MMap init: %v", err)}

	memobj := memObj {
		Size: memsize,
		Free: memsize - BLOCKSIZE,
		Ctl : &data,
//		Start: &(data[BLOCKSIZE:]),
	}

	return &memobj, nil
}


func (mem *memObj)Close()(err error) {

	b := *(mem.Ctl)
	err = unix.Munmap(b)

	if err != nil {return fmt.Errorf("Unmap: %v", err)}

	mem = nil
	return nil
}
