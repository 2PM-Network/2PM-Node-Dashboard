/*
 * Copyright 2021 Seven Seals Technology
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package user

import (
	"deltaboard-server/api/v1/response"
	"deltaboard-server/internal/service/current_user"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) (*response.Response, error) {
	currentUserInterface, _ := ctx.Get("CurrentUser")
	currentUser := currentUserInterface.(*current_user.CurrentUser)
	return &response.Response{}, currentUser.RemoveUserLoginSession(ctx.Request, ctx.Writer)
}
