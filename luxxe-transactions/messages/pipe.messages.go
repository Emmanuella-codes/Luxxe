package messages

import shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"

// -
const NOT_FOUND_TRANSACTION shared.PipeMessage = "not_found_transaction"
const FAIL_CREATE_TRANSACTION shared.PipeMessage = "fail_create_transaction"
const FAIL_VERIFY_TRANSACTION shared.PipeMessage = "fail_verify_transaction"
const FAIL_INITIALIZE_TRANSACTION shared.PipeMessage = "fail_initialize_transaction"
const FAIL_GET_TRANSACTION_HISTORY shared.PipeMessage = "fail_get_transaction_history"
const FAIL_UPDATE_TRANSACTION_STATUS shared.PipeMessage = "fail_update_transaction_status"

// +
const SUCCESS_CREATE_TRANSACTION shared.PipeMessage = "success_create_transaction"
const SUCCESS_VERIFY_TRANSACTION shared.PipeMessage = "success_verify_transaction"
const SUCCESS_GET_TRANSACTION_LIST shared.PipeMessage = "success_get_transaction_list"
const SUCCESS_GET_TRANSACTION_HISTORY shared.PipeMessage = "success_get_transaction_history"
