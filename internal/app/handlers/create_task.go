package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/GbSouza15/apiToDoGo/internal/authenticator"
	"github.com/google/uuid"
)

func (h handler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.SendResponse(404, []byte("Erro ao ler o corpo da requisição."), w)
		return
	}

	var newTask models.TaskCreate
	var taskId = uuid.NewString()

	if err := json.Unmarshal(body, &newTask); err != nil {
		response.SendResponse(500, []byte("Erro ao decodificação do JSON"), w)
		return
	}

	userId := authenticator.UserIDFromContext(r.Context())

	_, err = h.DB.Exec("INSERT INTO tdlist.tasks (id, title, description, user_id) VALUES ($1, $2, $3, $4)", taskId, newTask.Title, newTask.Description, userId)
	if err != nil {
		fmt.Println(err)
		fmt.Println(newTask)
		response.SendResponse(500, []byte("Erro ao criar tarefa."), w)
		return
	}

	response.SendResponse(201, []byte("Tarefa criada com sucesso."), w)
}
