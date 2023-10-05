// unit test for memLib
//
// Author: prr, azul software
// Date: 5 Oct 2023
// copyright 2023 (c) prr, azul software
//

package MemLib

import (
	"testing"
	)

func TestMemLib(t* testing.T) {

	// creating two blocks
	mem, err := InitMemLib(2)
	if err !=nil {t.Errorf("error initMemLib: %v", err)}

	bytsl := (*mem.Ctl)
	size := len(bytsl)

	if size != BLOCKSIZE * 2 {t.Errorf("error -- %d != 4096*2", size)}

	(*mem.Ctl)[BLOCKSIZE]='7'

	val := 	(*mem.Ctl)[BLOCKSIZE]

	if val != '7' {t.Errorf("error value is not 7!")}

	return
}
