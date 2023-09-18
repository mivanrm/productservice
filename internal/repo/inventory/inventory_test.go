package inventory

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/jmoiron/sqlx"
	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"
)

func Test_inventoryRepo_CreateInventory(t *testing.T) {

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	type args struct {
		inventory inventoryentity.Inventory
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				inventory: inventoryentity.Inventory{
					VariantID: 1,
					Amount:    10,
				},
			},
			mock: func() {
				mock.ExpectQuery(`INSERT INTO inventory (.+) RETURNING stock_id`).
					WithArgs(1, 10).
					WillReturnRows(sqlmock.NewRows([]string{"stock_id"}).AddRow(1))
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &inventoryRepo{
				db: db,
			}
			got, err := r.CreateInventory(tt.args.inventory)
			if (err != nil) != tt.wantErr {
				t.Errorf("inventoryRepo.CreateInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inventoryRepo.CreateInventory() = %v, want %v", got, tt.want)
			}
		})
	}
}
