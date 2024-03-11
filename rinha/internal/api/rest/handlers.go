package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"rinha/internal/banking"
	"rinha/internal/banking/service"
	"strconv"
)

func accountTransactionHandler(ctx context.Context, accountService service.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accountID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || accountID <= 0 || accountID > 5 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "user not found"}`))
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "failed to read request body"}`))
			return
		}

		var tx banking.RequestTransaction
		err = json.Unmarshal(b, &tx)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "malformed values"}`))
			return
		}

		var operation banking.TxType
		if len([]rune(tx.Type)) != 1 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "invalid transaction operation"}`))
			return
		}

		operation = banking.TxType([]rune(tx.Type)[0])
		if operation != banking.Credit && operation != banking.Debt {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "invalid transaction operation"}`))
			return
		}

		if tx.Amount <= 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "invalid amount"}`))
			return
		}

		if len(tx.Description) > 10 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "description exceeds the maximum length"}`))
			return
		}

		var resp service.ResponseTransaction
		switch operation {
		case banking.Credit:
			resp, err = accountService.Credit(ctx, accountID, tx.Amount, tx.Description)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{"error": "unable to process credit operation"}`))
				return
			}
		case banking.Debt:
			resp, err = accountService.Debt(ctx, accountID, tx.Amount, tx.Description)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{"error": "unable to process debt operation"}`))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func accountStatementHandler(ctx context.Context, accountService service.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accountID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || accountID <= 0 || accountID > 5 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "user not found"}`))
			return
		}

		accStmt, err := accountService.GenerateStatement(ctx, accountID)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(accStmt)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "failed to encode response into body"}`))
			return
		}
	}
}
