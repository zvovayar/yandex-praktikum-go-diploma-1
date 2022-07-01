package handlers

import "net/http"

func PostUserRegister(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: Регистрация пользователя
	// Хендлер: POST /api/user/register.
	// Регистрация производится по паре логин/пароль. Каждый логин должен быть уникальным. После успешной регистрации должна происходить автоматическая аутентификация пользователя.
	//

}

func PostUserLogin(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/login.
	// Аутентификация производится по паре логин/пароль
	//
}

func PostUserOrders(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: POST /api/user/orders.
	// Хендлер доступен только аутентифицированным пользователям. Номером заказа является последовательность цифр произвольной длины.
	// Номер заказа может быть проверен на корректность ввода с помощью алгоритма Луна
	//
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
