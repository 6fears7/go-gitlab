//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

// NamespacesService handles communication with the namespace related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/namespaces.html
type NamespacesService struct {
	client *Client
}

// Namespace represents a GitLab namespace.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/namespaces.html
type Namespace struct {
	ID                          int        `json:"id"`
	Name                        string     `json:"name"`
	Path                        string     `json:"path"`
	Kind                        string     `json:"kind"`
	FullPath                    string     `json:"full_path"`
	ParentID                    int        `json:"parent_id"`
	MembersCountWithDescendants int        `json:"members_count_with_descendants"`
	AvatarUrl                   *string    `json:"avatar_url"`
	Plan                        string     `json:"plan"`
	Trial                       bool       `json:"trial"`
	BillableMembersCount        int        `json:"billable_members_count"`
	WebUrl                      string     `json:"web_url"`
	TrialEndsOn                 *time.Time `json:"trial_ends_on"`
}

func (n Namespace) String() string {
	return Stringify(n)
}

// ListNamespacesOptions represents the available ListNamespaces() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/namespaces.html#list-namespaces
type ListNamespacesOptions struct {
	ListOptions
	Search    *string `url:"search,omitempty" json:"search,omitempty"`
	OwnedOnly *bool   `url:"owned_only,omitempty" json:"owned_only,omitempty"`
}

// ListNamespaces gets a list of projects accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/namespaces.html#list-namespaces
func (s *NamespacesService) ListNamespaces(opt *ListNamespacesOptions, options ...RequestOptionFunc) ([]*Namespace, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "namespaces", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n []*Namespace
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// SearchNamespace gets all namespaces that match your string in their name
// or path.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/namespaces.html#search-for-namespace
func (s *NamespacesService) SearchNamespace(query string, options ...RequestOptionFunc) ([]*Namespace, *Response, error) {
	var q struct {
		Search string `url:"search,omitempty" json:"search,omitempty"`
	}
	q.Search = query

	req, err := s.client.NewRequest(http.MethodGet, "namespaces", &q, options)
	if err != nil {
		return nil, nil, err
	}

	var n []*Namespace
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// GetNamespace gets a namespace by id.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/namespaces.html#get-namespace-by-id
func (s *NamespacesService) GetNamespace(id interface{}, options ...RequestOptionFunc) (*Namespace, *Response, error) {
	namespace, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("namespaces/%s", namespace)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Namespace)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}
