package faker

import (
	"slices"
	"testing"
)

func TestRegisterProvider(t *testing.T) {
	RegisterProvider("test_greeting", func(f *Faker) string {
		return "hello-" + f.Lexify("???")
	})

	f := NewWithSeedInt64(1)
	result := f.Provider("test_greeting")
	Expect(t, true, len(result) > 6)
	Expect(t, "hello-", result[:6])
}

func TestProviderUnknown(t *testing.T) {
	f := New()
	Expect(t, "", f.Provider("nonexistent_provider_xyz"))
}

func TestRegisteredProviders(t *testing.T) {
	RegisterProvider("test_list_provider", func(f *Faker) string {
		return "x"
	})
	names := RegisteredProviders()
	Expect(t, true, slices.Contains(names, "test_list_provider"))
}
