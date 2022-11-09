package database

import (
	"devfullcycle/gointensivo/internal/order/entity"
	"fmt"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	DB *gorm.DB
}

func ConectaComBancoDeDadosTeste() *gorm.DB {
	connectString := "host=postgres_test user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
	return DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db := ConectaComBancoDeDadosTeste()
	db.Exec("CREATE TABLE orders_test (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.DB = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	fmt.Println("Fim...")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.DB)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order

	suite.DB.First(&orderResult, order.ID)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
