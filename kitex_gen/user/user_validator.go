// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

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

func (p *DouyinUserRegisterRequest) IsValid() error {
	if len(p.Username) < int(2) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(6) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	_src := "[0-9A-Za-z]+"
	if ok, _ := regexp.MatchString(_src, p.Password); !ok {
		return fmt.Errorf("field Password pattern rule failed, current value: %v", p.Password)
	}
	return nil
}
func (p *DouyinUserRegisterResponse) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *DouyinUserLoginRequest) IsValid() error {
	if len(p.Username) < int(2) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(6) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *DouyinUserLoginResponse) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *DouyinUserRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *DouyinUserResponse) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *User) IsValid() error {
	if p.Id <= int64(0) {
		return fmt.Errorf("field Id gt rule failed, current value: %v", p.Id)
	}
	if p.FollowCount != nil {
		if *p.FollowCount <= int64(0) {
			return fmt.Errorf("field FollowCount gt rule failed, current value: %v", *p.FollowCount)
		}
	}
	if p.FollowerCount != nil {
		if *p.FollowerCount <= int64(0) {
			return fmt.Errorf("field FollowerCount gt rule failed, current value: %v", *p.FollowerCount)
		}
	}
	return nil
}