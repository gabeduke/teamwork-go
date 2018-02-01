package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
	"github.com/swill/teamwork"
)

var (
	conn       *teamwork.Connection
	filter     string
	projects   bool
	suggestMap []prompt.Suggest
)

func setAssignedTasks() {
	conn = TeamworkConnection()
	userID := viper.GetString("global.userId")

	// get all tasks
	taskOps := &teamwork.GetTasksOps{
		ResponsiblePartyIDs: userID,
	}

	t, _, err := conn.GetTasks(taskOps)
	if err != nil {
		log.Fatal(err)
	}

	for index := range t {

		s := strconv.Itoa(t[index].TaskListID)
		taskSuggestion := prompt.Suggest{
			Text:        t[index].ProjectName + ": " + t[index].Content,
			Description: s,
		}

		tasks = t
		suggestMap = append(suggestMap, taskSuggestion)
	}
	taskSuggestions = suggestMap
}

// GetTasks ...
func GetTasks() []prompt.Suggest {
	return taskSuggestions
}

// InitializeTeamworkData ...
func InitializeTeamworkData() {
	setAssignedTasks()
}

// TeamworkConnection ...
func TeamworkConnection() *teamwork.Connection {
	apiToken := viper.GetString("global.apiKey")

	// setup the teamwork connection
	conn, err := teamwork.Connect(apiToken)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}
