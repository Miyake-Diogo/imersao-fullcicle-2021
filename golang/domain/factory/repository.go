package factory
import (
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/domain/repository"
)


type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}