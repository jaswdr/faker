package faker

var (
	languages = []string{"Algerian Arabic", "Amharic", "Assamese", "Bavarian", "Bengali", "Bhojpuri", "Burmese", "Cebuano", "Chhattisgarhi", "Chittagonian", "Czech", "Deccan", "Dutch", "Eastern Punjabi", "Egyptian Arabic", "English", "French", "Gan Chinese", "German", "Greek", "Gujarati", "Hakka Chinese", "Hausa", "Hejazi Arabic", "Hindi", "Hungarian", "Igbo", "Indonesian", "Iranian Persian", "Italian", "Japanese", "Javanese", "Jin Chinese", "Kannada", "Kazakh", "Khmer", "Kinyarwanda", "Korean", "Magahi", "Maithili", "Malayalam", "Malaysian", "Mandarin Chinese", "Marathi", "Mesopotamian Arabic", "Min Bei Chinese", "Min Dong Chinese", "Min Nan Chinese", "Moroccan Arabic", "Nepali", "Nigerian Fulfulde", "North Levantine Arabic", "Northern Kurdish", "Northern Pashto", "Northern Uzbek", "Odia", "Polish", "Portuguese", "Romanian", "Rundi", "Russian", "Saʽidi Arabic", "Sanaani Spoken Arabic", "Saraiki", "Sindhi", "Sinhalese", "Somali", "South Azerbaijani", "South Levantine Arabic", "Southern Pashto", "Spanish", "Sudanese Arabic", "Sunda", "Sylheti", "Tagalog", "Taʽizzi-Adeni Arabic", "Tamil", "Telugu", "Thai", "Tunisian Arabic", "Turkish", "Ukrainian", "Urdu", "Uyghur", "Vietnamese", "Western Punjabi", "Wu Chinese", "Xiang Chinese", "Yoruba", "Yue Chinese", "Zulu",}
)

// Language is a faker struct for Language
type Language struct {
	Faker *Faker
}

// Language returns a fake language name for Language
func (l Language) Language() string {
	return l.Faker.RandomStringElement(languages)
}