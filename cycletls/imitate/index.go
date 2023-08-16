package imitate

import (
	"math/rand"
	"strings"
)

func shuffleExtension(extension string) string {
	extensions := strings.Split(extension, "-")
	for i := range extensions {
		j := rand.Intn(len(extensions))
		extensions[i], extensions[j] = extensions[j], extensions[i]
	}
	return strings.Join(extensions, "-")
}
