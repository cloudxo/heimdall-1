package data_collector

import (
	"git01.bravofly.com/n7/heimdall/cmd/model"
	"github.com/cloudflare/cloudflare-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_correctAdapting(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	now := time.Now().In(utc)
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())

	triggers := make([]model.WafTrigger, 0)
	triggers = append(triggers, newWafTrigger("host.it", "drop", now))
	triggers = append(triggers, newWafTrigger("host.it", "drop", now))
	triggers = append(triggers, newWafTrigger("host.it", "challenge", now))
	triggers = append(triggers, newWafTrigger("host.it", "challenge", now))
	triggers = append(triggers, newWafTrigger("host.it", "challenge", now))
	triggers = append(triggers, newWafTrigger("host.it", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.it", "jschallenge", now))

	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))
	triggers = append(triggers, newWafTrigger("host.com", "simulate", now))

	aggregate := model.NewAggregate(cloudflare.Zone{
		ID:   "123",
		Name: "zone",
	})

	aggregate.Totals[now] = model.NewCounters()

	collectWaf(triggers, utc, aggregate)

	assert.Equal(t, 2, aggregate.Totals[now].WafTrigger["host.it"].Drop.Value)
	assert.Equal(t, 7, aggregate.Totals[now].WafTrigger["host.com"].Simulate.Value)

}

func newWafTrigger(host, action string, now time.Time) model.WafTrigger {
	return model.WafTrigger{
		Host:       host,
		Action:     action,
		OccurredAt: now,
	}
}