package sentry

import "fmt"

type Rule struct {
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
