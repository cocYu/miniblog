// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package errno

var (
	// ErrUserAlreadyExists 用户已存在
	ErrUserAlreadyExists = &Errno{Code: "FailedOperation.UserAlreadyExist", Message: "User already exist.", HTTP: 400}
)
