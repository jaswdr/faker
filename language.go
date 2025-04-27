package faker

var (
	languages     = []string{"Algerian Arabic", "Amharic", "Assamese", "Bavarian", "Bengali", "Bhojpuri", "Burmese", "Cebuano", "Chhattisgarhi", "Chittagonian", "Czech", "Deccan", "Dutch", "Eastern Punjabi", "Egyptian Arabic", "English", "French", "Gan Chinese", "German", "Greek", "Gujarati", "Hakka Chinese", "Hausa", "Hejazi Arabic", "Hindi", "Hungarian", "Igbo", "Indonesian", "Iranian Persian", "Italian", "Japanese", "Javanese", "Jin Chinese", "Kannada", "Kazakh", "Khmer", "Kinyarwanda", "Korean", "Magahi", "Maithili", "Malayalam", "Malaysian", "Mandarin Chinese", "Marathi", "Mesopotamian Arabic", "Min Bei Chinese", "Min Dong Chinese", "Min Nan Chinese", "Moroccan Arabic", "Nepali", "Nigerian Fulfulde", "North Levantine Arabic", "Northern Kurdish", "Northern Pashto", "Northern Uzbek", "Odia", "Polish", "Portuguese", "Romanian", "Rundi", "Russian", "Saʽidi Arabic", "Sanaani Spoken Arabic", "Saraiki", "Sindhi", "Sinhalese", "Somali", "South Azerbaijani", "South Levantine Arabic", "Southern Pashto", "Spanish", "Sudanese Arabic", "Sunda", "Sylheti", "Tagalog", "Taʽizzi-Adeni Arabic", "Tamil", "Telugu", "Thai", "Tunisian Arabic", "Turkish", "Ukrainian", "Urdu", "Uyghur", "Vietnamese", "Western Punjabi", "Wu Chinese", "Xiang Chinese", "Yoruba", "Yue Chinese", "Zulu"}
	languagesAbbr = []string{"aa", "ab", "af", "am", "ar", "as", "ay", "az", "ba", "be", "bg", "bh", "bi", "bn", "bo", "br", "ca", "co", "cs", "cy", "da", "de", "dz", "el", "en", "eo", "es", "et", "eu", "fa", "fi", "fj", "fo", "fr", "fy", "ga", "gd", "gl", "gn", "gu", "ha", "he", "hi", "hr", "hu", "hy", "ia", "id", "ie", "ik", "in", "is", "it", "iu", "iw", "ja", "ji", "jw", "ka", "kk", "kl", "km", "kn", "ko", "ks", "ku", "ky", "la", "ln", "lo", "lt", "lv", "mg", "mi", "mk", "ml", "mn", "mo", "mr", "ms", "mt", "my", "na", "ne", "nl", "no", "oc", "om", "or", "pa", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "rw", "sa", "sd", "sg", "sh", "si", "sk", "sl", "sm", "sn", "so", "sq", "sr", "ss", "st", "su", "sv", "sw", "ta", "te", "tg", "th", "ti", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ug", "uk", "ur", "uz", "vi", "vo", "wo", "xh", "yi", "yo", "za", "zh", "zu"}
)

// Language is a faker struct for Language
type Language struct {
	Faker *Faker
}

// Language returns a fake language name for Language
func (l Language) Language() string {
	return l.Faker.RandomStringElement(languages)
}

// LanguageAbbr returns a fake language name for Language
func (l Language) LanguageAbbr() string {
	return l.Faker.RandomStringElement(languagesAbbr)
}

// ProgrammingLanguage returns a fake programming language for Language
func (l Language) ProgrammingLanguage() string {
	return l.Faker.ProgrammingLanguage().Name()
}
