package content

type Place struct {
	Year     int
	Location string
	Text     string
	Value    int
	Lat      float64
	Lng      float64
}

type Marker struct {
	Location string  `json:"location"`
	Text     string  `json:"text"`
	Years    string  `json:"years"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Radius   float64 `json:"radius"`
	Color    string  `json:"color"`
}

// Places is ordered oldest-first. Value is a rough "how much time I spent
// there" weight, which becomes the marker radius.
var Places = []Place{
	{Year: 2008, Location: "Riverside, California", Text: "Finished first year of university", Value: 9, Lat: 33.9461906, Lng: -117.5643322},
	{Year: 2008, Location: "Guadalajara, Mexico", Text: "Summer abroad in Mexico to learn Spanish", Value: 2, Lat: 20.6739406, Lng: -103.6307134},
	{Year: 2008, Location: "Tijuana, Mexico", Text: "Spanish practice", Value: 1, Lat: 32.4572678, Lng: -117.1166},
	{Year: 2009, Location: "Riverside, California", Text: "Second year of university", Value: 8, Lat: 33.9461906, Lng: -117.5643322},
	{Year: 2009, Location: "Santa Marta, Colombia", Text: "Scuba Diving Course", Value: 1, Lat: 11.2316066, Lng: -74.2235964},
	{Year: 2009, Location: "Ústí nad Labem, Czech Republic", Text: "First time in Europe", Value: 1, Lat: 50.6584116, Lng: 14.0354625},
	{Year: 2009, Location: "Buenos Aires, Argentina", Text: "Summer Abroad to learn Spanish", Value: 3, Lat: -34.6156548, Lng: -58.5156979},
	{Year: 2009, Location: "Mendoza, Argentina", Text: "Wine education", Value: 1, Lat: -32.8798371, Lng: -68.8642868},
	{Year: 2009, Location: "Tijuana, Mexico", Text: "Spanish practice", Value: 1, Lat: 32.4572678, Lng: -117.1166},
	{Year: 2010, Location: "Santa Barbara, California", Text: "UC transfer", Value: 5, Lat: 34.3987011, Lng: -119.8322553},
	{Year: 2010, Location: "Riverside, California", Text: "Third year of university", Value: 6, Lat: 33.9461906, Lng: -117.5643322},
	{Year: 2010, Location: "Cappadocia, Turkey", Text: "Family Vacation", Value: 1, Lat: 39.2278513, Lng: 34.1057059},
	{Year: 2010, Location: "Athens, Greece", Text: "Family Vacation", Value: 1, Lat: 37.9879889, Lng: 23.6939644},
	{Year: 2010, Location: "Rome, Italy", Text: "Family Vacation", Value: 1, Lat: 41.9059807, Lng: 12.4466053},
	{Year: 2010, Location: "Madrid, Spain", Text: "Family Vacation", Value: 1, Lat: 40.4378368, Lng: -3.844683},
	{Year: 2010, Location: "Tijuana, Mexico", Text: "Spanish practice", Value: 1, Lat: 32.4572678, Lng: -117.1166},
	{Year: 2011, Location: "Accra, Ghana", Text: "Semester abroad", Value: 5, Lat: 5.5913738, Lng: -0.2621296},
	{Year: 2011, Location: "Kumasi, Ghana", Text: "Field Trip", Value: 1, Lat: 6.6950214, Lng: -1.6437655},
	{Year: 2011, Location: "Lomé, Togo", Text: "Adventure", Value: 1, Lat: 6.1746457, Lng: 1.2175335},
	{Year: 2011, Location: "Cairo, Egypt", Text: "Winter Break travels", Value: 1, Lat: 30.0595546, Lng: 31.1108223},
	{Year: 2011, Location: "Amman, Jordan", Text: "Winter Break travels", Value: 1, Lat: 31.8359188, Lng: 35.618003},
	{Year: 2011, Location: "Granada, Spain", Text: "Semester abroad", Value: 4, Lat: 37.1810075, Lng: -3.632471},
	{Year: 2011, Location: "Tunis, Tunisia", Text: "Spring Break", Value: 1, Lat: 36.7950442, Lng: 9.9784737},
	{Year: 2011, Location: "Marrakesh, Morocco", Text: "Break after exams", Value: 1, Lat: 31.6131258, Lng: -8.0206542},
	{Year: 2012, Location: "Lisbon, Portugal", Text: "Long weekend", Value: 1, Lat: 38.7201028, Lng: -9.150459},
	{Year: 2012, Location: "El Fonoll, Spain", Text: "First job abroad", Value: 1, Lat: 41.5289786, Lng: 1.2268317},
	{Year: 2012, Location: "Madrid, Spain", Text: "Visits to friend", Value: 1, Lat: 40.4378368, Lng: -3.844683},
	{Year: 2012, Location: "Vancouver, Canada", Text: "Roadtrip", Value: 1, Lat: 49.2561342, Lng: -123.1311773},
	{Year: 2012, Location: "Chicago, USA", Text: "Roadtrip", Value: 1, Lat: 41.8587241, Lng: -87.7497106},
	{Year: 2012, Location: "Essen, Germany", Text: "Finished hitchhiking journey", Value: 1, Lat: 51.4409854, Lng: 6.9336189},
	{Year: 2012, Location: "Hong Kong", Text: "Temp work in Hong Kong", Value: 3, Lat: 22.3262127, Lng: 114.1684683},
	{Year: 2012, Location: "Macao", Text: "Work trip", Value: 1, Lat: 22.2002619, Lng: 113.5374163},
	{Year: 2012, Location: "Jakarta, Indonesia", Text: "Teaching contract", Value: 6, Lat: -6.2293796, Lng: 106.6647126},
	{Year: 2013, Location: "Jakarta, Indonesia", Text: "Teaching contract", Value: 3, Lat: -6.2293796, Lng: 106.6647126},
	{Year: 2013, Location: "Kuala Lumpur, Malaysia", Text: "Visiting old colleague", Value: 1, Lat: 3.1386741, Lng: 101.6045894},
	{Year: 2013, Location: "Banda Aceh, Indonesia", Text: "Travels", Value: 1, Lat: 5.5611859, Lng: 95.2875031},
	{Year: 2013, Location: "Bali, Indonesia", Text: "Travels", Value: 1, Lat: -8.3421174, Lng: 115.0658807},
	{Year: 2013, Location: "Singapore", Text: "Work related to get new visa", Value: 1, Lat: 1.3146649, Lng: 103.5146164},
	{Year: 2013, Location: "Tokyo, Japan", Text: "Vacation", Value: 1, Lat: 35.497202, Lng: 139.3664982},
	{Year: 2013, Location: "Nablus, Palestine", Text: "Work in a school", Value: 6, Lat: 32.2243442, Lng: 35.2270369},
	{Year: 2014, Location: "Dzerzhinsky, Russia", Text: "First job in Russia", Value: 8, Lat: 55.6275161, Lng: 37.7997683},
	{Year: 2014, Location: "Moscow, Russia", Text: "Better job in Russia", Value: 3, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2014, Location: "Mexico City, Mexico", Text: "Mandatory visa travels", Value: 1, Lat: 19.3909832, Lng: -99.3084179},
	{Year: 2014, Location: "Tijuana, Mexico", Text: "Spanish practice", Value: 1, Lat: 32.4572678, Lng: -117.1166},
	{Year: 2015, Location: "Ibiza, Spain", Text: "Holiday", Value: 1, Lat: 38.9073718, Lng: 1.3558305},
	{Year: 2015, Location: "Nizhny Novgorod, Russia", Text: "Holiday", Value: 1, Lat: 56.2920986, Lng: 43.7613651},
	{Year: 2015, Location: "Moscow, Russia", Text: "Better Work", Value: 9, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2015, Location: "Nice, France", Text: "Winter break", Value: 1, Lat: 43.7032898, Lng: 7.1704124},
	{Year: 2016, Location: "Berlin, Germany", Text: "Holiday", Value: 1, Lat: 52.5068042, Lng: 13.0951149},
	{Year: 2016, Location: "Kazan, Russia", Text: "Holiday", Value: 1, Lat: 55.7955843, Lng: 48.9085016},
	{Year: 2016, Location: "Moscow, Russia", Text: "Work", Value: 10, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2017, Location: "Moscow, Russia", Text: "Work", Value: 7, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2017, Location: "Bologna, Italy", Text: "Started Graduate School", Value: 5, Lat: 44.4992289, Lng: 11.249286},
	{Year: 2018, Location: "Dolomites, Italy", Text: "Holiday", Value: 1, Lat: 46.4102389, Lng: 11.8234356},
	{Year: 2018, Location: "Moscow, Russia", Text: "Carnegie Center", Value: 3, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2018, Location: "Bologna, Italy", Text: "Started Graduate School", Value: 5, Lat: 44.4992289, Lng: 11.249286},
	{Year: 2018, Location: "Oklahoma City, Oklahoma", Text: "Road Trip", Value: 1, Lat: 35.4942601, Lng: -97.7384269},
	{Year: 2018, Location: "Washington, DC", Text: "Second year of Graduate School", Value: 3, Lat: 38.8936499, Lng: -77.3098813},
	{Year: 2019, Location: "Washington, DC", Text: "Finished Graduate School", Value: 7, Lat: 38.8936499, Lng: -77.3098813},
	{Year: 2019, Location: "Boston, MA", Text: "Holiday", Value: 1, Lat: 42.3144858, Lng: -71.1350873},
	{Year: 2019, Location: "Moscow, Russia", Text: "Moved back to Moscow", Value: 3, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2019, Location: "Minsk, Belarus", Text: "Family Reunion", Value: 1, Lat: 53.8847295, Lng: 27.4285668},
	{Year: 2020, Location: "Moscow, Russia", Text: "Quarantine", Value: 11, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2020, Location: "Karelia, Russia", Text: "Holiday", Value: 1, Lat: 63.1728532, Lng: 32.2977844},
	{Year: 2021, Location: "Moscow, Russia", Text: "Quarantine", Value: 8, Lat: 55.754119, Lng: 37.6113672},
	{Year: 2021, Location: "Alpine, CA", Text: "Waiting for visa", Value: 3, Lat: 32.8441172, Lng: -116.8392662},
	{Year: 2021, Location: "Istanbul, Turkey", Text: "Job Hunting", Value: 1, Lat: 40.9996453, Lng: 27.765759},
	{Year: 2022, Location: "Istanbul, Turkey", Text: "Job Hunting", Value: 2, Lat: 40.9996453, Lng: 27.765759},
	{Year: 2022, Location: "London, UK", Text: "Work Travel", Value: 1, Lat: 51.528607, Lng: -0.4312161},
	{Year: 2022, Location: "Val d’Isère, France", Text: "Family Vacation", Value: 1, Lat: 45.4232007, Lng: 6.9325359},
	{Year: 2022, Location: "Geneva, Switzerland", Text: "Family Vacation", Value: 1, Lat: 46.2055315, Lng: 6.1223965},
	{Year: 2022, Location: "Phuket, Thailand", Text: "Work Travel", Value: 1, Lat: 7.8396548, Lng: 98.0752989},
	{Year: 2022, Location: "Riga, Latvia", Text: "Work Travel", Value: 1, Lat: 56.9718363, Lng: 23.9642722},
	{Year: 2022, Location: "Phoenix, Arizona", Text: "Data Science Job", Value: 9, Lat: 33.6045833, Lng: -112.7157717},
	{Year: 2023, Location: "Nashville, Tennessee", Text: "Developer Job", Value: 10, Lat: 36.1868042, Lng: -86.9503928},
	{Year: 2023, Location: "Istanbul, Turkey", Text: "Family Reunion", Value: 1, Lat: 40.9996453, Lng: 27.765759},
	{Year: 2023, Location: "Alpine, CA", Text: "Holiday", Value: 1, Lat: 32.8441172, Lng: -116.8392662},
	{Year: 2024, Location: "Nashville, Tennessee", Text: "Developer Job", Value: 9, Lat: 36.1868042, Lng: -86.9503928},
	{Year: 2024, Location: "Tbilisi, Georgia", Text: "Family Reunion and wedding", Value: 1, Lat: 41.7278607, Lng: 44.6419604},
	{Year: 2024, Location: "Buenos Aires, Argentina", Text: "Holiday", Value: 1, Lat: -34.6156548, Lng: -58.5156979},
	{Year: 2024, Location: "Raleigh, North Carolina", Text: "All Things Open Conference", Value: 1, Lat: 35.8238807, Lng: -78.6880527},
	{Year: 2024, Location: "Alpine, CA", Text: "Holiday", Value: 1, Lat: 32.8441172, Lng: -116.8392662},
	{Year: 2025, Location: "Nashville, Tennessee", Text: "Remote Work", Value: 9, Lat: 36.1868042, Lng: -86.9503928},
	{Year: 2025, Location: "Murfreesboro, Tennessee", Text: "Meeting supervisor", Value: 1, Lat: 35.8803208, Lng: -86.3874135},
	{Year: 2025, Location: "Little Italy, San Diego", Text: "Fancy holiday with wife", Value: 1, Lat: 32.7251527, Lng: -117.1735454},
	{Year: 2026, Location: "Little Italy, San Diego", Text: "Fancy holiday with wife", Value: 1, Lat: 32.7251527, Lng: -117.1735454},
}
