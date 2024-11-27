package object

import (
	"fmt"
	"hash/fnv"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/shopspring/decimal"
)

const STRING_OBJ = "STRING"

// type=mawu alternative=String
// The String object represents a string of characters.
// It is used to store and manipulate text.
// It is a sequence of characters, where each character is a Unicode code point.
// The String object is immutable, meaning that once a String object is created, it cannot be modified.
// However, you can create a new String object based on the original String object.
// The String object is a wrapper around the Go string type.
type String struct {
	Object
	Mappable
	Value string
}

func (string *String) Type() ObjectType { return STRING_OBJ }

func (string *String) String() string { return string.Value }

func (str *String) Method(method string, args []Object) (Object, bool) {
	switch method {
	case "peza":
		return str.methodFind(args)
	case "pezaZonse":
		return str.methodFindAll(args)
	case "format":
		return str.methodFormat(args)
	case "kumalizaNdi":
		return str.methodEndsWith(args)
	case "kutalika":
		return str.methodLength(args)
	case "gwirizana":
		return str.methodMatches(args)
	case "maloMwa":
		return str.methodReplace(args)
	case "gawa":
		return str.methodSplit(args)
	case "yayambaNdi":
		return str.methodStartsWith(args)
	case "toLowerCase":
		return str.methodToLowerCase(args)
	case "toUpperCase":
		return str.methodToUpperCase(args)
	case "kuMawu":
		return str.methodToString(args)
	case "kuNambala":
		return str.methodToNumber(args)
	case "chepetsa":
		return str.methodTrim(args)
	case "chepetsaKuMapeto":
		return str.methodTrimEnd(args)
	case "chepetsaKuchiyamba":
		return str.methodTrimStart(args)
	}

	return nil, false
}

func (string *String) MapKey() MapKey {
	h := fnv.New64a()
	h.Write([]byte(string.Value))
	return MapKey{Type: string.Type(), Value: h.Sum64()}
}

// method=peza args=[mawu{mawuOpeza}] return={mawu}
// This method finds the first occurrence of the given string in the string and returns it.
func (str *String) methodFind(args []Object) (Object, bool) {
	re := regexp.MustCompile(str.Value)

	found := re.FindStringSubmatch(args[0].(*String).Value)

	if len(found) > 0 {
		return &String{Value: found[1]}, true
	}

	return &String{}, true
}

// method=pezaZonse args=[mawu{mawuOpeza}] return={string[]}
// This method finds all occurrences of the given string in the string and returns them.
func (str *String) methodFindAll(args []Object) (Object, bool) {
	re := regexp.MustCompile(str.Value)
	list := &Array{}
	found := re.FindStringSubmatch(args[0].(*String).Value)

	for _, f := range found {
		list.Elements = append(list.Elements, &String{Value: f})
	}

	return list, true
}

// method=format args=[mawu{mawuOsintha}] return={mawu}
// This method formats the string with the given arguments.
func (str *String) methodFormat(args []Object) (Object, bool) {
	list := []interface{}{}

	for _, value := range args {
		list = append(list, value.String())
	}

	return &String{Value: fmt.Sprintf(str.Value, list...)}, true
}

// method=kumalizaNdi args=[mawu{mawuOsintha}] return={mawu}
// This method checks if the string ends with the given string.
func (str *String) methodEndsWith(args []Object) (Object, bool) {
	hasSuffix := strings.HasSuffix(str.Value, args[0].(*String).Value)

	return &Boolean{Value: hasSuffix}, true
}

// method=kutalika args=[] return={nambala}
// This method returns the length of the string.
func (str *String) methodLength(_ []Object) (Object, bool) {
	length := &Integer{Value: decimal.NewFromInt(int64(utf8.RuneCountInString(str.Value)))}

	return length, true
}

// method=gwirizana args=[mawu{mawuOsintha}] return={nambala}
// This method checks if the string matches the given regular expression.
func (str *String) methodMatches(args []Object) (Object, bool) {
	matches, err := regexp.Match(str.Value, []byte(args[0].(*String).Value))

	if err != nil {
		return &Error{Message: err.Error()}, false
	}

	return &Boolean{Value: matches}, true
}

// method=maloMwa args=[mawu{source},mawu{value}] return={mawu}
// This method replaces all occurrences of the given string with the new string.
func (str *String) methodReplace(args []Object) (Object, bool) {
	value := strings.Replace(str.Value, args[0].(*String).Value, args[1].(*String).Value, -1)

	return &String{Value: value}, true
}

// method=gawa args=[mawu{mawuOgawa}] return={string[]}
// This method splits the string by the given string and returns a list of strings.
func (str *String) methodSplit(args []Object) (Object, bool) {
	split := strings.Split(str.Value, args[0].(*String).Value)
	list := &Array{}

	for _, value := range split {
		list.Elements = append(list.Elements, &String{Value: value})
	}

	return list, true
}

// method=yayambaNdi args=[mawu{mawuOsintha}] return={mawu}
// This method checks if the string starts with the given string.
func (str *String) methodStartsWith(args []Object) (Object, bool) {
	hasPrefix := strings.HasPrefix(str.Value, args[0].(*String).Value)

	return &Boolean{Value: hasPrefix}, true
}

// method=toLowerCase args=[] return={mawu}
// This method converts the string to lowercase.
func (str *String) methodToLowerCase(_ []Object) (Object, bool) {
	return &String{Value: strings.ToLower(str.Value)}, true
}

// method=toUpperCase args=[] return={mawu}
// This method converts the string to uppercase.
func (str *String) methodToUpperCase(_ []Object) (Object, bool) {
	return &String{Value: strings.ToUpper(str.Value)}, true
}

// method=kuMawu args=[] return={mawu}
// This method converts the string to a string.
func (str *String) methodToString(_ []Object) (Object, bool) {
	return str, true
}

// method=kuNambala args=[] return={nambala}
// This method converts the string to a number.
func (str *String) methodToNumber(_ []Object) (Object, bool) {
	number, _ := decimal.NewFromString(str.Value)

	return &Integer{Value: number}, true
}

// method=chepetsa args=[] return={mawu}
// This method removes whitespace from the beginning and end of the string.
func (str *String) methodTrim(_ []Object) (Object, bool) {
	return &String{Value: strings.TrimSpace(str.Value)}, true
}

// method=chepetsaKuMapeto args=[] return={mawu}
// This method removes whitespace from the end of the string.
func (str *String) methodTrimEnd(_ []Object) (Object, bool) {
	return &String{Value: strings.TrimRight(str.Value, "\t\n\v\f\r ")}, true
}

// method=chepetsaKuchiyamba args=[] return={mawu}
// This method removes whitespace from the beginning of the string.
func (str *String) methodTrimStart(_ []Object) (Object, bool) {
	return &String{Value: strings.TrimLeft(str.Value, "\t\n\v\f\r ")}, true
}
