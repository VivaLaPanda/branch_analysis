package branching_model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BranchModel struct {
	branchTakens  map[string]int
	branchTotals  map[string]int
	branchIns     map[uint64]BranchInstruction
	historyLength int
}

type BranchInstruction struct {
	insAddr       uint64
	branchHistory []bool
}

func New(historyLength int) *BranchModel {
	branchModel := &BranchModel{}
	branchModel.historyLength = historyLength
	branchModel.branchTakens = make(map[string]int)
	branchModel.branchTotals = make(map[string]int)
	branchModel.branchIns = make(map[uint64]BranchInstruction)

	return branchModel
}

func (bm *BranchModel) ParseFile(filename string) *BranchModel {
	file, err := os.Open(filename) // just pass the file name
	if err != nil {
		check(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Grab a line from the file
		rowStr := scanner.Text()
		// If it's as blank line just skip it
		if rowStr == "" {
			continue
		}

		// Parse the line into the values that we need
		keyValStr := strings.Split(rowStr, ":")
		taken, b_err := strconv.ParseBool(keyValStr[0])
		addr, i_err := strconv.ParseUint(keyValStr[1], 16, 64)
		check(b_err)
		check(i_err)

		bm.tallyBranch(taken, addr)
	}

	return bm
}

func (bm *BranchModel) DisplayStatistics() {
	for key := range bm.branchTakens {
		branchRate := float64(bm.branchTakens[key]) / float64(bm.branchTotals[key])

		fmt.Printf("%v: %f\n", key, branchRate)
	}
}

func (bm *BranchModel) tallyBranch(taken bool, addr uint64) *BranchModel {
	branchIns := bm.branchIns[addr]

	// The branch address given has no data for it so far, so do some
	// initialization
	if branchIns.insAddr == 0 {
		// Declare the default branch history to be alternating true and false
		// of the length provided druing the bm initialization
		tempBrHistory := make([]bool, bm.historyLength)
		for i, _ := range tempBrHistory {
			if i > 0 {
				tempBrHistory[i] = !tempBrHistory[i-1]
			} else {
				tempBrHistory[i] = true
			}
		}

		branchIns = BranchInstruction{insAddr: addr, branchHistory: tempBrHistory}
	}

	// Record the change in state
	bm.branchIns[addr] = branchIns.takePath(taken, bm)
	return bm
}

func (bi BranchInstruction) takePath(taken bool, model *BranchModel) BranchInstruction {
	// Determine what state we will be in before and after the instruction
	// then stringfy it so we can use it as a map key
	stateIs := bi.branchHistory
	newState := append(bi.branchHistory[1:], taken)
	edge := stringifyEdge(stateIs, newState)

	// Record the relevant statistics
	model.branchTotals[edge]++
	if taken {
		model.branchTakens[edge]++
	}

	// Return the new state of the branch instruction
	bi.branchHistory = newState
	return bi
}

func stringifyEdge(stateIs []bool, newState []bool) string {
	return fmt.Sprintf("%v", stateIs)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
