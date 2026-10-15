package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // URL of the file you want to download
    fileURL := "https://github.com/crc-org/machine-driver-libvirt/releases/download/0.13.6/crc-driver-libvirt"

    // Send an HTTP GET request to the server
    response, err := http.Get(fileURL)
    if err != nil {
        fmt.Printf("Error making GET request: %v\n", err)
        return
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        fmt.Printf("Error: Status code %d\n", response.StatusCode)
        return
    }

    // Create or open the local file where you want to save the downloaded data
    outputFile, err := os.Create()
    if err != nil {
        fmt.Printf("Error creating output file: %v\n", err)
        return
    }
    defer outputFile.Close()

    // Copy the response body (file content) to the local file
    _, err = io.Copy(outputFile, response.Body)
    if err != nil {
        fmt.Printf("Error copying data to file: %v\n", err)
        return
    }

    fmt.Println("File downloaded successfully.")
}
