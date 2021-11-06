// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/comment"
	"artsign/ent/predicate"
	"artsign/ent/user"
	"artsign/ent/work"
	"fmt"
	"time"
)

// CategoryWhereInput represents a where input for filtering Category queries.
type CategoryWhereInput struct {
	Not *CategoryWhereInput   `json:"not,omitempty"`
	Or  []*CategoryWhereInput `json:"or,omitempty"`
	And []*CategoryWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "works" edge predicates.
	HasWorks     *bool             `json:"hasWorks,omitempty"`
	HasWorksWith []*WorkWhereInput `json:"hasWorksWith,omitempty"`
}

// Filter applies the CategoryWhereInput filter on the CategoryQuery builder.
func (i *CategoryWhereInput) Filter(q *CategoryQuery) (*CategoryQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering categories.
// An error is returned if the input is empty or invalid.
func (i *CategoryWhereInput) P() (predicate.Category, error) {
	var predicates []predicate.Category
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, category.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Category, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, category.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Category, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, category.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, category.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, category.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, category.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, category.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, category.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, category.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, category.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, category.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, category.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, category.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, category.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, category.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, category.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, category.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, category.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, category.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, category.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, category.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, category.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, category.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, category.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasWorks != nil {
		p := category.HasWorks()
		if !*i.HasWorks {
			p = category.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasWorksWith) > 0 {
		with := make([]predicate.Work, 0, len(i.HasWorksWith))
		for _, w := range i.HasWorksWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, category.HasWorksWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("artsign/ent: empty predicate CategoryWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return category.And(predicates...), nil
	}
}

// CommentWhereInput represents a where input for filtering Comment queries.
type CommentWhereInput struct {
	Not *CommentWhereInput   `json:"not,omitempty"`
	Or  []*CommentWhereInput `json:"or,omitempty"`
	And []*CommentWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "create_time" field predicates.
	CreateTime      *time.Time  `json:"createTime,omitempty"`
	CreateTimeNEQ   *time.Time  `json:"createTimeNEQ,omitempty"`
	CreateTimeIn    []time.Time `json:"createTimeIn,omitempty"`
	CreateTimeNotIn []time.Time `json:"createTimeNotIn,omitempty"`
	CreateTimeGT    *time.Time  `json:"createTimeGT,omitempty"`
	CreateTimeGTE   *time.Time  `json:"createTimeGTE,omitempty"`
	CreateTimeLT    *time.Time  `json:"createTimeLT,omitempty"`
	CreateTimeLTE   *time.Time  `json:"createTimeLTE,omitempty"`

	// "update_time" field predicates.
	UpdateTime      *time.Time  `json:"updateTime,omitempty"`
	UpdateTimeNEQ   *time.Time  `json:"updateTimeNEQ,omitempty"`
	UpdateTimeIn    []time.Time `json:"updateTimeIn,omitempty"`
	UpdateTimeNotIn []time.Time `json:"updateTimeNotIn,omitempty"`
	UpdateTimeGT    *time.Time  `json:"updateTimeGT,omitempty"`
	UpdateTimeGTE   *time.Time  `json:"updateTimeGTE,omitempty"`
	UpdateTimeLT    *time.Time  `json:"updateTimeLT,omitempty"`
	UpdateTimeLTE   *time.Time  `json:"updateTimeLTE,omitempty"`

	// "content" field predicates.
	Content             *string  `json:"content,omitempty"`
	ContentNEQ          *string  `json:"contentNEQ,omitempty"`
	ContentIn           []string `json:"contentIn,omitempty"`
	ContentNotIn        []string `json:"contentNotIn,omitempty"`
	ContentGT           *string  `json:"contentGT,omitempty"`
	ContentGTE          *string  `json:"contentGTE,omitempty"`
	ContentLT           *string  `json:"contentLT,omitempty"`
	ContentLTE          *string  `json:"contentLTE,omitempty"`
	ContentContains     *string  `json:"contentContains,omitempty"`
	ContentHasPrefix    *string  `json:"contentHasPrefix,omitempty"`
	ContentHasSuffix    *string  `json:"contentHasSuffix,omitempty"`
	ContentEqualFold    *string  `json:"contentEqualFold,omitempty"`
	ContentContainsFold *string  `json:"contentContainsFold,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`

	// "work" edge predicates.
	HasWork     *bool             `json:"hasWork,omitempty"`
	HasWorkWith []*WorkWhereInput `json:"hasWorkWith,omitempty"`
}

// Filter applies the CommentWhereInput filter on the CommentQuery builder.
func (i *CommentWhereInput) Filter(q *CommentQuery) (*CommentQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering comments.
// An error is returned if the input is empty or invalid.
func (i *CommentWhereInput) P() (predicate.Comment, error) {
	var predicates []predicate.Comment
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, comment.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Comment, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, comment.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Comment, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, comment.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, comment.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, comment.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, comment.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, comment.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, comment.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, comment.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, comment.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, comment.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, comment.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, comment.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, comment.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, comment.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, comment.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, comment.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, comment.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, comment.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, comment.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, comment.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, comment.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, comment.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, comment.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, comment.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, comment.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, comment.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.Content != nil {
		predicates = append(predicates, comment.ContentEQ(*i.Content))
	}
	if i.ContentNEQ != nil {
		predicates = append(predicates, comment.ContentNEQ(*i.ContentNEQ))
	}
	if len(i.ContentIn) > 0 {
		predicates = append(predicates, comment.ContentIn(i.ContentIn...))
	}
	if len(i.ContentNotIn) > 0 {
		predicates = append(predicates, comment.ContentNotIn(i.ContentNotIn...))
	}
	if i.ContentGT != nil {
		predicates = append(predicates, comment.ContentGT(*i.ContentGT))
	}
	if i.ContentGTE != nil {
		predicates = append(predicates, comment.ContentGTE(*i.ContentGTE))
	}
	if i.ContentLT != nil {
		predicates = append(predicates, comment.ContentLT(*i.ContentLT))
	}
	if i.ContentLTE != nil {
		predicates = append(predicates, comment.ContentLTE(*i.ContentLTE))
	}
	if i.ContentContains != nil {
		predicates = append(predicates, comment.ContentContains(*i.ContentContains))
	}
	if i.ContentHasPrefix != nil {
		predicates = append(predicates, comment.ContentHasPrefix(*i.ContentHasPrefix))
	}
	if i.ContentHasSuffix != nil {
		predicates = append(predicates, comment.ContentHasSuffix(*i.ContentHasSuffix))
	}
	if i.ContentEqualFold != nil {
		predicates = append(predicates, comment.ContentEqualFold(*i.ContentEqualFold))
	}
	if i.ContentContainsFold != nil {
		predicates = append(predicates, comment.ContentContainsFold(*i.ContentContainsFold))
	}

	if i.HasOwner != nil {
		p := comment.HasOwner()
		if !*i.HasOwner {
			p = comment.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, comment.HasOwnerWith(with...))
	}
	if i.HasWork != nil {
		p := comment.HasWork()
		if !*i.HasWork {
			p = comment.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasWorkWith) > 0 {
		with := make([]predicate.Work, 0, len(i.HasWorkWith))
		for _, w := range i.HasWorkWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, comment.HasWorkWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("artsign/ent: empty predicate CommentWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return comment.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "profile" field predicates.
	Profile             *string  `json:"profile,omitempty"`
	ProfileNEQ          *string  `json:"profileNEQ,omitempty"`
	ProfileIn           []string `json:"profileIn,omitempty"`
	ProfileNotIn        []string `json:"profileNotIn,omitempty"`
	ProfileGT           *string  `json:"profileGT,omitempty"`
	ProfileGTE          *string  `json:"profileGTE,omitempty"`
	ProfileLT           *string  `json:"profileLT,omitempty"`
	ProfileLTE          *string  `json:"profileLTE,omitempty"`
	ProfileContains     *string  `json:"profileContains,omitempty"`
	ProfileHasPrefix    *string  `json:"profileHasPrefix,omitempty"`
	ProfileHasSuffix    *string  `json:"profileHasSuffix,omitempty"`
	ProfileEqualFold    *string  `json:"profileEqualFold,omitempty"`
	ProfileContainsFold *string  `json:"profileContainsFold,omitempty"`

	// "works" edge predicates.
	HasWorks     *bool             `json:"hasWorks,omitempty"`
	HasWorksWith []*WorkWhereInput `json:"hasWorksWith,omitempty"`

	// "likes" edge predicates.
	HasLikes     *bool             `json:"hasLikes,omitempty"`
	HasLikesWith []*WorkWhereInput `json:"hasLikesWith,omitempty"`

	// "comments" edge predicates.
	HasComments     *bool                `json:"hasComments,omitempty"`
	HasCommentsWith []*CommentWhereInput `json:"hasCommentsWith,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}
	if i.Profile != nil {
		predicates = append(predicates, user.ProfileEQ(*i.Profile))
	}
	if i.ProfileNEQ != nil {
		predicates = append(predicates, user.ProfileNEQ(*i.ProfileNEQ))
	}
	if len(i.ProfileIn) > 0 {
		predicates = append(predicates, user.ProfileIn(i.ProfileIn...))
	}
	if len(i.ProfileNotIn) > 0 {
		predicates = append(predicates, user.ProfileNotIn(i.ProfileNotIn...))
	}
	if i.ProfileGT != nil {
		predicates = append(predicates, user.ProfileGT(*i.ProfileGT))
	}
	if i.ProfileGTE != nil {
		predicates = append(predicates, user.ProfileGTE(*i.ProfileGTE))
	}
	if i.ProfileLT != nil {
		predicates = append(predicates, user.ProfileLT(*i.ProfileLT))
	}
	if i.ProfileLTE != nil {
		predicates = append(predicates, user.ProfileLTE(*i.ProfileLTE))
	}
	if i.ProfileContains != nil {
		predicates = append(predicates, user.ProfileContains(*i.ProfileContains))
	}
	if i.ProfileHasPrefix != nil {
		predicates = append(predicates, user.ProfileHasPrefix(*i.ProfileHasPrefix))
	}
	if i.ProfileHasSuffix != nil {
		predicates = append(predicates, user.ProfileHasSuffix(*i.ProfileHasSuffix))
	}
	if i.ProfileEqualFold != nil {
		predicates = append(predicates, user.ProfileEqualFold(*i.ProfileEqualFold))
	}
	if i.ProfileContainsFold != nil {
		predicates = append(predicates, user.ProfileContainsFold(*i.ProfileContainsFold))
	}

	if i.HasWorks != nil {
		p := user.HasWorks()
		if !*i.HasWorks {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasWorksWith) > 0 {
		with := make([]predicate.Work, 0, len(i.HasWorksWith))
		for _, w := range i.HasWorksWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasWorksWith(with...))
	}
	if i.HasLikes != nil {
		p := user.HasLikes()
		if !*i.HasLikes {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasLikesWith) > 0 {
		with := make([]predicate.Work, 0, len(i.HasLikesWith))
		for _, w := range i.HasLikesWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasLikesWith(with...))
	}
	if i.HasComments != nil {
		p := user.HasComments()
		if !*i.HasComments {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCommentsWith) > 0 {
		with := make([]predicate.Comment, 0, len(i.HasCommentsWith))
		for _, w := range i.HasCommentsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasCommentsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("artsign/ent: empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// WorkWhereInput represents a where input for filtering Work queries.
type WorkWhereInput struct {
	Not *WorkWhereInput   `json:"not,omitempty"`
	Or  []*WorkWhereInput `json:"or,omitempty"`
	And []*WorkWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "title" field predicates.
	Title             *string  `json:"title,omitempty"`
	TitleNEQ          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGT           *string  `json:"titleGT,omitempty"`
	TitleGTE          *string  `json:"titleGTE,omitempty"`
	TitleLT           *string  `json:"titleLT,omitempty"`
	TitleLTE          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "image_url" field predicates.
	ImageURL             *string  `json:"imageURL,omitempty"`
	ImageURLNEQ          *string  `json:"imageURLNEQ,omitempty"`
	ImageURLIn           []string `json:"imageURLIn,omitempty"`
	ImageURLNotIn        []string `json:"imageURLNotIn,omitempty"`
	ImageURLGT           *string  `json:"imageURLGT,omitempty"`
	ImageURLGTE          *string  `json:"imageURLGTE,omitempty"`
	ImageURLLT           *string  `json:"imageURLLT,omitempty"`
	ImageURLLTE          *string  `json:"imageURLLTE,omitempty"`
	ImageURLContains     *string  `json:"imageURLContains,omitempty"`
	ImageURLHasPrefix    *string  `json:"imageURLHasPrefix,omitempty"`
	ImageURLHasSuffix    *string  `json:"imageURLHasSuffix,omitempty"`
	ImageURLEqualFold    *string  `json:"imageURLEqualFold,omitempty"`
	ImageURLContainsFold *string  `json:"imageURLContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "category" edge predicates.
	HasCategory     *bool                 `json:"hasCategory,omitempty"`
	HasCategoryWith []*CategoryWhereInput `json:"hasCategoryWith,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`

	// "likers" edge predicates.
	HasLikers     *bool             `json:"hasLikers,omitempty"`
	HasLikersWith []*UserWhereInput `json:"hasLikersWith,omitempty"`

	// "comments" edge predicates.
	HasComments     *bool                `json:"hasComments,omitempty"`
	HasCommentsWith []*CommentWhereInput `json:"hasCommentsWith,omitempty"`
}

// Filter applies the WorkWhereInput filter on the WorkQuery builder.
func (i *WorkWhereInput) Filter(q *WorkQuery) (*WorkQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering works.
// An error is returned if the input is empty or invalid.
func (i *WorkWhereInput) P() (predicate.Work, error) {
	var predicates []predicate.Work
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, work.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Work, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, work.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Work, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, work.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, work.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, work.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, work.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, work.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, work.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, work.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, work.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, work.IDLTE(*i.IDLTE))
	}
	if i.Title != nil {
		predicates = append(predicates, work.TitleEQ(*i.Title))
	}
	if i.TitleNEQ != nil {
		predicates = append(predicates, work.TitleNEQ(*i.TitleNEQ))
	}
	if len(i.TitleIn) > 0 {
		predicates = append(predicates, work.TitleIn(i.TitleIn...))
	}
	if len(i.TitleNotIn) > 0 {
		predicates = append(predicates, work.TitleNotIn(i.TitleNotIn...))
	}
	if i.TitleGT != nil {
		predicates = append(predicates, work.TitleGT(*i.TitleGT))
	}
	if i.TitleGTE != nil {
		predicates = append(predicates, work.TitleGTE(*i.TitleGTE))
	}
	if i.TitleLT != nil {
		predicates = append(predicates, work.TitleLT(*i.TitleLT))
	}
	if i.TitleLTE != nil {
		predicates = append(predicates, work.TitleLTE(*i.TitleLTE))
	}
	if i.TitleContains != nil {
		predicates = append(predicates, work.TitleContains(*i.TitleContains))
	}
	if i.TitleHasPrefix != nil {
		predicates = append(predicates, work.TitleHasPrefix(*i.TitleHasPrefix))
	}
	if i.TitleHasSuffix != nil {
		predicates = append(predicates, work.TitleHasSuffix(*i.TitleHasSuffix))
	}
	if i.TitleEqualFold != nil {
		predicates = append(predicates, work.TitleEqualFold(*i.TitleEqualFold))
	}
	if i.TitleContainsFold != nil {
		predicates = append(predicates, work.TitleContainsFold(*i.TitleContainsFold))
	}
	if i.Description != nil {
		predicates = append(predicates, work.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, work.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, work.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, work.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, work.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, work.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, work.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, work.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, work.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, work.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, work.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, work.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, work.DescriptionContainsFold(*i.DescriptionContainsFold))
	}
	if i.ImageURL != nil {
		predicates = append(predicates, work.ImageURLEQ(*i.ImageURL))
	}
	if i.ImageURLNEQ != nil {
		predicates = append(predicates, work.ImageURLNEQ(*i.ImageURLNEQ))
	}
	if len(i.ImageURLIn) > 0 {
		predicates = append(predicates, work.ImageURLIn(i.ImageURLIn...))
	}
	if len(i.ImageURLNotIn) > 0 {
		predicates = append(predicates, work.ImageURLNotIn(i.ImageURLNotIn...))
	}
	if i.ImageURLGT != nil {
		predicates = append(predicates, work.ImageURLGT(*i.ImageURLGT))
	}
	if i.ImageURLGTE != nil {
		predicates = append(predicates, work.ImageURLGTE(*i.ImageURLGTE))
	}
	if i.ImageURLLT != nil {
		predicates = append(predicates, work.ImageURLLT(*i.ImageURLLT))
	}
	if i.ImageURLLTE != nil {
		predicates = append(predicates, work.ImageURLLTE(*i.ImageURLLTE))
	}
	if i.ImageURLContains != nil {
		predicates = append(predicates, work.ImageURLContains(*i.ImageURLContains))
	}
	if i.ImageURLHasPrefix != nil {
		predicates = append(predicates, work.ImageURLHasPrefix(*i.ImageURLHasPrefix))
	}
	if i.ImageURLHasSuffix != nil {
		predicates = append(predicates, work.ImageURLHasSuffix(*i.ImageURLHasSuffix))
	}
	if i.ImageURLEqualFold != nil {
		predicates = append(predicates, work.ImageURLEqualFold(*i.ImageURLEqualFold))
	}
	if i.ImageURLContainsFold != nil {
		predicates = append(predicates, work.ImageURLContainsFold(*i.ImageURLContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, work.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, work.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, work.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, work.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, work.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, work.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, work.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, work.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, work.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, work.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, work.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, work.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, work.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, work.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, work.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, work.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasCategory != nil {
		p := work.HasCategory()
		if !*i.HasCategory {
			p = work.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCategoryWith) > 0 {
		with := make([]predicate.Category, 0, len(i.HasCategoryWith))
		for _, w := range i.HasCategoryWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, work.HasCategoryWith(with...))
	}
	if i.HasOwner != nil {
		p := work.HasOwner()
		if !*i.HasOwner {
			p = work.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, work.HasOwnerWith(with...))
	}
	if i.HasLikers != nil {
		p := work.HasLikers()
		if !*i.HasLikers {
			p = work.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasLikersWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasLikersWith))
		for _, w := range i.HasLikersWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, work.HasLikersWith(with...))
	}
	if i.HasComments != nil {
		p := work.HasComments()
		if !*i.HasComments {
			p = work.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCommentsWith) > 0 {
		with := make([]predicate.Comment, 0, len(i.HasCommentsWith))
		for _, w := range i.HasCommentsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, work.HasCommentsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("artsign/ent: empty predicate WorkWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return work.And(predicates...), nil
	}
}
