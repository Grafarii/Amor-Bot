const STORAGE_KEY = 'neuroquest_save';

const defaultState = () => ({
  xp: 0,
  level: 1,
  hearts: 3,
  maxHearts: 3,
  streak: 0,
  bestStreak: 0,
  units: UNITS.map(u => ({
    id: u.id,
    unlocked: u.id === 1,
    completed: false,
    lessonsDone: false,
    quizDone: false,
    matchDone: false,
    bossDone: false,
    progress: 0
  })),
  achievements: [],
  totalCorrect: 0,
  totalAnswered: 0
});

let state = loadState();
let currentUnit = null;
let currentLessonIndex = 0;
let quizQuestions = [];
let quizIndex = 0;
let quizMode = 'quiz';
let quizCorrectInRow = 0;
let quizNoLifeLost = true;
let bossQuestions = [];
let bossIndex = 0;
let bossHp = 100;
let matchPairs = [];
let matchSelected = null;
let matchMatched = 0;
let flashIndex = 0;

function loadState() {
  try {
    const saved = localStorage.getItem(STORAGE_KEY);
    if (saved) return { ...defaultState(), ...JSON.parse(saved) };
  } catch (_) {}
  return defaultState();
}

function saveState() {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(state));
}

function xpForLevel(lvl) {
  return lvl * 100;
}

function addXp(amount) {
  state.xp += amount;
  while (state.xp >= xpForLevel(state.level)) {
    state.xp -= xpForLevel(state.level);
    state.level++;
    showToast(`🎉 Nível ${state.level} alcançado!`);
    checkAchievement('level5', state.level >= 5);
    checkAchievement('level10', state.level >= 10);
  }
  updateStats();
  saveState();
}

function loseHeart() {
  if (state.hearts > 0) {
    state.hearts--;
    quizNoLifeLost = false;
    updateStats();
    saveState();
    if (state.hearts === 0) {
      showToast('💔 Sem vidas! Recuperando...');
      setTimeout(() => {
        state.hearts = state.maxHearts;
        state.streak = 0;
        updateStats();
        saveState();
      }, 1500);
    }
  }
}

function updateStats() {
  document.getElementById('xpDisplay').textContent = `${state.xp} XP`;
  document.getElementById('levelDisplay').textContent = `Nv. ${state.level}`;
  document.getElementById('streakDisplay').textContent = state.streak;
  const hearts = '❤️'.repeat(state.hearts) + '🖤'.repeat(state.maxHearts - state.hearts);
  document.getElementById('heartsDisplay').textContent = hearts;
}

function showScreen(id) {
  document.querySelectorAll('.screen').forEach(s => s.classList.remove('active'));
  document.getElementById(id).classList.add('active');
}

function showToast(msg) {
  const t = document.getElementById('toast');
  t.textContent = msg;
  t.classList.add('show');
  setTimeout(() => t.classList.remove('show'), 2500);
}

function checkAchievement(id, condition) {
  if (condition && !state.achievements.includes(id)) {
    state.achievements.push(id);
    const ach = ACHIEVEMENTS.find(a => a.id === id);
    if (ach) showToast(`🏅 Conquista: ${ach.title}!`);
    saveState();
  }
}

function getUnitState(id) {
  return state.units.find(u => u.id === id);
}

function updateUnitProgress(unitId) {
  const us = getUnitState(unitId);
  const unit = UNITS.find(u => u.id === unitId);
  let steps = 0;
  if (us.lessonsDone) steps++;
  if (us.quizDone) steps++;
  if (us.matchDone) steps++;
  if (us.bossDone) steps++;
  us.progress = Math.round((steps / 4) * 100);
  if (us.bossDone) {
    us.completed = true;
    const next = getUnitState(unitId + 1);
    if (next) next.unlocked = true;
    checkAchievement('unit1', unitId >= 1 && us.completed);
    checkAchievement('unit3', unitId >= 3 && us.completed);
    const allDone = state.units.every(u => u.completed);
    checkAchievement('unit6', allDone);
  }
  saveState();
}

function renderMap() {
  const map = document.getElementById('worldMap');
  map.innerHTML = UNITS.map(unit => {
    const us = getUnitState(unit.id);
    const cls = us.completed ? 'completed' : us.unlocked ? 'unlocked' : 'locked';
    return `
      <div class="world-node ${cls}" data-unit="${unit.id}">
        <div class="world-icon">${unit.icon}</div>
        <div class="world-info">
          <h3>Unidade ${unit.id}: ${unit.title}</h3>
          <p>${unit.description}</p>
        </div>
        <div class="world-progress">
          <div class="progress-ring">${us.progress}%</div>
        </div>
      </div>`;
  }).join('');

  map.querySelectorAll('.world-node.unlocked').forEach(node => {
    node.addEventListener('click', () => openUnitMenu(parseInt(node.dataset.unit)));
  });
}

