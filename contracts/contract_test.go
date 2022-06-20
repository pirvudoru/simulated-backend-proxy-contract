package contracts

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pirvudoru/simulated-backend-proxy/contracts/proxy"
	"github.com/pirvudoru/simulated-backend-proxy/contracts/target"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"testing"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

type Contracts struct {
	proxy        *proxy.Proxy
	proxyAddress common.Address

	target        *target.Target
	targetAddress common.Address
}

var simulatedChainId = big.NewInt(1337)

func TestMintThroughProxy(t *testing.T) {
	ownerWallet := createWallet()
	aliceWallet := createWallet()

	backend := createBackend(ownerWallet, aliceWallet)

	contracts := deployContracts(ownerWallet, backend)

	tx, err := contracts.proxy.Mint(auth(ownerWallet), aliceWallet.address)
	assert.Nil(t, err)
	assert.NotNil(t, tx)

	backend.Commit()

	mintEvents, err := contracts.target.FilterMint(nil, []common.Address{aliceWallet.address})
	mintEvents.Next()

	assert.Equal(t, mintEvents.Event.Amount, big.NewInt(42))
	assert.Equal(t, mintEvents.Event.Destination, aliceWallet.address)
}

func deployContracts(deployer *Wallet, backend *backends.SimulatedBackend) *Contracts {
	targetAddress, _, target, err := target.DeployTarget(auth(deployer), backend)
	if err != nil {
		log.Fatal(fmt.Errorf("failed deploying dPOR: %s", err))
	}

	proxyAddress, _, proxy, err := proxy.DeployProxy(auth(deployer), backend, targetAddress)
	if err != nil {
		log.Fatal(fmt.Errorf("failed deploying CxToken: %s", err))
	}

	target.AddMinter(auth(deployer), proxyAddress)

	return &Contracts{
		targetAddress: targetAddress,
		target:        target,
		proxyAddress:  proxyAddress,
		proxy:         proxy,
	}
}

func createWallet() *Wallet {
	privateKey, _ := crypto.GenerateKey()
	ownerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &Wallet{
		privateKey: privateKey,
		address:    ownerAddress,
	}
}

func createBackend(fundedWallets ...*Wallet) *backends.SimulatedBackend {
	balance := tokenAmount(10)

	genesisAlloc := map[common.Address]core.GenesisAccount{}
	for _, wallet := range fundedWallets {
		genesisAlloc[wallet.address] = core.GenesisAccount{
			Balance: balance,
		}
	}
	blockGasLimit := uint64(47123880000)

	backend := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	return backend
}

func auth(wallet *Wallet) *bind.TransactOpts {
	item, _ := bind.NewKeyedTransactorWithChainID(wallet.privateKey, simulatedChainId)

	return item
}

func tokenAmount(amount int) *big.Int {
	result := new(big.Int)
	result.SetString("1000000000000000000", 10)

	result.Mul(result, big.NewInt(int64(amount)))
	return result
}
