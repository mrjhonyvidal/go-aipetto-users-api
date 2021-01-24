package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {}

func (i *itemsService) GetItem() {
	panic("implement me")
}

func (i *itemsService) SaveItem() {
	panic("implement me")
}


