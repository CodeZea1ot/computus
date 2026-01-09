package testutil

// VerifiedEasterDates contains historically verified Easter Sunday dates.
// This can be used by all dependent tests to calculate moveable feasts.
var VerifiedEasterDates = map[int]string{
	1583: "1583-04-10", // first year Gregorian Easter
	1666: "1666-04-25", // latest possible Easter
	1693: "1693-03-22", // earliest possible Easter
	1818: "1818-03-22", // earliest possible Easter
	1900: "1900-04-15",
	1954: "1954-04-18",
	1970: "1970-03-29",
	1999: "1999-04-04",
	2000: "2000-04-23",
	2010: "2010-04-04",
	2016: "2016-03-27",
	2020: "2020-04-12",
	2021: "2021-04-04",
	2022: "2022-04-17",
	2023: "2023-04-09",
	2024: "2024-03-31",
	2025: "2025-04-20",
	2026: "2026-04-05",
	2038: "2038-04-25", // latest possible Easter
}
