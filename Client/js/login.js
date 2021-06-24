document.querySelector("#login-form").addEventListener("submit", event => {
	event.preventDefault();

	const options = {
		method: "POST",
		Headers: {
			"Content-Type": "form-data",
		},
	};

	const res = await fetch("/login", options);
	const user = await res.json();
	console.log(user)
})