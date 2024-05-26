package capf

import (
	"time"

	"github.com/ag7if/wendover/pkg/personnel"
)

type Capf160I struct {
	// Header of Page 1
	Name        string
	Grade       personnel.Grade
	CAPID       uint
	HomeUnit    personnel.HomeUnit
	DateOfBirth time.Time
	Height      uint
	Weight      uint
	HairColor   personnel.HairColor
	EyeColor    personnel.EyeColor
	Gender      personnel.Gender
	Allergies   string

	// Known Conditions
	DecreasedVision           bool // Decreased vision, glaucoma, contacts
	ChronicInjuries           bool // Chronic or recurring injuries
	EarInfections             bool // Ear infections, perforation
	ActivityRestrictions      bool // Activity, mobility restrictions
	DifficultyEqualizingEars  bool // Difficulty equalizing ears
	UseOfMobilityAid          bool // Use of cane, walker, wheelchair
	HearingLoss               bool // Hearing loss, hearing aid
	BackInjury                bool // Back or neck pain or injury
	NasalStuffiness           bool // Allergies, nasal stuffiness
	Migrate                   bool // Migraine or severe headaches
	Anaphylaxis               bool // Anaphylaxis, serious allergic reaction
	Dizziness                 bool // Dizziness or fainting spells
	Asthma                    bool // Asthma, emphysema (COPD)
	HeadInjury                bool // Head injury, unconsciousness
	Inhaler                   bool // Ever use an inhaler
	Epilepsy                  bool // Epilepsy or seizure
	ShortOfBreath             bool // Short of Breath with activity
	Stroke                    bool // Stroke, paralysis
	HeartAttack               bool // Heart Attack, chest pain, angina
	ThyroidProblems           bool // Thyroid problems (low or high)
	HeartMurmur               bool // Heart murmur, heart problems
	Diabetes                  bool // Diabetes, high or low blood sugars
	CongestiveHeartFailure    bool // Congestive heart failure
	Cancer                    bool // Cancer, leukemia
	IrregularHeartbeat        bool // Irregular or rapid heartbeat
	BloodDisease              bool // Blood disease, hemophilia
	HighOrLowBloodPressure    bool // High or low blood pressure
	MotionSickness            bool // Motion sickness
	StomachTrouble            bool // Stomach trouble, ulcers
	SpecialDiet               bool // Special diet, food allergies
	Hepatitis                 bool // Hepatitis or liver problems
	CurrentBedwetting         bool // Current bedwetting problems
	Diarrhea                  bool // Diarrhea, constipation
	ADD                       bool // ADD (Attention Deficit Disorder)
	Hernia                    bool // Hernia or rupture
	MentalIllness             bool // Mental illness (bipolar, other)
	KidneyDisease             bool // Kidney disease or stones
	Depression                bool // Depression, anxiety, suicidal
	ProstateProblems          bool // Prostate problems (men)
	HospitalAdmission         bool // Admission to the hospital
	FrequentUrination         bool // Frequent urination
	OtherChronicMentalIllness bool // Other chronic medical illnesses
	MenstralCramps            bool // Menstrual cramps (women)
	SleepDisorder             bool // Sleep disorder, sleep apnea
	BrokenBones               bool // Broken bone, joint problems
	SeriousInjury             bool // Serious Injury

	// Page 2
	DietaryRestrictions string
	PastSurgicalHistory string
	TetanusBooster      *time.Time
	HepatitisVaccine    *time.Time
	PneumoniaVaccine    *time.Time
	Varicella           *time.Time
	FluVaccine          *time.Time

	// Medication Information

}
