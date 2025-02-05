package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/entities"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringArray(arr []string) string {
	return arr[RandomInt(0, int64(len(arr)-1))]
}

func RandomTask() entities.Task {
	user := []string{
		"firstUser",
		"secondUser",
		"thirdUser",
	}
	task := entities.Task{
		ID:          uuid.New().String(),
		Content:     RandomString(8),
		CreatedDate: GetDate(),
		UserID:      RandomStringArray(user),
	}

	return task
}

func GetDate() string {
	return time.Now().Format(Conf.FormatDate)
}
