package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name      string
		params    *User
		returns   []any
		assertion func(int64, error)
	}{
		{
			name: "when given valid payload, returns no error",
			params: &User{
				FullName: "user test",
				Phone:    "+62812256112",
				Password: "dummy password",
			},
			returns: []any{int64(1), nil},
			assertion: func(id int64, err error) {
				gomock.Eq(int64(1)).Matches(id)
				gomock.Nil().Matches(err)
			},
		},
		{
			name: "when database returns an error, returns the error",
			params: &User{
				FullName: "user test",
				Phone:    "+62812256112",
				Password: "dummy password",
			},
			returns: []interface{}{int64(0), errors.New("database error")},
			assertion: func(id int64, err error) {
				gomock.Eq(int64(0)).Matches(id)
				gomock.Eq("database error").Matches(err.Error())
			},
		},
		{
			name: "when No FullName, returns validation error",
			params: &User{
				Phone:    "+62812256112",
				Password: "dummy password",
			},
			returns: []interface{}{int64(0), errors.New("full_name invalid")},
			assertion: func(id int64, err error) {
				gomock.Eq(int64(0)).Matches(id)
				gomock.Eq("full_name invalid").Matches(err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateUser(ctx, tt.params).Return(tt.returns...)

			id, err := mockRepo.CreateUser(ctx, tt.params)
			tt.assertion(id, err)
		})
	}
}

func TestFindUserByPhoneAndCountryCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name      string
		params    []string
		returns   []any
		assertion func(*User, error)
	}{
		{
			name:   "when given correct payload, returns correct user",
			params: []string{"813234911", "+62"},
			returns: []any{&User{
				ID:          int64(1),
				Phone:       "813234911",
				CountryCode: "+62",
				FullName:    "Dolor Ipsum",
			}, nil},
			assertion: func(u *User, err error) {
				gomock.Eq("813234911").Matches(u.Phone)
				gomock.Eq("Dolor Ipsum").Matches(u.FullName)
				gomock.Nil().Matches(err)
			},
		},
		{
			name:    "when user not found, returns nil",
			params:  []string{"999999999", "+62"},
			returns: []interface{}{nil, nil},
			assertion: func(u *User, err error) {
				gomock.Nil().Matches(u)
				gomock.Nil().Matches(err)
			},
		},
		{
			name:    "when database error occurs, returns error",
			params:  []string{"813234911", "+62"},
			returns: []interface{}{nil, errors.New("database error")},
			assertion: func(u *User, err error) {
				gomock.Nil().Matches(u)
				gomock.Not(gomock.Nil().Matches(err))
				gomock.Eq("database error").Matches(err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FindUserByPhoneAndCountryCode(ctx, tt.params[0], tt.params[1]).Return(tt.returns...)

			u, err := mockRepo.FindUserByPhoneAndCountryCode(ctx, tt.params[0], tt.params[1])
			tt.assertion(u, err)
		})
	}
}

func TestFindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name      string
		params    int64
		returns   []any
		assertion func(*User, error)
	}{
		{
			name:   "when given correct payload, returns correct user",
			params: int64(2),
			returns: []any{&User{
				ID:          int64(2),
				Phone:       "813234911",
				CountryCode: "+62",
				FullName:    "Dolor Ipsum",
			}, nil},
			assertion: func(u *User, err error) {
				gomock.Eq(int64(1)).Matches(u.ID)
				gomock.Eq("813234911").Matches(u.Phone)
				gomock.Eq("Dolor Ipsum").Matches(u.FullName)
				gomock.Nil().Matches(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FindUserByID(ctx, tt.params).Return(tt.returns...)

			u, err := mockRepo.FindUserByID(ctx, tt.params)
			tt.assertion(u, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name      string
		old       *User
		phone     string
		fullName  string
		returns   error
		assertion func(*User, error)
	}{
		{
			name: "when given correct payload, update user without error",
			old: &User{
				ID:       int64(1),
				Phone:    "8123202",
				FullName: "Dolor Ipsum",
			},
			phone:    "821239122",
			fullName: "New Name Dolor Ipusum",
			returns:  nil,
			assertion: func(u *User, err error) {
				gomock.Eq(int64(1)).Matches(u.ID)
				gomock.Eq("821239122").Matches(u.Phone)
				gomock.Eq("New Name Dolor Ipusum").Matches(u.FullName)
				gomock.Nil().Matches(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().UpdateUser(ctx, tt.old, tt.fullName, tt.phone)

			err := mockRepo.UpdateUser(ctx, tt.old, tt.fullName, tt.phone)
			tt.assertion(tt.old, err)
		})
	}
}

func TestCreateUserLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name      string
		param     *User
		returns   error
		assertion func(error)
	}{
		{
			name: "when given correct payload, update user without error",
			param: &User{
				ID:       int64(1),
				Phone:    "8123202",
				FullName: "Dolor Ipsum",
			},
			returns: nil,
			assertion: func(err error) {
				gomock.Nil().Matches(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateUserLog(ctx, tt.param).Return(tt.returns)

			err := mockRepo.CreateUserLog(ctx, tt.param)
			tt.assertion(err)
		})
	}
}
