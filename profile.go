package main

type UserProfile struct {
    Complexity int
    MinPasswordLength int
    MaxPasswordLength int
    FirstName     string
    LastName      string
    BirthDate     string
    Nicknames     []string
    FavoriteWords []string
}