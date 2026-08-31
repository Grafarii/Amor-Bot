# NeuroQuest 🧠

Jogo educativo para aprender **Neurociências e Neuropsicologia** (2º semestre — Psicologia).

## Jogar no celular

### Agora (link direto)
**https://temporary-fast-flurry-2znl2t9.vercel.app**

Abra esse link no Chrome ou Safari do celular e comece a jogar.

### URL permanente (recomendado)
O site já está publicado na branch `gh-pages`. Para ativar a URL fixa:

1. Acesse: https://github.com/Grafarii/Amor-Bot/settings/pages
2. Em **Source**, escolha: Branch `gh-pages` → pasta `/ (root)`
3. Salve e aguarde ~1 minuto

Depois use: **https://grafarii.github.io/Amor-Bot/**

### Instalar como app na tela inicial
- **iPhone:** Safari → Compartilhar → "Adicionar à Tela de Início"
- **Android:** Chrome → menu (⋮) → "Instalar app"

## Como jogar

1. Abra `index.html` no navegador (duplo clique ou servidor local)
2. Explore o **Mapa Neural** com 6 unidades do plano de ensino
3. Em cada unidade:
   - **Lições** — cartas interativas com o conteúdo
   - **Quiz** — perguntas de múltipla escolha
   - **Associação** — conecte termos e conceitos
   - **Chefe** — desafio final da unidade
4. Ganhe **XP**, suba de **nível**, mantenha sua **sequência** e desbloqueie **conquistas**
5. Use o **Modo Revisão** com flashcards quando quiser

## Conteúdo (6 unidades)

| Unidade | Tema |
|---------|------|
| 1 | Introdução às Neurociências |
| 2 | Introdução à Neuropsicologia |
| 3 | Linguagem, Atenção e Percepção |
| 4 | Memória e Funções Executivas |
| 5 | Transtornos do Neurodesenvolvimento |
| 6 | Quadros Neurológicos |

## Executar localmente

```bash
cd neuro-quest
python3 -m http.server 8080
```

Acesse: http://localhost:8080

O progresso é salvo automaticamente no navegador (localStorage).

## Tecnologias

HTML5, CSS3, JavaScript vanilla — sem dependências externas.
