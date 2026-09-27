const flashcardElement = document.querySelector("#flashcard");
const questionElement = document.querySelector("#question");
const answerElement = document.querySelector("#answer");
const nextButton = document.querySelector("#next-button");
const messageElement = document.querySelector("#message");

let flashcards = [];
let currentIndex = 0;
let answerVisible = false;

async function loadFlashcards() {
const response = await fetch("http://localhost:8080/");


if (!response.ok) {
    throw new Error(`Failed to load flashcards: ${response.status}`);
}

flashcards = await response.json();

if (flashcards.length === 0) {
    questionElement.textContent = "No flashcards available.";
    nextButton.disabled = true;
    return;
}

showCard();


}

function showCard() {
const card = flashcards[currentIndex];


questionElement.textContent = card.question;
answerElement.textContent = card.answer;

answerVisible = false;
answerElement.classList.add("hidden");

messageElement.textContent =
    `${currentIndex + 1} / ${flashcards.length}`;


}

flashcardElement.addEventListener("click", () => {
if (flashcards.length === 0) {
return;
}


answerVisible = !answerVisible;

answerElement.classList.toggle("hidden", !answerVisible);


});

nextButton.addEventListener("click", () => {
if (flashcards.length === 0) {
return;
}


currentIndex = (currentIndex + 1) % flashcards.length;

showCard();


});

loadFlashcards().catch((error) => {
console.error(error);


questionElement.textContent = "Failed to load flashcards.";
nextButton.disabled = true;


});

