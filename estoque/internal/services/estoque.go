package services

import (
	"estoque/internal/models"
	"fmt"
	"time"
)

type Estoque struct {
	items map[string]models.Item // items: Record<string, Item> no TS
	logs  []models.Log
}

// retorna um ponteiro de Estoque.

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item, userId string) error {
	if item.Quanity <= 0 {
		return fmt.Errorf("erro ao tentar adicionar estoque. ID: %s. Quantidade inválida", item.ID)
	}

	existingItem, exists := e.items[item.ID]

	if exists {
		item.Quanity += existingItem.Quanity
	}

	e.items[item.ID] = item

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		UserId:    userId,
		ItemId:    item.ID,
		Quantity:  item.Quanity,
		Reason:    "Repondo estoque acabando",
	})

	return nil
}

func (e *Estoque) ListItems() ([]models.Item, error) {
	var itemList []models.Item

	for _, item := range e.items {
		itemList = append(itemList, item)
	}

	return itemList, nil
}

func (e *Estoque) ListLogs() ([]models.Log, error) {

	return e.logs, nil
}

// func FindByName(data []models.Item, name string) ([]models.Item, error) {
// 	var result []models.Item
// 	for _, item := range data {
// 		if item.Name == name {
// 			result = append(result, item)
// 		}
// 	}
// 	if len(result) == 0 {
// 		return nil, fmt.Errorf("nenhum item com o nome '%s' foi encontrado", name)
// 	}
// 	return result, nil
// }

// Função genérica. Possibilita buscar por qualquer campo.

func Find[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}

	return result, nil
}
