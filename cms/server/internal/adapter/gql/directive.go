//go:generate go run github.com/99designs/gqlgen generate

package gql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func NewDirective() DirectiveRoot {
	return DirectiveRoot{
		OnlyOne: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			return next(ctx)

			/* This can not be implemented because of this issue: https://github.com/99designs/gqlgen/issues/2281
			   TODO: enable and check this code after the mentioned issue getting solved!

			// fc := graphql.GetFieldContext(ctx)
			inputMap, ok := obj.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("is not an object, onlyOne directive should be applied on objects only")
			}
			if violatingOneOfLogic(inputMap) {
				return nil, fmt.Errorf("should have only one property")
			}

			return next(ctx)*/
		},
	}
}

/*func violatingOneOfLogic(m map[string]interface{}) bool {
	return len(m) > 1
}*/
