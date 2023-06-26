package extensions

import "strings"

// ad *TAssignmentDefinition github.com/averyyan/bpmn-engine/bpmn/sepc/types.AssignmentDefinition
type TAssignmentDefinition struct {
	Assignee        string `xml:"assignee,attr" json:"assignee"`
	CandidateUsers  string `xml:"candidateUsers,attr" json:"candidate_users"`
	CandidateGroups string `xml:"candidateGroups,attr" json:"candidate_groups"`
}

// 获取用户ID
func (ad *TAssignmentDefinition) GetAssignee() string {
	return ad.Assignee
}

// 获取用户ID组
func (ad *TAssignmentDefinition) GetCandidateUsers() []string {
	users := strings.Split(ad.CandidateUsers, ",")
	for i, user := range users {
		users[i] = strings.TrimSpace(user)
	}
	return users
}

// 获取角色组
func (ad *TAssignmentDefinition) GetCandidateGroups() []string {
	groups := strings.Split(ad.CandidateGroups, ",")
	for i, group := range groups {
		groups[i] = strings.TrimSpace(group)
	}
	return groups
}
