package content

import (
	"fmt"
	"net/url"
	"strings"
)

// The CDMX trip guide. Coordinates were geocoded once against OpenStreetMap and
// baked in here, the same way geography.go carries its own; nothing is looked up
// at runtime.

const (
	LangEN = "en"
	LangRU = "ru"
)

// CDMXLang narrows whatever arrived in the query string to a language we have.
func CDMXLang(v string) string {
	if v == LangRU {
		return LangRU
	}
	return LangEN
}

// Spanish names and addresses are deliberately left untranslated in the Russian
// version: they have to match the sign on the door, the Google Maps result and
// whatever gets shown to a taxi driver. Only names that are English to begin
// with carry a NameRU.

type CDMXCategory struct {
	Key     string
	Label   string
	LabelRU string
	Icon    string
	Color   string
}

func (c CDMXCategory) localized(lang string) CDMXCategory {
	if lang == LangRU && c.LabelRU != "" {
		c.Label = c.LabelRU
	}
	return c
}

type CDMXPlace struct {
	Name     string
	NameRU   string
	Category string
	Address  string
	Note     string
	NoteRU   string
	URL      string // official site, only when one was verified to resolve
	Slug     string // set only for places with a /cdmx/{slug} page
	Booking  string // "" when you can just turn up; otherwise a CDMXBookings key
	Lead     string // how far ahead to book, shown on the badge
	LeadRU   string
	// MapsQuery overrides what gets sent to Google. Needed when Name is a label
	// rather than something Google can look up, e.g. "Home".
	MapsQuery string
	Lat       float64
	Lng       float64
}

func (p CDMXPlace) localized(lang string) CDMXPlace {
	if lang != LangRU {
		return p
	}
	if p.NameRU != "" {
		// Pin the Google query to the Latin name before the Russian one
		// overwrites it, or the map link searches for Cyrillic.
		if p.MapsQuery == "" {
			p.MapsQuery = p.Name
			if p.Address != "" {
				p.MapsQuery += ", " + p.Address
			}
		}
		p.Name = p.NameRU
	}
	if p.NoteRU != "" {
		p.Note = p.NoteRU
	}
	if p.LeadRU != "" {
		p.Lead = p.LeadRU
	}
	return p
}

// CDMXStrings is the page furniture that belongs to no particular place.
type CDMXStrings struct {
	Heading     string
	Lede        string
	BookHeading string
	BookLede    string
	All         string
	Details     string
	Maps        string
	Directions  string
	Website     string
}

func CDMXText(lang string) CDMXStrings {
	if lang == LangRU {
		return CDMXStrings{
			Heading:     "Мехико",
			Lede:        "Американская мама встречает русскую маму. Две сестры — свидетели.",
			BookHeading: "Это планируем заранее",
			BookLede:    "Всё остальное на этой странице можно решить в тот же день.",
			All:         "Все",
			Details:     "Подробнее",
			Maps:        "Google Карты",
			Directions:  "Маршрут",
			Website:     "Сайт",
		}
	}
	return CDMXStrings{
		Heading:     "Mexico City",
		Lede:        "American Mom meets Russian Mom. Two sisters witness.",
		BookHeading: "Plan these early",
		BookLede:    "Everything else on this page you can decide on the day.",
		All:         "All",
		Details:     "More details",
		Maps:        "Google Maps",
		Directions:  "Directions",
		Website:     "Website",
	}
}

// Three tiers, because "needs a reservation" covers two very different risks:
// missing out entirely, and merely queueing.
var CDMXBookings = []CDMXCategory{
	{Key: "required", Label: "Must book ahead", LabelRU: "Обязательно бронировать заранее", Icon: "🔴", Color: "#dc2626"},
	{Key: "advised", Label: "Book ahead to be safe", LabelRU: "Лучше забронировать заранее", Icon: "🟠", Color: "#ea580c"},
	{Key: "optional", Label: "Optional, skips the queue", LabelRU: "По желанию — чтобы не стоять в очереди", Icon: "🔵", Color: "#0891b2"},
}

