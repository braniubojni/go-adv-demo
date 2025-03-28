package link

import (
	"go/adv-demo/internal/stat"
	"go/adv-demo/pkg"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url:  url,
		Hash: RandStringRunes(6),
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = pkg.LETTERS[rand.Intn(len(pkg.LETTERS))]
	}

	return string(b)
}
