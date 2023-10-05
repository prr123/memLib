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
	numBlocks := uint64(2)
	mem, err := InitMemLib(numBlocks)
	if err !=nil {t.Errorf("error initMemLib: %v", err)}

	bytsl := (*mem.Ctl)
	size := len(bytsl)
	if size != BLOCKSIZE {t.Errorf("error -- Ctl %d != 4096", size)}

	datsl := (*mem.Start)
	size = len(datsl)
	if size != int((numBlocks - uint64(1))*BLOCKSIZE) {t.Errorf("error -- Data %d != 4096", size)}


	datsl[1]='7'
	if datsl[1] != '7' {t.Errorf("error value is not 7!")}

	err = mem.Close()
	if err !=nil {t.Errorf("error closeMemLib: %v", err)}
	return
}
