/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fakegithub

import (
	"fmt"

	"github.com/kubernetes/test-infra/ciongke/github"
)

type FakeClient struct {
	OrgMembers     []string
	IssueComments  map[int][]github.IssueComment
	IssueCommentID int
	PullRequests   map[int]*github.PullRequest
}

func (f *FakeClient) IsMember(org, user string) (bool, error) {
	for _, m := range f.OrgMembers {
		if m == user {
			return true, nil
		}
	}
	return false, nil
}

func (f *FakeClient) ListIssueComments(owner, repo string, number int) ([]github.IssueComment, error) {
	return append([]github.IssueComment{}, f.IssueComments[number]...), nil
}

func (f *FakeClient) CreateComment(owner, repo string, number int, comment string) error {
	f.IssueComments[number] = append(f.IssueComments[number], github.IssueComment{
		ID:   f.IssueCommentID,
		Body: comment,
	})
	f.IssueCommentID++
	return nil
}

func (f *FakeClient) DeleteComment(owner, repo string, ID int) error {
	for num, ics := range f.IssueComments {
		for i, ic := range ics {
			if ic.ID == ID {
				f.IssueComments[num] = append(ics[:i], ics[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("could not find issue comment %d", ID)
}

func (f *FakeClient) GetPullRequest(owner, repo string, number int) (*github.PullRequest, error) {
	return f.PullRequests[number], nil
}

func (f *FakeClient) CreateStatus(owner, repo, ref string, s github.Status) error {
	return nil
}
