package faker

import (
	"strconv"
	"strings"
	"sync"
)

var (
	cityPrefix = []string{"North", "East", "West", "South", "New", "Lake", "Port"}

	citySuffix = []string{"town", "ton", "land", "ville", "berg", "burgh", "borough", "bury", "view", "port", "mouth", "stad", "furt", "chester", "mouth", "fort", "haven", "side", "shire"}

	buildingNumber = []string{"####", "###", "##"}

	streetSuffix = []string{
		"Alley", "Avenue",
		"Branch", "Bridge", "Brook", "Brooks", "Burg", "Burgs", "Bypass",
		"Camp", "Canyon", "Cape", "Causeway", "Center", "Centers", "Circle", "Circles", "Cliff", "Cliffs", "Club", "Common", "Corner", "Corners", "Course", "Court", "Courts", "Cove", "Coves", "Creek", "Crescent", "Crest", "Crossing", "Crossroad", "Curve",
		"Dale", "Dam", "Divide", "Drive", "Drive", "Drives",
		"Estate", "Estates", "Expressway", "Extension", "Extensions",
		"Fall", "Falls", "Ferry", "Field", "Fields", "Flat", "Flats", "Ford", "Fords", "Forest", "Forge", "Forges", "Fork", "Forks", "Fort", "Freeway",
		"Garden", "Gardens", "Gateway", "Glen", "Glens", "Green", "Greens", "Grove", "Groves",
		"Harbor", "Harbors", "Haven", "Heights", "Highway", "Hill", "Hills", "Hollow",
		"Inlet", "Inlet", "Island", "Island", "Islands", "Islands", "Isle", "Isle",
		"Junction", "Junctions",
		"Key", "Keys", "Knoll", "Knolls",
		"Lake", "Lakes", "Land", "Landing", "Lane", "Light", "Lights", "Loaf", "Lock", "Locks", "Locks", "Lodge", "Lodge", "Loop",
		"Mall", "Manor", "Manors", "Meadow", "Meadows", "Mews", "Mill", "Mills", "Mission", "Mission", "Motorway", "Mount", "Mountain", "Mountain", "Mountains", "Mountains",
		"Neck",
		"Orchard", "Oval", "Overpass",
		"Park", "Parks", "Parkway", "Parkways", "Pass", "Passage", "Path", "Pike", "Pine", "Pines", "Place", "Plain", "Plains", "Plains", "Plaza", "Plaza", "Point", "Points", "Port", "Port", "Ports", "Ports", "Prairie", "Prairie",
		"Radial", "Ramp", "Ranch", "Rapid", "Rapids", "Rest", "Ridge", "Ridges", "River", "Road", "Road", "Roads", "Roads", "Route", "Row", "Rue", "Run",
		"Shoal", "Shoals", "Shore", "Shores", "Skyway", "Spring", "Springs", "Springs", "Spur", "Spurs", "Square", "Square", "Squares", "Squares", "Station", "Station", "Stravenue", "Stravenue", "Stream", "Stream", "Street", "Street", "Streets", "Summit", "Summit",
		"Terrace", "Throughway", "Trace", "Track", "Trafficway", "Trail", "Trail", "Tunnel", "Tunnel", "Turnpike", "Turnpike",
		"Underpass", "Union", "Unions",
		"Valley", "Valleys", "Via", "Viaduct", "View", "Views", "Village", "Village", "Villages", "Ville", "Vista", "Vista",
		"Walk", "Walks", "Wall", "Way", "Ways", "Well", "Wells",
	}

	postCode = []string{"#####", "#####-####"}

	state = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District of Columbia", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	stateAbbr = []string{"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME", "MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM", "NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VA", "VT", "WA", "WI", "WV", "WY"}

	country = []string{
		"Afghanistan", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antarctica (the territory South of 60 deg S)", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan",
		"Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia and Herzegovina", "Botswana", "Bouvet Island (Bouvetoya)", "Brazil", "British Indian Ocean Territory (Chagos Archipelago)", "British Virgin Islands", "Brunei Darussalam", "Bulgaria", "Burkina Faso", "Burundi",
		"Cambodia", "Cameroon", "Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China", "Christmas Island", "Cocos (Keeling) Islands", "Colombia", "Comoros", "Congo", "Cook Islands", "Costa Rica", "Cote d\"Ivoire", "Croatia", "Cuba", "Cyprus", "Czech Republic",
		"Denmark", "Djibouti", "Dominica", "Dominican Republic",
		"Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Ethiopia",
		"Faroe Islands", "Falkland Islands (Malvinas)", "Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern Territories",
		"Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana",
		"Haiti", "Heard Island and McDonald Islands", "Holy See (Vatican City State)", "Honduras", "Hong Kong", "Hungary",
		"Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy",
		"Jamaica", "Japan", "Jersey", "Jordan",
		"Kazakhstan", "Kenya", "Kiribati", "Korea", "Korea", "Kuwait", "Kyrgyz Republic",
		"Lao People\"s Democratic Republic", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libyan Arab Jamahiriya", "Liechtenstein", "Lithuania", "Luxembourg",
		"Macao", "Macedonia", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia", "Moldova", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar",
		"Namibia", "Nauru", "Nepal", "Netherlands Antilles", "Netherlands", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk Island", "Northern Mariana Islands", "Norway",
		"Oman",
		"Pakistan", "Palau", "Palestinian Territories", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Pitcairn Islands", "Poland", "Portugal", "Puerto Rico",
		"Qatar",
		"Reunion", "Romania", "Russian Federation", "Rwanda",
		"Saint Barthelemy", "Saint Helena", "Saint Kitts and Nevis", "Saint Lucia", "Saint Martin", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia (Slovak Republic)", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Georgia and the South Sandwich Islands", "Spain", "Sri Lanka", "Sudan", "Suriname", "Svalbard & Jan Mayen Islands", "Swaziland", "Sweden", "Switzerland", "Syrian Arab Republic",
		"Taiwan", "Tajikistan", "Tanzania", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu",
		"Uganda", "Ukraine", "United Arab Emirates", "United Kingdom", "United States of America", "United States Minor Outlying Islands", "United States Virgin Islands", "Uruguay", "Uzbekistan",
		"Vanuatu", "Venezuela", "Vietnam",
		"Wallis and Futuna", "Western Sahara",
		"Yemen",
		"Zambia", "Zimbabwe",
	}

	countryAbbr = []string{
		"ABW", "AFG", "AGO", "AIA", "ALA", "ALB", "AND", "ARE", "ARG", "ARM", "ASM", "ATA", "ATF", "ATG", "AUS", "AUT", "AZE", "BDI", "BEL", "BEN", "BES", "BFA", "BGD", "BGR", "BHR", "BHS", "BIH", "BLM", "BLR", "BLZ", "BMU", "BOL", "BRA", "BRB", "BRN", "BTN", "BVT", "BWA", "CAF", "CAN", "CCK", "CHE", "CHL", "CHN", "CIV", "CMR", "COD", "COG", "COK", "COL", "COM", "CPV", "CRI", "CUB", "CUW", "CXR", "CYM", "CYP", "CZE", "DEU", "DJI", "DMA", "DNK", "DOM", "DZA", "ECU", "EGY", "ERI", "ESH", "ESP", "EST", "ETH", "FIN", "FJI", "FLK", "FRA", "FRO",
		"FSM", "GAB", "GBR", "GEO", "GGY", "GHA", "GIB", "GIN", "GLP", "GMB", "GNB", "GNQ", "GRC", "GRD", "GRL", "GTM", "GUF", "GUM", "GUY", "HKG", "HMD", "HND", "HRV", "HTI", "HUN", "IDN", "IMN", "IND", "IOT", "IRL", "IRN", "IRQ", "ISL", "ISR", "ITA", "JAM", "JEY", "JOR", "JPN", "KAZ", "KEN", "KGZ", "KHM", "KIR", "KNA", "KOR", "KWT", "LAO", "LBN", "LBR", "LBY", "LCA", "LIE", "LKA", "LSO", "LTU", "LUX", "LVA", "MAC", "MAF", "MAR", "MCO", "MDA", "MDG", "MDV", "MEX", "MHL", "MKD", "MLI", "MLT", "MMR", "MNE", "MNG", "MNP", "MOZ", "MRT", "MSR", "MTQ", "MUS", "MWI",
		"MYS", "MYT", "NAM", "NCL", "NER", "NFK", "NGA", "NIC", "NIU", "NLD", "NOR", "NPL", "NRU", "NZL", "OMN", "PAK", "PAN", "PCN", "PER", "PHL", "PLW", "PNG", "POL", "PRI", "PRK", "PRT", "PRY", "PSE", "PYF", "QAT", "REU", "ROU", "RUS", "RWA", "SAU", "SDN", "SEN", "SGP", "SGS", "SHN", "SJM", "SLB", "SLE", "SLV", "SMR", "SOM", "SPM", "SRB", "SSD", "STP", "SUR", "SVK", "SVN", "SWE", "SWZ", "SXM", "SYC", "SYR", "TCA", "TCD", "TGO", "THA", "TJK", "TKL", "TKM", "TLS", "TON", "TTO", "TUN", "TUR", "TUV", "TWN", "TZA", "UGA", "UKR", "UMI", "URY", "USA", "UZB", "VAT",
		"VCT", "VEN", "VGB", "VIR", "VNM", "VUT", "WLF", "WSM", "YEM", "ZAF", "ZMB", "ZWE",
	}

	countryCode = []string{
		"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ",
		"BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ",
		"CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ",
		"DE", "DJ", "DK", "DM", "DO", "DZ",
		"EC", "EE", "EG", "EH", "ER", "ES", "ET",
		"FI", "FJ", "FK", "FM", "FO", "FR",
		"GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY",
		"HK", "HM", "HN", "HR", "HT", "HU",
		"ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT",
		"JE", "JM", "JO", "JP",
		"KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ",
		"LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY",
		"MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ",
		"NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ",
		"OM",
		"PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY",
		"QA",
		"RE", "RO", "RS", "RU", "RW",
		"SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ",
		"TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ",
		"UA", "UG", "UM", "US", "UY", "UZ",
		"VA", "VC", "VE", "VG", "VI", "VN", "VU",
		"WF", "WS",
		"YE", "YT",
		"ZA", "ZM", "ZW",
	}

	cityFormats = []string{
		"{{cityPrefix}} {{firstName}}{{citySuffix}}",
		"{{cityPrefix}} {{firstName}}",
		"{{firstName}}{{citySuffix}}",
		"{{lastName}}{{citySuffix}}",
	}

	streetNameFormats = []string{
		"{{firstName}} {{streetSuffix}}",
		"{{lastName}} {{streetSuffix}}",
	}

	streetAddressFormats = []string{
		"{{buildingNumber}} {{streetName}}",
		"{{buildingNumber}} {{streetName}} {{secondaryAddress}}",
	}

	addressFormats = []string{"{{streetAddress}}\n{{city}}, {{stateAbbr}} {{postCode}}"}

	secondaryAddressFormats = []string{"Apt. ###", "Suite ###"}
)

