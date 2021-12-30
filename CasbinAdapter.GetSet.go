package adapter

// SetRepo set database repo interface
func (adapter *CasbinAdapter) SetRepo(repo IRepo) *CasbinAdapter {
	adapter.repo = repo
	return adapter
}

// SetBatchSize set how many policies saved once
func (adapter *CasbinAdapter) SetBatchSize(size int) *CasbinAdapter {
	adapter.batchSize = size
	return adapter
}

func (adapter *CasbinAdapter) GetBatchSize() int {
	if adapter.batchSize == 0 {
		return 1000
	}
	return adapter.batchSize
}
