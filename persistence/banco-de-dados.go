package persistence

import (
	"encoding/json"
	"fmt"

	"github.com/andreyyrson/blockchain/block"
	"go.etcd.io/bbolt"
)

const (
	dbFileName = "blockchain.db" //falta criar um banco quando eu criar chamar o caminho do arquivo aqui
	bucketName = "blocks"
)

func SaveBlock(b *block.Block) error {
	db, err := bbolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return fmt.Errorf("erro ao abrir o banco de dados: %v", err)
	}
	defer db.Close()

	return db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("erro ao criar o bucket: %v", err)
		}

		blockData, err := json.Marshal(b)
		if err != nil {
			return fmt.Errorf("erro ao serializar o bloco: %v", err)
		}

		err = bucket.Put([]byte(fmt.Sprintf("%d", b.Index)), blockData)
		if err != nil {
			return fmt.Errorf("erro ao salvar o bloco: %v", err)
		}

		return nil
	})
}

func LoadBlockchain() ([]*block.Block, error) {
	var blocks []*block.Block

	db, err := bbolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o banco de dados: %v", err)
	}
	defer db.Close()

	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket não encontrado")
		}

		return bucket.ForEach(func(k, v []byte) error {
			var block block.Block
			if err := json.Unmarshal(v, &block); err != nil {
				return fmt.Errorf("erro ao desserializar o bloco: %v", err)
			}
			blocks = append(blocks, &block)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return blocks, nil
}
