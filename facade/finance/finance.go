package finance

type Store interface {
	GetBalance() float32
	SetBalance(amount float32)
}

type FinStorage interface {
	StoreAmount(store Store, amount float32)
}

type FinStorageFacade struct {
	Storage FinStorage
}

func NewFinStorageFacade(storage FinStorage) *FinStorageFacade {
	return &FinStorageFacade{
		Storage: storage,
	}
}

func (f *FinStorageFacade) StoreAmount(store Store, amount float32) {
	f.Storage.StoreAmount(store, amount)
}