function openUnitMenu(unitId) {
  currentUnit = UNITS.find(u => u.id === unitId);
  const us = getUnitState(unitId);

  if (!us.lessonsDone) {
    startLesson(unitId);
  } else if (!us.quizDone) {
    startQuiz(unitId);
  } else if (!us.matchDone) {
    startMatch(unitId);
  } else if (!us.bossDone) {
    startBoss(unitId);
  } else {
    showToast('✅ Unidade já completa! Escolha outra ou use Revisão.');
    showScreen('screen-map');
  }
}

function startLesson(unitId) {
  currentUnit = UNITS.find(u => u.id === unitId);
  currentLessonIndex = 0;
  showScreen('screen-lesson');
  renderLessonCard();
}

function renderLessonCard() {
  const lesson = currentUnit.lessons[currentLessonIndex];
  document.getElementById('lessonProgress').textContent =
    `Carta ${currentLessonIndex + 1} de ${currentUnit.lessons.length}`;
  document.getElementById('lessonCard').innerHTML = `<h3>${lesson.title}</h3>${lesson.content}`;
  document.getElementById('btnPrevCard').disabled = currentLessonIndex === 0;
  document.getElementById('btnNextCard').textContent =
    currentLessonIndex === currentUnit.lessons.length - 1 ? 'Iniciar Quiz →' : 'Próximo';
}

function nextLessonCard() {
  if (currentLessonIndex < currentUnit.lessons.length - 1) {
    currentLessonIndex++;
    renderLessonCard();
  } else {
    const us = getUnitState(currentUnit.id);
    us.lessonsDone = true;
    updateUnitProgress(currentUnit.id);
    checkAchievement('first_lesson', true);
    addXp(30);
    showToast('📖 Lições concluídas! +30 XP');
    startQuiz(currentUnit.id);
  }
}

function startQuiz(unitId) {
  currentUnit = UNITS.find(u => u.id === unitId);
  quizQuestions = prepareQuestions(currentUnit.quiz).sort(() => Math.random() - 0.5);
  quizIndex = 0;
  quizMode = 'quiz';
  quizCorrectInRow = 0;
  quizNoLifeLost = true;
  showScreen('screen-quiz');
  document.getElementById('quizTitle').textContent = `Quiz — ${currentUnit.title}`;
  renderQuizQuestion();
}

function renderQuizQuestion() {
  const q = quizQuestions[quizIndex];
  document.getElementById('quizCounter').textContent = `${quizIndex + 1}/${quizQuestions.length}`;
  document.getElementById('quizXpFill').style.width = `${((quizIndex) / quizQuestions.length) * 100}%`;
  document.getElementById('quizQuestion').textContent = q.q;
  document.getElementById('quizFeedback').className = 'quiz-feedback';
  document.getElementById('quizFeedback').innerHTML = '';

  const opts = document.getElementById('quizOptions');
  opts.innerHTML = q.options.map((opt, i) =>
    `<button class="quiz-option" data-idx="${i}">${opt}</button>`
  ).join('');

  opts.querySelectorAll('.quiz-option').forEach(btn => {
    btn.addEventListener('click', () => answerQuiz(parseInt(btn.dataset.idx)));
  });
}

function answerQuiz(selected) {
  const q = quizQuestions[quizIndex];
  const opts = document.querySelectorAll('#quizOptions .quiz-option');
  opts.forEach((btn, i) => {
    btn.classList.add('disabled');
    if (i === q.answer) btn.classList.add('correct');
    else if (i === selected) btn.classList.add('wrong');
  });

  const fb = document.getElementById('quizFeedback');
  state.totalAnswered++;

  if (selected === q.answer) {
    state.totalCorrect++;
    state.streak++;
    quizCorrectInRow++;
    if (state.streak > state.bestStreak) state.bestStreak = state.streak;
    addXp(15);
    fb.className = 'quiz-feedback show correct-fb';
    fb.innerHTML = `✅ Correto! +15 XP<br><small>${q.explain}</small>`;
    checkAchievement('streak5', state.streak >= 5);
    checkAchievement('streak10', state.streak >= 10);
  } else {
    state.streak = 0;
    loseHeart();
    fb.className = 'quiz-feedback show wrong-fb';
    fb.innerHTML = `❌ Errado. Resposta: <strong>${q.options[q.answer]}</strong><br><small>${q.explain}</small>`;
  }

  saveState();
  setTimeout(() => {
    quizIndex++;
    if (quizIndex >= quizQuestions.length) {
      finishQuiz();
    } else {
      renderQuizQuestion();
    }
  }, 2200);
}