// Order here is the order of the filter buttons and of the list sections.
var CDMXCategories = []CDMXCategory{
	{Key: "home", Label: "Home", LabelRU: "Дом", Icon: "🏠"},
	{Key: "transport", Label: "Getting around", LabelRU: "Транспорт", Icon: "✈️"},
	{Key: "parks", Label: "Parks & walks", LabelRU: "Парки и прогулки", Icon: "🌳"},
	{Key: "cafes", Label: "Cafes", LabelRU: "Кофейни", Icon: "☕"},
	{Key: "dining", Label: "Dining", LabelRU: "Рестораны", Icon: "🍽️"},
	{Key: "russian", Label: "Russian food", LabelRU: "Русская еда", Icon: "🇷🇺"},
	{Key: "shopping", Label: "Shopping", LabelRU: "Шопинг", Icon: "🛍️"},
	{Key: "books", Label: "Bookstores", LabelRU: "Книжные магазины", Icon: "📚"},
	{Key: "sights", Label: "Main sights", LabelRU: "Главные достопримечательности", Icon: "🏛️"},
	{Key: "daytrips", Label: "Day trips", LabelRU: "Поездки за город", Icon: "🗿"},
	{Key: "entertainment", Label: "Entertainment", LabelRU: "Развлечения", Icon: "🎭"},
}

