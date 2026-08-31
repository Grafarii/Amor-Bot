function shuffleQuestion(question) {
  const items = question.options.map((text, originalIndex) => ({ text, originalIndex }));
  for (let i = items.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [items[i], items[j]] = [items[j], items[i]];
  }
  return {
    ...question,
    options: items.map(item => item.text),
    answer: items.findIndex(item => item.originalIndex === question.answer)
  };
}

function prepareQuestions(questions) {
  return questions.map(shuffleQuestion);
}
