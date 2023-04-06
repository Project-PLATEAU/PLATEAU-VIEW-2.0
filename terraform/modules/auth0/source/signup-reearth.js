const axios = require("axios");

/**
* Handler that will be called during the execution of a PostUserRegistration flow.
*
* @param {Event} event - Details about the context and user that has registered.
* @param {PostUserRegistrationAPI} api - Methods and utilities to help change the behavior after a signup.
*/
exports.onExecutePostUserRegistration = async (event) => {
	try {
    const res = await axios.post(event.secrets.api, {
			email: event.user.email,
			sub: event.user.user_id,
			username: event.user.username || event.user.nickname || event.user.email,
			secret: event.secrets.secret
		});
  } catch (err) {
  	console.error(err)
  }
};