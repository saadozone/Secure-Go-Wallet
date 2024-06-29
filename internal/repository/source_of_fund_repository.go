package repository

import (
	"github.com/saadozone/gin-gorm-rest/internal/model"
	"gorm.io/gorm"
)

// Define constants for source of funds
const (
	BankTransfer = 1
	CreditCard   = 2
	Cash         = 3
)

type SourceOfFundRepository interface {
	FindById(id int) (*model.SourceOfFund, error)
}

type sourceOfFundRepository struct {
	db *gorm.DB
}

type SRConfig struct {
	DB *gorm.DB
}

// NewSourceOfFundRepository creates a new instance of SourceOfFundRepository.
func NewSourceOfFundRepository(c *SRConfig) SourceOfFundRepository {
	return &sourceOfFundRepository{
		db: c.DB,
	}
}

// FindById retrieves a source of fund by its ID.
func (r *sourceOfFundRepository) FindById(id int) (*model.SourceOfFund, error) {
	var sourceOfFund *model.SourceOfFund

	// Handle hardcoded values
	switch id {
	case BankTransfer:
		sourceOfFund = &model.SourceOfFund{ID: BankTransfer, Name: "Bank Transfer"}
	case CreditCard:
		sourceOfFund = &model.SourceOfFund{ID: CreditCard, Name: "Credit Card"}
	case Cash:
		sourceOfFund = &model.SourceOfFund{ID: Cash, Name: "Cash"}
	default:
		return nil, gorm.ErrRecordNotFound
	}

	return sourceOfFund, nil
}
