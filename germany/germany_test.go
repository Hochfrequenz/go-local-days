package germany_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-local-days/germany"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

/*********************
 Add Local Day Tests
**********************/

// Test_Add_Local_Day_Normal_GET tests that adding works as expected when both source time and target time are in UTC+1 (MEZ/CET).
func (s *Suite) Test_Add_Local_Day_Normal_GET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.AddLocalDays(date, 3)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC)))
}

// Test_Add_Local_Day_Normal_CEST tests that adding works as expected when both source time and target time are in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Add_Local_Day_Normal_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.AddLocalDays(date, 3)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 4, 0, 0, 0, 0, time.UTC)))
}

// Test_Add_Local_Day_Normal_CET_To_CEST_Transition tests that adding works as expected when the source time is in UTC+1 (MEZ/CET) and the target is in UTC+2 (MESZ/CEST); This is the 1 local day=23h case.
func (s *Suite) Test_Add_Local_Day_Normal_CET_To_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC)
	actual := berlin.AddLocalDays(date, 1)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 27, 23, 0, 0, 0, time.UTC)))
}

// Test_Add_Local_Day_Normal_CEST_To_CET_Transition tests that adding works as expected when the source time is in UTC+2 (MESZ/CEST) and the target is in UTC+1 (MEZ/CET); This is the 1 local day=25h case.
func (s *Suite) Test_Add_Local_Day_Normal_CEST_To_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := berlin.AddLocalDays(date, 1)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 31, 1, 0, 0, 0, time.UTC)))
}

/*******************
 Start of Local Day
*******************/

// Test_Start_Of_Day_In_Localy_Normal_CET tests that midnight is found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_Day_In_Localy_Normal_CET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Day_In_Localy_Normal_CET tests that midnight is found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Day_In_Localy_Normal_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 5, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CET_To_CEST_Transition tests that midnight is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Day_In_Localy_CET_To_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 27, 12, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 26, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CEST_To_CET_Transition tests that midnight is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Day_In_Localy_CEST_To_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 29, 22, 0, 0, 0, time.UTC)))
}

/************************
Start of Next Local Day
************************/

// Test_Start_Of_Next_Day_In_Localy_Normal_CET tests that the next day is found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_Next_Day_In_Localy_Normal_CET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 1, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_Normal_CET tests that the next day is found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Next_Day_In_Localy_Normal_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 1, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CET_To_CEST_Transition tests that the next day is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Next_Day_In_Localy_CET_To_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 27, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CEST_To_CET_Transition tests that the next day is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Next_Day_In_Localy_CEST_To_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 30, 23, 0, 0, 0, time.UTC)))
}

/***********************
 Start of Local Month
***********************/

// Test_Start_Of_Local_Month_Normal_CET tests that the start of the month of found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_Local_Month_Normal_CET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Local_Month_Normal_CEST tests that the start of the month of found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Local_Month_Normal_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 5, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Local_Month_CET_To_CEST_Transition tests that start of the month is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Local_Month_CET_To_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 31, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 9, 30, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CEST_To_CET_Transition tests that start of the month is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Local_Month_CEST_To_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 2, 28, 23, 0, 0, 0, time.UTC)))
}

/**************************
Start of Next Local Month
**************************/

// Test_Start_Of_Next_Local_Month_Normal_CET_Same_Year tests that the start of the next month of found in UTC+1 (MEZ/CET) within the same year.
func (s *Suite) Test_Start_Of_Next_Local_Month_Normal_CET_Same_Year() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 31, 23, 0, 0, 0, time.UTC)))
}

func (s *Suite) Test_Start_Of_Next_Local_Month_Normal_CET_Same_Year_2259() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2021, 12, 31, 22, 59, 59, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Local_Month_Normal_CET_Next_Year tests that the start of the next month of found in UTC+1 (MEZ/CET) over a year.
func (s *Suite) Test_Start_Of_Next_Local_Month_Normal_CET_Next_Year() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2021, 12, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Local_Month_Normal_CEST tests that the start of the next month of found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Next_Local_Month_Normal_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 30, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Local_Month_CET_To_CEST_Transition tests that start of the next month is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Next_Local_Month_CET_To_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Localy_CEST_To_CET_Transition tests that start of the next month is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Next_Month_In_Localy_CEST_To_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.StartOfNextLocalMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 31, 23, 0, 0, 0, time.UTC)))
}