var CDMXPlaces = []CDMXPlace{
	{
		Category: "home", Name: "Home", NameRU: "Дом",
		Address:   "C. Ometusco 30, Hipódromo, Cuauhtémoc, 06100 Ciudad de México",
		MapsQuery: "C. Ometusco 30, Hipódromo, Cuauhtémoc, 06100 Ciudad de México",
		Lat:       19.4080863, Lng: -99.1729379,
	},
	{
		Category: "transport", Name: "Mexico City International Airport",
		NameRU:  "Международный аэропорт Мехико",
		Address: "Aeropuerto Internacional Benito Juárez, Venustiano Carranza, 15700 Ciudad de México",
		Note:    "Uber, or a prepaid airport taxi — buy the ticket at a kiosk inside the terminal before walking outside.",
		NoteRU:  "Uber или предоплаченное такси из аэропорта — билет покупайте в киоске внутри терминала, до выхода на улицу.",
		Slug:    "getting-around",
		Lat:     19.4342349, Lng: -99.0733121,
	},

	{
		Category: "parks", Name: "Avenida Ámsterdam",
		Address: "Hipódromo, 06100 Ciudad de México",
		Note:    "The morning walk — a leafy loop that follows the old racetrack.",
		NoteRU:  "Утренняя прогулка — зелёное кольцо по контуру старого ипподрома.",
		Lat:     19.4093600, Lng: -99.1701017,
	},
	{
		Category: "parks", Name: "Parque México",
		Address: "Hipódromo, Cuauhtémoc, 06100 Ciudad de México",
		Note:    "Fountains, dog walkers and street performers.",
		NoteRU:  "Фонтаны, собачники и уличные артисты.",
		URL:     "https://parquemexico.com.mx/",
		Lat:     19.4118592, Lng: -99.1698053,
	},
	{
		Category: "parks", Name: "Parque España",
		Address: "Colonia Condesa, Cuauhtémoc, 06140 Ciudad de México",
		Note:    "Fountains, dog walkers and street performers.",
		NoteRU:  "Фонтаны, собачники и уличные артисты.",
		Lat:     19.4149354, Lng: -99.1714106,
	},

	{
		Category: "cafes", Name: "Maque",
		Address: "C. Ozuluama 4, Hipódromo, Cuauhtémoc, 06100 Ciudad de México",
		Lat:     19.4103892, Lng: -99.1710885,
	},
	{
		Category: "cafes", Name: "BUNA",
		Address: "Ámsterdam 285, Colonia Condesa, Cuauhtémoc, 06100 Ciudad de México",
		URL:     "https://buna.mx/",
		Lat:     19.4131402, Lng: -99.1669681,
	},
	{
		// Approximate: OSM has no node at Tamaulipas 60, so this sits on the
		// right block rather than the doorway. Same for Kolobok below.
		Category: "cafes", Name: "Blend Station",
		Address: "Av. Tamaulipas 60, Condesa, 06140 Ciudad de México",
		URL:     "https://blendstation.com.mx/",
		Lat:     19.4099518, Lng: -99.1752665,
	},
	{
		// MapsQuery skips the "(Melanie)" tag, which Google can't look up.
		Category: "cafes", Name: "Panadería Rosetta (Melanie)",
		Address:   "Av. Medellín 21, Roma Nte., Cuauhtémoc, 06700 Ciudad de México",
		MapsQuery: "Panadería Rosetta, Av. Medellín 21, Roma Nte., Ciudad de México",
		URL:       "https://rosetta.com.mx/",
		Lat:       19.4215303, Lng: -99.1672094,
	},

	{
		Category: "dining", Name: "Lardo",
		Address: "Agustín Melgar 6, Colonia Condesa, Cuauhtémoc, 06140 Ciudad de México",
		URL:     "https://www.opentable.com.mx/restaurant/profile/185659?ref=16420",
		Booking: "advised", Lead: "Weekend brunch especially — bookable on OpenTable",
		LeadRU: "Особенно на бранч в выходные — бронь через OpenTable",
		Lat:    19.4179021, Lng: -99.1752714,
	},
	{
		Category: "dining", Name: "Taquería Orinoco",
		Address: "Av. Álvaro Obregón 100, Roma Nte., Cuauhtémoc, 06700 Ciudad de México",
		URL:     "https://taqueriaorinoco.com/",
		Lat:     19.4181560, Lng: -99.1592321,
	},
	{
		Category: "dining", Name: "Contramar",
		Address: "Durango 200, Roma Nte., Cuauhtémoc, 06700 Ciudad de México",
		Note:    "Lunch is the thing here. Walk-ins can wait for a table, but the wait is real.",
		NoteRU:  "Сюда идут на обед. Без брони пустят, но ждать столик придётся по-настоящему.",
		URL:     "http://www.contramar.com.mx/",
		Booking: "advised", Lead: "OpenTable opens ~30 days out and fills fast",
		LeadRU: "Бронь на OpenTable открывается примерно за 30 дней и разбирается быстро",
		Lat:    19.4196110, Lng: -99.1672259,
	},

	{
		Category: "russian", Name: "Boris Delicatessen",
		Address: "Carlos B. Zetina 12, Hipódromo Condesa, Cuauhtémoc, 06170 Ciudad de México",
		Note:    "Russian food store.",
		NoteRU:  "Магазин русских продуктов.",
		Lat:     19.4094635, Lng: -99.1818453,
	},
	{
		// Approximate — street-level, see the Blend Station note above.
		Category: "russian", Name: "Kolobok Escandón",
		Address: "Av. José Martí 160, Escandón I Secc, Miguel Hidalgo, 11800 Ciudad de México",
		URL:     "https://www.kolobok.com.mx/",
		Lat:     19.4033995, Lng: -99.1818056,
	},

	{
		Category: "shopping", Name: "Reforma 222",
		Address: "Av. P.º de la Reforma 222, Juárez, Cuauhtémoc, 06600 Ciudad de México",
		Lat:     19.4286929, Lng: -99.1615866,
	},
	{
		Category: "shopping", Name: "Antara Fashion Hall",
		Address: "Av. Ejército Nacional, Granada, Miguel Hidalgo, 11520 Ciudad de México",
		URL:     "https://www.antara.com.mx/",
		Lat:     19.4389154, Lng: -99.2022744,
	},
	{
		Category: "shopping", Name: "Avenida Presidente Masaryk",
		Address: "Polanco, Miguel Hidalgo, 11530 Ciudad de México",
		Note:    "The upmarket shopping street, a short walk from Antara.",
		NoteRU:  "Улица дорогих магазинов, в двух шагах от Antara.",
		Lat:     19.4323078, Lng: -99.2044299,
	},

	{
		Category: "books", Name: "Centro Cultural Bella Época",
		Address: "Av. Tamaulipas 202, Hipódromo, Cuauhtémoc, 06100 Ciudad de México",
		Note:    "The Rosario Castellanos bookshop, in a converted cinema.",
		NoteRU:  "Книжный магазин «Росарио Кастельянос» в здании бывшего кинотеатра.",
		URL:     "https://www.fondodeculturaeconomica.com/",
		Lat:     19.4073267, Lng: -99.1774222,
	},
	{
		Category: "books", Name: "Cafebrería El Péndulo Condesa",
		Address: "Av. Nuevo León 115, Colonia Condesa, Cuauhtémoc, 06140 Ciudad de México",
		Note:    "Bookshop and cafe in one.",
		NoteRU:  "Книжный магазин и кафе в одном месте.",
		URL:     "https://www.elpendulo.com/",
		Lat:     19.4105440, Lng: -99.1729781,
	},
	{
		Category: "books", Name: "Biblioteca Vasconcelos",
		Address: "Eje 1 Nte. S/N, Buenavista, Cuauhtémoc, 06350 Ciudad de México",
		Note:    "Must-see public library.",
		NoteRU:  "Публичная библиотека, которую обязательно стоит увидеть.",
		URL:     "https://www.bibliotecavasconcelos.gob.mx/",
		Lat:     19.4474857, Lng: -99.1508070,
	},

	{
		Category: "sights", Name: "Chapultepec Castle", NameRU: "Замок Чапультепек",
		Address: "Bosque de Chapultepec I Secc, Miguel Hidalgo, 11580 Ciudad de México",
		Note:    "Huge park with a castle and great views — the city's central park. Closed Mondays.",
		NoteRU:  "Огромный парк с замком и прекрасными видами — главный парк города. По понедельникам закрыт.",
		URL:     "https://mnh.inah.gob.mx/",
		Booking: "optional", Lead: "General admission is walk-up; ahead only to skip the line",
		LeadRU: "Обычный билет можно купить на месте; заранее — только чтобы не стоять в очереди",
		Lat:    19.4204632, Lng: -99.1820866,
	},
	{
		Category: "sights", Name: "National Museum of Anthropology",
		NameRU:  "Национальный музей антропологии",
		Address: "Av. P.º de la Reforma s/n, Bosque de Chapultepec I Secc, Miguel Hidalgo, 11560 Ciudad de México",
		URL:     "https://mna.inah.gob.mx/",
		Booking: "optional", Lead: "Walk-up is fine, but the queue can hit an hour",
		LeadRU: "Можно прийти без билета, но очередь бывает до часа",
		Lat:    19.4261524, Lng: -99.1866516,
	},
	{
		Category: "sights", Name: "Zócalo & the Historic Center",
		NameRU:  "Сокало и исторический центр",
		Address: "Plaza de la Constitución, Centro Histórico, Cuauhtémoc, Ciudad de México",
		Lat:     19.4326468, Lng: -99.1331985,
	},
	{
		Category: "sights", Name: "Frida Kahlo Museum", NameRU: "Музей Фриды Кало",
		Address: "Londres 247, Del Carmen, Coyoacán, 04100 Ciudad de México",
		Note:    "In Coyoacán. Online sales only — you cannot buy at the door.",
		NoteRU:  "В районе Койоакан. Билеты только онлайн — на входе не продают.",
		URL:     "https://www.museofridakahlo.org.mx/",
		Slug:    "frida",
		Booking: "required", Lead: "Book the moment dates are set — sells out weeks ahead",
		LeadRU: "Бронировать сразу, как определимся с датами — билеты разбирают за недели",
		Lat:    19.3551412, Lng: -99.1623564,
	},
	{
		// The friend-credit suffixes need a MapsQuery: Google can't look up
		// "Cuicuilco (Helen)".
		Category: "sights", Name: "Cuicuilco (Helen)",
		MapsQuery: "Zona Arqueológica de Cuicuilco, Ciudad de México",
		Address:   "Zona Arqueológica de Cuicuilco, Tlalpan, Ciudad de México",
		Note:      "Round pyramid, inside the city and almost never crowded.",
		NoteRU:    "Круглая пирамида — в черте города, и почти никогда нет толп.",
		URL:       "https://www.inah.gob.mx/zonas/zona-arqueologica-cuicuilco",
		Lat:       19.3009293, Lng: -99.1826980,
	},

	{
		Category: "daytrips", Name: "Teotihuacan Pyramids", NameRU: "Пирамиды Теотиуакана",
		Address: "San Juan Teotihuacán, State of Mexico",
		URL:     "https://www.teotihuacan.inah.gob.mx/",
		Slug:    "teotihuacan",
		Booking: "optional", Lead: "Walk-up works; buying ahead saves a long queue",
		LeadRU: "Можно купить на месте; заранее — чтобы не стоять в длинной очереди",
		Lat:    19.6923666, Lng: -98.8436073,
	},
	{
		Category: "daytrips", Name: "Xochicalco (Helen)",
		MapsQuery: "Zona Arqueológica de Xochicalco, Morelos",
		Address:   "Zona Arqueológica de Xochicalco, Temixco, Morelos",
		Note:      "Hilltop city, UNESCO-listed, a fraction of Teotihuacan's traffic.",
		NoteRU:    "Древний город на вершине холма, в списке ЮНЕСКО; посетителей несравнимо меньше, чем в Теотиуакане.",
		URL:       "https://www.inah.gob.mx/zonas/zona-arqueologica-de-xochicalco",
		Lat:       18.8040439, Lng: -99.2947104,
	},
	{
		Category: "daytrips", Name: "Malinalco (Helen)",
		MapsQuery: "Zona Arqueológica de Malinalco, Estado de México",
		Address:   "Zona Arqueológica de Malinalco, Malinalco, Estado de México",
		Note:      "Temple cut into the rock face, above an old town worth the walk itself — chapels, market, ordinary Mexican life.",
		NoteRU:    "Храм, вырубленный в скале, над старым городком, ради которого стоит остаться, — часовни, рынок, обычная мексиканская жизнь.",
		URL:       "https://www.inah.gob.mx/zonas/135-zona-arqueologica-malinalco",
		Lat:       18.9534825, Lng: -99.5032195,
	},
	{
		Category: "daytrips", Name: "Tula (Helen)",
		MapsQuery: "Zona Arqueológica de Tula, Tula de Allende, Hidalgo",
		Address:   "Zona Arqueológica de Tula, Tula de Allende, Hidalgo",
		Note:      "The Toltec capital and its stone warriors, north of the city.",
		NoteRU:    "Столица тольтеков и её каменные воины, к северу от города.",
		URL:       "https://www.inah.gob.mx/zonas/zona-arqueologica-y-museo-de-sitio-de-tula",
		Lat:       20.0638169, Lng: -99.3411080,
	},

	{
		Category: "entertainment", Name: "Jazzatlán Capital",
		Address: "Guanajuato 239, Roma Nte., Cuauhtémoc, 06700 Ciudad de México",
		Note:    "Live jazz, a walk from the house.",
		NoteRU:  "Живой джаз, до которого от дома можно дойти пешком.",
		URL:     "https://jazzatlan.club/",
		Slug:    "jazzatlan",
		Booking: "required", Lead: "Small room — book once we pick a night",
		LeadRU: "Зал маленький — бронируем, как только выберем вечер",
		Lat:    19.4158966, Lng: -99.1649284,
	},
	{
		Category: "entertainment", Name: "Plaza Garibaldi",
		Address: "Centro, Cuauhtémoc, 06010 Ciudad de México",
		Note:    "Mariachi bands.",
		NoteRU:  "Оркестры мариачи.",
		Lat:     19.4407117, Lng: -99.1389623,
	},
	{
		Category: "entertainment", Name: "Arena México",
		Address: "Dr. Lavista 189, Doctores, Cuauhtémoc, 06720 Ciudad de México",
		Note:    "Lucha libre. CMLL runs Tuesday, Friday and Sunday; Friday is the big one.",
		NoteRU:  "Луча либре — мексиканский рестлинг в масках. CMLL выступает по вторникам, пятницам и воскресеньям; пятница — главный вечер.",
		URL:     "https://www.cmll.com/",
		Booking: "advised", Lead: "Friday sells out — buy on Ticketmaster MX ahead",
		LeadRU: "На пятницу всё раскупают — берите заранее на Ticketmaster MX",
		Lat:    19.4246141, Lng: -99.1520014,
	},
	{
		Category: "entertainment", Name: "Embarcadero Nuevo Nativitas",
		Address: "C. del Mercado, San Jerónimo, Xochimilco, 16090 Ciudad de México",
		Note:    "The trajinera boats. No booking needed — hire at the pier, priced per boat, haggling expected.",
		NoteRU:  "Лодки-трахинеры. Бронировать не нужно — нанимаете прямо на причале, цена за лодку целиком, торговаться здесь в порядке вещей.",
		Slug:    "boats",
		Lat:     19.2520558, Lng: -99.0933501,
	},
}

