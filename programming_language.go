package faker

// ProgrammingLanguage is a faker struct for ProgrammingLanguage
type ProgrammingLanguage struct {
	Faker *Faker
}

// Common programming language names
const (
	ABAP         = "ABAP"
	ALGOL        = "ALGOL"
	APL          = "APL"
	ASP          = "ASP / ASP.NET"
	ActionScript = "ActionScript"
	Ada          = "Ada"
	Alice        = "Alice"
	Assembly     = "Assembly Language"
	Awk          = "Awk"
	BBBasic      = "BBC Basic"
	C            = "C"
	COBOL        = "COBOL"
	CPP          = "C++"
	CSS          = "Cascading Style Sheets"
	CSharp       = "C#"
	D            = "D"
	Delphi       = "Delphi"
	Dreamweaver  = "Dreamweaver"
	Elixir       = "Elixir"
	Erlang       = "Erlang"
	FORTH        = "FORTH"
	FORTRAN      = "FORTRAN"
	FSharp       = "F#"
	Go           = "Go"
	HTML         = "HTML"
	Haskell      = "Haskell"
	IDL          = "IDL"
	INTERCAL     = "INTERCAL"
	Java         = "Java"
	JavaScript   = "JavaScript"
	Jquery       = "jQuery"
	LaTex        = "LaTeX"
	LabVIEW      = "LabVIEW"
	Lisp         = "Lisp"
	Logo         = "Logo"
	ML           = "ML"
	MQL          = "MetaQuotes Language"
	MSAccess     = "MS Access"
	Modula3      = "Modula-3"
	MySQL        = "MySQL"
	NXTG         = "NXT-G"
	OCaml        = "OCaml"
	ObjectC      = "Objective-C"
	PHP          = "PHP"
	PLI          = "PL/I"
	PLSQL        = "PL/SQL"
	PROLOG       = "PROLOG"
	Pascal       = "Pascal"
	Perl         = "Perl"
	PostScript   = "PostScript"
	PostgreSQL   = "PostgreSQL"
	PureData     = "Pure Data"
	Python       = "Python"
	R            = "R"
	RapidWeaver  = "RapidWeaver"
	RavenDB      = "RavenDB"
	Rexx         = "Rexx"
	Ruby         = "Ruby"
	RubyOnRails  = "Ruby on Rails"
	Rust         = "Rust"
	SAS          = "SAS"
	SGML         = "SGML"
	SMIL         = "SMIL"
	SNOBOL       = "SNOBOL"
	SPlus        = "S-PLUS"
	SQL          = "SQL"
	SSI          = "SSI"
	Scala        = "Scala"
	Sed          = "Sed"
	Shellscript  = "Shellscript"
	Simular      = "Simula"
	Smalltalk    = "Smalltalk"
	Stata        = "Stata"
	Swift        = "Swift"
	TclTk        = "Tcl/Tk"
	UML          = "Unified Modeling Language"
	VHDL         = "VHDL"
	VRML         = "VRML"
	Verilog      = "Verilog"
	VisualBasic  = "Visual Basic"
	VisualFoxPro = "Visual FoxPro"
	WAPWML       = "WAP/WML"
	XML          = "XML"
	XSL          = "XSL"
)

