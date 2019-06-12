package atoi

type State int

const (
	Skipping	itoa

)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0;
	}

	isParsing := false;

	for _, r := range str {
		if !isParsing && r == ' ' {
			continue;
		}

		if r == '-'
	}

	return 0;
}