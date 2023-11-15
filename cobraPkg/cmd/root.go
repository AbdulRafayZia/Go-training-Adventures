package cmd
import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/spf13/cobra"
    "github.com/spf13/hello/pkg"
)
var rootCmd = &cobra.Command{
    Use:   "hello",
    Short: "A brief description of your application",
}
func Execute() error {
    return rootCmd.Execute()
}
func init() {
    rootCmd.AddCommand(fileCmd)
    fileCmd.Flags().IntP("routines", "r", 1, "Number of routines")
    fileCmd.Flags().String("filepath", "file1.txt", "File to read.")
    // Add other commands as needed
}
var fileCmd = &cobra.Command{
    Use:   "process_file",
    Short: "To process a file and create the summary.",
    Run: func(cmd *cobra.Command, args []string) {
        routines, _ := cmd.Flags().GetInt("routines")
        filePath, _ := cmd.Flags().GetString("filepath")
        fmt.Println("routines are ", routines)
        processFile(filePath, routines)
    },
}
func processFile(filePath string, routines int) {
    channal := make(chan pkg.Summary)
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fileData := string(content)
    chunk := len(fileData) / routines
    startIndex := 0
    endIndex := chunk
    for iterations := 0; iterations < routines; iterations++ {
        go pkg.Counts(fileData[startIndex:endIndex], channal)
        startIndex = endIndex
        endIndex += chunk
    }
    for iterations := 0; iterations < routines; iterations++ {
        counts := <-channal
        fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
        fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
        fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
        fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
    }
}