func TestDeterministicStringsWithSeed(t *testing.T) {
    seed := int64(12345)
    f1 := NewWithSeedInt64(seed)
    f2 := NewWithSeedInt64(seed)

    // If Fill generates strings, create a test struct
    type S struct {
        A string
        B int
        C string
    }

    var s1, s2 S
    f1.Fill(&s1)
    f2.Fill(&s2)
    if !reflect.DeepEqual(s1, s2) {
        t.Fatalf("expected same results for same seed; got %#v and %#v", s1, s2)
    }
}
// lcgString generates deterministic ASCII strings using a tiny LCG.
// seed: initial seed, n: length
func lcgString(seed uint64, n int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const m = uint64(1<<63 - 1) // not used as modulus for speed; we'll use full 64-bit wrap
    x := seed
    b := make([]byte, n)
    for i := 0; i < n; i++ {
        // simple LCG: x = (a*x + c)
        x = x*6364136223846793005 + 1442695040888963407
        b[i] = letters[(x>>1)%uint64(len(letters))]
    }
    _ = m
    return string(b)
}
package faker_test

import (
    "math/rand"
    "testing"

    jf "github.com/jaswdr/faker" // adjust import path if the module path is different
)

// simple deterministic provider using the RNG instance
func deterministicStringProvider(r *rand.Rand, n int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, n)
    for i := 0; i < n; i++ {
        b[i] = letters[r.Int63()%int64(len(letters))]
    }
    return string(b)
}

func TestDeterministicStringsWithSeed_PublicProvider(t *testing.T) {
    seed := int64(123456789)

    // create two faker instances with same seed
    f1 := jf.NewWithSeedInt64(seed)
    f2 := jf.NewWithSeedInt64(seed)

    // Try to set deterministic string provider in each (adapt to API)
    // Many libs expose SetStringProvider or WithStringProvider; adjust if needed.
    // The provider signature may differ; this is an example.
    if spSetter, ok := interface{}(f1).(interface{ SetStringProvider(func(*rand.Rand, int) string) }); ok {
        spSetter.SetStringProvider(deterministicStringProvider)
    } else {
        t.Log("f1.SetStringProvider not available; skipping public-provider test")
        t.SkipNow()
    }
    if spSetter2, ok := interface{}(f2).(interface{ SetStringProvider(func(*rand.Rand, int) string) }); ok {
        spSetter2.SetStringProvider(deterministicStringProvider)
    }

    // Generate a sequence of strings from both instances and ensure equality
    for i := 0; i < 50; i++ {
        s1 := f1.Lorem().Sentence(1) // adapt the call to whichever string API your faker provides
        s2 := f2.Lorem().Sentence(1)
        if s1 != s2 {
            t.Fatalf("strings differ at iteration %d: %q vs %q", i, s1, s2)
        }
    }
}
import { enhancePromptWithGroq } from './src/promptEnhancer.js';
const enhanced = await enhancePromptWithGroq("Draw a class diagram for a bookstore inventory system.");
console.log(enhanced);

