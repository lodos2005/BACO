package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    profile := GetUserProfile()
    wordlist := GenerateWordlist(profile)
    SaveWordlist(wordlist,profile)
}

func GetUserProfile() UserProfile {
    reader := bufio.NewReader(os.Stdin)
    profile := UserProfile{}

    fmt.Println("=== BACO: Parola Listesi Oluşturucu ===")

    fmt.Print("Ad: ")
    profile.FirstName, _ = reader.ReadString('\n')

    fmt.Print("Soyad: ")
    profile.LastName, _ = reader.ReadString('\n')

    fmt.Print("Doğum Tarihi (DDMMYYYY): ")
    profile.BirthDate, _ = reader.ReadString('\n')

    fmt.Print("Takma Adlar (virgülle ayrılmış): ")
    nicknames, _ := reader.ReadString('\n')
    profile.Nicknames = strings.Split(strings.TrimSpace(nicknames), ",")

    fmt.Print("Sevdiği Kelimeler (virgülle ayrılmış): ")
    words, _ := reader.ReadString('\n')

    fmt.Print("Minimum parola uzunluğu: ")
    fmt.Fscanf(reader, "%d\n", &profile.MinPasswordLength)

    fmt.Print("Maximum parola uzunluğu: ")
    fmt.Fscanf(reader, "%d\n", &profile.MaxPasswordLength)

    fmt.Print("BACO Seviyesi (1-5): ")
    fmt.Fscanf(reader, "%d\n", &profile.Complexity)
    //check if profile.Complexity is >5
    if profile.Complexity > 5 {
        profile.Complexity = 5
    }
    //check if Complexity is zero
    if profile.Complexity == 0 {
        profile.Complexity = 1
    }
    
    //check if profile.MinPasswordLength is >0 also check if empty

    if profile.MinPasswordLength < 0 || profile.MinPasswordLength == 0 {
        profile.MinPasswordLength = 0
    }
    //check if profile.MaxPasswordLength is >32
    if profile.MaxPasswordLength > 32 || profile.MaxPasswordLength == 0  {
        profile.MaxPasswordLength = 32
    }
    //check if profile.MinPasswordLength is > profile.MaxPasswordLength
    if profile.MinPasswordLength > profile.MaxPasswordLength {
        profile.MinPasswordLength = profile.MaxPasswordLength
    }
    
    

    profile.FavoriteWords = strings.Split(strings.TrimSpace(words), ",")

    profile.FirstName = strings.TrimSpace(profile.FirstName)
    profile.LastName = strings.TrimSpace(profile.LastName)
    profile.BirthDate = strings.TrimSpace(profile.BirthDate)

    return profile
}