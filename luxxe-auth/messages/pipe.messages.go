package messages

import shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"

const NOT_REGISTERED_EMAIL shared.PipeMessage = "not_registered_email"
const NOT_FOUND_USER shared.PipeMessage = "not_found_user"
const INCORRECT_EXPIRED_OTP shared.PipeMessage = "incorrect_expired_otp"
const EXISTING_ACCOUNT_REGISTERED_EMAIL shared.PipeMessage = "existing_account_registered_email"
const INCORRECT_PASSWORD_EMAIL shared.PipeMessage = "incorrect_password_email"
const CANNOT_SEND_MAIL_TO_ANOTHER_USERS_ACCOUNT shared.PipeMessage = "cannot_send_mail_to_another_users_account"
const CANNOT_VERIFY_ANOTHER_USERS_EMAIL shared.PipeMessage = "cannot_verify_another_users_email"

// +
const SENT_OTP_EMAIL shared.PipeMessage = "sent_otp_email"
const SUCCESS_CHANGE_PASSWORD shared.PipeMessage = "success_change_password"
const CREATED_NEW_ACCOUNT shared.PipeMessage = "created_new_account"
const SUCCESS_SIGN_IN shared.PipeMessage = "success_sign_in"
const SUCCESS_MAIL_VERIFIED shared.PipeMessage = "success_sign_in"