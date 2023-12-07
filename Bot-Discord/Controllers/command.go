package command

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	re := regexp.MustCompile(`!([^|]+) ([^|]+) ([^|]+)`)

	matches := re.FindStringSubmatch(m.Content)
	if len(matches) == 0 {
		if strings.ToLower(m.Content) == "!ping"{
			// Get the current timestamp
			timestamp := m.Timestamp

			// Calculate the latency
			latency := time.Since(timestamp)

			// Respond with the latency
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! Latency: %s", latency.String()))
		}
		return
	}

	// Convert case-insensitive comparison
	cmd := strings.ToLower(matches[1])
	item := strings.ToLower(matches[2])
	value := strings.ToLower(matches[3])

	switch cmd {
	case "buy":
		if item == "cid" {
			// Handle buy cid case
			// Attempt to convert value to an integer
			number, err := strconv.Atoi(value)

			// Check if conversion was successful
			if err == nil {
				// value is a valid integer
				if number < 1 {
					s.ChannelMessageSend(m.ChannelID, "Invalid quantity. Please provide a quantity greater than or equal to 1.")
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Are you sure you want to buy %d?", number))
				}
			} else {
				// lowerValue is not a valid integer
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`Invalid input. Please use the format "!%s %s 1" and provide a valid numeric value.`, matches[1], matches[2]))
			}
		}
	case "addball":
		// Handle addball case
	case "useball":
		// Handle useball case
	case "help":
		// Handle help case
	case "buyball":
		// Handle buyball case
	default:
		// If cmd doesn't match any of the above cases
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Type 'help' for assistance.")
	}
}
