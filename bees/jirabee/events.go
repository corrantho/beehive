/*
 *    Copyright (C) 2021 Anthony Corrado
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
 *      Anthony Corrado <anthony@synetz.fr>
 */

// Package jirabee is a Bee that can interface with Jira
package jirabee

import (
	"encoding/json"
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/muesli/beehive/bees"
)

// JiraEvent represents a Jira Event
type JiraEvent struct {
	WebhookEvent string      `json:"webhookEvent"`
	Issue        *jira.Issue `json:"issue"`
}

func (mod *JiraBee) handleJiraEvent(data []byte) (*JiraEvent, error) {
	jiraEvent := &JiraEvent{}
	err := json.Unmarshal(data, &jiraEvent)
	if err != nil {
		return nil, fmt.Errorf("Error during JiraEvent Unmarshal: %v", err)
	}

	switch jiraEvent.WebhookEvent {

	case "jira:issue_created":
		mod.handleIssueCreatedEvent(jiraEvent)

	default:
		return jiraEvent, fmt.Errorf("Unhandled event: %s", jiraEvent.WebhookEvent)
	}

	return jiraEvent, nil
}

func (mod *JiraBee) handleIssueCreatedEvent(data *JiraEvent) error {
	ev := bees.Event{
		Bee: mod.Name(),
		Options: []bees.Placeholder{
			{
				Name:  "remote_addr",
				Type:  "address",
				Value: nil,
			},
			{
				Name:  "url",
				Type:  "url",
				Value: nil,
			},
		},
	}

	mod.eventChan <- ev
	return nil
}
