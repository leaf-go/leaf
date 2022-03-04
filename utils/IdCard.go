package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	idCardExpr = `^(\d{6})(\d{8})[0-9xX]{1,4}$`
)

type idCard struct {
	CardNo         string
	Birthday       string
	FormatBirthDay string
	Year           int
	Month          int
	Day            int
}

func NewIdCard(cardNo string) *idCard {
	return &idCard{
		CardNo: cardNo,
	}
}

func (card *idCard) Parse() (err error) {
	reg, err := regexp.Compile(idCardExpr)
	if err != nil {
		return
	}

	ok := reg.MatchString(card.CardNo)
	if !ok {
		return errors.New("not is a idcard")
	}

	matchs := reg.FindStringSubmatch(card.CardNo)
	if len(matchs) < 2 {
		return errors.New("not is a idcard")
	}

	card.Birthday = matchs[2]
	card.FormatBirthDay = card.Birthday[0:4] + "-" + card.Birthday[4:6] + "-" + card.Birthday[6:8]
	card.Year, _ = strconv.Atoi(card.Birthday[0:4])
	card.Month, _ = strconv.Atoi(card.Birthday[4:6])
	card.Day, _ = strconv.Atoi(card.Birthday[6:8])

	return nil
}

func (card idCard) Age() int {
	now := time.Now()
	age := now.Year() - card.Year // 如果计算虚岁需这样：age := now_year - idcard_year+1
	nowTime := int(now.Month())*100 + now.Day()
	cardTime := card.Month*100 + card.Day

	if nowTime < cardTime {
		age -= 1
	}

	return age
}

// Gender 性别 0男 1女
func (card idCard) Gender() int {
	count := strings.Count(card.CardNo, "") - 1
	var bit int
	if count == 18 {
		bit, _ = strconv.Atoi(card.CardNo[16:17])
	} else {
		bit, _  = strconv.Atoi(card.CardNo[14:15])
	}

	if bit%2 == 1 {
		return 0
	}

	return 1
}
