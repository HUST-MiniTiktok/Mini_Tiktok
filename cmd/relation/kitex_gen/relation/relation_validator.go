// Code generated by Validator v0.1.4. DO NOT EDIT.

package relation

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *User) IsValid() error {
	return nil
}
func (p *FriendUser) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("field User not valid, %w", err)
		}
	}
	return nil
}
func (p *RelationActionRequest) IsValid() error {
	return nil
}
func (p *RelationActionResponse) IsValid() error {
	return nil
}
func (p *RelationFollowListRequest) IsValid() error {
	return nil
}
func (p *RelationFollowListResponse) IsValid() error {
	return nil
}
func (p *RelationFollowerListRequest) IsValid() error {
	return nil
}
func (p *RelationFollowerListResponse) IsValid() error {
	return nil
}
func (p *RelationFriendListRequest) IsValid() error {
	return nil
}
func (p *RelationFriendListResponse) IsValid() error {
	return nil
}
