package maps

import (
	"fmt"
)

func maps() {
	website := map[int]string{
		101: "google.com",
		102: "youtube.com",
	}
	fmt.Println(website)
	fmt.Println(website[101])
	website[103] = "facebook.com"
	website[104] = "instagram.com"
	fmt.Println(website)
	delete(website, 102)
	fmt.Println(website)
	oldwebsite := make(map[int]string, 4)

	oldwebsite[101] = "twitter.com"
	oldwebsite[102] = "yahoo.com"
	oldwebsite[103] = "opera.com"
	oldwebsite[104] = "nf.com"
	fmt.Println(oldwebsite)
	for key, value := range oldwebsite {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}
}