// Template cache for optimized processing
var (
	templateCache   = make(map[string][]templatePart)
	templateCacheMu sync.RWMutex
)

// templatePart represents a part of a template - either literal text or a placeholder
type templatePart struct {
	isPlaceholder bool
	text          string
}

// parseTemplate parses a template string into parts for efficient processing
func parseTemplate(template string) []templatePart {
	// Check cache first
	templateCacheMu.RLock()
	if parts, exists := templateCache[template]; exists {
		templateCacheMu.RUnlock()
		return parts
	}
	templateCacheMu.RUnlock()

	// Parse the template
	var parts []templatePart
	remaining := template

	for len(remaining) > 0 {
		start := strings.Index(remaining, "{{")
		if start == -1 {
			// No more placeholders, add the rest as literal
			if len(remaining) > 0 {
				parts = append(parts, templatePart{false, remaining})
			}
			break
		}

		// Add literal text before placeholder
		if start > 0 {
			parts = append(parts, templatePart{false, remaining[:start]})
		}

		// Find end of placeholder
		end := strings.Index(remaining[start:], "}}")
		if end == -1 {
			// Malformed template, treat rest as literal
			parts = append(parts, templatePart{false, remaining})
			break
		}

		// Add placeholder
		placeholder := remaining[start+2 : start+end]
		parts = append(parts, templatePart{true, placeholder})

		remaining = remaining[start+end+2:]
	}

	// Cache the parsed template
	templateCacheMu.Lock()
	templateCache[template] = parts
	templateCacheMu.Unlock()

	return parts
}

