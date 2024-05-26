package reports

import (
	"github.com/ag7if/go-files"

	"github.com/ag7if/wendover/internal/clients/registration_zone"
)

func IngestRegistrationZoneReport(f files.File) error {
	_, err := registration_zone.ParseParticipantsFromRZExport(f)

	return err
}
