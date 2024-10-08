// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Member represents a Trello member.
// https://developers.trello.com/reference/#member-object
type Member struct {
	client          *Client
	ID              string   `json:"id"`
	Username        string   `json:"username,omitempty"`
	FullName        string   `json:"fullName,omitempty"`
	Initials        string   `json:"initials,omitempty"`
	AvatarHash      string   `json:"avatarHash,omitempty"`
	AvatarURL       string   `json:"avatarUrl,omitempty"`
	Email           string   `json:"email,omitempty"`
	IDBoards        []string `json:"idBoards,omitempty"`
	IDOrganizations []string `json:"idOrganizations,omitempty"`
}

// GetMember takes a member id and Arguments and returns a Member or an error.
func (c *Client) GetMember(memberID string, extraArgs ...Arguments) (member *Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("members/%s", memberID)
	err = c.Get(path, args, &member)
	if err == nil {
		member.SetClient(c)
	}
	return
}

// GetMyMember returns Member for the user authenticating the API call
func (c *Client) GetMyMember(args Arguments) (member *Member, err error) {
	path := "members/me"
	err = c.Get(path, args, &member)
	if err == nil {
		member.SetClient(c)
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the organization or an error.
func (o *Organization) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("organizations/%s/members", o.ID)
	err = o.client.Get(path, args, &members)
	for i := range members {
		members[i].SetClient(o.client)
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Board or an error.
func (b *Board) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Get(path, args, &members)
	for i := range members {
		members[i].SetClient(b.client)
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Card or an error.
func (c *Card) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("cards/%s/members", c.ID)
	err = c.client.Get(path, args, &members)
	for i := range members {
		members[i].SetClient(c.client)
	}
	return
}

// SetClient can be used to override this Member's internal connection to the
// Trello API. Normally, this is set automatically after API calls.
func (m *Member) SetClient(newClient *Client) {
	m.client = newClient
}
