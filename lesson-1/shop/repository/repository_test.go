package repository

import (
	"gb-go-architecture/lesson-1/shop/models"
	"testing"
)

func testItemsFirst() (item *models.Item) {
	for _, value := range testItems {
		// we just need any Item
		item = value
		break
	}
	return item
}

func TestGetItem(t *testing.T) {
	db := NewMapDB()
	expectedItem, err := db.CreateItem(testItemsFirst())
	if err != nil {
		t.Fatal("unexpected error during setup test: ", err)
	}

	t.Run("Get existing", func(t *testing.T) {
		result, err := db.GetItem(expectedItem.ID)
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if expectedItem.ID != result.ID {
			t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, result.ID)
		}
		if expectedItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if expectedItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	})

	t.Run("Get not existing", func(t *testing.T) {
		_, err := db.GetItem(-1)
		if err != ErrNotFound {
			t.Errorf("unexpected error raised: expected %v, result: %v ", ErrNotFound, err)
		}
	})
}

func TestCreateItem(t *testing.T) {
	db := NewMapDB()

	for _, item := range testItems {
		expectedItem, err := db.CreateItem(item)
		if err != nil {
			t.Error("unexpected error during create: ", err)
		}

		result, err := db.GetItem(expectedItem.ID)
		if err != nil {
			t.Error("unexpected error during get: ", err)
		}

		if expectedItem.ID != result.ID {
			t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, result.ID)
		}
		if expectedItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if expectedItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	}
}

func TestListItem(t *testing.T) {
	db := NewMapDB()

	// fill repo with data
	for _, item := range testItems {
		_, err := db.CreateItem(item)
		if err != nil {
			t.Fatal("unexpected error during setup test: ", err)
		}
	}

	t.Run("Empty filter", func(t *testing.T) {
		result, err := db.ListItems(&ItemFilter{})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		// check if enough items returned
		if len(testItems) != len(result) {
			t.Errorf("not enough items: expected %d, result %d", len(testItems), len(result))
		}

		// check if not any "extra" items returned
		for _, item := range result {
			if _, ok := testItems[item.Name]; !ok {
				t.Errorf("unexpected error: item with name %s did not exists in test data", item.Name)
			}

		}
	})

	t.Run("Price gt filter", func(t *testing.T) {
		priceLeft := int64(20)
		result, err := db.ListItems(&ItemFilter{PriceLeft: &priceLeft})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if len(result) == 0 {
			t.Errorf("not enough items: expected gt 0, result: %d", len(result))
		}

		for _, item := range result {
			if priceLeft > item.Price {
				t.Errorf("invalid price: expected gt %d, result %d", priceLeft, item.Price)
			}
		}
	})

	t.Run("Price lt filter", func(t *testing.T) {
		priceRight := int64(40)
		result, err := db.ListItems(&ItemFilter{PriceRight: &priceRight})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if len(result) == 0 {
			t.Errorf("not enough items: expected gt 0, result: %d", len(result))
		}

		for _, item := range result {
			if priceRight < item.Price {
				t.Errorf("invalid price: expected lt %d, result %d", priceRight, item.Price)
			}
		}
	})

	t.Run("Price gt+lt filter", func(t *testing.T) {
		priceLeft := int64(20)
		priceRight := int64(40)
		result, err := db.ListItems(&ItemFilter{
			PriceLeft:  &priceLeft,
			PriceRight: &priceRight})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if len(result) == 0 {
			t.Errorf("not enough items: expected gt 0, result: %d", len(result))
		}

		for _, item := range result {
			if priceRight < item.Price || priceLeft > item.Price {
				t.Errorf("invalid price: expected between %d and %d, result %d", priceLeft, priceRight, item.Price)
			}
		}
	})

	t.Run("Limit+offset", func(t *testing.T) {
		limitMax := 4
		limitMin := 2

		resultLimitMax, err := db.ListItems(&ItemFilter{
			Limit: limitMax,
		})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if limitMax != len(resultLimitMax) {
			t.Errorf("not enough items: expected %d, result: %d", limitMax, len(resultLimitMax))
		}

		expectedOffset0 := resultLimitMax[0:limitMin]
		resultOffset0, err := db.ListItems(&ItemFilter{
			Limit:  limitMin,
			Offset: 0,
		})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if len(expectedOffset0) != len(resultOffset0) {
			t.Errorf("not enough items on page 1: expected %d, result: %d", len(expectedOffset0), len(resultOffset0))
		}
		for idx, item := range resultOffset0 {
			if resultLimitMax[idx] != item {
				t.Errorf("invalid item on page 1: expected %v, result %v", resultLimitMax[idx], item)
			}
		}

		expectedOffset1 := resultLimitMax[limitMin : limitMin*2]
		resultOffset1, err := db.ListItems(&ItemFilter{
			Limit:  limitMin,
			Offset: limitMin,
		})
		if err != nil {
			t.Error("unexpected error: ", err)
		}

		if len(expectedOffset1) != len(resultOffset1) {
			t.Errorf("not enough items on page 2: expected %d, result: %d", len(expectedOffset1), len(resultOffset1))
		}
		for idx, item := range resultOffset1 {
			if resultLimitMax[idx+limitMin] != item {
				t.Errorf("invalid item on page 2: expected %v, result %v", resultLimitMax[idx+limitMin], item)
			}
		}
	})

}

func TestDeleteItem(t *testing.T) {
	db := NewMapDB()
	expectedItem, err := db.CreateItem(testItemsFirst())
	if err != nil {
		t.Fatal("unexpected error during setup test: ", err)
	}

	t.Run("Delete existing", func(t *testing.T) {
		err := db.DeleteItem(expectedItem.ID)
		if err != nil {
			t.Error("unexpected error during delete: ", err)
		}

		// item should not exists anymore
		_, err = db.GetItem(expectedItem.ID)
		if err != ErrNotFound {
			t.Errorf("unexpected error raised: expected %v, result: %v ", ErrNotFound, err)
		}

	})

	t.Run("Delete not existing", func(t *testing.T) {
		err := db.DeleteItem(-1)
		if err != ErrNotFound {
			t.Errorf("unexpected error raised: expected %v, result: %v ", ErrNotFound, err)
		}
	})
}

func TestUpdateItem(t *testing.T) {
	db := NewMapDB()
	expectedItem, err := db.CreateItem(testItemsFirst())
	if err != nil {
		t.Fatal("unexpected error during setup test: ", err)
	}

	t.Run("Update existing", func(t *testing.T) {
		newItem := &models.Item{ID: expectedItem.ID, Name: expectedItem.Name, Price: 100}
		_, err := db.UpdateItem(newItem)
		if err != nil {
			t.Error("unexpected error during update: ", err)
		}

		result, err := db.GetItem(expectedItem.ID)
		if err != nil {
			t.Error("unexpected error during get: ", err)
		}

		if expectedItem.ID != result.ID {
			t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, result.ID)
		}
		if newItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if newItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	})

	t.Run("Update not existing", func(t *testing.T) {
		newItem := &models.Item{ID: -1, Name: expectedItem.Name, Price: 100}
		_, err := db.UpdateItem(newItem)
		if err != ErrNotFound {
			t.Errorf("unexpected error raised: expected %v, result: %v ", ErrNotFound, err)
		}
	})

}
