package registration_zone

import (
	"strconv"
	"strings"
	"time"

	"github.com/ag7if/go-files"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"

	"github.com/ag7if/wendover/pkg/personnel"
)

const nhqTimeFormat = `02 Jan 2006`

// TODO: Make this configurable
var (
	capidCol                   = mustColumnNameToIndex("A")
	lastNameCol                = mustColumnNameToIndex("G")
	firstNameCol               = mustColumnNameToIndex("H")
	middleNameCol              = mustColumnNameToIndex("I")
	memberTypeCol              = mustColumnNameToIndex("U")
	gradeCol                   = mustColumnNameToIndex("F")
	genderCol                  = mustColumnNameToIndex("M")
	ageAtEventStartCol         = mustColumnNameToIndex("P")
	ageAtEventEndCol           = mustColumnNameToIndex("Q")
	shirtSizeCol               = mustColumnNameToIndex("T")
	membershipExpiresCol       = mustColumnNameToIndex("V")
	eServicesEmailCol          = mustColumnNameToIndex("Z")
	regionCol                  = mustColumnNameToIndex("J")
	wingCol                    = mustColumnNameToIndex("K")
	unitNumberCol              = mustColumnNameToIndex("L")
	unitApprovalDateCol        = mustColumnNameToIndex("AQ")
	cadetParentPrimaryPhoneCol = mustColumnNameToIndex("AW")
	cadetParentPrimaryEmailCol = mustColumnNameToIndex("AZ")
	unitCCNameCol              = mustColumnNameToIndex("BA")
	unitCCEmailCol             = mustColumnNameToIndex("BB")
	wingCCNameCol              = mustColumnNameToIndex("BC")
	wingCCEmailCol             = mustColumnNameToIndex("BD")
	cpptExpiresCol             = mustColumnNameToIndex("BN")
	icutDateCol                = mustColumnNameToIndex("BQ")
	capDlExpiresCol            = mustColumnNameToIndex("BF")
)

func mustColumnNameToIndex(col string) int {
	idx, err := excelize.ColumnNameToNumber(col)
	if err != nil {
		panic(errors.WithStack(err))
	}

	return idx - 1
}

func parseNullableTime(s string) (*time.Time, error) {
	timeStr := strings.TrimSpace(s)
	if timeStr == "" || strings.ToUpper(timeStr) == "NOT COMPLETE" || strings.ToUpper(timeStr) == "NEEDS CPPT" {
		return nil, nil
	}

	t, err := time.Parse(nhqTimeFormat, strings.TrimSpace(s))
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to parse time: %s", s)
	}

	return &t, nil
}

func readRow(row []string) (personnel.Participant, error) {
	lastName := row[lastNameCol]
	firstName := row[firstNameCol]
	middleName := row[middleNameCol]
	memberType := personnel.ParseMemberType(row[memberTypeCol])
	gender := personnel.ParseGender(row[genderCol])
	shirtSize := personnel.ParseShirtSize(row[shirtSizeCol])
	eServicesEmail := row[eServicesEmailCol]
	cadetParentPrimaryPhone := row[cadetParentPrimaryPhoneCol]
	cadetParentPrimaryEmail := row[cadetParentPrimaryEmailCol]
	unitCCName := row[unitCCNameCol]
	unitCCEmail := row[unitCCEmailCol]
	wingCCName := row[wingCCNameCol]
	wingCCEmail := row[wingCCEmailCol]

	capid, err := strconv.Atoi(row[capidCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse CAPID: %s", row[capidCol])
	}

	grade, err := personnel.ParseGrade(row[gradeCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse grade: %s", row[gradeCol])
	}

	ageAtEventStart, err := strconv.Atoi(row[ageAtEventStartCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse age at event start: %s", row[ageAtEventStartCol])
	}

	ageAtEventEnd, err := strconv.Atoi(row[ageAtEventEndCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse age at event end: %s", row[ageAtEventEndCol])
	}

	membershipExpires, err := time.Parse(nhqTimeFormat, strings.TrimSpace(row[membershipExpiresCol]))
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse membership expiration date: %s", row[membershipExpiresCol])
	}

	region, err := personnel.ParseRegion(row[regionCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse region: %s", row[regionCol])
	}

	wing, err := personnel.ParseWing(row[wingCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse wing: %s", row[wingCol])
	}

	unitNumber, err := strconv.Atoi(row[unitNumberCol])
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "failed to parse unit number: %s", row[unitNumberCol])
	}

	homeUnit, err := personnel.NewHomeUnit(region, wing, uint(unitNumber), "")
	if err != nil {
		return personnel.Participant{}, errors.WithMessagef(err, "invalid home unit charter number: %s-%s-%03d", region, wing, unitNumber)
	}

	unitApprovalDate, err := parseNullableTime(row[unitApprovalDateCol])
	if err != nil {
		return personnel.Participant{}, err
	}

	cpptExpires, err := parseNullableTime(row[cpptExpiresCol])
	if err != nil {
		return personnel.Participant{}, err
	}

	icutDate, err := parseNullableTime(row[icutDateCol])
	if err != nil {
		return personnel.Participant{}, err
	}

	capDlExpires, err := parseNullableTime(row[capDlExpiresCol])

	participant := personnel.NewParticipant(
		uuid.Nil,
		uint(capid),
		lastName,
		firstName,
		middleName,
		memberType,
		grade,
		gender,
		uint(ageAtEventStart),
		uint(ageAtEventEnd),
		shirtSize,
		membershipExpires,
		eServicesEmail,
		homeUnit,
		unitApprovalDate,
		cadetParentPrimaryPhone,
		cadetParentPrimaryEmail,
		unitCCName,
		unitCCEmail,
		wingCCName,
		wingCCEmail,
		cpptExpires,
		icutDate,
		capDlExpires,
	)

	return participant, nil
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
		// Skip header row
		if i == 0 {
			continue
		}
		p, err := readRow(row)
		if err != nil {
			log.Warn().Err(err).Str("path", file.FullPath()).Int("row", i).Msg("Failed to process row. Skipping.")
		}

		participants = append(participants, p)
	}

	return participants, nil
}
