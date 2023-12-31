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

//const BLOCKSIZE = 4096

type memObj struct {
	BlkSize uint64
	Size uint64 // memory in blocks
	Unit uint64
	Free uint64
	all *[]byte
	Start *[]byte
	Ctl *[]byte
}

func InitMemLib(blocks uint64) (mem *memObj, err error){

	blksize := unix.Getpagesize()
	memsize := blocks*uint64(blksize)

	data , err := unix.Mmap(
		-1,                                  // required by MAP_ANONYMOUS
		0,                                   // offset from file descriptor start, required by MAP_ANONYMOUS
		int(memsize),                                // how much memory
		unix.PROT_READ|unix.PROT_WRITE,      // protection on memory
		unix.MAP_ANONYMOUS|unix.MAP_PRIVATE, // private so other processes don't see the changes, anonymous so that nothing gets synced to the file
	)

	if err != nil {return nil, fmt.Errorf("MMap init: %v", err)}

	ctlslice := data[:blksize]
	newslice := data[blksize:]


	memobj := memObj {
		BlkSize: uint64(blksize),
		Size: memsize,
		Free: memsize - uint64(blksize),
		all: &data,       // pointer to total byte slice
		Ctl : &ctlslice,  // pointer to control block slice
		Start: &newslice, // pointer to data slice
	}

	return &memobj, nil
}


func (mem *memObj)Close()(err error) {

	b := *(mem.all)
	err = unix.Munmap(b)

	if err != nil {return fmt.Errorf("Unmap: %v", err)}

	mem = nil
	return nil
}

func GetBlockSize()(size int){

	return unix.Getpagesize()
}
