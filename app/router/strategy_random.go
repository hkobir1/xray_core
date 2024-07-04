package router

import (
	"context"

	"github.com/hkobir1/xray_core/app/observatory"
	"github.com/hkobir1/xray_core/common"
	"github.com/hkobir1/xray_core/common/dice"
	"github.com/hkobir1/xray_core/core"
	"github.com/hkobir1/xray_core/features/extension"
)

// RandomStrategy represents a random balancing strategy
type RandomStrategy struct {
	FallbackTag string

	ctx         context.Context
	observatory extension.Observatory
}

func (s *RandomStrategy) InjectContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *RandomStrategy) GetPrincipleTarget(strings []string) []string {
	return strings
}

func (s *RandomStrategy) PickOutbound(candidates []string) string {
	if len(s.FallbackTag) > 0 && s.observatory == nil {
		common.Must(core.RequireFeatures(s.ctx, func(observatory extension.Observatory) error {
			s.observatory = observatory
			return nil
		}))
	}
	if s.observatory != nil {
		observeReport, err := s.observatory.GetObservation(s.ctx)
		if err == nil {
			aliveTags := make([]string, 0)
			if result, ok := observeReport.(*observatory.ObservationResult); ok {
				status := result.Status
				statusMap := make(map[string]*observatory.OutboundStatus)
				for _, outboundStatus := range status {
					statusMap[outboundStatus.OutboundTag] = outboundStatus
				}
				for _, candidate := range candidates {
					if outboundStatus, found := statusMap[candidate]; found {
						if outboundStatus.Alive {
							aliveTags = append(aliveTags, candidate)
						}
					} else {
						// unfound candidate is considered alive
						aliveTags = append(aliveTags, candidate)
					}
				}
				candidates = aliveTags
			}
		}
	}

	count := len(candidates)
	if count == 0 {
		// goes to fallbackTag
		return ""
	}
	return candidates[dice.Roll(count)]
}