/****************
 Local Weekday
***************/

// Test_Local_Weekday tests that weekday is correct in UTC+1 (CET/MEZ) adn UTC+2 (CEST/MESZ)
func (s *Suite) Test_Local_Weekday() {
	berlin := germany.NewGermanLocalDaysCalculator()
	// CET
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 11, 15, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Tuesday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 11, 16, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Wednesday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 11, 16, 23, 0, 0, 0, time.UTC)), is.EqualTo(time.Thursday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Sunday))

	// CEST
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Wednesday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 6, 16, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Thursday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 6, 16, 22, 0, 0, 0, time.UTC)), is.EqualTo(time.Friday))
	then.AssertThat(s.T(), berlin.GetLocalWeekday(time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC)), is.EqualTo(time.Sunday))
}

/****************************
Start of Next Local Weekday
****************************/

// Test_Start_of_Next_Local_Weekday_CET tests that start of the next weekday is found when only acting in UTC+1.
func (s *Suite) Test_Start_of_Next_Local_Weekday_CET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 11, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.NextLocalWeekday(date, time.Friday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 11, 17, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_Local_Weekday_CET tests that start of the next weekday is found when only acting in UTC+1 when the specified day is the same weekday but not midnight
func (s *Suite) Test_Start_of_Next_Local_Weekday_CET_Plus7() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 11, 15, 12, 0, 0, 0, time.UTC)
	actual := berlin.NextLocalWeekday(date, time.Tuesday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 11, 21, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_Local_Weekday_CET tests that start of the next weekday is found when only acting in UTC+2.
func (s *Suite) Test_Start_of_Next_Local_Weekday_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := berlin.NextLocalWeekday(date, time.Friday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 16, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_Local_Weekday_CET_to_CEST_Transition tests that start of the next weekday is found when transitioning from UTC+1 to UTC+2.
func (s *Suite) Test_Start_of_Next_Local_Weekday_CET_to_CEST_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 3, 25, 0, 0, 0, 0, time.UTC)
	actual := berlin.NextLocalWeekday(date, time.Wednesday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 29, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_Local_Weekday_CEST_to_CET_Transition tests that start of the next weekday is found when transitioning from UTC+2 to UTC+1.
func (s *Suite) Test_Start_of_Next_Local_Weekday_CEST_to_CET_Transition() {
	berlin := germany.NewGermanLocalDaysCalculator()
	date := time.Date(2022, 10, 27, 0, 0, 0, 0, time.UTC)
	actual := berlin.NextLocalWeekday(date, time.Monday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 30, 23, 0, 0, 0, time.UTC)))
}

/*****************
 Local Midnight
*****************/

// Test_Is_Local_Midnight_CET tests the determination of Local midnight in UTC+1 (MEZ/CET).
func (s *Suite) Test_Is_Local_Midnight_CET() {
	berlin := germany.NewGermanLocalDaysCalculator()
	midnight := time.Date(2022, 11, 4, 23, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), berlin.IsLocalMidnight(midnight), is.True())

	notMidnight := time.Date(2022, 11, 5, 0, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), berlin.IsLocalMidnight(notMidnight), is.False())
}

// Test_Is_Local_Midnight_CEST tests the determination of Local midnight in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Is_Local_Midnight_CEST() {
	berlin := germany.NewGermanLocalDaysCalculator()
	midnight := time.Date(2022, 5, 4, 22, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), berlin.IsLocalMidnight(midnight), is.True())

	notMidnight := time.Date(2022, 5, 5, 0, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), berlin.IsLocalMidnight(notMidnight), is.False())
}

// ----------------------------
// test framework boiler plate
// ---------------------------
type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
