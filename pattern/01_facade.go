package main

import "fmt"

type EmailService struct{}

// Отправка e-mail
func (e *EmailService) SendEmail(to, subject, message string) {
	fmt.Printf("Отправка e-mail клиенту %s, тема %s,  сообщение: %s\n", to, subject, message)
}

type SMSService struct{}

// Отправка СМС
func (s *SMSService) SendSMS(to, message string) {
	fmt.Printf("Отправка СМС клиенту %s, сообщение: %s\n", to, message)
}

type PushService struct{}

// Отправка push-уведомления
func (p *PushService) SendPushNotification(to, message string) {
	fmt.Printf("Отправка push-уведомления клиенту %s, сообщение: %s\n", to, message)
}

// Фасад, предоставляющий простой интерфейс для отправки уведомлений
type NotificationFacade struct {
	EmailService            *EmailService
	SMSService              *SMSService
	PushNotificationService *PushService
}

// Метод, который использует три способа отправки уведомлений
func (n *NotificationFacade) SendNotification(to, subject, message string) {
	n.EmailService.SendEmail(to, subject, message)
	n.SMSService.SendSMS(to, message)
	n.PushNotificationService.SendPushNotification(to, message)
}

func main() {
	emailService := &EmailService{}
	smsService := &SMSService{}
	pushNotificationService := &PushService{}
	// Создание фасада
	notificationFacade := &NotificationFacade{
		EmailService:            emailService,
		SMSService:              smsService,
		PushNotificationService: pushNotificationService,
	}
	// Использование метода фасада
	notificationFacade.SendNotification("Иванов Иван Иванович", "Приветствие", "Привет")
}

/* Паттерн "фасад" слудет применять в том случае, когда клиент взаимодействует со сложной модульной системой. С помощью "фасада" можно
облегчить взаимодействие путём предоставления простого интерфейса для использования функциональности системы
Плюсы паттерна:
- Упрощает сложность системы для клиента, предоставляя простой интерфейс.

Минусы паттерна:
- Может привести к избыточности, если разрабатывается слишком много методов в фасаде.
- Может ограничивать гибкость, т.к. добавление нового функционала требует изменений в фасаде.
*/