// processTemplate efficiently processes a template with given replacements
func processTemplate(template string, replacements map[string]func() string) string {
	parts := parseTemplate(template)

	if len(parts) == 0 {
		return template
	}

	var builder strings.Builder
	estimatedSize := len(template) + 50 // Estimate final size
	builder.Grow(estimatedSize)

	for _, part := range parts {
		if part.isPlaceholder {
			if fn, exists := replacements[part.text]; exists {
				builder.WriteString(fn())
			} else {
				// Keep placeholder if no replacement found
				builder.WriteString("{{")
				builder.WriteString(part.text)
				builder.WriteString("}}")
			}
		} else {
			builder.WriteString(part.text)
		}
	}

	return builder.String()
}

// Address is a faker struct for Address
type Address struct {
	Faker *Faker
}

// CityPrefix returns a fake city prefix for Address
func (a Address) CityPrefix() string {
	return a.Faker.RandomStringElement(cityPrefix)
}

// SecondaryAddress returns a fake secondary address for Address
func (a Address) SecondaryAddress() string {
	format := a.Faker.RandomStringElement(secondaryAddressFormats)
	return a.Faker.Bothify(format)
}

// State returns a fake state for Address
func (a Address) State() string {
	return a.Faker.RandomStringElement(state)
}

// StateAbbr returns a fake state abbreviation for Address
func (a Address) StateAbbr() string {
	return a.Faker.RandomStringElement(stateAbbr)
}

