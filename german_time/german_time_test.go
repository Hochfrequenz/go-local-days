package german_time_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-german-time/german_time"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

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

/*********************
 Add German Day Tests
**********************/

// Test_Add_German_Day_Normal_GET tests that adding works as expected when both source time and target time are in UTC+1 (MEZ/CET).
func (s *Suite) Test_Add_German_Day_Normal_GET() {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.AddGermanDays(date, 3)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC)))
}

// Test_Add_German_Day_Normal_CEST tests that adding works as expected when both source time and target time are in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Add_German_Day_Normal_CEST() {
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.AddGermanDays(date, 3)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 4, 0, 0, 0, 0, time.UTC)))
}

// Test_Add_German_Day_Normal_CET_To_CEST_Transition tests that adding works as expected when the source time is in UTC+1 (MEZ/CET) and the target is in UTC+2 (MESZ/CEST); This is the 1 local day=23h case.
func (s *Suite) Test_Add_German_Day_Normal_CET_To_CEST_Transition() {
	date := time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC)
	actual := german_time.AddGermanDays(date, 1)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 27, 23, 0, 0, 0, time.UTC)))
}

// Test_Add_German_Day_Normal_CEST_To_CET_Transition tests that adding works as expected when the source time is in UTC+2 (MESZ/CEST) and the target is in UTC+1 (MEZ/CET); This is the 1 local day=25h case.
func (s *Suite) Test_Add_German_Day_Normal_CEST_To_CET_Transition() {
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := german_time.AddGermanDays(date, 1)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 31, 1, 0, 0, 0, time.UTC)))
}

/*******************
 Start of German Day
*******************/

// Test_Start_Of_Day_In_Germany_Normal_CET tests that midnight is found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_Day_In_Germany_Normal_CET() {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Day_In_Germany_Normal_CET tests that midnight is found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Day_In_Germany_Normal_CEST() {
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 5, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CET_To_CEST_Transition tests that midnight is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Day_In_Germany_CET_To_CEST_Transition() {
	date := time.Date(2022, 3, 27, 12, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 26, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CEST_To_CET_Transition tests that midnight is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Day_In_Germany_CEST_To_CET_Transition() {
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 29, 22, 0, 0, 0, time.UTC)))
}

/************************
Start of Next German Day
************************/

// Test_Start_Of_Next_Day_In_Germany_Normal_CET tests that the next day is found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_Next_Day_In_Germany_Normal_CET() {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 1, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_Normal_CET tests that the next day is found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Next_Day_In_Germany_Normal_CEST() {
	date := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 1, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CET_To_CEST_Transition tests that the next day is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Next_Day_In_Germany_CET_To_CEST_Transition() {
	date := time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 27, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CEST_To_CET_Transition tests that the next day is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Next_Day_In_Germany_CEST_To_CET_Transition() {
	date := time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanDay(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 30, 23, 0, 0, 0, time.UTC)))
}

/***********************
 Start of German Month
***********************/

// Test_Start_Of_German_Month_Normal_CET tests that the start of the month of found in UTC+1 (MEZ/CET).
func (s *Suite) Test_Start_Of_German_Month_Normal_CET() {
	date := time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_German_Month_Normal_CEST tests that the start of the month of found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_German_Month_Normal_CEST() {
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 5, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_German_Month_CET_To_CEST_Transition tests that start of the month is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_German_Month_CET_To_CEST_Transition() {
	date := time.Date(2022, 10, 31, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 9, 30, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CEST_To_CET_Transition tests that start of the month is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_German_Month_CEST_To_CET_Transition() {
	date := time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 2, 28, 23, 0, 0, 0, time.UTC)))
}

/**************************
Start of Next German Month
**************************/

// Test_Start_Of_Next_German_Month_Normal_CET_Same_Year tests that the start of the next month of found in UTC+1 (MEZ/CET) within the same year.
func (s *Suite) Test_Start_Of_Next_German_Month_Normal_CET_Same_Year() {
	date := time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 1, 31, 23, 0, 0, 0, time.UTC)))
}

