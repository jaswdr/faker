package faker

import "fmt"

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const videoLength = 11

type YouTube struct {
	Faker *Faker
}

// GenerateVideoID returns an 11 characters long string,
// which consists of both upper and lower case alphabets and numeric values.
// It is used to define a YouTube video uniquely.
func (y YouTube) GenerateVideoID() (videoID string) {
	b := make([]byte, videoLength)
	for i := range b {
		b[i] = charset[y.Faker.RandomNumber(len(charset))]
	}
	return string(b)
}

// GenerateFullURL returns a fake, standard youtube video URL
func (y YouTube) GenerateFullURL() string {
	return fmt.Sprintf("www.youtube.com/watch?v=%s", y.GenerateVideoID())
}

// GenerateShareURL returns a fake, share youtube video URL
func (y YouTube) GenerateShareURL() string {
	return fmt.Sprintf("youtu.be/%s", y.GenerateVideoID())
}

// GenerateEmbededURL returns a fake, embeded youtube video URL
func (y YouTube) GenerateEmbededURL() string {
	return fmt.Sprintf("www.youtube.com/embed/%s", y.GenerateVideoID())
}
