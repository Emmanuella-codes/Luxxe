package messages

import shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"

// -
const FAILURE_USER_NOT_FETCHED shared.PipeMessage = "failure_user_fetched"
const NOT_FOUND_USER shared.PipeMessage = "not_found_user"
const FAIL_GET_USER shared.PipeMessage = "fail_get_user"
const FAIL_UPDATE_USER shared.PipeMessage = "fail_update_user"

// +
const SUCCESS_GET_USER shared.PipeMessage = "success_get_user"
const SUCCESS_UPDATE_USER shared.PipeMessage = "success_update_user"
const SUCCESS_USER_FETCHED shared.PipeMessage = "success_user_fetched"
