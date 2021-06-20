$("#login-form").on("submit", event => {
	event.preventDefault();

	const options = {
		method: "POST",
		Headers: {
			"Content-Type": "form-data",
		},
	};
	const res = await fetch("/login", options);
});
