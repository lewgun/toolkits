// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
package main 
import "fmt"
func main() { 
        input := "The quick brown 狐 jumped over the lazy 犬" 
        // Get Unicode code points. 
        n := 0
        rune := make([]rune, len(input))
        for _, r := range input { 
                rune[n] = r
                n++
        } 
        rune = rune[0:n]
        // Reverse 
        for i := 0; i < n/2; i++ { 
                rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
        } 
        // Convert back to UTF-8. 
        output := string(rune)
        fmt.Println(output)
}
