const jar = {};
let User;

// Get cookies and user profile
window.onload = () => {
	const cookies = document.cookie.split("; ");
	for (let cookie of cookies) {
		const [key, value] = cookie.split("=");
		jar[key] = value;
	}

	try {
		const res = await fetch("/users/get");
		User = await res.json();
		console.log(User);
	} catch (err) {
		console.error(err);
	}
};
