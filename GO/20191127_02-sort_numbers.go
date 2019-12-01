package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "strconv"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//使用者輸入提示
	fmt.Println("--------------------")	
	fmt.Print("Enter some number (ex:1 22 333 1234 12)\n")
	fmt.Print("Get output like:\n")
	fmt.Print("Sort reverse by value:1234 333 22 12 1\n")	
	fmt.Print("Sort by len of strings:1 22 12 333 1234\n")
	fmt.Println("--------------------")			
	fmt.Print("Your input:")

	//輸入處理
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text=text[0:len(text)-2]
	fmt.Println("--------------------")	

	a1 :=[]int{}
	a2 := strings.Split(text, " ")           //讀取的字串用空格分開塞入a2
	a3 :=[]string{}

	for i :=0;i < len(a2);i++{               //轉換成數字塞入a1中
		N,_:=strconv.Atoi(a2[i])
		a1 = append(a1,N)
	}
    //排序處理
	sort.Sort(sort.Reverse(sort.IntSlice(a1))) //由大到小對a1中的int排序
	sort.Sort(ByLength(a2))                    //依照字串長度由小到大對a2排序

	for i :=0;i < len(a1);i++{                //把反轉過a1的int改成string塞入a3中
	S:=strconv.Itoa(a1[i])
	a3 = append(a3,S)
	}

	//印出結果
	fmt.Println("Sort reverse by value:", strings.Join(a3, " "))  //用join的方式 用空格將a3的元素印出
	fmt.Println("Sort by len of strings:", strings.Join(a2, " ")) //用join的方式 用空格將a2的元素印出
}
