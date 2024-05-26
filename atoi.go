

package main
import "fmt"


func findDup(str string) bool{
    
    dup := map[string]string{}
    for _, char := range str{
        if dup[string(char)] == string(char){
        return true
        }
        dup[string(char)] = string(char)
    }
    return false
}
func AtoiBase(s1, s2 string) int{
    
    Base := len(s2)
    res := 0
       index := 0
    for i:= 0; i < len(s1); i++{
     
        for j:= 0; j < len(s2); j++{
            if s1[i] == s2[j] {
                index = j
            }
        }
        res = (res * Base) + index
    }
    return res
}
func main() {
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	  fmt.Println(findDup("00123456789ABCDEF"))
  }