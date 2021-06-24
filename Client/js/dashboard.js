const jar = {};

// Get cookies
window.onload = () => {
	const cookies = document.cookie.split("; ");
	for (let cookie of cookies) {
		const [key, value] = cookie.split("=");
		jar[key] = value
	}
}