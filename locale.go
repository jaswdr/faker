package faker

// LocaleCode identifies a locale for localized data generation.
type LocaleCode string

const (
	// LocaleEnUS is the default English (United States) locale.
	LocaleEnUS LocaleCode = "en_US"
	// LocaleDeDE is the German (Germany) locale.
	LocaleDeDE LocaleCode = "de_DE"
)

// LocaleData holds localized name data for a locale.
type LocaleData struct {
	Code            LocaleCode
	FirstNameMale   []string
	FirstNameFemale []string
	LastName        []string
}

var localeRegistry = map[LocaleCode]LocaleData{
	LocaleDeDE: localeDeDE,
}

// RegisterLocale adds or replaces locale data for custom locale support.
func RegisterLocale(data LocaleData) {
	localeRegistry[data.Code] = data
}

func localeDataFor(code LocaleCode) *LocaleData {
	if code == "" || code == LocaleEnUS {
		return nil
	}
	data, ok := localeRegistry[code]
	if !ok {
		return nil
	}
	return &data
}

// Locale returns the active locale code for this Faker instance.
func (f Faker) Locale() LocaleCode {
	if f.locale == "" {
		return LocaleEnUS
	}
	return f.locale
}

// WithLocale returns a copy of the Faker configured with the given locale.
func (f Faker) WithLocale(code LocaleCode) Faker {
	f.locale = code
	return f
}

// NewWithLocale returns a new Faker instance with the given locale.
func NewWithLocale(code LocaleCode) Faker {
	return New().WithLocale(code)
}
