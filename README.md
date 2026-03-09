# CalculatorGO-Postgres
Калькулятор, принимает json с полями: first_number, sign, second_number

Выполняет указанную операцию, если знак валидный (+, -, *, /) и нет ошибки

Возвращает json с полями: result, error (если произошла ошибка)

Также записывает операцию в postgresql таблицу calculator с полями id, first_number, sign, second_number, result
