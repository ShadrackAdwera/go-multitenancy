package api

import (
	"fmt"
	"net/http"

	db "github.com/ShadrackAdwera/go-multitenancy/internal/db/sqlc"
	"github.com/ShadrackAdwera/go-utils/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateTenantArgs struct {
	Username   string `json:"username" binding:"required,alphanum"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	TenantName string `json:"tenant_name" binding:"required,min=6"`
}

func (srv *Server) createTenant(ctx *gin.Context) {
	var createTenantArgs CreateTenantArgs

	if err := ctx.ShouldBindJSON(&createTenantArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	pw, _ := utils.HashPassword(createTenantArgs.Password)

	tenantUser, err := srv.store.CreateTenantUserTx(ctx, db.TenantUserTxInput{
		TenantData: db.CreateTenantParams{
			Name: createTenantArgs.TenantName,
			Logo: pgtype.Text{
				String: "",
				Valid:  false,
			},
		},
		UserData: db.CreateUserArgs{
			Username: createTenantArgs.Username,
			Email:    createTenantArgs.Email,
			Password: pw,
		},
	})

	// distribute to message queue to create tenant db and add to connection pool

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("tenant %s has been created", tenantUser.Tenant.Name),
	})
}