// MapsURL searches Google by name and address rather than by coordinate, so the
// result is the business card — hours, photos, reviews — instead of a bare
// dropped pin. It also papers over the two approximate geocodes above.
func (p CDMXPlace) MapsURL() string {
	q := p.MapsQuery
	if q == "" {
		q = p.Name
		if p.Address != "" {
			q += ", " + p.Address
		}
	}
	return "https://www.google.com/maps/search/?api=1&query=" + strings.ReplaceAll(url.QueryEscape(q), "+", "%20")
}

// DirectionsURL goes straight to turn-by-turn, where the exact coordinate is
// what you want.
func (p CDMXPlace) DirectionsURL() string {
	return fmt.Sprintf("https://www.google.com/maps/dir/?api=1&destination=%v,%v", p.Lat, p.Lng)
}

// CDMXListPlace carries the place's index in CDMXPlaces, which is the id the
// map and the list card agree on. Grouping the list by category loses that
// index, so it rides along.
type CDMXListPlace struct {
	CDMXPlace
	ID int
}

type CDMXSection struct {
	CDMXCategory
	Places []CDMXListPlace
}

// CDMXSections groups the places for the written-out list, keeping the category
// order above and dropping any category that ended up empty.
func CDMXSections(lang string) []CDMXSection {
	var sections []CDMXSection
	for _, c := range CDMXCategories {
		var places []CDMXListPlace
		for i, p := range CDMXPlaces {
			if p.Category == c.Key {
				places = append(places, CDMXListPlace{CDMXPlace: p.localized(lang), ID: i})
			}
		}
		if len(places) > 0 {
			sections = append(sections, CDMXSection{CDMXCategory: c.localized(lang), Places: places})
		}
	}
	return sections
}

