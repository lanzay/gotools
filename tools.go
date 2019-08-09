package tools

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Min(args ...int) int {

	min := args[0]
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(args ...int) int {

	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func LeftRuneValid(s string, lenS int) string {

	by := []byte(s)
	posE := Min(lenS, len(s))
	for posE := posE; posE > 0; posE-- {
		if utf8.Valid(by[:posE]) {
			return string(by[:posE])
		}
	}
	return ""

	if len(by) > 1 {
		lenS = Min(lenS-1, len(s))
		if int(by[lenS-2]) >= 224 {
			lenS -= 2
		} else if int(by[lenS-1]) >= 192 && int(by[lenS-1]) < 224 {
			lenS -= 1
		}
		return s[0:lenS]
	}
	return s
}

//Получить левую часть строки до целого количества RUNE - не ломая кодировку
func LeftRuneMax(str string, l int) string {

	if len(str) > l {
		ch := []rune(str)
		result := make([]rune, 0, l+1)
		for i := range ch {
			result = append(result, ch[i])
			if len(string(result)) > l {
				resultStr := string(result[:i])
				return resultStr
			}
		}
	}
	return str
}

//Получает подстроку байт по указанному началу и окончанию строки
//также задается смещение границ относительо поисковых строк
func SubByte(body []byte, startStr string, startShift int, finStr string, finShift int) []byte {

	if body == nil {
		return nil
	}
	start := bytes.Index(body, []byte(startStr)) + startShift
	if start <= 0 {
		return nil
	}

	fin := bytes.Index(body[start:], []byte(finStr)) + finShift
	if (start-startShift) == -1 || (fin-finShift) == -1 {
		return nil
	}
	sub := body[start : start+fin]
	return sub

}

func ErrPanic(err error) {

	if err != nil {
		log.Panicln("[F]", err)
	}
}
func ErrPrint(err error) {

	if err != nil {
		log.Println("[F]", err)
	}
}
func ErrPanicDebug(err error) {

	if err != nil {
		log.Panicln("[F] DEB ", err)
	}
}

func Atoi(str string) (int, error) {

	str = strings.ReplaceAll(str, " ", "")
	return strconv.Atoi(str)
}

func AtoiMust(str string) int {

	x, _ := Atoi(str)
	return x
}

func EscapeSQL(str string) string {

	str = strings.ReplaceAll(str, "`", "'")
	str = strings.ReplaceAll(str, "'", "''")
	str = strings.ReplaceAll(str, "--", "-")
	return str

}
func LeftEscapeSQL(str string, l int) string {

	tmp := strings.ReplaceAll(str, "`", "'")
	tmp = strings.ReplaceAll(tmp, "'", "''")
	tmp = strings.ReplaceAll(tmp, "--", "-")
	tmp = LeftRuneMax(tmp, l)

	if len(tmp) == l {
		tmp = strings.ReplaceAll(tmp, "'", "")
		tmp = strings.ReplaceAll(tmp, "`", "")
	}

	return tmp

}