// CitySuffix returns a fake city suffix for Address
func (a Address) CitySuffix() string {
	return a.Faker.RandomStringElement(citySuffix)
}

// StreetSuffix returns a fake street suffix for Address
func (a Address) StreetSuffix() string {
	return a.Faker.RandomStringElement(streetSuffix)
}

// BuildingNumber returns a fake building number for Address
func (a Address) BuildingNumber() (bn string) {
	pattern := a.Faker.RandomStringElement(buildingNumber)
	return a.Faker.Numerify(pattern)
}

// City returns a fake city for Address
func (a Address) City() string {
	template := a.Faker.RandomStringElement(cityFormats)
	p := a.Faker.Person()

	replacements := map[string]func() string{
		"cityPrefix": a.CityPrefix,
		"firstName":  p.FirstName,
		"lastName":   p.LastName,
		"citySuffix": a.CitySuffix,
	}

	return processTemplate(template, replacements)
}

// StreetName returns a fake street name for Address
func (a Address) StreetName() string {
	template := a.Faker.RandomStringElement(streetNameFormats)
	p := a.Faker.Person()

	replacements := map[string]func() string{
		"firstName":    p.FirstName,
		"lastName":     p.LastName,
		"streetSuffix": a.StreetSuffix,
	}

	return processTemplate(template, replacements)
}

// StreetAddress returns a fake street address for Address
func (a Address) StreetAddress() string {
	template := a.Faker.RandomStringElement(streetAddressFormats)

	replacements := map[string]func() string{
		"buildingNumber":   a.BuildingNumber,
		"streetName":       a.StreetName,
		"secondaryAddress": a.SecondaryAddress,
	}

	return processTemplate(template, replacements)
}

// PostCode returns a fake postal code for Address
func (a Address) PostCode() string {
	format := a.Faker.RandomStringElement(postCode)
	return a.Faker.Bothify(format)
}

// Address returns a fake Address
func (a Address) Address() string {
	template := a.Faker.RandomStringElement(addressFormats)

	replacements := map[string]func() string{
		"streetAddress": a.StreetAddress,
		"city":          a.City,
		"stateAbbr":     a.StateAbbr,
		"postCode":      a.PostCode,
	}

	return processTemplate(template, replacements)
}

// Country returns a fake country for Address
func (a Address) Country() string {
	return a.Faker.RandomStringElement(country)
}

// CountryAbbr returns a fake country abbreviation for Address
func (a Address) CountryAbbr() string {
	return a.Faker.RandomStringElement(countryAbbr)
}

// CountryCode returns a fake country code for Address
func (a Address) CountryCode() string {
	return a.Faker.RandomStringElement(countryCode)
}

// Latitude returns a fake latitude for Address
func (a Address) Latitude() (latitude float64) {
	latitude, _ = strconv.ParseFloat(a.Faker.Numerify("##.######"), 64)
	return
}

// Longitude returns a fake longitude for Address
func (a Address) Longitude() (longitude float64) {
	longitude, _ = strconv.ParseFloat(a.Faker.Numerify("##.######"), 64)
	return
}