var languageVersions = map[string][]string{
	ABAP:         {"7.40", "7.50", "7.52", "7.54", "7.56", "7.58"},
	ALGOL:        {"58", "60", "68"},
	APL:          {"1.0", "2.0", "3.0", "4.0", "5.0"},
	ASP:          {"1.0", "2.0", "3.0", "4.0", "5.0"},
	ActionScript: {"1.0", "2.0", "3.0"},
	Ada:          {"83", "95", "2005", "2012", "2022"},
	Alice:        {"2.0", "2.2", "2.3", "3.0"},
	Assembly:     {"x86", "x64", "ARM", "MIPS"},
	Awk:          {"1.0", "2.0", "3.0", "4.0", "5.0"},
	BBBasic:      {"1.0", "2.0", "3.0", "4.0", "5.0"},
	C:            {"89", "90", "99", "11", "17", "23"},
	COBOL:        {"85", "2002", "2014", "2023"},
	CPP:          {"11", "14", "17", "20", "23"},
	CSS:          {"1", "2", "3", "4"},
	CSharp:       {"7.0", "8.0", "9.0", "10.0", "11.0"},
	D:            {"1.0", "2.0", "2.1", "2.2", "2.3"},
	Delphi:       {"1.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "10.0", "11.0"},
	Dreamweaver:  {"1.0", "2.0", "3.0", "4.0", "MX", "8", "CS3", "CS4", "CS5", "CS6", "CC"},
	Elixir:       {"1.0", "1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1.10", "1.11", "1.12", "1.13", "1.14"},
	Erlang:       {"R13", "R14", "R15", "R16", "R17", "R18", "R19", "R20", "R21", "R22", "R23", "R24", "R25", "R26"},
	FORTH:        {"79", "83", "94", "2012"},
	FORTRAN:      {"66", "77", "90", "95", "2003", "2008", "2018"},
	FSharp:       {"1.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0"},
	Go:           {"1.17", "1.18", "1.19", "1.20", "1.21", "1.22", "1.23", "1.24"},
	HTML:         {"2.0", "3.2", "4.0", "4.01", "5.0", "5.1", "5.2", "5.3"},
	Haskell:      {"98", "2010", "2020"},
	IDL:          {"6.0", "6.1", "6.2", "6.3", "6.4", "7.0", "7.1", "8.0", "8.1", "8.2", "8.3", "8.4", "8.5", "8.6", "8.7", "8.8"},
	INTERCAL:     {"1.0", "2.0", "3.0"},
	Java:         {"8", "11", "17", "19", "20"},
	JavaScript:   {"ES5", "ES6", "ES2017", "ES2018", "ES2019", "ES2020", "ES2021", "ES2022"},
	Jquery:       {"1.0", "1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1.10", "1.11", "1.12", "2.0", "2.1", "2.2", "3.0", "3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "3.7"},
	LaTex:        {"2.09", "2e", "3.0"},
	LabVIEW:      {"1.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "8.2", "8.5", "8.6", "2009", "2010", "2011", "2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022", "2023"},
	Lisp:         {"1.0", "1.5", "2.0", "2.5", "3.0"},
	Logo:         {"1.0", "2.0", "3.0", "4.0", "5.0", "6.0"},
	ML:           {"1.0", "2.0", "3.0", "4.0", "5.0"},
	MQL:          {"4", "5"},
	MSAccess:     {"2000", "2003", "2007", "2010", "2013", "2016", "2019"},
	Modula3:      {"1.0", "2.0", "3.0"},
	MySQL:        {"5.6", "5.7", "8.0"},
	NXTG:         {"1.0", "2.0"},
	OCaml:        {"4.0", "4.1", "4.2", "4.3", "5.0"},
	ObjectC:      {"1.0", "2.0", "2.1", "2.2", "2.3"},
	PHP:          {"7.4", "8.0", "8.1", "8.2"},
	PLI:          {"1.0", "2.0", "3.0"},
	PLSQL:        {"8i", "9i", "10g", "11g", "12c", "18c", "19c", "21c"},
	PROLOG:       {"1.0", "2.0", "3.0", "4.0", "5.0"},
	Pascal:       {"5.5", "6.0", "7.0", "7.1", "7.2"},
	Perl:         {"5.8", "5.10", "5.12", "5.14", "5.16", "5.18", "5.20", "5.22", "5.24", "5.26", "5.28", "5.30", "5.32"},
	PostScript:   {"1", "2", "3"},
	PostgreSQL:   {"9.6", "10", "11", "12", "13", "14", "15", "16"},
	PureData:     {"0.37", "0.38", "0.39", "0.40", "0.41"},
	Python:       {"2.7", "3.7", "3.8", "3.9", "3.10", "3.11"},
	R:            {"3.0", "3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "4.0", "4.1", "4.2"},
	RapidWeaver:  {"5.0", "6.0", "7.0", "8.0"},
	RavenDB:      {"3.5", "4.0", "4.1", "4.2", "5.0", "5.1", "5.2"},
	Rexx:         {"3.6", "4.0", "4.1", "4.2"},
	Ruby:         {"2.6", "2.7", "3.0", "3.1", "3.2"},
	RubyOnRails:  {"5.0", "5.1", "5.2", "6.0", "6.1", "7.0", "7.1"},
	Rust:         {"1.65", "1.66", "1.67", "1.68", "1.69", "1.70"},
	SAS:          {"9.1", "9.2", "9.3", "9.4"},
	SGML:         {"ISO 8879:1986"},
	SMIL:         {"1.0", "2.0", "3.0"},
	SNOBOL:       {"2", "3", "4"},
	SPlus:        {"3.4", "4.0", "6.0", "8.0"},
	SQL:          {"92", "99", "2003", "2006", "2008", "2011", "2016", "2019"},
	SSI:          {"2.0", "2.2", "2.4"},
	Scala:        {"2.11", "2.12", "2.13", "3.0", "3.1", "3.2", "3.3"},
	Sed:          {"4.0", "4.1", "4.2", "4.3"},
	Shellscript:  {"bash3", "bash4", "bash5"},
	Simular:      {"67", "87"},
	Smalltalk:    {"5.0", "5.5", "6.0", "7.0", "8.0"},
	Stata:        {"14", "15", "16", "17", "18"},
	Swift:        {"4.0", "4.2", "5.0", "5.1", "5.2", "5.3", "5.4", "5.7"},
	TclTk:        {"8.0", "8.1", "8.2", "8.3", "8.4", "8.5", "8.6", "8.7"},
	UML:          {"1.1", "1.2", "1.3", "1.4", "1.5", "2.0", "2.5"},
	VHDL:         {"87", "93", "2000", "2002", "2008", "2019"},
	VRML:         {"1.0", "2.0"},
	Verilog:      {"95", "2001", "2005"},
	VisualBasic:  {"6.0", "2005", "2008", "2010", "2012", "2015", "2017", "2019"},
	VisualFoxPro: {"6.0", "7.0", "8.0", "9.0"},
	WAPWML:       {"1.1", "1.2", "1.3", "2.0"},
	XML:          {"1.0", "1.1"},
	XSL:          {"1.0", "1.1", "2.0", "3.0"},
}

// Name returns a random programming language name
func (pl ProgrammingLanguage) Name() string {
	names := make([]string, 0, len(languageVersions))
	for lang := range languageVersions {
		names = append(names, lang)
	}
	return pl.Faker.RandomStringElement(names)
}

// Version returns a random version for the given language
func (pl ProgrammingLanguage) Version(language string) string {
	return pl.Faker.RandomStringElement(languageVersions[language])
}

// VariableName generates a random valid programming language identifier
func (pl ProgrammingLanguage) VariableName() string {
	return pl.VariableNameWithLength(pl.Faker.IntBetween(1, 11))
}

// VariableNameWithLength generates a random valid programming language identifier with the specified length
func (pl ProgrammingLanguage) VariableNameWithLength(length int) string {
	if length <= 0 {
		length = 1
	}

	// First character must be a letter or underscore
	firstCharSet := "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	restCharSet := firstCharSet + "0123456789"

	result := make([]byte, length)
	result[0] = firstCharSet[pl.Faker.IntBetween(0, len(firstCharSet)-1)]

	for i := 1; i < length; i++ {
		result[i] = restCharSet[pl.Faker.IntBetween(0, len(restCharSet)-1)]
	}

	return string(result)
}
