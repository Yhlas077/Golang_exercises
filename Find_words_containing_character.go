package main

import "fmt"

func main() {
    a := []string{"leet", "code"}
    b := byte('e') 
    fmt.Println(findWordsContaining(a, b))
}

func findWordsContaining(words []string, x byte) []int {
    result := []int{}

    for i := 0; i < len(words); i++ {
        currentWord := words[i]
        
        for j := 0; j < len(currentWord); j++ {
            if currentWord[j] == x {
                result = append(result, i)
                break 
            }
        }
    }
    return result
}