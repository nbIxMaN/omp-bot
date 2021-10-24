package life

import (
	"github.com/ozonmp/omp-bot/internal/errors/insurance/life"
	"github.com/ozonmp/omp-bot/internal/model/insurance"
)

type DummyLifeService struct {
	lifeStorage []insurance.Life
	nextId      uint64
}

func NewDummyLifeService() *DummyLifeService {
	mapLifeService := DummyLifeService{
		lifeStorage: []insurance.Life{},
		nextId:      1,
	}
	for _, v := range AllLifes {
		mapLifeService.Create(v)
	}
	return &mapLifeService
}

func (mapLifeService *DummyLifeService) findById(LifeID uint64) (int, error) {
	for i, lifeObject := range mapLifeService.lifeStorage {
		if lifeObject.Id == LifeID {
			return i, nil
		}
	}
	return -1, life.NotFoundError()
}

func (mapLifeService *DummyLifeService) Describe(LifeID uint64) (*insurance.Life, error) {
	if i, err := mapLifeService.findById(LifeID); err == nil {
		return &mapLifeService.lifeStorage[i], nil
	} else {
		return nil, err
	}
}

func (mapLifeService *DummyLifeService) List(cursor uint64, limit uint64) ([]insurance.Life, error) {

	lifeStorageLast := uint64(len(mapLifeService.lifeStorage))

	if lifeStorageLast == 0 {
		return mapLifeService.lifeStorage, life.LastPageError()
	}

	if cursor < 0 || cursor >= uint64(len(mapLifeService.lifeStorage)) {
		return nil, life.IncorrectPositionError()
	}

	var err error
	last := cursor + limit

	if last >= lifeStorageLast {
		last = lifeStorageLast
		err = life.LastPageError()
	}

	return mapLifeService.lifeStorage[cursor:last], err
}

func (mapLifeService *DummyLifeService) Create(life insurance.Life) (uint64, error) {
	life.Id = mapLifeService.nextId
	mapLifeService.nextId++
	mapLifeService.lifeStorage = append(mapLifeService.lifeStorage, life)
	return life.Id, nil
}

func (mapLifeService *DummyLifeService) Update(LifeID uint64, life insurance.Life) error {
	if i, err := mapLifeService.findById(LifeID); err == nil {
		life.Id = LifeID
		mapLifeService.lifeStorage[i] = life
		return nil
	} else {
		return err
	}
}

func (mapLifeService *DummyLifeService) Remove(LifeID uint64) (bool, error) {
	if i, err := mapLifeService.findById(LifeID); err == nil {
		mapLifeService.lifeStorage = append(mapLifeService.lifeStorage[:i],
			mapLifeService.lifeStorage[i+1:]...)
		return true, nil
	} else {
		return false, err
	}
}
