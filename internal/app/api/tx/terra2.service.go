package tx

import (
	"fmt"

	"github.com/terraswap/terraswap-service/internal/app/api/utils/responser"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

var _ Service = &phoenixServiceImpl{}

type phoenixServiceImpl struct {
	mixinImpl
}

func newService(r Repository) Service {
	return &phoenixServiceImpl{
		mixinImpl{r},
	}
}

func (s *phoenixServiceImpl) GetSwapTxs(from, to, amount, sender, max_spread, belief_price string, deadline uint64, hop_count int) ([][]*terraswap.UnsignedTx, *responser.ErrorResponse) {

	terraAmount, err := s.convertToTerraAmount(amount, from)
	if err != nil {
		msg := fmt.Sprintf("cannot convert amount(%s) for %s", amount, from)
		res := responser.GetBadRequest(msg, "")
		return nil, &res
	}

	utxs := [][]*terraswap.UnsignedTx{}
	paths := s.repo.GetRoutes(from, to)

	for _, path := range paths {
		var utx []*terraswap.UnsignedTx
		pathLen := len(path)

		if pathLen-1 > hop_count {
			break
		}

		if pathLen == 1 {
			utx = s.getSwapTx(from, to, terraAmount, sender, max_spread, belief_price, deadline)
		} else {
			var err error
			utx, err = s.getRouteSwapTx(from, terraAmount, sender, path, deadline)
			if err != nil {
				continue
			}
		}
		utxs = append(utxs, utx)
	}

	return utxs, nil
}
