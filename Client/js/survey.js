// Submit survey

$(document).ready(() => {
	console.log("Ready");
});

$("#survey-form").on("submit", event => {
	console.log("Form submitted");

	event.preventDefault();
	const questions = [];

	for (let i = 0; $(`input[name='qstn${i}']`).length != 0; i++) {
		// Loops over each input node and returns all the selected values as an array
		questions.push(getInputValues("qstn" + i));
	}

	sendSurvey(questions);
});

function getInputValues(name) {
	const result = [];

	for (let elem of document.getElementsByName(name)) {
		const type = elem.dataset.type;

		if ((type == "single" || type == "multiple") && elem.checked) {
			result.push(elem.value);
		}

		if (type == "text" && elem.value != "") {
			result.push(elem.value);
		}
	}

	return result;
}

async function sendSurvey(questions) {
	console.log("New request");

	const data = JSON.stringify(questions);
	const options = {
		method: "POST",
		headers: {
			"Content-Type": "text",
		},
		body: data,
	};

	const urlParams = new URLSearchParams(window.location.search);
	const surveyID = urlParams.get("id"); // Change

	const res = await fetch(`/submitSurveyData?id=${surveyID}`, options);
	if (res.ok) {
		console.log("Request success");
		showResultScreen();
	} else {
		window.location = `/error/${res.status}`;
	}
}

function showResultScreen() {
	// Show success message
}
