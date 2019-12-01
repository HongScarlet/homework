package main
import "fmt"
import "temperature"  // import $GOPATH/src/temperature下的temperature.go

var X string
var T float32

func main() {
	fmt.Println("Please enter temp and unit (ex: 100 C): ")
	fmt.Scanln(&T, &X)
	// fmt.Scanf  讀取使用者輸入(數值 & 單位)
	switch X {
	case "C":
		temperature.C2F(T)   	// 攝氏轉華氏
	case "F":
		temperature.F2C(T)   	// 華氏轉攝氏
	default:
		fmt.Println("Error")    // 輸入非C或F則顯示Error
	}
}

// import的function必須放置在gopath下的src內
// 本題中即是把temperature.go放在 $GOPATH/src/temperature之下 
