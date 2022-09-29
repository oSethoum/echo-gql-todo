package auth

import (
	"context"
	"todo/ent/privacy"
)

func QueryPrivacy(node string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		// grab values from context
		return privacy.Skip
	})
}

func MutationPrivacy(node string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		// grab values from context
		return privacy.Skip
	})
}
