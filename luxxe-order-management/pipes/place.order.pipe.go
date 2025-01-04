package pipes

import (
	"context"
	"fmt"

	cart_messages "github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	order_messages "github.com/Emmanuella-codes/Luxxe/luxxe-order-management/messages"
	product_messages "github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	user_messages "github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	order_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/order"
	product_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	transaction_dtos "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	transaction_messages "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/messages"
	transaction_pipes "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/pipes"
	transaction_services "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/services"
	transaction_types "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/types"
)

type PlaceOrderResponse struct {
	shouldRedirect 							bool
	Transaction    							*entities.Transaction
	PaystackTransactionResponse *transaction_types.PaystackInitializeTransactionResponseData
}

func PlaceOrderPipe(ctx context.Context, dto *dtos.PlaceOrderDTO) *shared.PipeRes[PlaceOrderResponse] {
	// checks order exists and status is pending
	// checks for the cartID within and checks if the products quantity is sufficient to place order
	// initialize paystack transaction
	// on success, reduce the product quantity that were in the cart/order

	user, err := repo_user.UserRepo.QueryByID(ctx, dto.UserID)
	if err != nil {
		return &shared.PipeRes[PlaceOrderResponse]{
			Success: false,
			Message: user_messages.FAIL_GET_USER,
		}
	}
	
	order, err := order_repo.OrderRepo.QueryByID(ctx, dto.OrderID)
	if err != nil || order.OrderStatus != entities.OrderStatusPending {
		return &shared.PipeRes[PlaceOrderResponse]{
			Success: false,
			Message: order_messages.NOT_FOUND_ORDER,
		}
	}

	cart, err := cart_repo.CartRepo.QueryByID(ctx, order.CartID.Hex())
	if err != nil {
		return &shared.PipeRes[PlaceOrderResponse]{
			Success: false,
			Message: cart_messages.NOT_FOUND_CART,
		}
	}

	for _, item := range cart.Items {
		product, err := product_repo.ProductRepo.QueryByID(ctx, item.ProductID.Hex())
		if err != nil {
			return &shared.PipeRes[PlaceOrderResponse]{
				Success: false,
				Message: product_messages.NOT_FOUND_PRODUCT,
			}
		}

		if product.Quantity < item.Quantity {
			return &shared.PipeRes[PlaceOrderResponse]{
				Success: false,
				Message: product_messages.INSUFFICIENT_PRODUCT_QUANTITY,
			}
		}
	}

	transactionDTO := &transaction_dtos.CreateTransactionDTO{
		PmtTransactionCtx:          	entities.PmtTxCtxOrder,
		PmtTransactionCtxID:        	order.ID.Hex(),
		PmtTransactionInitiatorCtx: 	entities.PmtTxInitCtxUser,
		PmtTransactionInitiatorCtxID: user.ID.Hex(),
		PaymentClient:                dto.PaymentClient,
		Amount:                       order.CartTotal,
		Meta:                         map[string]interface{}{"orderID": order.ID.Hex()},
	}

	txPipeRes := transaction_pipes.CreateTransactionPipe(ctx, transactionDTO)
	if !txPipeRes.Success {
		return &shared.PipeRes[PlaceOrderResponse]{
			Success: false,
			Message: transaction_messages.FAIL_CREATE_TRANSACTION,
		}
	}

	transaction := txPipeRes.Data

	response := PlaceOrderResponse{
		Transaction: transaction,
		shouldRedirect: true,
	}

	if dto.PaymentClient == entities.PmtClientPaystack {
		pystkRes, err := transaction_services.InitializePaystackTransactionService(transaction, order)
		if err != nil || pystkRes.AuthorizationUrl == "" {
			fmt.Printf("Paystack initialization failed: %v", err)
			return &shared.PipeRes[PlaceOrderResponse]{
				Success: false,
				Message: transaction_messages.FAIL_INITIALIZE_TRANSACTION,
			}
		}

		response.PaystackTransactionResponse = pystkRes
	}

	order.OrderStatus = entities.OrderStatusProcessing
	_, err = order_repo.OrderRepo.UpdateOrder(ctx, order)
	if err != nil {
		fmt.Printf("failed to update order status: %v", err)
	}

	return &shared.PipeRes[PlaceOrderResponse]{
		Success: true,
		Message: order_messages.SUCCESS_PLACE_ORDER,
		Data:    &response,
	}
}
