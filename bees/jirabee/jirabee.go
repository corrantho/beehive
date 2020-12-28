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

		mod.handleCreateIssueAction(project, reporterEmail, assigneeEmail, issueType, issueSummary, issueDescription)

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

func (mod *JiraBee) handleCreateIssueAction(project string, reporterEmail string, assigneeEmail string, issueType string, issueSummary string, issueDescription string) {
	fmt.Println("handleCreateIssue has been called")
	// If reporterEmail is not empty, we search for the AccountID of the user

	// If assigneeEmail is not empty, we search for the AccountID of the user

	// Create issue

}
