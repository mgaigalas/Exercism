package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	newLine := "\n"
	var sb strings.Builder
	sb.WriteString(border)
	sb.WriteString(newLine)
	sb.WriteString(welcomeMsg)
	sb.WriteString(newLine)
	sb.WriteString(border)
	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	str := strings.ReplaceAll(oldMsg, "*", "")
	str = strings.ReplaceAll(str, "\n", "")
	return strings.TrimSpace(str)
}
