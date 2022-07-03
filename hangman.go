package main

import (
    "fmt"
    "os"
    "os/exec"
    "io/ioutil"
    "strings"
    "runtime"
    "bufio"
    "math/rand"
    "time"
)

var clear map[string]func()

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

var words []string
var word string

var guessedIndexes []bool

var wrong int = 0
const maxWrong int = 4

// ---

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        
        cmd.Run()
    }
}

func Clear() {
    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

// ---

func loadWords() {
    stream, err := ioutil.ReadFile("words.txt")
    
    if err != nil {
        panic(err)
    }
    
    str := string(stream)
    words = strings.Split(str, " ")
}

func printHangman(wrong int) {
    // 1 - head, 2 - head and body, 3 - head body and arms, 4 - complete
    
    switch wrong {
        case 0:
            fmt.Println("____")
            fmt.Println("|  |")
            fmt.Println("|")
            fmt.Println("|")
            fmt.Println("|")
            fmt.Println("|")
        break
        
        case 1:
            fmt.Println("____")
            fmt.Println("|  |")
            fmt.Println("|  o")
            fmt.Println("|")
            fmt.Println("|")
            fmt.Println("|")
        break
        
        case 2:
            fmt.Println("____")
            fmt.Println("|  |")
            fmt.Println("|  o")
            fmt.Println("|  |")
            fmt.Println("|")
            fmt.Println("|")
        break
        
        case 3:
            fmt.Println("____")
            fmt.Println("|  |")
            fmt.Println("|  o")
            fmt.Println("| /|\\ ")
            fmt.Println("|")
            fmt.Println("|")
        break
        
        case 4:
            fmt.Println("____")
            fmt.Println("|  |")
            fmt.Println("|  o")
            fmt.Println("| /|\\ ")
            fmt.Println("| / \\ ")
            fmt.Println("|")
        break
    }
}

func printLetters(word string, guessed []bool) {
    for i, c := range word {
        if guessed[i] == true {
            fmt.Printf(" %c ", c)
        } else {
            fmt.Printf(" _ ")
        }
    }
    
    fmt.Println()
}

func processInput(input string) {
    if len(input) != 1 {
        return
    }
    
    if strings.Contains(word, input) {
        guessedIndexes[strings.Index(word, input)] = true
    } else {
        wrong++
    }
    
    return
}

func gameOver() {
    Clear()
    
    fmt.Println("Hangman\n")
    printHangman(wrong)
    
    fmt.Println()
    
    printLetters(word, guessedIndexes)
    fmt.Println()
    
    fmt.Println("Game Over!")
    fmt.Printf("The word was %s.", word)
    fmt.Println()
    
    os.Exit(0)
}

func checkWin() bool {
    for i := 0; i < len(guessedIndexes); i++ {
        if guessedIndexes[i] == false {
            return false // didn't found all the letters
        }
    }
    
    return true
}

func doWin() {
    Clear()
    
    fmt.Println("Hangman\n")
    printHangman(wrong)
    
    fmt.Println()
    
    printLetters(word, guessedIndexes)
    fmt.Println()
    
    fmt.Println("You won!")
    
    os.Exit(0)
}

func main() {
    loadWords()
    
    rand.Seed(time.Now().UnixNano())
    
    word = words[rand.Intn(len(words))]
    guessedIndexes = make([]bool, len(word))
    
    for i := 0; i < len(guessedIndexes); i++ {
        guessedIndexes[i] = false
    }
    
    for {
        Clear()
        
        if wrong == maxWrong {
            gameOver()
        }
        
        if checkWin() {
            doWin()
        }
        
        fmt.Println("Hangman\n")
        printHangman(wrong)
        
        fmt.Println()
        printLetters(word, guessedIndexes)
        fmt.Println()
        
        // ---
        
        fmt.Printf("Enter letter: ")
        scanner.Scan()
        
        input := scanner.Text()
        
        processInput(input)
    }
}
