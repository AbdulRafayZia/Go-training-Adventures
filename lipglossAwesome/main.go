package main
import ( 
	"github.com/charmbracelet/lipgloss"
	"fmt"
)

var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("5")).
    Background(lipgloss.Color("9")).
    PaddingTop(4).
    PaddingLeft(4).
    Width(22)
func main(){
	fmt.Println(style.Render("Hello, kitty"))
}
