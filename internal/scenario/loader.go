package scenario

import "github.com/IAmRiteshKoushik/smriti/internal/common"

func Load(path string) (*Scenario, error) {
	return common.Load[Scenario](path)
}
