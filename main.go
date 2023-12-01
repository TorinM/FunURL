package main

import (
    "fmt"
    "flag"
)

func printHelpMsg() {
    fmt.Println("Usage: ")
    fmt.Println("  -url <url>   URL to shorten")
}

func main() {
    url := flag.String("url", "n/a", "URL to shorten")
    is_help := flag.Bool("help", false, "Show help")
    flag.Parse()

    if *is_help {
        printHelpMsg()
    }
    
    if *url == "n/a" {
        fmt.Println("URL is required, use -url <url>")
        return
    } else {
        fmt.Println("URL: ", *url)
    }
}

