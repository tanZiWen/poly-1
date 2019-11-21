/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package side_chain_manager

import (
	"fmt"
	"github.com/ontio/multi-chain/common"
	cstates "github.com/ontio/multi-chain/core/states"
	"github.com/ontio/multi-chain/native"
	"github.com/ontio/multi-chain/native/service/utils"
	"math"
)

func getRegisterSideChain(native *native.NativeService, chanid uint64) (*SideChain, error) {
	contract := utils.SideChainManagerContractAddress
	chainidByte := utils.GetUint64Bytes(chanid)

	sideChainStore, err := native.GetCacheDB().Get(utils.ConcatKey(contract, []byte(REGISTER_SIDE_CHAIN_REQUEST),
		chainidByte))
	if err != nil {
		return nil, fmt.Errorf("getRegisterSideChain,get registerSideChainRequestStore error: %v", err)
	}
	sideChain := &SideChain{
		ChainId: math.MaxUint64,
	}
	if sideChainStore != nil {
		sideChainBytes, err := cstates.GetValueFromRawStorageItem(sideChainStore)
		if err != nil {
			return nil, fmt.Errorf("getRegisterSideChain, deserialize from raw storage item err:%v", err)
		}
		if err := sideChain.Deserialization(common.NewZeroCopySource(sideChainBytes)); err != nil {
			return nil, fmt.Errorf("getRegisterSideChain, deserialize sideChain error: %v", err)
		}
	}
	return sideChain, nil
}

func putRegisterSideChain(native *native.NativeService, sideChain *SideChain) error {
	contract := utils.SideChainManagerContractAddress
	chainidByte := utils.GetUint64Bytes(sideChain.ChainId)

	sink := common.NewZeroCopySink(nil)
	err := sideChain.Serialization(sink)
	if err != nil {
		return fmt.Errorf("putRegisterSideChain, sideChain.Serialization error: %v", err)
	}

	native.GetCacheDB().Put(utils.ConcatKey(contract, []byte(REGISTER_SIDE_CHAIN_REQUEST), chainidByte),
		cstates.GenRawStorageItem(sink.Bytes()))
	return nil
}

func GetSideChain(native *native.NativeService, chainID uint64) (*SideChain, error) {
	contract := utils.SideChainManagerContractAddress
	chainIDByte := utils.GetUint64Bytes(chainID)

	sideChainStore, err := native.GetCacheDB().Get(utils.ConcatKey(contract, []byte(SIDE_CHAIN),
		chainIDByte))
	if err != nil {
		return nil, fmt.Errorf("getSideChain,get registerSideChainRequestStore error: %v", err)
	}
	sideChain := &SideChain{
		ChainId: math.MaxUint64,
	}
	if sideChainStore != nil {
		sideChainBytes, err := cstates.GetValueFromRawStorageItem(sideChainStore)
		if err != nil {
			return nil, fmt.Errorf("getSideChain, deserialize from raw storage item err:%v", err)
		}
		if err := sideChain.Deserialization(common.NewZeroCopySource(sideChainBytes)); err != nil {
			return nil, fmt.Errorf("getSideChain, deserialize sideChain error: %v", err)
		}
	}
	return sideChain, nil
}

func putSideChain(native *native.NativeService, sideChain *SideChain) error {
	contract := utils.SideChainManagerContractAddress
	chainidByte := utils.GetUint64Bytes(sideChain.ChainId)

	sink := common.NewZeroCopySink(nil)
	err := sideChain.Serialization(sink)
	if err != nil {
		return fmt.Errorf("putSideChain, sideChain.Serialization error: %v", err)
	}

	native.GetCacheDB().Put(utils.ConcatKey(contract, []byte(SIDE_CHAIN), chainidByte),
		cstates.GenRawStorageItem(sink.Bytes()))
	return nil
}

func getUpdateSideChain(native *native.NativeService, chanid uint64) (*SideChain, error) {
	contract := utils.SideChainManagerContractAddress
	chainidByte := utils.GetUint64Bytes(chanid)

	sideChainStore, err := native.GetCacheDB().Get(utils.ConcatKey(contract, []byte(UPDATE_SIDE_CHAIN_REQUEST),
		chainidByte))
	if err != nil {
		return nil, fmt.Errorf("getUpdateSideChain,get registerSideChainRequestStore error: %v", err)
	}
	sideChain := new(SideChain)
	if sideChainStore != nil {
		sideChainBytes, err := cstates.GetValueFromRawStorageItem(sideChainStore)
		if err != nil {
			return nil, fmt.Errorf("getUpdateSideChain, deserialize from raw storage item err:%v", err)
		}
		if err := sideChain.Deserialization(common.NewZeroCopySource(sideChainBytes)); err != nil {
			return nil, fmt.Errorf("getUpdateSideChain, deserialize sideChain error: %v", err)
		}
	}
	return sideChain, nil
}

func putUpdateSideChain(native *native.NativeService, sideChain *SideChain) error {
	contract := utils.SideChainManagerContractAddress
	chainidByte := utils.GetUint64Bytes(sideChain.ChainId)

	sink := common.NewZeroCopySink(nil)
	err := sideChain.Serialization(sink)
	if err != nil {
		return fmt.Errorf("putUpdateSideChain, sideChain.Serialization error: %v", err)
	}

	native.GetCacheDB().Put(utils.ConcatKey(contract, []byte(UPDATE_SIDE_CHAIN_REQUEST), chainidByte),
		cstates.GenRawStorageItem(sink.Bytes()))
	return nil
}