function finishQuiz() {
  const us = getUnitState(currentUnit.id);
  us.quizDone = true;
  updateUnitProgress(currentUnit.id);
  checkAchievement('first_quiz', true);
  if (quizNoLifeLost) checkAchievement('no_life_lost', true);
  addXp(25);
  showToast('✅ Quiz completo! +25 XP bônus');
  startMatch(currentUnit.id);
}

function startMatch(unitId) {
  currentUnit = UNITS.find(u => u.id === unitId);
  matchPairs = [...currentUnit.matching].sort(() => Math.random() - 0.5);
  matchSelected = null;
  matchMatched = 0;
  showScreen('screen-match');
  document.getElementById('matchTitle').textContent = `Associe — ${currentUnit.title}`;
  renderMatch();
}

function renderMatch() {
  const leftItems = matchPairs.map((p, i) => ({ text: p.term, idx: i }));
  const rightItems = matchPairs.map((p, i) => ({ text: p.match, idx: i }))
    .sort(() => Math.random() - 0.5);

  document.getElementById('matchScore').textContent = `${matchMatched}/${matchPairs.length}`;

  document.getElementById('matchLeft').innerHTML = leftItems.map(item =>
    `<div class="match-item" data-side="left" data-idx="${item.idx}">${item.text}</div>`
  ).join('');

  document.getElementById('matchRight').innerHTML = rightItems.map(item =>
    `<div class="match-item" data-side="right" data-idx="${item.idx}">${item.text}</div>`
  ).join('');

  document.querySelectorAll('.match-item:not(.matched)').forEach(el => {
    el.addEventListener('click', () => handleMatchClick(el));
  });
}

function handleMatchClick(el) {
  if (el.classList.contains('matched')) return;

  if (!matchSelected) {
    matchSelected = el;
    el.classList.add('selected');
    return;
  }

  if (matchSelected === el) {
    el.classList.remove('selected');
    matchSelected = null;
    return;
  }

  if (matchSelected.dataset.side === el.dataset.side) {
    matchSelected.classList.remove('selected');
    matchSelected = el;
    el.classList.add('selected');
    return;
  }

  const idx1 = parseInt(matchSelected.dataset.idx);
  const idx2 = parseInt(el.dataset.idx);

  if (idx1 === idx2) {
    matchSelected.classList.add('matched');
    el.classList.add('matched');
    matchSelected.classList.remove('selected');
    matchSelected = null;
    matchMatched++;
    addXp(10);
    document.getElementById('matchScore').textContent = `${matchMatched}/${matchPairs.length}`;

    if (matchMatched >= matchPairs.length) {
      setTimeout(finishMatch, 600);
    }
  } else {
    matchSelected.classList.add('wrong-flash');
    el.classList.add('wrong-flash');
    loseHeart();
    state.streak = 0;
    saveState();
    setTimeout(() => {
      matchSelected.classList.remove('selected', 'wrong-flash');
      el.classList.remove('wrong-flash');
      matchSelected = null;
    }, 600);
  }
}

function finishMatch() {
  const us = getUnitState(currentUnit.id);
  us.matchDone = true;
  updateUnitProgress(currentUnit.id);
  checkAchievement('match_master', true);
  addXp(20);
  showToast('🔗 Associação completa! +20 XP');
  startBoss(currentUnit.id);
}

function startBoss(unitId) {
  currentUnit = UNITS.find(u => u.id === unitId);
  bossQuestions = prepareQuestions(currentUnit.boss.questions);
  bossIndex = 0;
  bossHp = 100;
  showScreen('screen-boss');
  document.getElementById('bossIntro').style.display = 'block';
  document.getElementById('bossBattle').style.display = 'none';
  document.getElementById('bossVictory').style.display = 'none';
  document.getElementById('bossName').textContent = currentUnit.boss.name;
  document.getElementById('bossDesc').textContent = currentUnit.boss.desc;
}

function fightBoss() {
  document.getElementById('bossIntro').style.display = 'none';
  document.getElementById('bossBattle').style.display = 'block';
  renderBossQuestion();
}

function renderBossQuestion() {
  const q = bossQuestions[bossIndex];
  document.getElementById('bossHpFill').style.width = `${bossHp}%`;
  document.getElementById('bossQuestion').textContent = q.q;
  document.getElementById('bossFeedback').className = 'quiz-feedback';
  document.getElementById('bossFeedback').innerHTML = '';

  const opts = document.getElementById('bossOptions');
  opts.innerHTML = q.options.map((opt, i) =>
    `<button class="quiz-option" data-idx="${i}">${opt}</button>`
  ).join('');

  opts.querySelectorAll('.quiz-option').forEach(btn => {
    btn.addEventListener('click', () => answerBoss(parseInt(btn.dataset.idx)));
  });
}

