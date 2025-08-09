package customerSubscriptionsController

import (
	"fmt"
	"net/http"

	//customerHttpContract "project1.v0/api_gateway/internal/customers-api/controller/subscriptions/test-contracts"

	//"github.com/google/uuid"
	//customerDomainContract "project1.v0/contracts/domain/customer"
	//customerBrokerContract "project1.v0/contracts/transport/broker/customer"

	customerHttpContract "project1.v0/contracts/transport/http/customer"
	mapper "project1.v0/mappers/http_broker/customer"

	mdw "project1.v0/pkg/http/middleware"
	"project1.v0/pkg/http/resp"
)

func NewCustomerSubsciptionsApiController(deps CustomerSubsciptionsApiControllerDeps) {

	handler := CustomerSubsciptionsApiController{
		service:   deps.CustomerSubscriptionService,
		validator: deps.ValidatorX,
		decoder:   deps.DecoderService,
	}

	///////////
	//WEB

	deps.Router.Handle(fmt.Sprintf("POST /%s",
		customerHttpContract.PathSubscriptions,
	), mdw.Chain(
		mdw.HttpValidationWithCtx[
			customerHttpContract.CustomerSubscriptionCreateRequest,
			struct{}, // customerHttpContract.CustomerSubscriptionCreateQuery,
			struct{}, // customerHttpContract.CustomerSubscriptionCreatePath,
		](
			handler.validator,
			handler.decoder,
		),
	)(http.HandlerFunc(handler.CreateSubscription())))

	deps.Router.Handle(
		fmt.Sprintf("PATCH /%s/{%s}",
			customerHttpContract.PathSubscriptions,
			customerHttpContract.PathID,
		), mdw.Chain(
			mdw.HttpValidationWithCtx[
				customerHttpContract.CustomerSubscriptionUpdateRequest,
				struct{},
				customerHttpContract.CustomerSubscriptionUpdatePath,
			](
				handler.validator,
				handler.decoder,
			),
		)(http.HandlerFunc(handler.UpdateSubscription())))

	deps.Router.Handle(fmt.Sprintf("GET /%s",
		customerHttpContract.PathSubscriptions,
	), mdw.Chain(
		mdw.HttpValidationWithCtx[
			struct{},
			customerHttpContract.SubscriptionListQuery,
			struct{},
		](
			handler.validator,
			handler.decoder,
		),
	)(http.HandlerFunc(handler.ListSubscriptions())))

	deps.Router.Handle(fmt.Sprintf("GET /%s/{%s}",
		customerHttpContract.PathSubscriptions,
		customerHttpContract.PathID,
	), mdw.Chain(
		mdw.HttpValidationWithCtx[
			struct{},
			struct{},
			customerHttpContract.CustomerSubscriptionIDPath,
		](
			handler.validator,
			handler.decoder,
		),
	)(http.HandlerFunc(handler.GetSubscription())))

	deps.Router.Handle(fmt.Sprintf("DELETE /%s/{%s}",
		customerHttpContract.PathSubscriptions,
		customerHttpContract.PathID,
	), mdw.Chain(
		mdw.HttpValidationWithCtx[
			struct{},
			struct{},
			customerHttpContract.CustomerSubscriptionIDPath,
		](
			handler.validator,
			handler.decoder,
		),
	)(http.HandlerFunc(handler.DeleteSubscription())))

	deps.Router.Handle(fmt.Sprintf("GET /%s/%s",
		customerHttpContract.PathSubscriptions,
		customerHttpContract.PathTotalCost,
	), mdw.Chain(
		mdw.HttpValidationWithCtx[
			struct{},
			customerHttpContract.SubscriptionTotalQuery,
			struct{},
		](
			handler.validator,
			handler.decoder,
		),
	)(http.HandlerFunc(handler.SubscriptionsTotalCost())))

	///////////
	// NATIVE
}

// CreateSubscription godoc
// @Summary Create subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param payload body customerHttpContract.CustomerSubscriptionCreateRequest true "Payload to create subscription"
// @Success 201 {object} customerHttpContract.CustomerSubscriptionCreateResponse
// @Failure 400 {object} httpContractCommons.ErrorResponse
// @Failure 500 {object} httpContractCommons.ErrorResponse
// @Router /subscriptions [post]
func (h *CustomerSubsciptionsApiController) CreateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := mdw.GetFromContext[customerHttpContract.CustomerSubscriptionCreateRequest](r.Context())

		brokerReq := mapper.CustomerSubscriptionCreateRequest_HttpToBroker(body)

		brokerResp, err := h.service.Create(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		respPayload := &customerHttpContract.CustomerSubscriptionCreateResponse{
			ID: brokerResp.ID,
		}
		resp.Json(w, respPayload, 201)

	}
}

