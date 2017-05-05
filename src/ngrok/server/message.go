package server

import (
	"encoding/json"
	"errors"
	"ngrok/conn"
	hasura "ngrok/hasura"
	"ngrok/log"
)

var ProjectsMessage = map[string]string{}

func fetchMessage(projectName string) (message string, err error) {
	var userData []NgrokUsers
	m := hasura.HasuraQuery{
		Type: "select",
		Args: hasura.HasuraArgument{
			Table:   "ngrok_users",
			Columns: []string{"message"},
			Where: map[string]string{
				"project_name": projectName,
			},
		},
	}
	_, body, err := hasura.SendQuery(m)
	if err != nil {
		log.Warn("Error in fetching all messages %s", err.Error())
		return
	}
	json.Unmarshal(body, &userData)
	if len(userData) == 0 {
		log.Info("No Projects Found")
		err = errors.New("No Projects Found")
		return
	}
	message = userData[0].Message
	return
}

func GetProjectMessage(c conn.Conn, ProjectName string) (message string, err error) {
	c.Info("Getting Message for Project %s", ProjectName)
	message, ok := ProjectsMessage[ProjectName]
	if !ok {
		c.Info("Project Not present internally")
		message, err = fetchMessage(ProjectName)
		if err != nil {
			return
		}
		ProjectsMessage[ProjectName] = message
		return
	}
	return
}

func fetchAllProjects() {
	var userData []NgrokUsers
	m := hasura.HasuraQuery{
		Type: "select",
		Args: hasura.HasuraArgument{
			Table:   "ngrok_users",
			Columns: []string{"message", "project_name"},
		},
	}
	_, body, err := hasura.SendQuery(m)
	if err != nil {
		log.Warn("Error in fetching all messages %s", err.Error())
		return
	}
	json.Unmarshal(body, &userData)
	if len(userData) == 0 {
		log.Warn("No Projects Found")
		return
	}
	for _, project := range userData {
		ProjectsMessage[project.ProjectName] = project.Message
	}
	log.Info("Messages Updated")
}
