package tx

import (
	"fmt"

	"github.com/terraswap/terraswap-service/internal/app/api/utils/responser"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

var _ Service = &classicServiceImpl{}

type classicServiceImpl struct {
	mixinImpl
}

func newClassicService(r Repository) Service {
	return &classicServiceImpl{
		mixinImpl{r},
	}
}

func (s *classicServiceImpl) GetSwapTxs(from, to, amount, sender, max_spread, belief_price string, hop_count int) ([][]*terraswap.UnsignedTx, *responser.ErrorResponse) {

	terraAmount, err := s.convertToTerraAmount(amount, from)
	if err != nil {
		msg := fmt.Sprintf("cannot convert amount(%s) for %s", amount, from)
		res := responser.GetBadRequest(msg, "")
		return nil, &res
	}

	utxs := [][]*terraswap.UnsignedTx{}
	if terraswap.IsNativeToken(from) && terraswap.IsNativeToken(to) {
		if utx, err := s.getRouteSwapTx(from, terraAmount, sender, max_spread, belief_price, []string{to}); err == nil {
			utxs = append(utxs, utx)
		}
	}

	paths := s.repo.GetRoutes(from, to)
	for _, path := range paths {
		var utx []*terraswap.UnsignedTx
		pathLen := len(path)

		if pathLen-1 > hop_count {
			break
		}

		if pathLen == 1 {
			utx = s.getSwapTx(from, to, terraAmount, sender, max_spread, belief_price)
		} else {
			var err error
			utx, err = s.getRouteSwapTx(from, terraAmount, sender, max_spread, belief_price, path)
			if err != nil {
				continue
			}
		}
		utxs = append(utxs, utx)
	}

	if len(utxs) == 0 {
		msg := fmt.Sprintf("cannot find a path(%s, %s)", from, to)
		res := responser.NotFound(msg, "")
		return nil, &res
	}

	return utxs, nil
}
