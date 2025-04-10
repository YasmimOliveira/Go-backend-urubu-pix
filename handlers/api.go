package handlers

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "backend-urubu-do-pix/models"
    "backend-urubu-do-pix/utils"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.DepositDate = time.Now().Format(time.RFC3339)
    user.Balance = user.InitialAmount

    err = user.Save()
    if err != nil {
        http.Error(w, "Erro ao salvar usuário no banco de dados", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user, err := models.GetByID(id)
    if err != nil {
        http.Error(w, "Erro ao buscar usuário no banco de dados", http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, "Usuário não encontrado", http.StatusNotFound)
        return
    }

    depositDate, err := user.GetDepositDateAsTime()
    if err != nil {
        http.Error(w, "Erro ao converter data do depósito", http.StatusInternalServerError)
        return
    }

    daysPassed := utils.DaysSince(depositDate)
    balance := utils.CalculateBalance(user.InitialAmount, daysPassed, 0.3333)

    user.Balance = balance
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user, err := models.GetByID(id)
    if err != nil {
        http.Error(w, "Erro ao buscar usuário no banco de dados", http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, "Usuário não encontrado", http.StatusNotFound)
        return
    }

    depositDate, err := user.GetDepositDateAsTime()
    if err != nil {
        http.Error(w, "Erro ao converter data do depósito", http.StatusInternalServerError)
        return
    }

    daysPassed := utils.DaysSince(depositDate)
    if daysPassed < 30 {
        http.Error(w, "Ainda não se passaram 30 dias desde o depósito", http.StatusForbidden)
        return
    }

    query := `DELETE FROM users WHERE id = ?`
    _, err = utils.DB.Exec(query, id)
    if err != nil {
        http.Error(w, "Erro ao remover usuário do banco de dados", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Saque realizado com sucesso!"))
}