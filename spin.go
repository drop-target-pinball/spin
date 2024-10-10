package spin

import "fmt"

func Banner() string {
	return fmt.Sprintf("%v %v (%v) git%v", ProgName, Version, BuildDate, Commit)
}