// CDMXToBook drives the planning callout: everything needing a booking, in
// tier order, so the things that sell out sit at the top.
func CDMXToBook(lang string) []CDMXSection {
	var sections []CDMXSection
	for _, b := range CDMXBookings {
		var places []CDMXListPlace
		for i, p := range CDMXPlaces {
			if p.Booking == b.Key {
				places = append(places, CDMXListPlace{CDMXPlace: p.localized(lang), ID: i})
			}
		}
		if len(places) > 0 {
			sections = append(sections, CDMXSection{CDMXCategory: b.localized(lang), Places: places})
		}
	}
	return sections
}

func (p CDMXPlace) BookingLabel(lang string) string {
	for _, b := range CDMXBookings {
		if b.Key == p.Booking {
			return b.localized(lang).Label
		}
	}
	return ""
}

// CDMXMarker is the JSON shape handed to the map. It flattens the category's
// icon and color onto each place so the JS never has to join the two lists.
type CDMXMarker struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Icon     string  `json:"icon"`
	Address  string  `json:"address"`
	Note     string  `json:"note"`
	Slug     string  `json:"slug"`
	MapsURL  string  `json:"mapsUrl"`
	Booking  string  `json:"booking"`
	BookText string  `json:"bookText"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

func CDMXMarkers(lang string) []CDMXMarker {
	icons := map[string]CDMXCategory{}
	for _, c := range CDMXCategories {
		icons[c.Key] = c
	}

	markers := make([]CDMXMarker, len(CDMXPlaces))
	for i, raw := range CDMXPlaces {
		p := raw.localized(lang)
		c := icons[p.Category]
		markers[i] = CDMXMarker{
			ID: i, Name: p.Name, Category: p.Category,
			Icon:    c.Icon,
			Address: p.Address, Note: p.Note, Slug: p.Slug,
			MapsURL: p.MapsURL(),
			Booking: p.Booking, BookText: p.BookingLabel(lang),
			Lat: p.Lat, Lng: p.Lng,
		}
	}
	return markers
}

type CDMXDetail struct {
	Slug    string
	Title   string
	TitleRU string
	Body    []string
	BodyRU  []string
}

func (d CDMXDetail) localized(lang string) CDMXDetail {
	if lang != LangRU {
		return d
	}
	if d.TitleRU != "" {
		d.Title = d.TitleRU
	}
	if len(d.BodyRU) > 0 {
		d.Body = d.BodyRU
	}
	return d
}

// Sub-pages for the places that need more than a line. Everything here comes
// from the trip notes; anything marked TODO is a deliberate blank rather than a
// guess.
var CDMXDetails = []CDMXDetail{
	{
		Slug:    "getting-around",
		Title:   "Getting around",
		TitleRU: "Транспорт",
		Body: []string{
			"From the airport: Uber, or a prepaid airport taxi. Buy the taxi ticket at one of the kiosks inside the terminal before you walk outside — pay at the kiosk, not at the car.",
			"Ecobici is the city bike-share. Stations are dense in Condesa, Roma and the Centro, which covers most of this list.",
			"TODO: how we're signing up for Ecobici, and whether it's worth it for a short trip.",
			"TODO: metro vs Uber for the longer hops (Coyoacán, Xochimilco, Teotihuacan).",
		},
		BodyRU: []string{
			"Из аэропорта: Uber или предоплаченное такси. Билет на такси покупайте в одном из киосков внутри терминала, ещё до выхода на улицу — платить нужно в киоске, а не водителю у машины.",
			"Ecobici — городской велопрокат. Станций много в Кондесе, Роме и в центре, а это большая часть мест из этого списка.",
			"TODO: как оформляем Ecobici и стоит ли оно того для короткой поездки.",
			"TODO: метро или Uber для дальних поездок (Койоакан, Сочимилько, Теотиуакан).",
		},
	},
	{
		Slug:    "boats",
		Title:   "The Xochimilco boats",
		TitleRU: "Лодки в Сочимилько",
		Body: []string{
			"Embarcadero Nuevo Nativitas is the pier we're aiming for, in Xochimilco in the far south of the city. The trajineras are the painted flat-bottomed boats; you hire the whole boat, not a seat.",
			"It's the longest trip on this list from the house — worth treating as its own outing rather than squeezing it beside something else.",
			"Nothing to book. You hire at the pier, the price is per boat rather than per person, and haggling is expected. Nuevo Nativitas has a visitors' centre with the official price list posted and a government rep on hand, so check the board before agreeing to anything.",
			"Buy from the boatmen directly rather than from the touts who approach you on the way in.",
			"If it ends up being only one or two of us, a collectivo trajinera runs between piers for a fraction of the price.",
			"TODO: how long we want to go out for, and whether we bring our own food and drink or buy from the boats that come alongside.",
		},
		BodyRU: []string{
			"Нам нужен причал Embarcadero Nuevo Nativitas в Сочимилько, на самом юге города. Трахинеры — это те самые расписные плоскодонные лодки; нанимают лодку целиком, а не место на ней.",
			"Это самая дальняя поездка из всех в списке — лучше выделить под неё отдельный выход, а не втискивать между другими делами.",
			"Бронировать ничего не нужно. Лодку нанимают на причале, цена — за лодку, а не за человека, и торговаться здесь в порядке вещей. В Nuevo Nativitas есть центр для посетителей, где вывешен официальный прейскурант и дежурит представитель властей, так что сначала смотрим на доску с ценами и только потом договариваемся.",
			"Договаривайтесь напрямую с лодочниками, а не с зазывалами, которые подходят по дороге к причалу.",
			"Если в итоге поедут один-два человека, между причалами ходит коллективная трахинера — это стоит в разы дешевле.",
			"TODO: на сколько часов выходим и берём ли еду и напитки с собой или покупаем с лодок, которые подплывают по пути.",
		},
	},
	{
		Slug:    "frida",
		Title:   "Frida Kahlo Museum",
		TitleRU: "Музей Фриды Кало",
		Body: []string{
			"The Casa Azul, in Coyoacán. This is the one thing on the list we can genuinely miss out on: tickets are sold online only, there is no ticket window at the door, and they sell out weeks ahead.",
			"Entry runs on 30-minute timed slots. The slot is your entry window, not how long you get inside, so plan to arrive 10–15 minutes early.",
			"Official tickets are at boletos.museofridakahlo.org.mx. Book as soon as the dates are fixed — during holidays and high season people are booking a month or more out.",
			"Coyoacán itself is worth the trip beyond the museum, so it pairs well with an afternoon down there.",
			"TODO: book, then put the date and time slot here so everyone can see it.",
		},
		BodyRU: []string{
			"«Синий дом» в Койоакане. Это единственное место в списке, куда мы реально можем не попасть: билеты продаются только онлайн, кассы на входе нет, и всё разбирают за недели вперёд.",
			"Вход по получасовым сеансам. Сеанс — это окно, в которое вас впустят, а не время, которое вы проведёте внутри, так что приходите за 10–15 минут.",
			"Официальные билеты — на boletos.museofridakahlo.org.mx. Бронируем сразу, как определимся с датами: в праздники и в высокий сезон берут за месяц и раньше.",
			"Сам Койоакан стоит поездки и без музея, так что имеет смысл заложить на этот район полдня.",
			"TODO: забронировать и записать сюда дату и время сеанса, чтобы все видели.",
		},
	},
	{
		Slug:    "teotihuacan",
		Title:   "Teotihuacan Pyramids",
		TitleRU: "Пирамиды Теотиуакана",
		Body: []string{
			"Outside Mexico City, in the State of Mexico — a day trip rather than an afternoon.",
			"No reservation is required; you can buy at the gate. Buying ahead on the INAH portal (boletos.inah.gob.mx) mainly saves queue time, and foreign cards are often rejected there, so don't count on it working.",
			"TODO: how we're getting there (tour, bus from Terminal Norte, or a driver) and what time we want to leave.",
		},
		BodyRU: []string{
			"За пределами Мехико, в штате Мехико — это поездка на целый день, а не на полдня.",
			"Бронировать не обязательно, билеты продают на входе. Покупка заранее на портале INAH (boletos.inah.gob.mx) в основном экономит время в очереди, но иностранные карты там часто не проходят, так что рассчитывать на это не стоит.",
			"TODO: как добираемся (экскурсия, автобус с Terminal Norte или водитель) и во сколько выезжаем.",
		},
	},
	{
		Slug:    "jazzatlan",
		Title:   "Jazzatlán Capital",
		TitleRU: "Jazzatlán Capital",
		Body: []string{
			"Live jazz in Roma Norte, a walk from the house. Open Tuesday to Sunday from 5pm; shows start at 8pm.",
			"It's a small room and shows sell out, so book — especially Thursday to Sunday. They're not on OpenTable, so it's a phone call: +52 55 5459 2840.",
			"Cover runs roughly 100–600 pesos depending on who's playing. Arrive well before 8pm if you want a good seat.",
			"TODO: pick a night, check who's playing, and book.",
		},
		BodyRU: []string{
			"Живой джаз в Роме Норте, от дома можно дойти пешком. Работает со вторника по воскресенье с 17:00, концерты начинаются в 20:00.",
			"Зал маленький, и билеты разбирают, поэтому бронируем — особенно с четверга по воскресенье. На OpenTable их нет, так что только по телефону: +52 55 5459 2840.",
			"Вход стоит примерно 100–600 песо в зависимости от того, кто играет. Если хотите хорошие места, приходите заметно раньше 20:00.",
			"TODO: выбрать вечер, посмотреть, кто играет, и забронировать.",
		},
	},
}

func CDMXDetailBySlug(slug, lang string) (CDMXDetail, bool) {
	for _, d := range CDMXDetails {
		if d.Slug == slug {
			return d.localized(lang), true
		}
	}
	return CDMXDetail{}, false
}

func CDMXPlaceBySlug(slug, lang string) (CDMXPlace, bool) {
	for _, p := range CDMXPlaces {
		if p.Slug == slug {
			return p.localized(lang), true
		}
	}
	return CDMXPlace{}, false
}
