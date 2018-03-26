package faker

import (
	"strings"
	"testing"
)

func TestAreaCode(t *testing.T) {
	p := New().Phone()
	Expect(t, true, len(p.AreaCode()) == 3)
}

func TestExchangeCode(t *testing.T) {
	p := New().Phone()
	Expect(t, true, len(p.ExchangeCode()) == 3)
}

func TestNumber(t *testing.T) {
	a := New().Phone()
	number := a.Number()
	Expect(t, true, len(number) > 0)
	Expect(t, false, strings.Contains(number, "{{areaCode}}"))
	Expect(t, false, strings.Contains(number, "{{exchangeCode}}"))
	Expect(t, false, strings.Contains(number, "#"))
	Expect(t, false, strings.Contains(number, "{{"))
	Expect(t, false, strings.Contains(number, "}}"))
}

func TestTollFreeAreaCode(t *testing.T) {
	a := New().Phone()
	code := a.TollFreeAreaCode()
	Expect(t, true, len(code) > 0)
}

func TestTollFreeNumber(t *testing.T) {
	a := New().Phone()
	number := a.ToolFreeNumber()
	Expect(t, true, len(number) > 0)
	Expect(t, false, strings.Contains(number, "{{tollFreeAreaCode}}"))
	Expect(t, false, strings.Contains(number, "{{exchangeCode}}"))
	Expect(t, false, strings.Contains(number, "#"))
	Expect(t, false, strings.Contains(number, "{{"))
	Expect(t, false, strings.Contains(number, "}}"))
}

func TestE164Number(t *testing.T) {
	a := New().Phone()
	number := a.E164Number()
	Expect(t, true, len(number) > 0)
}
