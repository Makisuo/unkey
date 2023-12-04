// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: find_workspace.sql

package database

import (
	"context"
)

const findWorkspace = `-- name: FindWorkspace :one
SELECT
    id, name, slug, tenant_id, stripe_customer_id, stripe_subscription_id, plan, quota_max_active_keys, usage_active_keys, quota_max_verifications, usage_verifications, last_usage_update, billing_period_start, billing_period_end, trial_ends, features, beta_features, subscriptions, plan_locked_until, plan_changed
FROM
    ` + "`" + `workspaces` + "`" + `
WHERE
    id = ?
`

func (q *Queries) FindWorkspace(ctx context.Context, id string) (Workspace, error) {
	row := q.db.QueryRowContext(ctx, findWorkspace, id)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.TenantID,
		&i.StripeCustomerID,
		&i.StripeSubscriptionID,
		&i.Plan,
		&i.QuotaMaxActiveKeys,
		&i.UsageActiveKeys,
		&i.QuotaMaxVerifications,
		&i.UsageVerifications,
		&i.LastUsageUpdate,
		&i.BillingPeriodStart,
		&i.BillingPeriodEnd,
		&i.TrialEnds,
		&i.Features,
		&i.BetaFeatures,
		&i.Subscriptions,
		&i.PlanLockedUntil,
		&i.PlanChanged,
	)
	return i, err
}
