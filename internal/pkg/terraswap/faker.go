package terraswap

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func NewFakeToken() Token {
	token := Token{}
	faker.FakeData(&token)
	return token
}

func NewFakeAssetInfo() AssetInfo {
	idx := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	ai := AssetInfo{}
	if err := faker.FakeData(&ai); err != nil {
		panic(err)
	}

	if idx%2 == 0 {
		ai.NativeToken = nil
	} else {
		ai.Token = nil
	}

	return ai
}

func NewFakePair() Pair {
	return Pair{
		AssetInfos:     []AssetInfo{NewFakeAssetInfo(), NewFakeAssetInfo()},
		LiquidityToken: faker.UUIDHyphenated(),
		ContractAddr:   faker.UUIDHyphenated(),
	}
}

func NewFakePairs(len uint) []Pair {
	pairs := []Pair{}
	for i := uint(0); i < len; i++ {
		pairs = append(pairs, NewFakePair())
	}
	return pairs
}
