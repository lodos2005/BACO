package main

import (
    "fmt"
    "os"
    "strings"
)

func GenerateWordlist(profile UserProfile) []string {
    var wordlist []string

    // Temel kombinasyonlar
    wordlist = append(wordlist, profile.FirstName)
    wordlist = append(wordlist, profile.LastName)
    wordlist = append(wordlist, profile.FirstName+profile.LastName)
    wordlist = append(wordlist, profile.LastName+profile.FirstName)

    // Assume profile.BirthDate is in the format "ddmmyy"
    if len(profile.BirthDate) == 8 {
        day := profile.BirthDate[:2]   // "dd"
        month := profile.BirthDate[2:4] // "mm"
        year := profile.BirthDate[4:8]  // "yyyy"
        yearshort := profile.BirthDate[6:8]  // "yy"
    
        if (profile.Complexity > 0){
            wordlist = append(wordlist, profile.FirstName+profile.BirthDate)
            wordlist = append(wordlist, profile.LastName+profile.BirthDate)
        }
        if (profile.Complexity > 1){
            wordlist = append(wordlist, profile.FirstName+year)
            wordlist = append(wordlist, profile.LastName+year)
        }
        if (profile.Complexity > 2){
            wordlist = append(wordlist, profile.FirstName+day)
            wordlist = append(wordlist, profile.FirstName+month)
            wordlist = append(wordlist, profile.LastName+day)
            wordlist = append(wordlist, profile.LastName+month)
        }
        if (profile.Complexity > 3){
            wordlist = append(wordlist, profile.FirstName+yearshort)
            wordlist = append(wordlist, profile.LastName+yearshort)
        }
        if (profile.Complexity > 4){
            wordlist = append(wordlist, profile.FirstName+yearshort+yearshort)
            wordlist = append(wordlist, profile.LastName+yearshort+yearshort)
        }

    
        // Takma adlar ve sevdiği kelimelerle kombinasyonlar

        for _, nickname := range profile.Nicknames {
            wordlist = append(wordlist, nickname)
            
            if (profile.Complexity > 0){
                wordlist = append(wordlist, nickname+profile.BirthDate)
            }
            if (profile.Complexity > 1){
                wordlist = append(wordlist, nickname+year)
            }
            if (profile.Complexity > 2){
                wordlist = append(wordlist, nickname+day)
                wordlist = append(wordlist, nickname+month)
            }
            if (profile.Complexity > 3){
                wordlist = append(wordlist, nickname+yearshort)
            }
            if (profile.Complexity > 4){
                wordlist = append(wordlist, nickname+yearshort+yearshort)
            }
        }

        for _, word := range profile.FavoriteWords {
            wordlist = append(wordlist, word)
            if (profile.Complexity > 0){
                wordlist = append(wordlist, word+profile.BirthDate)
            }
            if (profile.Complexity > 1){
                wordlist = append(wordlist, word+year)
            }
            if (profile.Complexity > 2){
                wordlist = append(wordlist, word+day)
                wordlist = append(wordlist, word+month)
            }
            if (profile.Complexity > 3){
                wordlist = append(wordlist, word+yearshort)
            }
            if (profile.Complexity > 4){
                wordlist = append(wordlist, word+yearshort+yearshort)
            }
        }
    }else{

        for _, word := range profile.FavoriteWords {
            wordlist = append(wordlist, word)
        }

        for _, nickname := range profile.Nicknames {
            wordlist = append(wordlist, nickname)
        }
    }

    


    var extendedEl33tWordlist []string

    // El33t düzenlemesi ekle
    for _, word := range wordlist {
        leetVariations := leetSpeakVariations(word)
        extendedEl33tWordlist = append(extendedEl33tWordlist, leetVariations...)
    }

    wordlist = append(wordlist, extendedEl33tWordlist...)


    // Özel karakter ve sayı eklemeleri
    var minyear int = 1900
    var maxyear int = 2050
    var numbers []string
    var specialChars []string
    if (profile.Complexity > 0){
        specialChars = []string{}
        minyear = 1990
        maxyear = 2025
    }
    if (profile.Complexity > 1){
        specialChars = []string{"!", "@", "#", "."}
        minyear = 1950
        maxyear = 2035

    }
    if (profile.Complexity > 2){
        specialChars = []string{"!", "@", "#", ".", "$", "%", "&", "*",}
        minyear = 1900
        maxyear = 2050        
        numbers = append(numbers, "1071","1299","1453","123","123456","123654") // Adding the predefined numbers

    }
    if (profile.Complexity > 3){
        specialChars = []string{"!", "@", "#", "$", "%", "&", "*",".",",","+"}
        minyear = 1900
        maxyear = 2050
        numbers = append(numbers, "1","2","3","4","5","6","7","8","9","0", "1071","1299","1453","123","123456","123654") // Adding the predefined numbers

    }
    if (profile.Complexity > 4){        
        specialChars = []string{"!", "@", "#", "$", "%", "&", "*",".",",","+","_","-","?"}
        minyear = 1900
        maxyear = 2050        
        numbers = append(numbers, "1","2","3","4","5","6","7","8","9","0", "1071","1299","1453","123","123456","123654","123123","123456789","123123","987654321") // Adding the predefined numbers
    }
   
    for year := minyear; year <= maxyear; year++ { 
        numbers = append(numbers, fmt.Sprintf("%d", year))
    }

    var extendedNumWordlist []string

    for _, word := range wordlist {
        for _, num := range numbers {
            extendedNumWordlist = append(extendedNumWordlist, word+num)
        }
    }
    wordlist = append(wordlist, extendedNumWordlist...)
    //remote memory extendedNumWordlist
    //free memory
    extendedNumWordlist = nil


    var extendedSpecNumWordlist []string

    for _, word := range wordlist {
        for _, char := range specialChars {
            extendedSpecNumWordlist = append(extendedSpecNumWordlist, word+char)
        }
    }
    wordlist = append(wordlist, extendedSpecNumWordlist...)
    //remote memory extendedSpecNumWordlist
    //free memory
    extendedSpecNumWordlist = nil


    if (profile.Complexity > 3){ 
        for _, word := range wordlist {
            if len(word) > 0 && ((word[0] >= 'A' && word[0] <= 'Z') || (word[0] >= 'a' && word[0] <= 'z')) {
                firstChar := rune(word[0])
                if firstChar >= 'A' && firstChar <= 'Z' {
                    word = strings.ToLower(string(firstChar)) + word[1:]
                } else if firstChar >= 'a' && firstChar <= 'z' {
                    word = strings.ToUpper(string(firstChar)) + word[1:]
                }
            }
            wordlist = append(wordlist, word)
        }
    }
    
    // Filter wordlist to ensure passwords meet length requirements
    var filteredWordlist []string
    minPasswordLength := profile.MinPasswordLength
    maxPasswordLength := profile.MaxPasswordLength

    wordMap := make(map[string]bool)

    for _, word := range wordlist {
        if len(word) >= minPasswordLength && len(word) <= maxPasswordLength {
            if _, exists := wordMap[word]; !exists {
                wordMap[word] = true
                filteredWordlist = append(filteredWordlist, word)
            }
        }
    }

    return filteredWordlist

}

