package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/charmbracelet/lipgloss"
)

// Struct to hold the analysis result of a password
type Result struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

// Styles for terminal output
var (
	accent = lipgloss.Color("#c73e1d")

	title = lipgloss.NewStyle().Bold(true).Foreground(accent)
	dim   = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	label = dim.Width(12)

	green  = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	red    = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)

// analyze checks the password for various characteristics and returns a Result struct
func analyze(password string) Result {
	r := Result{Length: len([]rune(password))}

	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			r.Uppercase = true
		case unicode.IsLower(c):
			r.Lowercase = true
		case unicode.IsNumber(c):
			r.Numbers = true
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			r.Symbols = true
		}
	}
	return r
}

// tags
func tag(ok bool) string {
	if ok {
		return dim.Render("[") + green.Render("✓") + dim.Render("]")
	}
	return dim.Render("[") + red.Render("✗") + dim.Render("]")
}

// row formats a line of output with a label and its corresponding value
func row(name, value string) string {
	return "  " + label.Render(name) + value
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: hardpass <password>")
		os.Exit(2)
	}

	r := analyze(os.Args[1])
	longEnough := r.Length >= 12

	score := 0
	for _, ok := range []bool{r.Uppercase, r.Lowercase, r.Numbers, r.Symbols, longEnough} {
		if ok {
			score++
		}
	}

	style, verdict := red, "weak"
	if score >= 5 {
		style, verdict = green, "strong"
	} else if score >= 3 {
		style, verdict = yellow, "medium"
	}
	bar := style.Render(strings.Repeat("█", score)) + dim.Render(strings.Repeat("░", 5-score))

	line := dim.Render(strings.Repeat("─", 38))

	fmt.Println()
	fmt.Println("\n  " + title.Render("HardPass") + dim.Render("  password strength analyzer"))
	fmt.Println("  " + line)
	fmt.Println(row("length", tag(longEnough)+dim.Render(fmt.Sprintf("  %d chars (min 12)", r.Length))))
	fmt.Println(row("uppercase", tag(r.Uppercase)))
	fmt.Println(row("lowercase", tag(r.Lowercase)))
	fmt.Println(row("numbers", tag(r.Numbers)))
	fmt.Println(row("symbols", tag(r.Symbols)))
	fmt.Println("  " + line)
	fmt.Println(row("strength", bar+"  "+style.Render(verdict)))
	fmt.Println(" ")
}