func (s *Suite) Test_Start_Of_Next_German_Month_Normal_CET_Same_Year_2259() {
	date := time.Date(2021, 12, 31, 22, 59, 59, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_German_Month_Normal_CET_Next_Year tests that the start of the next month of found in UTC+1 (MEZ/CET) over a year.
func (s *Suite) Test_Start_Of_Next_German_Month_Normal_CET_Next_Year() {
	date := time.Date(2021, 12, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2021, 12, 31, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_German_Month_Normal_CEST tests that the start of the next month of found in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Start_Of_Next_German_Month_Normal_CEST() {
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 30, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_German_Month_CET_To_CEST_Transition tests that start of the next month is found when starting in UTC+1 and ending in UTC+2.
func (s *Suite) Test_Start_Of_Next_German_Month_CET_To_CEST_Transition() {
	date := time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 31, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_Of_Next_Day_In_Germany_CEST_To_CET_Transition tests that start of the next month is found when starting in UTC+2 and ending in UTC+1.
func (s *Suite) Test_Start_Of_Next_Month_In_Germany_CEST_To_CET_Transition() {
	date := time.Date(2022, 10, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.StartOfNextGermanMonth(date)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 31, 23, 0, 0, 0, time.UTC)))
}

/****************************
Start of Next German Weekday
****************************/

// Test_Start_of_Next_German_Weekday_CET tests that start of the next weekday is found when only acting in UTC+1.
func (s *Suite) Test_Start_of_Next_German_Weekday_CET() {
	date := time.Date(2022, 11, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.NextGermanWeekday(date, time.Friday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 11, 17, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_German_Weekday_CET tests that start of the next weekday is found when only acting in UTC+1 when the specified day is the same weekday but not midnight
func (s *Suite) Test_Start_of_Next_German_Weekday_CET_Plus7() {
	date := time.Date(2022, 11, 15, 12, 0, 0, 0, time.UTC)
	actual := german_time.NextGermanWeekday(date, time.Tuesday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 11, 21, 23, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_German_Weekday_CET tests that start of the next weekday is found when only acting in UTC+2.
func (s *Suite) Test_Start_of_Next_German_Weekday_CEST() {
	date := time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)
	actual := german_time.NextGermanWeekday(date, time.Friday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 6, 16, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_German_Weekday_CET_to_CEST_Transition tests that start of the next weekday is found when transitioning from UTC+1 to UTC+2.
func (s *Suite) Test_Start_of_Next_German_Weekday_CET_to_CEST_Transition() {
	date := time.Date(2022, 3, 25, 0, 0, 0, 0, time.UTC)
	actual := german_time.NextGermanWeekday(date, time.Wednesday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 3, 29, 22, 0, 0, 0, time.UTC)))
}

// Test_Start_of_Next_German_Weekday_CEST_to_CET_Transition tests that start of the next weekday is found when transitioning from UTC+2 to UTC+1.
func (s *Suite) Test_Start_of_Next_German_Weekday_CEST_to_CET_Transition() {
	date := time.Date(2022, 10, 27, 0, 0, 0, 0, time.UTC)
	actual := german_time.NextGermanWeekday(date, time.Monday)
	then.AssertThat(s.T(), actual, is.EqualTo(time.Date(2022, 10, 30, 23, 0, 0, 0, time.UTC)))
}

/*****************
 German Midnight
*****************/

// Test_Is_German_Midnight_CET tests the determination of German midnight in UTC+1 (MEZ/CET).
func (s *Suite) Test_Is_German_Midnight_CET() {
	midnight := time.Date(2022, 11, 4, 23, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), german_time.IsGermanMidnight(midnight), is.True())

	notMidnight := time.Date(2022, 11, 5, 0, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), german_time.IsGermanMidnight(notMidnight), is.False())
}

// Test_Is_German_Midnight_CEST tests the determination of German midnight in UTC+2 (MESZ/CEST).
func (s *Suite) Test_Is_German_Midnight_CEST() {
	midnight := time.Date(2022, 5, 4, 22, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), german_time.IsGermanMidnight(midnight), is.True())

	notMidnight := time.Date(2022, 5, 5, 0, 0, 0, 0, time.UTC)
	then.AssertThat(s.T(), german_time.IsGermanMidnight(notMidnight), is.False())
}
