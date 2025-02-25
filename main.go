package main

import (
	"fmt"

	"github.com/andreyyrson/blockchain/blockchain"
	"github.com/andreyyrson/blockchain/transaction"
	"github.com/andreyyrson/blockchain/wallet"
)

func main() {
	// Cria uma nova carteira
	myWallet := wallet.NewWallet()
	fmt.Println("Carteira Criada:")
	fmt.Printf("Endereço: %s\n", myWallet.Address)

	// Assina uma mensagem
	message := []byte("Hello, Blockchain!")
	signature, err := myWallet.SignData(message)
	if err != nil {
		fmt.Println("Erro ao assinar mensagem:", err)
		return
	}
	fmt.Println("\nMensagem Assinada:")
	fmt.Printf("Assinatura: %s\n", signature)

	// Verifica a assinatura
	isValid := wallet.VerifySignature(myWallet.PublicKey, message, signature)
	fmt.Println("\nVerificação de Assinatura:")
	if isValid {
		fmt.Println("Assinatura válida!")
	} else {
		fmt.Println("Assinatura inválida!")
	}

	// Cria uma nova blockchain com dificuldade 4 e recompensa de 100
	bc := blockchain.NewBlockchain(4, 100)

	// Cria duas carteiras
	wallet1 := wallet.NewWallet()
	wallet2 := wallet.NewWallet()

	// Adiciona transações pendentes
	bc.AddTransaction(transaction.NewTransaction(wallet1.Address, wallet2.Address, 50))
	bc.AddTransaction(transaction.NewTransaction(wallet2.Address, wallet1.Address, 25))

	// Minera as transações pendentes
	fmt.Println("Mineração em andamento...")
	bc.MinePendingTransactions(wallet1.Address)

	// Exibe o saldo das carteiras
	fmt.Println("\nSaldo da Carteira 1:", bc.GetBalance(wallet1.Address))
	fmt.Println("Saldo da Carteira 2:", bc.GetBalance(wallet2.Address))

	// Valida a blockchain
	fmt.Println("\nBlockchain válida?", bc.IsChainValid())
}
