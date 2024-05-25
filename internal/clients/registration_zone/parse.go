package registration_zone

import (
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"

	"github.com/ag7if/wendover/pkg/personnel"
)

const (
	capidCol = iota
	lastNameCol
	firstNameCol
	middleNameCol
	memberTypeCol
	gradeCol
	genderCol
	ageAtEventStartCol
	ageAtEventEndCol
	shirtSizeCol
	membershipExpiresCol
	eServicesEmailCol
	unitApprovalDateCol
	cadetParentPrimaryPhoneCol
	cadetParentPrimaryEmailCol
	unitCCNameCol
	unitCCEmailCol
	wingCCNameCol
	wingCCEmailCol
	cpptExpiresCol
	icutDateCol
	capDlExpiresCol
)

var columns = map[int]string{
	capidCol:                   "A",
	lastNameCol:                "G",
	firstNameCol:               "H",
	middleNameCol:              "I",
	memberTypeCol:              "U",
	gradeCol:                   "F",
	genderCol:                  "M",
	ageAtEventStartCol:         "P",
	ageAtEventEndCol:           "Q",
	shirtSizeCol:               "T",
	membershipExpiresCol:       "V",
	eServicesEmailCol:          "Z",
	unitApprovalDateCol:        "AQ",
	cadetParentPrimaryPhoneCol: "AW",
	cadetParentPrimaryEmailCol: "AZ",
	unitCCNameCol:              "BA",
	unitCCEmailCol:             "BB",
	wingCCNameCol:              "BC",
	wingCCEmailCol:             "BD",
	cpptExpiresCol:             "BN",
	icutDateCol:                "BQ",
	capDlExpiresCol:            "BF",
}

func lookUpColumnIndex(col int) int {
	idx, err := excelize.ColumnNameToNumber(columns[col])
	if err != nil {
		panic(errors.WithStack(err))
	}

	return idx
}

func readRow(row []string) (personnel.Participant, error) {
	capid := row[lookUpColumnIndex(capidCol)]

}

func ParseParticipantsFromRZExport(file files.File) ([]personnel.Participant, error) {
	f, err := excelize.OpenFile(file.FullPath())
	if err != nil {
		log.Error().Err(err).Str("path", file.FullPath()).Msg("Failed to open spreadsheet.")
		return nil, errors.WithStack(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error().Err(err).Str("path", file.FullPath()).Msg("Error thrown when attempting to close spreadsheet.")
		}
	}()

	var participants []personnel.Participant
	sheetName := f.GetSheetList()[0]
	rows, err := f.GetRows(sheetName)
	for i, row := range rows {
		p, err := readRow(row)
		if err != nil {
			log.Warn().Err(err).Str("path", file.FullPath()).Int("row", i).Msg("Failed to parse row. Skipping.")
		}

		participants = append(participants, p)
	}

	return participants, nil
}
