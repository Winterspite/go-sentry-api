package sentry

import (
	"fmt"
	"time"
)

const (
	FirstSeenEventCondition      = "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition"
	FirstSeenEventConditionText  = "A new issue is created"
	RegressionEventCondition     = "sentry.rules.conditions.regression_event.RegressionEventCondition"
	RegressionEventConditionText = "The issue changes state from resolved to unresolved"
	NotifyEmailAction            = "sentry.mail.actions.NotifyEmailAction"
	NotifyEmailActionText        = "Send a notification to %s"
	TargetIssueOwners            = "IssueOwners"
	NotifySlackAction            = "sentry.integrations.slack.notify_action.SlackNotifyServiceAction"
)

type RuleCondition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RuleAction struct {
	ID               string `json:"id"`
	TargetType       string `json:"targetType,omitempty"`
	TargetIdentifier string `json:"targetIdentifier,omitempty"`
	Name             string `json:"name"`
	Workspace        string `json:"workspace,omitempty"`
	Channel          string `json:"channel,omitempty"`
	Tags             string `json:"tags,omitempty"`
	ChannelID        string `json:"channel_id,omitempty"`
}

type Rule struct {
	ID          string          `json:"id"`
	Conditions  []RuleCondition `json:"conditions"`
	Filters     []interface{}   `json:"filters"`
	Actions     []RuleAction    `json:"actions"`
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	DateCreated time.Time       `json:"dateCreated"`
	Owner       string          `json:"owner"`
	CreatedBy   interface{}     `json:"createdBy"`
	Environment string          `json:"environment"`
	Projects    []string        `json:"projects"`
}

// GetRules will fetch all alert rules for the specified organization and project.
func (c *Client) GetRules(o Organization, p Project) ([]Rule, error) {
	var rules []Rule

	err := c.do("GET", fmt.Sprintf("projects/%s/%s/rules", *o.Slug, *p.Slug), &rules, nil)

	return rules, err
}

func (c *Client) GetRuleByID(o Organization, p Project, id int) (*Rule, error) {
	var rule Rule

	err := c.do("GET", fmt.Sprintf("projects/%s/%s/rules/%d/", *o.Slug, *p.Slug, id), &rule, nil)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}
