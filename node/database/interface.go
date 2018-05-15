package database

import (
	"github.com/gelembjuk/democoin/lib/utils"
)

type DBManager interface {
	SetConfig(config DatabaseConfig) error
	SetLogger(logger *utils.LoggerMan) error

	InitDatabase() error
	CheckDBExists() (bool, error)

	OpenConnection(reason string) error
	CloseConnection() error
	IsConnectionOpen() bool

	GetBlockchainObject() (BlockchainInterface, error)
	GetTransactionsObject() (TranactionsInterface, error)
	GetUnapprovedTransactionsObject() (UnapprovedTransactionsInterface, error)
	GetUnspentOutputsObject() (UnspentOutputsInterface, error)
	GetNodesObject() (NodesInterface, error)
}

type DatabaseConnection interface {
	Close() error
}
type ForEachKeyIteratorInterface func(key, value []byte)

type CursorInterface interface {
	Next() ([]byte, []byte, error)
	Count() (int, error)
}

type BlockchainInterface interface {
	InitDB() error

	GetTopBlock() ([]byte, error)
	GetBlock(hash []byte) ([]byte, error)
	PutBlockOnTop(hash []byte, blockdata []byte) error
	PutBlock(hash []byte, blockdata []byte) error
	CheckBlockExists(hash []byte) (bool, error)
	DeleteBlock(hash []byte) error
	SaveTopHash(hash []byte) error
	GetTopHash() ([]byte, error)
	SaveFirstHash(hash []byte) error
	GetFirstHash() ([]byte, error)
}

type TranactionsInterface interface {
	InitDB() error
	TruncateDB() error
	PutTXToBlockLink(txID []byte, blockHash []byte) error
	GetBlockHashForTX(txID []byte) ([]byte, error)
	DeleteTXToBlockLink(txID []byte) error
	PutTXSpentOutputs(txID []byte, outputs []byte) error
	GetTXSpentOutputs(txID []byte) ([]byte, error)
	DeleteTXSpentData(txID []byte) error
}

type UnapprovedTransactionsInterface interface {
	InitDB() error
	TruncateDB() error
	GetCursor() (CursorInterface, error)

	GetTransaction(txID []byte) ([]byte, error)
	PutTransaction(txID []byte, txdata []byte) error
	DeleteTransaction(txID []byte) error
}

type UnspentOutputsInterface interface {
	InitDB() error
	TruncateDB() error
	GetCursor() (CursorInterface, error)

	GetDataForTransaction(txID []byte) ([]byte, error)
	DeleteDataForTransaction(txID []byte) error
	PutDataForTransaction(txID []byte, txData []byte) error
}

type NodesInterface interface {
	InitDB() error
	ForEach(callback ForEachKeyIteratorInterface) error

	PutNode(nodeID []byte, nodeData []byte) error
	DeleteNode(nodeID []byte) error
}
