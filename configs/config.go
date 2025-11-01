package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type botConfig struct {
	Token              string
	OwnerID            string
	OpenCategoryName   string
	ClosedCategoryName string
}

type commandConfig struct {
	DeveloperOnlyGuildID string
}

// MFConfig for Muffin Framework
type MFConfig struct {
	Bot     botConfig
	Command commandConfig
}

var instance *MFConfig

// GetConfig gets MFConfig instance
func GetConfig() *MFConfig {
	if instance == nil {
		godotenv.Load()
		instance = &MFConfig{}
		setConfig(instance)
	}

	return instance
}

func getRequiredValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("[MF] required '%s' value not found in .env", key))
	}

	return value
}

func getValue(key string) string {
	return os.Getenv(key)
}

func setConfig(config *MFConfig) {
	config.Bot = botConfig{
		Token:              getRequiredValue("BOT_TOKEN"),
		OwnerID:            getRequiredValue("BOT_OWNER_ID"),
		OpenCategoryName:   getRequiredValue("BOT_OPEN_CATEGORY_NAME"),
		ClosedCategoryName: getRequiredValue("BOT_CLOSED_CATEGORY_NAME"),
	}

	config.Command = commandConfig{
		DeveloperOnlyGuildID: getValue("COMMAND_DEVELOPER_ONLY_GUILD_ID"),
	}
}
