package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/businesslogic"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

type Token struct {
	Token string `json:"token"`
}

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

	claims := make(map[string]interface{})
	claims["user_id"] = newuser.Login
	claims["exp"] = jwtauth.ExpireIn(time.Minute * time.Duration(config.ConfigCLS.TokenTimountMinutes))

	config.LoggerCLS.Sugar().Debugf("claims=%v", claims)

	_, tokenString, _ := TokenAuth.Encode(claims)

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("user registered: " + newuser.Login))
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

	config.LoggerCLS.Sugar().Debugf("user=%v", user)

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
	claims := make(map[string]interface{})
	claims["user_id"] = user.Login
	claims["exp"] = jwtauth.ExpireIn(time.Minute * time.Duration(config.ConfigCLS.TokenTimountMinutes))

	config.LoggerCLS.Sugar().Debugf("user logged in claims=%v", claims)

	_, tokenString, _ := TokenAuth.Encode(claims)

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	w.Header().Set("Content-Type", "text/plain")

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

	b := make([]byte, 300)

	// decode orderstring
	n, _ := r.Body.Read(b)
	ordercode = string(b[:n])
	config.LoggerCLS.Sugar().Debugf("body=%v ordercode=%v", b[:n], ordercode)

	// decode climes from JWT
	_, claims, _ := jwtauth.FromContext(r.Context())
	config.LoggerCLS.Debug(fmt.Sprintf("JWT for user %v recieved", claims["user_id"]))

	// call business logic
	bs := new(businesslogic.BusinessSession)
	err := bs.LoadOrder(ordercode, fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail load order code </h1>" + ordercode + " " + err.Error()))
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
	// Хендлер доступен только авторизованному пользователю.
	// Номера заказа в выдаче должны быть отсортированы по времени загрузки от самых старых к самым новым
	//

	// load data from service
	// call business logic

	// decode climes from JWT
	_, claims, _ := jwtauth.FromContext(r.Context())
	config.LoggerCLS.Debug(fmt.Sprintf("JWT for user %v recieved", claims["user_id"]))

	bs := new(businesslogic.BusinessSession)
	json, err := bs.GetOrders(fmt.Sprintf("%v", claims["user_id"]))

	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail get orders </h1>"))
		return
	}

	// return answer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(json))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: GET /api/user/balance.
	// Хендлер доступен только авторизованному пользователю.
	// В ответе должны содержаться данные о текущей сумме баллов лояльности,
	// а также сумме использованных за весь период регистрации баллов
	//
	// load data from service
	// call business logic

	// decode climes from JWT
	_, claims, _ := jwtauth.FromContext(r.Context())
	config.LoggerCLS.Debug(fmt.Sprintf("JWT for user %v recieved", claims["user_id"]))

	bs := new(businesslogic.BusinessSession)
	jsonb, err := bs.GetBalance(fmt.Sprintf("%v", claims["user_id"]))

	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail get balance </h1>"))
		return
	}

	// return answer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonb)
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func PostUserBalanceWithdraw(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/balance/withdraw
	// Хендлер доступен только авторизованному пользователю.
	// Номер заказа представляет собой гипотетический номер нового заказа пользователя, в счёт оплаты которого списываются баллы.
	// Примечание: для успешного списания достаточно успешной регистрации запроса,
	// никаких внешних систем начисления не предусмотрено и не требуется реализовывать
	//
	var withdraw storage.Withdraw

	// decode JSON
	config.LoggerCLS.Sugar().Debugf("Body=%v", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&withdraw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		config.LoggerCLS.Info(err.Error())
		return
	}

	config.LoggerCLS.Sugar().Debugf("withdraw=%v", withdraw)

	// decode climes from JWT
	_, claims, _ := jwtauth.FromContext(r.Context())
	config.LoggerCLS.Debug(fmt.Sprintf("JWT for user %v recieved", claims["user_id"]))

	// call business logic
	bs := new(businesslogic.BusinessSession)
	err := bs.Withdraw(withdraw, fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail withdraw </h1>"))
		return
	}

	// return answer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("<h1>Withdraw registered </h1>"))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}

func GetUserBalanceWithdrawals(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: GET /api/user/balance/withdrawals.
	// Хендлер доступен только авторизованному пользователю.
	// Факты выводов в выдаче должны быть отсортированы по времени вывода от самых старых к самым новым
	//

	// decode climes from JWT
	_, claims, _ := jwtauth.FromContext(r.Context())
	config.LoggerCLS.Debug(fmt.Sprintf("JWT for user %v recieved", claims["user_id"]))

	bs := new(businesslogic.BusinessSession)
	json, err := bs.GetWithdrawals(fmt.Sprintf("%v", claims["user_id"]))

	if err != nil {
		config.LoggerCLS.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("<h1>Fail get withdrawals balance </h1>"))
		return
	}

	// return answer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(json))
	if err != nil {
		config.LoggerCLS.Info(err.Error())
	}
}
