package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/businesslogic"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

func PostUserRegister(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: Регистрация пользователя
	// Хендлер: POST /api/user/register.
	// Регистрация производится по паре логин/пароль. Каждый логин должен быть уникальным.
	// После успешной регистрации должна происходить автоматическая аутентификация пользователя.
	// Content-Type: application/json
	// {
	//  "login": "<login>",
	//     "password": "<password>"
	// }
	// Возможные коды ответа:
	// 200 — пользователь успешно зарегистрирован и аутентифицирован;
	// 400 — неверный формат запроса;
	// 409 — логин уже занят;
	// 500 — внутренняя ошибка сервера.

	var newuser storage.User

	// decode JSON
	config.LoggerCLS.Sugar().Debugf("Body=%v", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		config.LoggerCLS.Info(err.Error())
		return
	}

	config.LoggerCLS.Sugar().Debugf("newuser=%v", newuser)

	// call business logic
	bs := new(businesslogic.BusinessSession)
	err := bs.RegisterNewUser(newuser)
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail create new user </h1>" + newuser.Login))
		return
	}

	// return answer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("<h1>Created new user </h1>" + newuser.Login))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func PostUserLogin(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/login.
	// Аутентификация производится по паре логин/пароль
	//
	var user storage.User

	// decode JSON
	config.LoggerCLS.Sugar().Debugf("Body=%v", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		config.LoggerCLS.Info(err.Error())
		return
	}

	config.LoggerCLS.Sugar().Debugf("newuser=%v", user)

	// call business logic
	bs := new(businesslogic.BusinessSession)
	err := bs.UserLogin(user)
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail login user </h1>" + user.Login))
		return
	}

	// return answer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("<h1>User logged in </h1>" + user.Login))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func PostUserOrders(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/orders.
	// Хендлер доступен только аутентифицированным пользователям.
	// Номером заказа является последовательность цифр произвольной длины.
	// Номер заказа может быть проверен на корректность ввода с помощью алгоритма Луна
	//

	var ordercode string

	b := make([]byte, 100)

	// decode orderstring
	r.Body.Read(b)
	ordercode = string(b)
	config.LoggerCLS.Sugar().Debugf("body=%v ordercode=%v", b, ordercode)

	// call business logic
	bs := new(businesslogic.BusinessSession)
	err := bs.LoadOrder(ordercode)
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail load order code </h1>" + ordercode))
		return
	}

	// return answer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("<h1>Loaded order code </h1>" + ordercode))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func GetUserOrders(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: GET /api/user/orders.
	// Хендлер доступен только авторизованному пользователю. Номера заказа в выдаче должны быть отсортированы по времени загрузки от самых старых к самым новым
	//
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: GET /api/user/balance.
	// Хендлер доступен только авторизованному пользователю. В ответе должны содержаться данные о текущей сумме баллов лояльности, а также сумме использованных за весь период регистрации баллов
	//
}

func PostUserBalanceWithdraw(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/balance/withdraw
	// Хендлер доступен только авторизованному пользователю. Номер заказа представляет собой гипотетический номер нового заказа пользователя, в счёт оплаты которого списываются баллы.
	// Примечание: для успешного списания достаточно успешной регистрации запроса, никаких внешних систем начисления не предусмотрено и не требуется реализовывать
	//
}

func GetUserBalanceWithdrawals(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: GET /api/user/balance/withdrawals.
	// Хендлер доступен только авторизованному пользователю. Факты выводов в выдаче должны быть отсортированы по времени вывода от самых старых к самым новым
	//
}