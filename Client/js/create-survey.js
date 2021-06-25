let numQuestions = 0;

function createQuestion() {
	const prompt = document.getElementById("prompt").value;
	const type = document.getElementById("type-select").value;
	const current = numQuestions;

	const container = document.getElementById("questions");
	const div = document.createElement("div");
	div.classList.add("question");
	div.setAttribute("id", "qstn" + current);

	// Sub
	const sub = document.createElement("sub");
	sub.textContent = `Question ${current + 1}:`;
	div.appendChild(sub);

	// Title
	const h2 = document.createElement("h2");
	h2.textContent = prompt;
	div.appendChild(h2);

	// Answers
	const answers = document.createElement("div");
	answers.setAttribute("id", "ans" + current);
	answers.classList.add("answers");

	// Radio and Checkbox
	if (type === "radio" || type === "checkbox") {
		div.appendChild(answers);

		// Text input
		const input = document.createElement("input");
		input.setAttribute("type", "text");
		input.setAttribute("id", "input" + current);
		div.appendChild(input);

		// Add input button
		const btn = document.createElement("button");
		btn.textContent = "Click to add option";
		btn.onclick = () => {
			addOption(answers.id, input.value, type);
			input.value = "";
		};
		div.appendChild(btn);
	}

	// Text
	if (type === "text") {
		const input = document.createElement("input");
		input.setAttribute("type", "text");
		answers.appendChild(input);
		div.appendChild(answers);
	}

	container.appendChild(div);
	numQuestions++;
}

function addOption(id, value, type) {
	const div = document.getElementById(id);

	// Input
	const input = document.createElement("input");
	input.setAttribute("type", type);
	input.setAttribute("name", id);
	div.appendChild(input);

	// Label
	const label = document.createElement("label");
	label.setAttribute("for", id);
	label.textContent = value;
	div.appendChild(label);

	div.appendChild(document.createElement("br"));
}

function getSurveyInputs() {
	let allQuestions = [];
	const lookup = {
		checkbox: "multiple",
		radio: "single",
		text: "text",
	};

	for (let qstn of document.querySelectorAll(".question")) {
		const firstInput = document.querySelector(`#${qstn.id} .answers input`);
		if (firstInput == null) {
			continue;
		}

		let question = {
			Optional: false,
			Prompt: document.querySelector(`#${qstn.id} h2`).textContent,
			Class: lookup[firstInput.type],
			Options: [],
		};

		for (let label of document.querySelectorAll(`#${qstn.id} .answers label`)) {
			question.Options.push(label.textContent);
		}

		allQuestions.push(question);
	}

	return allQuestions;
}

async function publishCreatedSurvey() {
	try {
		const data = {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(getSurveyInputs()),
		};

		const res = await fetch("/users/create", data);
		if (res.ok) console.log("success");
	} catch (err) {
		console.error(err);
	}
}
