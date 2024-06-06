package postgresrepo

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/ag7if/wendover/internal/config"
	"github.com/ag7if/wendover/internal/database"
	"github.com/ag7if/wendover/pkg/org"
)

type ActivityUnitTestSuite struct {
	suite.Suite
	repo     *PostgresRepository
	activity org.Activity
}

func (au *ActivityUnitTestSuite) SetupSuite() {
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}

	au.repo = NewPostgresRepository(db)

	activity := org.NewActivity(
		uuid.Nil,
		"test",
		"Test Activity",
		"",
		time.Now(),
		time.Now(),
		0,
		0,
		0,
		0,
	)

	activity, err = au.repo.InsertActivity(activity)
	if err != nil {
		panic(err)
	}

	if activity.ID() == uuid.Nil {
		panic("activity UUID not set")
	}

	au.activity = activity
}

func (au *ActivityUnitTestSuite) TearDownSuite() {
	key := au.activity.Key()
	err := au.repo.DeleteActivity(key)
	if err != nil {
		panic(err)
	}
}

func (au *ActivityUnitTestSuite) TestHierarchyInsert() {
	var err error

	group := org.NewActivityUnit(uuid.Nil, "Group")
	sq1 := org.NewActivityUnit(uuid.Nil, "Squadron 1")
	sq2 := org.NewActivityUnit(uuid.Nil, "Squadron 2")
	aflt := org.NewActivityUnit(uuid.Nil, "Alpha Flight")
	bflt := org.NewActivityUnit(uuid.Nil, "Bravo Flight")
	cflt := org.NewActivityUnit(uuid.Nil, "Charlie Flight")
	dflt := org.NewActivityUnit(uuid.Nil, "Delta Flight")
	eflt := org.NewActivityUnit(uuid.Nil, "Echo Flight")
	fflt := org.NewActivityUnit(uuid.Nil, "Foxtrot Flight")

	sq1.AddSubordinateUnit(aflt)
	sq1.AddSubordinateUnit(bflt)
	sq1.AddSubordinateUnit(cflt)

	sq2.AddSubordinateUnit(dflt)
	sq2.AddSubordinateUnit(eflt)
	sq2.AddSubordinateUnit(fflt)

	group.AddSubordinateUnit(sq1)
	group.AddSubordinateUnit(sq2)

	igp, err := au.repo.InsertActivityHierachy(au.activity.ID(), group)
	assert.NoError(au.T(), err)

	assert.NotEqual(au.T(), uuid.Nil, igp.ID())
	assert.Equal(au.T(), "Group", igp.UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[0].ID())
	assert.Equal(au.T(), "Squadron 1", igp.SubordinateUnits()[0].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[0].SubordinateUnits()[0].ID())
	assert.Equal(au.T(), "Alpha Flight", igp.SubordinateUnits()[0].SubordinateUnits()[0].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[0].SubordinateUnits()[1].ID())
	assert.Equal(au.T(), "Bravo Flight", igp.SubordinateUnits()[0].SubordinateUnits()[1].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[0].SubordinateUnits()[2].ID())
	assert.Equal(au.T(), "Charlie Flight", igp.SubordinateUnits()[0].SubordinateUnits()[2].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[1].ID())
	assert.Equal(au.T(), "Squadron 2", igp.SubordinateUnits()[1].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[1].SubordinateUnits()[0].ID())
	assert.Equal(au.T(), "Delta Flight", igp.SubordinateUnits()[1].SubordinateUnits()[0].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[1].SubordinateUnits()[1].ID())
	assert.Equal(au.T(), "Echo Flight", igp.SubordinateUnits()[1].SubordinateUnits()[1].UnitName())

	assert.NotEqual(au.T(), uuid.Nil, igp.SubordinateUnits()[1].SubordinateUnits()[2].ID())
	assert.Equal(au.T(), "Foxtrot Flight", igp.SubordinateUnits()[1].SubordinateUnits()[2].UnitName())
}

func TestActivityUnit(t *testing.T) {
	suite.Run(t, new(ActivityUnitTestSuite))
}
