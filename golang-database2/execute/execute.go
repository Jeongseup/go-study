package execute

import (
	"fmt"
	"golang-database2/constants"
	"strings"
)

// const (
// 	PREPARE_UNRECONIZED_STATEMENT = 999
// 	PREPARE_SUCCESS               = 0

// 	STATEMENT_INSERT = 1
// 	STATEMENT_SELECT = 2
// )

func PrepareStatement(input string, statementType *int) int {
	if strings.Compare("insert", input[:6]) == 0 {
		*statementType = constants.STATEMENT_INSERT
		return constants.PREPARE_SUCCESS
	} else if strings.Compare("select", input[:6]) == 0 {
		*statementType = constants.STATEMENT_SELECT
		return constants.PREPARE_SUCCESS
	}

	return constants.PREPARE_UNRECOGNIZED_STATEMENT
}

func ExcuteStatement(statementType int) {
	switch statementType {
	case constants.STATEMENT_INSERT:
		fmt.Println("this is where we will do an insert")
		return
	case constants.STATEMENT_SELECT:
		fmt.Println("this is where we will do a select")
		return
	}
}
