package service

import (
	"api"
	"api/pkg/repository"
	"math/rand"
	"time"
)

type ShortService struct {
	repo repository.Short
}

func NewShortService(repo repository.Short) *ShortService {
	return &ShortService{
		repo: repo,
	}
}
func (s *ShortService) Create(user api.Short) (string, error) {
	user.ShortCode = CreateShortCode()
	user.ShortURL = "http://localhost/" + CreateShortCode()
	return s.repo.Create(user)
}

func (s *ShortService) GetUrl(shortUrl string) (string, error) {
	return s.repo.GetUrl(shortUrl)
}

func CreateShortCode() string {
	uprReg := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowReg := []rune("abcdefghijklmnopqrstuvwxyz")
	numbers := []rune("1234567890")
	underlsnding := []rune("_")
	short := []rune("")
	rez := ""
	var min, max int = 1, 5
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 11; i++ {

		n := min + rand.Intn(max-min)
		//var i int = 0
		switch n {
		case 1:
			short = append(short, uprReg[rand.Intn(len(uprReg))])

		case 2:

			short = append(short, lowReg[rand.Intn(len(lowReg))])
			i++

		case 3:
			short = append(short, numbers[rand.Intn(len(numbers))])

		case 4:
			short = append(short, underlsnding[rand.Intn(len(underlsnding))])

		}
	}
	for _, a := range short {
		rez = rez + string(a)
	}
	return rez
}
