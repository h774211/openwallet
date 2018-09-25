/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package manager

import (
	"fmt"
)

func (wm *WalletManager) RescanBlockHeight(symbol string, startHeight uint64, endHeight uint64) error {

	assetsMgr, err := GetAssetsManager(symbol)
	if err != nil {
		return err
	}

	scanner := assetsMgr.GetBlockScanner()

	if scanner == nil {
		return fmt.Errorf("%s is not support block scan", symbol)
	}

	if startHeight <= endHeight {
		for i := startHeight;i<=endHeight;i++ {
			err := scanner.ScanBlock(i)
			if err != nil {
				continue
			}
		}
	} else {
		return fmt.Errorf("start block height: %d is greater than end block height: %d", startHeight, endHeight)
	}

	return nil
}