function answerBoss(selected) {
  const q = bossQuestions[bossIndex];
  const opts = document.querySelectorAll('#bossOptions .quiz-option');
  opts.forEach((btn, i) => {
    btn.classList.add('disabled');
    if (i === q.answer) btn.classList.add('correct');
    else if (i === selected) btn.classList.add('wrong');
  });

  const fb = document.getElementById('bossFeedback');

  if (selected === q.answer) {
    bossHp -= 100 / bossQuestions.length;
    state.streak++;
    addXp(20);
    fb.className = 'quiz-feedback show correct-fb';
    fb.innerHTML = `✅ Dano causado! +20 XP<br><small>${q.explain}</small>`;
  } else {
    loseHeart();
    state.streak = 0;
    fb.className = 'quiz-feedback show wrong-fb';
    fb.innerHTML = `❌ O chefe contra-ataca! Resposta: <strong>${q.options[q.answer]}</strong><br><small>${q.explain}</small>`;
  }

  saveState();
  setTimeout(() => {
    bossIndex++;
    if (bossIndex >= bossQuestions.length || bossHp <= 0) {
      if (bossHp <= 0 || bossIndex >= bossQuestions.length) {
        victoryBoss();
      } else {
        renderBossQuestion();
      }
    } else {
      renderBossQuestion();
    }
  }, 2200);
}

function victoryBoss() {
  document.getElementById('bossBattle').style.display = 'none';
  document.getElementById('bossVictory').style.display = 'block';
  const us = getUnitState(currentUnit.id);
  us.bossDone = true;
  updateUnitProgress(currentUnit.id);
  checkAchievement('first_boss', true);
  addXp(50);
  document.getElementById('victoryMessage').textContent =
    `Você dominou "${currentUnit.title}"!`;
  document.getElementById('victoryXp').textContent = '+50 XP bônus de chefe!';
}

function renderAchievements() {
  const grid = document.getElementById('achievementsGrid');
  grid.innerHTML = ACHIEVEMENTS.map(a => {
    const earned = state.achievements.includes(a.id);
    return `
      <div class="achievement-card ${earned ? 'earned' : ''}">
        <div class="ach-icon">${earned ? a.icon : '🔒'}</div>
        <h4>${a.title}</h4>
        <p>${a.desc}</p>
      </div>`;
  }).join('');
}

function renderAchievementsPreview() {
  const preview = document.getElementById('achievementsPreview');
  preview.innerHTML = ACHIEVEMENTS.slice(0, 6).map(a => {
    const earned = state.achievements.includes(a.id);
    return `<div class="ach-badge ${earned ? 'earned' : ''}">${a.icon} ${a.title}</div>`;
  }).join('');
}

function startReview() {
  flashIndex = Math.floor(Math.random() * FLASHCARDS.length);
  showScreen('screen-review');
  renderFlashcard();
}

function renderFlashcard() {
  const card = FLASHCARDS[flashIndex];
  document.getElementById('flashcardFront').textContent = card.front;
  document.getElementById('flashcardBack').textContent = card.back;
  document.getElementById('flashcardInner').classList.remove('flipped');
}

function nextFlashcard() {
  flashIndex = Math.floor(Math.random() * FLASHCARDS.length);
  renderFlashcard();
}

function flipFlashcard() {
  document.getElementById('flashcardInner').classList.toggle('flipped');
}

function hasProgress() {
  return state.units.some(u => u.lessonsDone || u.quizDone || u.progress > 0);
}

document.addEventListener('DOMContentLoaded', () => {
  updateStats();
  renderAchievementsPreview();

  if (hasProgress()) {
    document.getElementById('btnContinue').style.display = 'inline-block';
  }

  document.getElementById('btnStart').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnContinue').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnBackLesson').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnPrevCard').addEventListener('click', () => {
    if (currentLessonIndex > 0) {
      currentLessonIndex--;
      renderLessonCard();
    }
  });

  document.getElementById('btnNextCard').addEventListener('click', nextLessonCard);

  document.getElementById('btnBackQuiz').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnBackMatch').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnFightBoss').addEventListener('click', fightBoss);

  document.getElementById('btnVictoryContinue').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnAchievements').addEventListener('click', () => {
    renderAchievements();
    showScreen('screen-achievements');
  });

  document.getElementById('btnBackAchievements').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnReview').addEventListener('click', startReview);

  document.getElementById('btnBackReview').addEventListener('click', () => {
    renderMap();
    showScreen('screen-map');
  });

  document.getElementById('btnFlipCard').addEventListener('click', flipFlashcard);
  document.getElementById('flashcard').addEventListener('click', flipFlashcard);
  document.getElementById('btnNextFlash').addEventListener('click', nextFlashcard);
});
