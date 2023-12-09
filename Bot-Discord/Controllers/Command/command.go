package command

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
	"time"
	"encoding/json"
	
	"github.com/bwmarrin/discordgo"
	"github.com/Galangrs/Bot-Discord/Config"
	"github.com/Galangrs/Bot-Discord/Controllers/Fetch"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	re := regexp.MustCompile(`!([^|]+) ([^|]+) ([^|]+)`)

	matches := re.FindStringSubmatch(m.Content)
	if len(matches) == 0 {
		if strings.ToLower(m.Content) == "!ping" {
			// Get the current timestamp
			timestamp := m.Timestamp

			// Calculate the latency
			latency := time.Since(timestamp)

			// Respond with the latency
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! Latency: %s", latency.String()))
		} else if strings.ToLower(m.Content) == "!help" {
			helpMessage(s, m, config.Owner())
		} else if strings.ToLower(m.Content) == "!register"{
			// Hadle register case
			res, err := fetching.SendHTTPRequest("POST", config.URL()+"/register", []byte(`{"id": "`+m.Author.ID+`"}`))
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
			} else {
				// Assuming res is a JSON response like {"name": "id cannot be null"}
				var responseMap map[string]interface{}
				if err := json.Unmarshal(res, &responseMap); err != nil {
					fmt.Println("Error parsing JSON:", err)
					return
				}

				// Access the "name" fields from the response
				name, ok := responseMap["name"].(string)
				if !ok {
					fmt.Println("Error extracting name from JSON")
					return
				}

				// Build the response string
				resStr := fmt.Sprintf("%s", name)

				// Send the message to the Discord channel
				s.ChannelMessageSend(m.ChannelID, resStr)
			}	
		}
		return
	}

	// Convert case-insensitive comparison
	cmd := strings.ToLower(matches[1])

	switch cmd {
	case "buy":
		item := strings.ToLower(matches[2])
		value := strings.ToLower(matches[3])
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
					res, err := fetching.SendHTTPRequest("PUT", config.URL()+"/buycid", []byte(`{"id": "`+m.Author.ID+`","quantity":"`+value+`"}`))
					if err != nil {
						// Assuming res is a JSON response like {"name": "Internal Server Error"}
						var responseMap map[string]interface{}
						if err := json.Unmarshal([]byte(res), &responseMap); err != nil {
							fmt.Println("Error parsing JSON:", err)
							return
						}

						// Access the "name" field from the response
						name, ok := responseMap["name"].(string)
						if !ok {
							fmt.Println("Error extracting name from JSON")
							return
						}

						// Build the response string
						resStr := name
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", resStr))
					} else {
						// Assuming res is a JSON response like {"name": "uniqcode\nuniqcode"}
						var responseMap map[string]interface{}
						if err := json.Unmarshal(res, &responseMap); err != nil {
							fmt.Println("Error parsing JSON:", err)
							return
						}
		
						// Access the "name" fields from the response
						name, ok := responseMap["name"].(string)
						if !ok {
							fmt.Println("Error extracting name from JSON")
							return
						}
		
						// Build the response string
						resStr := fmt.Sprintf("%s", name)

						// Send the message to the Discord channel
						channel, err := s.UserChannelCreate(m.Author.ID)
						if err != nil {
							fmt.Println("Error creating channel:", err)
							s.ChannelMessageSend(
								m.ChannelID,
								"Something went wrong while sending the DM!",
							)
							return
						}
						// Create a File structure with the text content
						fileToSend := &discordgo.File{
							Name:   "file.txt",
							Reader: strings.NewReader(resStr),
						}

						// Send the file to the Discord channel
						s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
							Content: "CID items at here",
							Files:   []*discordgo.File{fileToSend},
						})
						s.ChannelMessageSend(m.ChannelID, "Succes buy item")
					}	
				}
			} else {
				// lowerValue is not a valid integer
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`Invalid input. Please use the format "!%s %s 1" and provide a valid numeric value.`, matches[1], matches[2]))
			}
		}
	case "addbal":
		// Handle addbal case
		idClient := strings.ToLower(matches[2])
		value := strings.ToLower(matches[3])
		addbalMessage(s, m, config.Owner(), idClient, value)
	case "delbal":
		// Handle delbal case
		idClient := strings.ToLower(matches[2])
		value := strings.ToLower(matches[3])
		delbalMessage(s, m, config.Owner(), idClient, value)
	case "help":
		// Handle help case
		helpMessage(s, m, config.Owner())
	default:
		// If cmd doesn't match any of the above cases
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Type 'help' for assistance.")
	}
}

func helpMessage(s *discordgo.Session, m *discordgo.MessageCreate, ownerID string) {
	if m.Author.ID == ownerID {
		s.ChannelMessageSend(m.ChannelID, `
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5

**Command: !addbal <id client> <value>**
- Add a balance to a client's inventory.
	Example: !addbal 1234567890 10

**Command: !delbal <id client> <value>**
- Use a balance from a client's inventory.
	Example: !delbal 1234567890 5

**Command: !register**
- register your account discord
	Example: !register

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Replace <id client> with the client ID, and <value> with the desired value when using !addbal or !delbal.
- Use "!addbal" to add a balance to a client's inventory. For example, "!addbal 1234567890 10" adds 10 to the balance for the client with ID 1234567890.
- Use "!delbal" to deduct a balance from a client's inventory. For example, "!delbal 1234567890 5" deducts 5 from the balance for the client with ID 1234567890.
- Use "!register" to register account at store
`)		
	} else {
		s.ChannelMessageSend(m.ChannelID, `
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5
	
**Command: !register**
- register your account discord
	Example: !register

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Use "!register" to register account at store
`)
	}
}

func addbalMessage(s *discordgo.Session, m *discordgo.MessageCreate, ownerID string, idClient string,value string) {
	if m.Author.ID == ownerID {
		res, err := fetching.SendHTTPRequest("PUT", config.URL()+"/addbal", []byte(`{"id": "`+idClient+`","value":"`+value+`"}`))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
		} else {
			// Assuming res is a JSON response like {"name": "uniqcode\nuniqcode"}
			var responseMap map[string]interface{}
			if err := json.Unmarshal(res, &responseMap); err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			// Access the "name" fields from the response
			name, ok := responseMap["name"].(string)
			if !ok {
				fmt.Println("Error extracting name from JSON")
				return
			}

			// Build the response string
			resStr := fmt.Sprintf("%s", name)

			// Send the message to the Discord channel
			s.ChannelMessageSend(m.ChannelID, resStr)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "you no have promise in command this")
	}
}

func delbalMessage(s *discordgo.Session, m *discordgo.MessageCreate, ownerID string, idClient string,value string) {
	if m.Author.ID == ownerID {
		res, err := fetching.SendHTTPRequest("PUT", config.URL()+"/delbal", []byte(`{"id": "`+idClient+`","value":"`+value+`"}`))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
		} else {
			// Assuming res is a JSON response like {"name": "uniqcode\nuniqcode"}
			var responseMap map[string]interface{}
			if err := json.Unmarshal(res, &responseMap); err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			// Access the "name" fields from the response
			name, ok := responseMap["name"].(string)
			if !ok {
				fmt.Println("Error extracting name from JSON")
				return
			}

			// Build the response string
			resStr := fmt.Sprintf("%s", name)

			// Send the message to the Discord channel
			s.ChannelMessageSend(m.ChannelID, resStr)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "you no have promise in command this")
	}
}