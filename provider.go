package faker

import "sync"

// ProviderFunc generates a string value using the given Faker instance.
type ProviderFunc func(f *Faker) string

var (
	customProviders = make(map[string]ProviderFunc)
	providerMu      sync.RWMutex
)

// RegisterProvider registers a custom named provider function.
func RegisterProvider(name string, fn ProviderFunc) {
	providerMu.Lock()
	customProviders[name] = fn
	providerMu.Unlock()
}

// Provider invokes a registered custom provider by name.
// Returns an empty string if the provider is not registered.
func (f Faker) Provider(name string) string {
	providerMu.RLock()
	fn, ok := customProviders[name]
	providerMu.RUnlock()
	if !ok {
		return ""
	}
	return fn(&f)
}

// RegisteredProviders returns the names of all registered custom providers.
func RegisteredProviders() []string {
	providerMu.RLock()
	defer providerMu.RUnlock()
	names := make([]string, 0, len(customProviders))
	for name := range customProviders {
		names = append(names, name)
	}
	return names
}
