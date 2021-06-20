const cookies = {};

// Get cookies
$(document).ready(() => {
	const split = document.cookie.split("=");
	for (let i = 0; i < split.length; i += 2) {
		cookies[split[i]] = split[i + 1];
	}
});
