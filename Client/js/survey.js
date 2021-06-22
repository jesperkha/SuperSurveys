// Submit survey

$("#survey-form").on("submit", event => {
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
	const urlParams = new URLSearchParams(window.location.search);
	const surveyId = urlParams.get("id");

	const options = {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			Data: questions,
			Id: surveyId,
		}),
	};

	try {
		const res = await fetch(`/survey`, options);
		if (res.ok) console.log("Success");
		else window.location = `/error/${res.status}`;
	} catch (err) {
		console.error(err);
		window.location = `/error/500`;
	}
}