func leetSpeakVariations(word string) []string {
    var variations []string
    leetMap := map[rune]rune{
        'a': '4',
        'e': '3',
        'i': '1',
        'o': '0',
        'l': '1',
        's': '5',
        't': '7',
        '1':'!',
        '2':'\'',
        '3':'^',
        '4':'+',
        '5':'%',
        '6':'&',
        '7':'/',
        '8':'(',
        '9':')',
        '0':'=',
    }

    runes := []rune(strings.ToLower(word))

    var indexes [][]interface{}
    for i, r := range runes {
        if leetVersion, ok := leetMap[r]; ok {
            indexes = append(indexes, []interface{}{i, r, leetVersion})
        }
    }

    totalVars := 1 << len(indexes)
    for i := 0; i < totalVars; i++ {
        tempWord := make([]rune, len(runes))
        copy(tempWord, runes)
        for j, index := range indexes {
            mask := 1 << j
            if i&mask != 0 {
                tempWord[index[0].(int)] = index[2].(rune) // değiştirilmiş karakter
            }
        }
        variations = append(variations, string(tempWord))
    }

    return variations
}

func SaveWordlist(wordlist []string, profile UserProfile) {
    file, err := os.Create("wordlist.txt")
    if err != nil {
        fmt.Println("Dosya oluşturulamadı:", err)
        return
    }
    defer file.Close()

    for _, word := range wordlist {
        _, err := file.WriteString(word + "\n")
        if err != nil {
            fmt.Println("Dosyaya yazılamadı:", err)
            return
        }
    }

    fmt.Printf("\nBACO Level: %d \n'wordlist.txt' dosyasına kaydedildi.\n", profile.Complexity)
}