// GetSubscription godoc
// @Summary Get subscription by id
// @Tags subscriptions
// @Produce json
// @Param id path string true "subscription ID (UUID)" example(60601fee-2bf1-4721-ae6f-7636e79a0cba)
// @Success 200 {object} customerHttpContract.CustomerSubscriptionResponse
// @Failure 400 {object} httpContractCommons.ErrorResponse
// @Failure 404 {object} httpContractCommons.ErrorResponse
// @Failure 500 {object} httpContractCommons.ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *CustomerSubsciptionsApiController) GetSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// id := r.PathValue(customerHttpContract.PathID)
		// _, err := uuid.Parse(id)
		// if err != nil {
		// 	code := customerHttpContract.HttpStatusForError(customerDomainContract.ErrInvalidID)
		// 	resp.WriteError(w, customerDomainContract.ErrInvalidID.Error(), code)
		// 	return
		// }
		path, _ := mdw.GetFromContext[customerHttpContract.CustomerSubscriptionIDPath](r.Context())

		brokerReq := mapper.CustomerSubscriptionIDRequest_HttpToBroker(path)

		subscription, err := h.service.Get(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		resp.Json(w, subscription, 200)
	}
}

// ListSubscriptions godoc
// @Summary Get list of subscriptions by user_id with optional filters
// @Tags subscriptions
// @Produce json
// @Param user_id query string true "User ID (UUID)" example(60601fee-2bf1-4721-ae6f-7636e79a0cba)
// @Param service_name query string false "Partial match of service name" example(yandex)
// @Param limit query int false "Page size (max: 100)" default(20) minimum(1) maximum(100)
// @Param offset query int false "Offset for pagination" default(0) minimum(0)
// @Success 200 {object} customerHttpContract.CustomerSubscriptionsListResponse
// @Failure 400 {object} httpContractCommons.ErrorResponse
// @Failure 500 {object} httpContractCommons.ErrorResponse
// @Router /subscriptions [get]
func (h *CustomerSubsciptionsApiController) ListSubscriptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query, _ := mdw.GetFromContext[customerHttpContract.SubscriptionListQuery](r.Context())

		brokerReq := mapper.CustomerSubscriptionListRequest_HttpToBroker(query)

		subs, err := h.service.List(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		resp.Json(w, subs.Subscriptions, 200)
	}
}

// DeleteSubscription godoc
// @Summary Delete subscription by ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "subscription ID (UUID)" example(4b70f8d6-c702-4e6e-9c65-2ae0e3bb0e5c)
// @Success 204 "no content"
// @Failure 400 {object} httpContractCommons.ErrorResponse
// @Failure 500 {object} httpContractCommons.ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *CustomerSubsciptionsApiController) DeleteSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// id := r.PathValue(customerHttpContract.PathID)
		// _, err := uuid.Parse(id)
		// if err != nil {
		// 	code := customerHttpContract.HttpStatusForError(customerDomainContract.ErrInvalidID)
		// 	resp.WriteError(w, customerDomainContract.ErrInvalidID.Error(), code)
		// 	return
		// }
		path, _ := mdw.GetFromContext[customerHttpContract.CustomerSubscriptionIDPath](r.Context())

		brokerReq := mapper.CustomerSubscriptionIDRequest_HttpToBroker(path)

		_, err := h.service.Delete(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204, without body
	}
}

// UpdateSubscription godoc
// @Summary Update subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "subscription ID (UUID)" example(4b70f8d6-c702-4e6e-9c65-2ae0e3bb0e5c)
// @Param payload body customerHttpContract.CustomerSubscriptionUpdateRequest true "Fields to update"
// @Success 204 "no content"
// @Failure 400 {object} httpContractCommons.ErrorResponse
// @Failure 500 {object} httpContractCommons.ErrorResponse
// @Router /subscriptions/{id} [patch]
func (h *CustomerSubsciptionsApiController) UpdateSubscription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := mdw.GetFromContext[customerHttpContract.CustomerSubscriptionUpdateRequest](r.Context())
		path, _ := mdw.GetFromContext[customerHttpContract.CustomerSubscriptionUpdatePath](r.Context())

		brokerReq := mapper.CustomerSubscriptionUpdateRequest_HttpToBroker(*path.ID, body)

		_, err := h.service.Update(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *CustomerSubsciptionsApiController) SubscriptionsTotalCost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query, _ := mdw.GetFromContext[customerHttpContract.SubscriptionTotalQuery](r.Context())

		brokerReq := mapper.CustomerSubscriptionTotalCostRequest_HttpToBroker(query)

		brokerResp, err := h.service.TotalCost(brokerReq)
		if err != nil {
			code := customerHttpContract.HttpStatusForError(err)
			resp.WriteError(w, err.Error(), code)
			return
		}
		resp.Json(w, brokerResp, 200)
	}
}
