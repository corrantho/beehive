/*
 *    Copyright (C) 2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Nicolas Martin <penguwingithub@gmail.com>
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// Package jirabee is a Bee that can interface with Jira
package jirabee

import (
	"bytes"
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/muesli/beehive/bees"
)

// JiraBee is a Bee that can interface with Jira
type JiraBee struct {
	bees.Bee

	eventChan chan bees.Event
	client    *jira.Client

	url      string
	username string
	password string
}

// Action triggers the actions passed to it.
func (mod *JiraBee) Action(action bees.Action) []bees.Placeholder {
	outs := []bees.Placeholder{}

	switch action.Name {

	case "create_issue":
		var project string
		var reporterEmail string
		var assigneeEmail string
		var issueType string
		var issueSummary string
		var issueDescription string

		action.Options.Bind("project", &project)
		action.Options.Bind("reporter_email", &reporterEmail)
		action.Options.Bind("assignee_email", &assigneeEmail)
		action.Options.Bind("issue_type", &issueType)
		action.Options.Bind("issue_summary", &issueSummary)
		action.Options.Bind("issue_description", &issueDescription)

		issueKey, err := mod.handleCreateIssueAction(project, reporterEmail, assigneeEmail, issueType, issueSummary, issueDescription)
		if err != nil {
			mod.LogErrorf("Error during handleCreateIssueAction: %v", err)
		} else {
			mod.Logf("Issue created: %s", *issueKey)
		}

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// Run executes the Bee's event loop.
func (mod *JiraBee) Run(eventChan chan bees.Event) {

	tp := jira.BasicAuthTransport{
		Username: mod.username,
		Password: mod.password,
	}

	var err error

	mod.client, err = jira.NewClient(tp.Client(), mod.url)
	if err != nil {
		mod.LogErrorf("Failed to create JIRA client: %v", err)
	}

}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *JiraBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)

	options.Bind("url", &mod.url)
	options.Bind("username", &mod.username)
	options.Bind("password", &mod.password)
}

func (mod *JiraBee) handleCreateIssueAction(project string, reporterEmail string, assigneeEmail string, issueType string, issueSummary string, issueDescription string) (*string, error) {

	// Create issue
	i := jira.Issue{
		Fields: &jira.IssueFields{
			Description: issueDescription,
			Type: jira.IssueType{
				Name: issueType,
			},
			Project: jira.Project{
				Key: project,
			},
			Summary: issueSummary,
		},
	}

	// If reporterEmail is not empty, we search for the AccountID of the user
	if len(reporterEmail) > 0 {
		reporterUser, err := mod.getJiraUser(reporterEmail)
		if err != nil {
			return nil, fmt.Errorf("Error when trying to get reporter user: %v", err)
		}

		i.Fields.Reporter = &jira.User{
			AccountID: reporterUser.AccountID,
		}
	}

	// If assigneeEmail is not empty, we search for the AccountID of the user
	if len(assigneeEmail) > 0 {
		assigneeUser, err := mod.getJiraUser(assigneeEmail)
		if err != nil {
			return nil, fmt.Errorf("Error when trying to get assignee user: %v", err)
		}

		i.Fields.Assignee = &jira.User{
			AccountID: assigneeUser.AccountID,
		}
	}

	// Call Issue service
	issueCreated, jiraResponse, err := mod.client.Issue.Create(&i)
	if err != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(jiraResponse.Body)
		jiraResponseBody := buf.String()

		return nil, fmt.Errorf("Error when trying to create an issue: \n%v\n%v", err, jiraResponseBody)
	}

	return &issueCreated.Key, nil
}

func (mod *JiraBee) getJiraUser(email string) (*jira.User, error) {
	if len(email) > 0 {
		usersFound, _, err := mod.client.User.Find(email)

		if err != nil {
			return nil, err
		}

		if len(usersFound) != 1 {
			return nil, fmt.Errorf("Zero or more than one user found with email address %s", email)
		}

		return &usersFound[0], nil
	}
	return nil, nil
}
