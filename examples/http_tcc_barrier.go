/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/go-resty/resty/v2"
)

func init() {
	AddCommand("http_tcc_barrier", func() string {
		logger.Debugf("tcc transaction begin")
		gid := dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)
		err := dtmcli.TccGlobalTransaction(dtmutil.DefaultHTTPServer, gid, func(tcc *dtmcli.Tcc) (*resty.Response, error) {
			resp, err := tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TccBTransOutTry",
				busi.Busi+"/TccBTransOutConfirm", busi.Busi+"/TccBTransOutCancel")
			if err != nil {
				return resp, err
			}
			return tcc.CallBranch(&busi.TransReq{Amount: 30}, busi.Busi+"/TccBTransInTry", busi.Busi+"/TccBTransInConfirm", busi.Busi+"/TccBTransInCancel")
		})
		logger.FatalIfError(err)
		return gid
	})
}
