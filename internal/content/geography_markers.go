package content

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var MarkersByYear = markersByYear()

func markersByYear() map[string][]Marker {
	byYear := map[string][]Marker{"all": Markers(AllYears)}
	for _, y := range Years() {
		byYear[strconv.Itoa(y.Year)] = Markers(y.Year)
	}
	return byYear
}

// The palette runs from this year's shade to the oldest year's.
const (
	newestShade = "#2E86AB"
	oldestShade = "#000356"
)

// mapColors is generated to span exactly the years present in the data: index 0
// is the most recent year, the last index the oldest. A fixed-length palette
// had to grow by hand every time the data outlived it, and silently collapsed
// the oldest years onto a single shade when nobody remembered to.
//
// Package-level vars initialize in dependency order rather than source order,
// so reading EndYear here is fine even though it is declared below.
var mapColors = shades(newestShade, oldestShade, EndYear-StartYear+1)

// shades interpolates n colors evenly from a to b, including both endpoints.
func shades(a, b string, n int) []string {
	ar, ag, ab := hexToRGB(a)
	br, bg, bb := hexToRGB(b)

	if n < 2 {
		return []string{a}
	}

	out := make([]string, n)
	for i := range out {
		t := float64(i) / float64(n-1)
		out[i] = fmt.Sprintf("#%02X%02X%02X",
			lerp(ar, br, t), lerp(ag, bg, t), lerp(ab, bb, t))
	}
	return out
}

// lerp blends one color channel, rounding to the nearest whole step.
func lerp(a, b uint8, t float64) uint8 {
	return uint8(math.Round(float64(a) + (float64(b)-float64(a))*t))
}

// hexToRGB splits "#rrggbb" into channels. The only inputs are the constants
// above, so a malformed one is a bug to surface at startup rather than an error
// to thread through package initialization.
func hexToRGB(s string) (uint8, uint8, uint8) {
	digits := strings.TrimPrefix(s, "#")
	v, err := strconv.ParseUint(digits, 16, 32)
	if err != nil || len(digits) != 6 {
		panic("content: malformed palette color " + s)
	}
	return uint8(v >> 16), uint8(v >> 8), uint8(v)
}

// StartYear and EndYear bracket the data. Deriving them rather than hardcoding
// means adding a Place in a new year extends the year picker and re-bases the
// colors on its own.
var StartYear, EndYear = yearRange()

func yearRange() (int, int) {
	lo, hi := Places[0].Year, Places[0].Year
	for _, p := range Places[1:] {
		lo = min(lo, p.Year)
		hi = max(hi, p.Year)
	}
	return lo, hi
}

// YearOption is one button in the year picker. Carrying the shade alongside the
// year lets the button be tinted to match the markers it selects.
type YearOption struct {
	Year  int
	Color string
}

// Years lists every year in the data, oldest first, for the year picker.
func Years() []YearOption {
	years := make([]YearOption, 0, EndYear-StartYear+1)
	for y := StartYear; y <= EndYear; y++ {
		years = append(years, YearOption{Year: y, Color: colorFor(y)})
	}
	return years
}

// colorFor shades a year by age. The palette is sized to the data, so the clamp
// is belt-and-braces against a year outside the range Places covers.
func colorFor(year int) string {
	return mapColors[min(max(EndYear-year, 0), len(mapColors)-1)]
}

// AllYears selects every year at once.
const AllYears = 0

const (
	minRadius = 4.0
	maxRadius = 40.0

	// spread stops the scaling from dividing by zero when every location in a
	// year has the same value. It also compresses the range a little, which
	// keeps single-entry years from all drawing at maximum size.
	spread = 3.0
)

// Markers groups a year's places into one marker per location, sized by total
// value and shaded by the most recent year that location was visited.
func Markers(year int) []Marker {
	type group struct {
		place  Place
		value  int
		texts  []string
		years  []string
		newest int
	}

	// order tracks insertion order separately, because ranging over a map in Go
	// yields keys in a random order and the output would shuffle every request.
	var order []string
	groups := make(map[string]*group)

	for _, p := range Places {
		if year != AllYears && p.Year != year {
			continue
		}
		g, ok := groups[p.Location]
		if !ok {
			g = &group{place: p}
			groups[p.Location] = g
			order = append(order, p.Location)
		}
		g.value += p.Value
		g.years = append(g.years, strconv.Itoa(p.Year))
		g.newest = max(g.newest, p.Year)

		// Repeat visits often carry the same note; listing it once reads better.
		if !slices.Contains(g.texts, p.Text) {
			g.texts = append(g.texts, p.Text)
		}
	}

	if len(order) == 0 {
		return nil
	}

	lo, hi := groups[order[0]].value, groups[order[0]].value
	for _, loc := range order[1:] {
		lo = min(lo, groups[loc].value)
		hi = max(hi, groups[loc].value)
	}

	markers := make([]Marker, 0, len(order))
	for _, loc := range order {
		g := groups[loc]
		scale := float64(g.value-lo) / (float64(hi-lo) + spread)

		markers = append(markers, Marker{
			Location: g.place.Location,
			Text:     strings.Join(g.texts, ", "),
			Years:    strings.Join(g.years, ", "),
			Lat:      g.place.Lat,
			Lng:      g.place.Lng,
			Radius:   minRadius + scale*(maxRadius-minRadius),
			Color:    colorFor(g.newest),
		})
	}
	return markers
}
