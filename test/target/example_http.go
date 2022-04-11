package target

import (
	"damascus/internal/invoice"
	"damascus/internal/order"
	"damascus/internal/plan"
	"damascus/internal/va"
	"damascus/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	s   Service
	ivs invoice.Service
	vas va.Service
	pls plan.Service
	ors order.Service
}

func NewHandler(s Service, ivs invoice.Service, vas va.Service, pls plan.Service, ors order.Service) *handler {
	return &handler{
		s,
		ivs,
		vas,
		pls,
		ors,
	}
}

func (h *handler) RegisterRoute(r *mux.Router) {
	r.HandleFunc("/supported_payments", h.supportedPayments).Methods("GET")
	r.HandleFunc("/checkout", h.checkout).Methods("POST")
}

func (h *handler) supportedPayments(w http.ResponseWriter, r *http.Request) {
	result := h.s.GetSupportedPayments()

	response.Success(w, http.StatusOK, result)
}

type CheckoutRequest struct {
	StudentID              int    `json:"student_id"`
	SaleableType           string `json:"saleable_type"`
	SaleableID             int    `json:"saleable_id"`
	SupportedPaymentPrefix string `json:"supported_payment_prefix"`
	SupportedPaymentItemID int    `json:"supported_payment_item_id"`
}

func (h *handler) checkout(w http.ResponseWriter, r *http.Request) {

	var cr CheckoutRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		response.Success(w, http.StatusOK, "Invalid payload")
		return
	}
	defer r.Body.Close()

	if cr.SaleableType != "plans" {
		response.Error(w, http.StatusNotImplemented, "Sorry we are not supporting another purchase instead of plans")
		return
	}

	supPaymentItem := h.s.GetSupportedPaymentItem(cr.SupportedPaymentItemID)

	purchasedItem := h.pls.GetPlan(cr.SaleableID)
	// TODO: convert this to transaction mode

	// create order
	// create payment_virtual_accounts
	// create payment
	// create invoice

	order := order.CreateOrder{
		StudentID:    cr.StudentID,
		SaleableType: cr.SaleableType,
		SaleableID:   cr.SaleableID,
	}
	orderInsertedID, err := h.ors.CreateOrder(order)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Failed when create order")
	}

	va := va.PaymentVirtualAccount{
		ExternalID:    "testing",
		BankCode:      supPaymentItem.Name,
		AccountNumber: "testing",
		Amount:        purchasedItem.Price,
		Status:        va.PaymentVirtualAccountStatus_Pending,
	}

	vaInsertedID, err := h.vas.CreateVA(va)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Failed when create va")
	}

	payment := Payment{
		PaymentableType: cr.SupportedPaymentPrefix,
		PaymentableID:   vaInsertedID,
		Status:          "PENDING",
	}

	paymentInsertedID, err := h.s.CreatePayment(payment)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Failed when create payment")
	}

	invoice := invoice.CreateInvoice{
		OrderID:   orderInsertedID,
		PaymentID: paymentInsertedID,
	}

	err = h.ivs.CreateInvoice(invoice)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Failed when create invoice")
	}

	// TODO: response with http success with no content
	message := map[string]string{
		"message": "checkout success",
	}
	response.Success(w, http.StatusOK, message)
}
