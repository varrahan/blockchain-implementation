package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) GetBlocks_Get_Handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(ctx.Blockchain.Blocks)
}
