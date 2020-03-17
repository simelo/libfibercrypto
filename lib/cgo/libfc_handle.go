package main

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>


  #include "fctypes.h"
*/
import "C"

type Handle uint64

var (
	handlesCounter uint64 = 0
	handleMap             = make(map[Handle]interface{})
)

func registerHandle(obj interface{}) C.Handle {
	handlesCounter++
	handle := handlesCounter
	handleMap[Handle(handle)] = obj
	return (C.Handle)(handle)
}

func lookupHandle(handle C.Handle) (interface{}, bool) {
	obj, ok := handleMap[Handle(handle)]
	return obj, ok
}

func overwriteHandle(handle C.Handle, obj interface{}) bool {
	_, ok := handleMap[Handle(handle)]
	if ok {
		handleMap[Handle(handle)] = obj
		return true
	}
	return false
}

func closeHandle(handle Handle) {
	delete(handleMap, handle)
}

//export FC_handle_close
func FC_handle_close(handle C.Handle) {
	closeHandle(Handle(handle))
}
func registerStringsHandle(obj []string) C.Strings__Handle {
	return (C.Strings__Handle)(registerHandle(obj))
}

func lookupStringsHandle(handle C.Strings__Handle) ([]string, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).([]string); isOK {
			return obj, true
		}
	}
	return nil, false
}

func lookupAltcoinPluginHandle(handle C.AltcoinPlugin__Handle) (*core.AltcoinPlugin, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.AltcoinPlugin); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerAltcoinPluginHandle(obj *core.AltcoinPlugin) C.AltcoinPlugin__Handle {
	return (C.AltcoinPlugin__Handle)(registerHandle(obj))
}

func lookupWalletHandle(handle C.Wallet__Handle) (*core.Wallet, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Wallet); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletHandle(obj *core.Wallet) C.Wallet__Handle {
	return (C.Wallet__Handle)(registerHandle(obj))
}

func lookupTxnSignerHandle(handle C.TxnSigner__Handle) (*core.TxnSigner, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TxnSigner); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTxnSignerHandle(obj *core.TxnSigner) C.TxnSigner__Handle {
	return (C.TxnSigner__Handle)(registerHandle(obj))
}

func lookupCryptoAccountHandle(handle C.CryptoAccount__Handle) (*core.CryptoAccount, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.CryptoAccount); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerCryptoAccountHandle(obj *core.CryptoAccount) C.CryptoAccount__Handle {
	return (C.CryptoAccount__Handle)(registerHandle(obj))
}

func lookupAddressHandle(handle C.Address__Handle) (*core.Address, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Address); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAddressHandle(obj *core.Address) C.Address__Handle {
	return (C.Address__Handle)(registerHandle(obj))
}

func lookupGenericOutputHandle(handle C.GenericOutput__Handle) (*util.GenericOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.GenericOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerGenericOutputHandle(obj *util.GenericOutput) C.GenericOutput__Handle {
	return (C.GenericOutput__Handle)(registerHandle(obj))
}

func lookupTxnSignerIteratorHandle(handle C.TxnSignerIterator__Handle) (*core.TxnSignerIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TxnSignerIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTxnSignerIteratorHandle(obj *core.TxnSignerIterator) C.TxnSignerIterator__Handle {
	return (C.TxnSignerIterator__Handle)(registerHandle(obj))
}

func lookupTransactionHandle(handle C.Transaction__Handle) (*core.Transaction, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Transaction); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionHandle(obj *core.Transaction) C.Transaction__Handle {
	return (C.Transaction__Handle)(registerHandle(obj))
}

func lookupKeyValueStoreHandle(handle C.KeyValueStore__Handle) (*core.KeyValueStore, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.KeyValueStore); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerKeyValueStoreHandle(obj *core.KeyValueStore) C.KeyValueStore__Handle {
	return (C.KeyValueStore__Handle)(registerHandle(obj))
}

func lookupSimpleWalletOutputHandle(handle C.SimpleWalletOutput__Handle) (*util.SimpleWalletOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.SimpleWalletOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerSimpleWalletOutputHandle(obj *util.SimpleWalletOutput) C.SimpleWalletOutput__Handle {
	return (C.SimpleWalletOutput__Handle)(registerHandle(obj))
}

func lookupTransactionOutputHandle(handle C.TransactionOutput__Handle) (*core.TransactionOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TransactionOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionOutputHandle(obj *core.TransactionOutput) C.TransactionOutput__Handle {
	return (C.TransactionOutput__Handle)(registerHandle(obj))
}

func lookupSimpleWalletAddressHandle(handle C.SimpleWalletAddress__Handle) (*util.SimpleWalletAddress, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.SimpleWalletAddress); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerSimpleWalletAddressHandle(obj *util.SimpleWalletAddress) C.SimpleWalletAddress__Handle {
	return (C.SimpleWalletAddress__Handle)(registerHandle(obj))
}

func lookupBalanceSnapshotHandle(handle C.BalanceSnapshot__Handle) (*util.BalanceSnapshot, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.BalanceSnapshot); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBalanceSnapshotHandle(obj *util.BalanceSnapshot) C.BalanceSnapshot__Handle {
	return (C.BalanceSnapshot__Handle)(registerHandle(obj))
}
