package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"math/bits"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n请输入密码>>> ")
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("接收到密码>>> ", readString)
		//if isValidPassword(strings.TrimSpace(readString)) {
		//	fmt.Print("密码正确!")
		//	break
		//}
		//

		//if err = isPwd(strings.TrimSpace(readString)); err == nil {
		//	fmt.Print("密码正确!")
		//	break
		//}
		//log.Println(err)

		if err = isRightPwd(strings.TrimSpace(readString)); err == nil {
			fmt.Print("密码正确!")
			break
		}
		log.Println(err)
	}
}

func isPwd(pwd string) error {
	if len(pwd) < 8 {
		return errors.New("密码长度低于八位，请重新输入")
	}

	var low, up, num, pun int8
	for _, ch := range pwd {
		switch {
		case unicode.IsDigit(ch):
			num++
		case unicode.IsLower(ch):
			low++
		case unicode.IsUpper(ch):
			up++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			pun++
		}
	}

	var kindCount int8
	if low > 0 {
		kindCount++
	}
	if up > 0 {
		kindCount++
	}
	if num > 0 {
		kindCount++
	}
	if pun > 0 {
		kindCount++
	}
	if kindCount < 3 {
		return errors.New("密码需包含数字、大写字母、小写字母、特殊符号等任意三种")
	}

	return nil
}

func isValidPassword(password string) bool {
	// 正则表达式，匹配包含数字、大写字母、小写字母、特殊符号的任意三种组合，且长度大于等于8位
	//regex := regexp.MustCompile(`^(?=(.*\d))(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*[\p{P}\p{S}]))(?!(.*\s)).{8,}$`)
	//regex := regexp2.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A_Z])(?=.*[)`, 0)
	//regex := regexp2.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[\p{P}\p{S}])(?!.*[=]).{8,}$`, 0)
	regex := regexp2.MustCompile(`^(?=(.*\d))(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*[\p{P}\p{S}]))(?:.*[a-zA-Z\d\p{P}\p{S}]){8,}$`, 0)
	matchString, err := regex.MatchString(password)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return matchString
}

func isRightPwd(pwd string) error {
	if len(pwd) < 8 {
		return errors.New("密码长度小于八位")
	}

	var bit uint8
	if regexp.MustCompile(`\d`).MatchString(pwd) {
		bit |= 1 << 0
	}

	if regexp.MustCompile(`[a-z]`).MatchString(pwd) {
		bit |= 1 << 1
	}

	if regexp.MustCompile(`[A-Z]`).MatchString(pwd) {
		bit |= 1 << 2
	}

	if regexp.MustCompile(`[\p{P}\p{S}]`).MatchString(pwd) {
		bit |= 1 << 3
	}

	fmt.Println("bit:", bit, "bitsOne:", bits.OnesCount8(bit))

	if bits.OnesCount8(bit) < 3 {
		return errors.New("密码至少包含三种")
	}
	return nil
}
