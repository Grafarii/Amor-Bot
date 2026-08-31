const UNITS = [
  {
    id: 1,
    title: "Introdução às Neurociências",
    icon: "🔬",
    color: "#7c3aed",
    description: "Definição, lobos cerebrais, circuitos córtico-subcorticais e história",
    lessons: [
      {
        title: "O que são Neurociências?",
        content: `<p>As <strong>Neurociências</strong> são o estudo científico do sistema nervoso — desde a estrutura molecular até o comportamento complexo.</p>
        <div class="highlight"><strong>Subáreas principais:</strong> Neuroanatomia, Neurofisiologia, Neuropsicologia, Neuroimagem, Psicofarmacologia e Neurociência Cognitiva.</div>
        <p>Interface direta com a <strong>Psicologia</strong>: compreender como o cérebro produz mente, emoção, cognição e comportamento.</p>`
      },
      {
        title: "Organização Cortical",
        content: `<p>O <strong>córtex cerebral</strong> é dividido em áreas especializadas que processam diferentes funções psicológicas.</p>
        <ul>
          <li><strong>Córtex motor</strong> — controle de movimentos voluntários</li>
          <li><strong>Córtex sensorial</strong> — processamento de estímulos</li>
          <li><strong>Áreas de associação</strong> — integração de informações (pensamento, linguagem)</li>
        </ul>
        <p>O princípio da <strong>especialização funcional</strong> (localização) mostra que funções específicas têm regiões dedicadas, mas com interação entre áreas.</p>`
      },
      {
        title: "Lobos Cerebrais",
        content: `<ul>
          <li><strong>Frontal</strong> — funções executivas, planejamento, personalidade, controle motor, linguagem (Broca)</li>
          <li><strong>Parietal</strong> — integração sensorial, atenção espacial, praxia</li>
          <li><strong>Temporal</strong> — audição, memória (hipocampo), linguagem (Wernicke), emoção (amígdala)</li>
          <li><strong>Occipital</strong> — processamento visual primário</li>
        </ul>
        <div class="highlight">Cada lobo tem funções predominantes, mas trabalham em <strong>redes neurais integradas</strong>.</div>`
      },
      {
        title: "Circuitos Córtico-Subcorticais",
        content: `<p><strong>Tálamo:</strong> "Central de retransmissão" — filtra e direciona informações sensoriais ao córtex.</p>
        <p><strong>Gânglios da Base:</strong> Iniciam e modulam movimentos; envolvidos em TDAH, Parkinson e Tourette.</p>
        <p><strong>Cerebelo:</strong> Coordenação motora, equilíbrio, timing e também funções cognitivas e emocionais.</p>
        <p><strong>Sistema Límbico:</strong> Amígdala e hipocampo — emoção e memória.</p>`
      },
      {
        title: "História das Neurociências",
        content: `<ul>
          <li><strong>Séc. XIX:</strong> Broca (1861) — área da fala; Wernicke (1874) — compreensão da linguagem</li>
          <li><strong>Phineas Gage (1848):</strong> Lesão frontal → mudança de personalidade</li>
          <li><strong>Paul Broca & Carl Wernicke:</strong> Fundaram a localização funcional</li>
          <li><strong>Alexander Luria:</strong> Neuropsicologia clínica sistemática (URSS)</li>
          <li><strong>Séc. XX-XXI:</strong> Neuroimagem (TCE, RMf), plasticidade neural</li>
        </ul>`
      }
    ],
    quiz: [
      { q: "Qual lobo cerebral está mais associado às funções executivas e planejamento?", options: ["Occipital", "Frontal", "Temporal", "Parietal"], answer: 1, explain: "O lobo frontal, especialmente o córtex pré-frontal, é o centro das funções executivas." },
      { q: "O tálamo funciona principalmente como:", options: ["Armazenamento da memória","Retransmissão sensorial central","Regulação motora direta","Processamento visual inicial"], answer: 1, explain: "O tálamo recebe informações sensoriais e as direciona ao córtex apropriado." },
      { q: "A área de Wernicke está localizada no lobo:", options: ["Frontal", "Parietal", "Temporal", "Occipital"], answer: 2, explain: "A área de Wernicke fica no lobo temporal esquerdo e é responsável pela compreensão da linguagem." },
      { q: "Os gânglios da base estão envolvidos em:", options: ["Processamento da visão","Modulação dos movimentos","Processamento da audição","Controle do equilíbrio postural"], answer: 1, explain: "Os gânglios da base iniciam e modulam movimentos voluntários." },
      { q: "Phineas Gage teve lesão no lobo frontal, demonstrando relação entre:", options: ["Visão e lobo occipital posterior","Personalidade e córtex frontal","Memória e lobo temporal medial","Linguagem e lobo parietal superior"], answer: 1, explain: "Após lesão frontal, Gage teve alterações marcantes de personalidade e comportamento." }
    ],
    matching: [
      { term: "Lobo Frontal", match: "Funções executivas" },
      { term: "Lobo Occipital", match: "Visão" },
      { term: "Hipocampo", match: "Memória" },
      { term: "Cerebelo", match: "Coordenação motora" },
      { term: "Tálamo", match: "Retransmissão sensorial" }
    ],
    boss: {
      name: "Guardião Neural",
      desc: "Prove que domina os fundamentos das Neurociências!",
      questions: [
        { q: "Qual estrutura NÃO é subcortical?", options: ["Amígdala profunda","Hipocampo medial","Córtex pré-frontal","Tálamo dorsal"], answer: 2, explain: "O córtex pré-frontal é cortical; amígdala, hipocampo e tálamo são subcorticais." },
        { q: "A neurociência cognitiva estuda principalmente:", options: ["Neuroanatomia estrutural do cérebro","Cérebro, cognição e comportamento","Farmacologia dos psicofármacos","Cirurgias do sistema nervoso central"], answer: 1, explain: "Integra neurobiologia com processos cognitivos e comportamentais." },
        { q: "Luria contribuiu para:", options: ["Descoberta da estrutura do DNA","Neuropsicologia clínica sistemática","Teoria behaviorista radical clássica","Mapeamento eletroencefalográfico"], answer: 1, explain: "Alexander Luria foi pioneiro na avaliação neuropsicológica clínica." },
        { q: "O lobo parietal está associado a:", options: ["Processamento auditivo primário","Integração sensorial e atenção","Processamento visual primário","Produção da fala articulada"], answer: 1, explain: "O parietal integra informações sensoriais e coordena atenção espacial." },
        { q: "Circuitos córtico-subcorticais envolvem:", options: ["Apenas estruturas corticais","Interação córtex-subcortical","Somente vias da medula","Apenas nervos periféricos"], answer: 1, explain: "Funções complexas dependem de loops entre córtex e subcorticais." }
      ]
    }
  },
  {
    id: 2,
    title: "Introdução à Neuropsicologia",
    icon: "🩺",
    color: "#6366f1",
    description: "Fundamentos, história, Brasil e casos clássicos",
    lessons: [
      {
        title: "O que é Neuropsicologia?",
        content: `<p><strong>Neuropsicologia</strong> é a área que estuda a relação entre o funcionamento cerebral e o comportamento cognitivo, emocional e social.</p>
        <div class="highlight">Objetivo clínico: avaliar, diagnosticar e reabilitar alterações cognitivas resultantes de lesões ou disfunções cerebrais.</div>
        <p>O neuropsicólogo utiliza testes padronizados, entrevistas e observação para mapear o perfil cognitivo do paciente.</p>`
      },
      {
        title: "História da Neuropsicologia",
        content: `<ul>
          <li><strong>Broca & Wernicke:</strong> Modelos de localização da linguagem</li>
          <li><strong>Luria (1902-1977):</strong> Abordagem sistêmico-dinâmica; baterias de avaliação</li>
          <li><strong>Rey, Halstead, Luria-Nebraska:</strong> Desenvolvimento de testes neuropsicológicos</li>
          <li><strong>Neuroimagem moderna:</strong> Validou e refinou modelos clíssicos</li>
        </ul>`
      },
      {
        title: "Neuropsicologia no Brasil",
        content: `<p>A Neuropsicologia brasileira cresceu significativamente nas últimas décadas, com formação em pós-graduação e grupos de pesquisa.</p>
        <div class="highlight"><strong>Resolução CFP nº 002/2004:</strong> Reconhece a Neuropsicologia como especialidade profissional do psicólogo no Brasil.</div>
        <p>Referências nacionais: Fuentes, Malloy-Diniz, Miotto, Santos, Haase, entre outros.</p>`
      },
      {
        title: "Caso Phineas Gage",
        content: `<p>Em 1848, Phineas Gage, ferroviário, sofreu lesão por barra de ferro atravessando o lobo <strong>frontal</strong>.</p>
        <p>Surpreendentemente sobreviveu, mas teve mudanças drásticas de personalidade: impulsividade, irritabilidade, dificuldade de planejamento.</p>
        <div class="highlight">Este caso histórico demonstrou que lesões cerebrais podem alterar comportamento e personalidade — base da neuropsicologia moderna.</div>`
      },
      {
        title: "Caso Paciente H.M. (Henry Molaison)",
        content: `<p>H.M. teve lobectomia bilateral do <strong>hipocampo</strong> para tratar epilepsia (1953).</p>
        <p>Desenvolveu <strong>amnésia anterógrada severa</strong> — incapaz de formar novas memórias episódicas, mas memórias antigas e habilidades procedimentais preservadas.</p>
        <div class="highlight">H.M. provou que o hipocampo é essencial para a consolidação de novas memórias declarativas.</div>`
      }
    ],
    quiz: [
      { q: "A Resolução CFP nº 002/2004 reconhece a Neuropsicologia como:", options: ["Curso superior de medicina","Especialidade do psicólogo","Área da biologia molecular","Subcampo da fisioterapia"], answer: 1, explain: "A resolução oficializa a Neuropsicologia como especialidade profissional do psicólogo." },
      { q: "Phineas Gage teve lesão principalmente no:", options: ["Hipocampo", "Lobo frontal", "Cerebelo", "Occipital"], answer: 1, explain: "A barra de ferro lesionou o lobo frontal, alterando personalidade e comportamento." },
      { q: "O paciente H.M. tinha amnésia:", options: ["Retrógrada isolada","Anterógrada severa","Totalmente inexistente","Apenas de estímulos sensoriais"], answer: 1, explain: "H.M. não conseguia formar novas memórias episódicas após a cirurgia." },
      { q: "Alexander Luria é conhecido por:", options: ["Teoria psicanalítica freudiana clássica","Neuropsicologia clínica sistemática","Escola behaviorista clássica radical","Abordagem humanista integral clínica"], answer: 1, explain: "Luria desenvolveu abordagem sistêmico-dinâmica e baterias de avaliação." },
      { q: "A neuropsicologia clínica tem como foco:", options: ["Pesquisa laboratorial básica","Avaliação e reabilitação cognitiva","Procedimentos cirúrgicos cerebrais","Prescrição de psicofármacos"], answer: 1, explain: "O neuropsicólogo clínico avalia, diagnostica e propõe reabilitação cognitiva." }
    ],
    matching: [
      { term: "Phineas Gage", match: "Lesão frontal / personalidade" },
      { term: "Paciente H.M.", match: "Amnésia anterógrada" },
      { term: "Resolução CFP 002/2004", match: "Especialidade do psicólogo" },
      { term: "Luria", match: "Abordagem sistêmico-dinâmica" },
      { term: "Neuropsicologia", match: "Cérebro e comportamento" }
    ],
    boss: {
      name: "Mestre Clínico",
      desc: "Demonstre domínio dos fundamentos da Neuropsicologia!",
      questions: [
        { q: "A amnésia anterógrada significa dificuldade em:", options: ["Lembrar eventos passados","Formar novas memórias","Produzir fala articulada","Processar estímulos visuais"], answer: 1, explain: "Anterógrada = incapacidade de consolidar novas memórias após o evento." },
        { q: "Qual estrutura foi removida em H.M.?", options: ["Amígdala bilateral","Hipocampo bilateral","Cerebelo bilateral","Tálamo bilateral"], answer: 1, explain: "A lobectomia bilateral do hipocampo causou a amnésia de H.M." },
        { q: "Neuropsicologia integra conhecimentos de:", options: ["Filosofia continental apenas","Neurociência e psicologia","Estatística matemática apenas","Pedagogia escolar apenas"], answer: 1, explain: "É a interface entre neurobiologia e processos psicológicos." },
        { q: "A avaliação neuropsicológica utiliza:", options: ["Somente entrevista clínica","Testes padronizados e observação","Apenas exames de neuroimagem","Somente exames laboratoriais"], answer: 1, explain: "Combina instrumentos padronizados, entrevista e raciocínio clínico." },
        { q: "H.M. preservou memória:", options: ["Episódica recente intacta","Procedimental e antiga","Nenhuma forma de memória","Somente emocional recente"], answer: 1, explain: "H.M. mantinha memórias antigas e aprendia habilidades motoras (procedimental)." }
      ]
    }
  },
  {
    id: 3,
    title: "Linguagem, Atenção e Percepção",
    icon: "💬",
    color: "#3b82f6",
    description: "Broca, Wernicke, afasias, atenção, agnosias",
    lessons: [
      {
        title: "Linguagem e Substratos Neurais",
        content: `<p><strong>Área de Broca</strong> (frontal esquerdo, giro frontal inferior): produção da fala — fala não fluente em lesão.</p>
        <p><strong>Área de Wernicke</strong> (temporal esquerdo, giro temporal superior): compreensão da linguagem — fala fluente sem sentido em lesão.</p>
        <p><strong>Fascículo arqueado:</strong> conecta Broca e Wernicke — lesão causa afasia de condução.</p>
        <p><strong>Modelo de Geschwind:</strong> integra áreas de linguagem com regiões visuais e auditivas.</p>`
      },
      {
        title: "Afasias, Alexias e Agrafias",
        content: `<ul>
          <li><strong>Afasia de Broca:</strong> fala não fluente, compreensão relativamente preservada</li>
          <li><strong>Afasia de Wernicke:</strong> fala fluente sem sentido, compreensão comprometida</li>
          <li><strong>Afasia global:</strong> produção e compreensão gravemente afetadas</li>
          <li><strong>Alexia:</strong> incapacidade de ler (sem cegueira)</li>
          <li><strong>Agrafia:</strong> incapacidade de escrever</li>
        </ul>`
      },
      {
        title: "Atenção",
        content: `<p><strong>Tipos:</strong> Focalizada/seletiva, sustentada, dividida e alternada.</p>
        <p><strong>Modelos:</strong></p>
        <ul>
          <li>Filter de Broadbent — seleção precoce</li>
          <li>Attenuation de Treisman — seleção tardia</li>
          <li>Modelo de Posner — orienting, alerting, executive</li>
        </ul>
        <p><strong>Substratos:</strong> Córtex parietal posterior, córtex pré-frontal, sistema reticular activador ascendente (SRAA).</p>`
      },
      {
        title: "Percepção Sensorial",
        content: `<p><strong>Visual:</strong> Retina → nervo óptico → tálamo (LGN) → córtex visual primário (V1, occipital)</p>
        <p><strong>Auditiva:</strong> Cóclea → tronco → colículo inferior → tálamo (MGN) → córtex auditivo temporal</p>
        <p><strong>Somatossensorial:</strong> Receptores → medula → tálamo → córtex somatosensorial (parietal)</p>
        <p>Vias específicas processam qualidade do estímulo; áreas de associação integram significado.</p>`
      },
      {
        title: "Agnosias",
        content: `<p><strong>Agnosia</strong> = incapacidade de reconhecer estímulos apesar de sensação preservada.</p>
        <ul>
          <li><strong>Visual:</strong> Prosopagnosia (rostos), simultagnosia, achromatopsia</li>
          <li><strong>Auditiva:</strong> Incapacidade de reconhecer sons/música</li>
          <li><strong>Tátil:</strong> Astereognosia — não reconhece objetos pelo tato</li>
        </ul>
        <div class="highlight">Agnosias indicam lesão em áreas de associação, não nos órgãos sensoriais.</div>`
      }
    ],
    quiz: [
      { q: "Afasia de Broca caracteriza-se por:", options: ["Fala fluente sem compreensão da fala","Fala não fluente, compreensão preservada","Perda total da audição bilateral","Perda total da visão binocular"], answer: 1, explain: "Broca afeta produção; compreensão geralmente está melhor preservada." },
      { q: "A área de Wernicke está envolvida em:", options: ["Produção motora da fala","Compreensão da linguagem","Processamento visual primário","Controle do equilíbrio"], answer: 1, explain: "Wernicke processa a compreensão da linguagem falada e escrita." },
      { q: "Prosopagnosia é dificuldade em reconhecer:", options: ["Cores e tonalidades","Rostos de pessoas conhecidas","Sons ambientais complexos","Objetos pelo tato apenas"], answer: 1, explain: "Prosopagnosia = agnosia visual para rostos." },
      { q: "O SRAA (sistema reticular) está relacionado a:", options: ["Processamento visual primário","Vigília e atenção global","Processamento auditivo primário","Motricidade fina dos dedos"], answer: 1, explain: "O SRAA regula estados de alerta e vigília." },
      { q: "Afasia de Wernicke apresenta fala:", options: ["Não fluente e telegráfica","Fluente, porém sem sentido","Mutismo completo e duradouro","Normal e bem articulada"], answer: 1, explain: "Pacientes falam fluentemente, mas com parafasias e neologismos." }
    ],
    matching: [
      { term: "Broca", match: "Produção da fala" },
      { term: "Wernicke", match: "Compreensão da linguagem" },
      { term: "Prosopagnosia", match: "Não reconhece rostos" },
      { term: "Atenção seletiva", match: "Foco em estímulo relevante" },
      { term: "V1 (córtex visual)", match: "Processamento visual primário" }
    ],
    boss: {
      name: "Oráculo da Linguagem",
      desc: "Prove seu domínio sobre linguagem, atenção e percepção!",
      questions: [
        { q: "Agrafia está associada a dificuldade em:", options: ["Ler textos escritos","Escrever palavras e frases","Ouvir sons ambientais","Executar movimentos de marcha"], answer: 1, explain: "Agrafia = incapacidade de escrever, frequentemente com alexia." },
        { q: "Atenção dividida envolve:", options: ["Focar em um único estímulo relevante","Processar vários estímulos ao mesmo tempo","Reduzir o nível de vigília global","Armazenar memórias antigas episódicas"], answer: 1, explain: "Atenção dividida = capacidade de monitorar e responder a múltiplas fontes." },
        { q: "Lesão no fascículo arqueado causa:", options: ["Afasia motora de Broca","Afasia de condução","Surdez neurossensorial","Cegueira cortical total"], answer: 1, explain: "Desconecta Wernicke de Broca — compreende mas não repete bem." },
        { q: "Astereognosia é agnosia:", options: ["Do tipo visual","Do tipo auditiva","Do tipo tátil","Do tipo olfativa"], answer: 2, explain: "Incapacidade de reconhecer objetos pelo tato com sensibilidade preservada." },
        { q: "O córtex parietal posterior participa de:", options: ["Audição primária localizada","Atenção espacial integrada","Produção motora da fala","Visão de cores isolada"], answer: 1, explain: "Parietal posterior integra informação espacial e atenção." }
      ]
    }
  },
  {
    id: 4,
    title: "Memória e Funções Executivas",
    icon: "🧩",
    color: "#10b981",
    description: "Hipocampo, amnésias, córtex pré-frontal",
    lessons: [
      {
        title: "Classificação da Memória",
        content: `<ul>
          <li><strong>Sensorial → Curto prazo → Longo prazo</strong></li>
          <li><strong>Declarativa</strong> (episódica + semântica) vs <strong>Procedimental</strong></li>
          <li><strong>Memória de trabalho</strong> (Baddeley): executiva, loop fonológico, esboço visuoespacial</li>
          <li><strong>Implícita vs Explícita</strong></li>
        </ul>
        <div class="highlight"><strong>Episódica:</strong> eventos pessoais | <strong>Semântica:</strong> conhecimentos gerais | <strong>Procedimental:</strong> habilidades</div>`
      },
      {
        title: "Substratos Neurais da Memória",
        content: `<p><strong>Hipocampo:</strong> Consolidação de memórias declarativas (episódicas e semânticas).</p>
        <p><strong>Amígdala:</strong> Modulação emocional da memória — memórias emocionais mais fortes.</p>
        <p><strong>Córtex:</strong> Armazenamento de longo prazo (memórias distribuídas).</p>
        <p><strong>Modelo de memória de trabalho:</strong> Córtex pré-frontal como centro executivo.</p>`
      },
      {
        title: "Consolidação e Amnésias",
        content: `<p><strong>Consolidação:</strong> Processo pelo qual memórias instáveis se tornam permanentes (hipocampo → córtex).</p>
        <p><strong>Reconsolidação:</strong> Memórias reativadas podem ser modificadas antes de reconsolidar.</p>
        <ul>
          <li><strong>Anterógrada:</strong> não forma novas memórias (H.M.)</li>
          <li><strong>Retrógrada:</strong> perde memórias anteriores ao evento</li>
          <li><strong>Síndrome de Korsakoff:</strong> amnésia por deficiência de tiamina (álcool)</li>
        </ul>`
      },
      {
        title: "Funções Executivas",
        content: `<p>Processos de alto nível gerenciados pelo <strong>córtex pré-frontal</strong>:</p>
        <ul>
          <li><strong>Planejamento</strong> — organizar ações futuras</li>
          <li><strong>Inibição</strong> — suprimir respostas inadequadas</li>
          <li><strong>Flexibilidade cognitiva</strong> — alternar entre tarefas/regras</li>
          <li><strong>Memória de trabalho</strong> — manter e manipular informação</li>
          <li><strong>Tomada de decisão</strong></li>
        </ul>
        <p>Modelo de Miyake et al.: três componentes — atualização, alternância, inibição.</p>`
      },
      {
        title: "Déficits Executivos",
        content: `<p><strong>Lesões frontais:</strong> Desinibição, perseveração, dificuldade de planejamento (síndrome disexecutiva).</p>
        <p><strong>Condições neuropsiquiátricas:</strong> TDAH, esquizofrenia, TEA, demência frontotemporal.</p>
        <div class="highlight">Testes comuns: Torre de Hanoi, Stroop, Trail Making, Wisconsin Card Sorting, Fluência Verbal.</div>`
      }
    ],
    quiz: [
      { q: "O hipocampo é essencial para:", options: ["Coordenação motora fina dos dedos","Consolidação da memória declarativa","Processamento visual primário occipital","Produção articulada da fala motora"], answer: 1, explain: "O hipocampo consolida novas memórias episódicas e semânticas." },
      { q: "Amnésia retrógrada significa:", options: ["Não formar novas memórias","Perder memórias anteriores ao evento","Esquecer apenas nomes de rostos","Perder habilidades motoras aprendidas"], answer: 1, explain: "Retrógrada = perda de memórias formadas antes da lesão/dano." },
      { q: "Funções executivas estão principalmente no:", options: ["Lobo occipital posterior","Córtex pré-frontal dorsal","Cerebelo hemisférico","Tronco encefálico medular"], answer: 1, explain: "O PFC é o centro das funções executivas." },
      { q: "Memória procedimental é do tipo:", options: ["Declarativa e explícita","Implícita de habilidades","Sensorial de curto prazo","Episódica autobiográfica"], answer: 1, explain: "Procedimental = habilidades como andar de bicicleta, implícita." },
      { q: "Síndrome de Korsakoff está associada a:", options: ["Trauma cranioencefálico grave","Deficiência de tiamina por álcool","Crises epilépticas recorrentes","Transtorno do espectro autista"], answer: 1, explain: "Korsakoff resulta de deficiência de tiamina, comum em alcoolismo crônico." }
    ],
    matching: [
      { term: "Hipocampo", match: "Consolidação de memória" },
      { term: "Amígdala", match: "Memória emocional" },
      { term: "Córtex pré-frontal", match: "Funções executivas" },
      { term: "Amnésia anterógrada", match: "Não forma novas memórias" },
      { term: "Memória de trabalho", match: "Manter info temporariamente" }
    ],
    boss: {
      name: "Titã da Cognição",
      desc: "Enfrente o desafio final sobre memória e funções executivas!",
      questions: [
        { q: "Perseveração em lesões frontais significa:", options: ["Esquecer todos os eventos","Repetir respostas inadequadamente","Falar de forma excessivamente fluente","Ver imagens sem estímulo real"], answer: 1, explain: "Perseveração = dificuldade em mudar de resposta/comportamento." },
        { q: "Teste de Stroop avalia principalmente:", options: ["Memória visual de curto prazo","Inibição e controle atencional","Processamento auditivo básico","Motricidade grossa e equilíbrio"], answer: 1, explain: "Stroop mede interferência e capacidade de inibir resposta automática." },
        { q: "Reconsolidação ocorre quando:", options: ["A memória é formada pela 1ª vez","Memória reativada pode ser alterada","O hipocampo é removido cirurgicamente","A pessoa entra em sono profundo"], answer: 1, explain: "Ao reativar uma memória, ela entra em estado labil antes de reconsolidar." },
        { q: "Memória semântica armazena:", options: ["Eventos autobiográficos pessoais","Conhecimentos gerais e fatos","Habilidades motoras automatizadas","Sensações momentâneas sensoriais"], answer: 1, explain: "Semântica = conhecimento factual descontextualizado." },
        { q: "Flexibilidade cognitiva é testada por:", options: ["Teste de Raven progressivo","Wisconsin Card Sorting Test","Audiometria tonal liminar","Eletroencefalograma de repouso"], answer: 1, explain: "WCST avalia alternância entre categorias — flexibilidade cognitiva." }
      ]
    }
  },
  {
    id: 5,
    title: "Transtornos do Neurodesenvolvimento",
    icon: "🌱",
    color: "#f59e0b",
    description: "TDAH, dislexia, discalculia e TEA",
    lessons: [
      {
        title: "TDAH — Déficit de Atenção e Hiperatividade",
        content: `<p><strong>Características:</strong> Desatenção, hiperatividade e impulsividade (DSM-5).</p>
        <p><strong>Neurobiologia:</strong> Disfunção em circuitos fronto-estriatais e fronto-cerebelares; déficits de funções executivas.</p>
        <p><strong>Neuroimagem:</strong> Redução de volume em PFC, gânglios da base; alterações de conectividade.</p>
        <div class="highlight">Perfil neuropsicológico: dificuldade de atenção sustentada, memória de trabalho, inibição e planejamento.</div>`
      },
      {
        title: "Dislexia e Discalculia",
        content: `<p><strong>Dislexia (TEA de leitura):</strong> Dificuldade específica em decodificação/leitura apesar de inteligência preservada.</p>
        <p>Alterações em áreas temporo-parietais esquerdas (processamento fonológico).</p>
        <p><strong>Discalculia (TEA de matemática):</strong> Dificuldade em compreensão de quantidades, cálculo e conceitos numéricos.</p>
        <p>Associada a disfunção em parietal (intra-parietal) e rede fronto-parietal.</p>`
      },
      {
        title: "Transtorno do Espectro Autista (TEA)",
        content: `<p><strong>Características centrais (DSM-5):</strong></p>
        <ul>
          <li>Déficits na comunicação social e interação</li>
          <li>Padrões restritos e repetitivos de comportamento/interesses</li>
        </ul>
        <p><strong>Neuropsicologia:</strong> Dificuldades em cognição social, teoria da mente, flexibilidade cognitiva; habilidades visuoespaciais podem ser preservadas ou superiores.</p>
        <p>Neurobiologia: conectividade atípica, alterações em amígdala, PFC e cerebelo.</p>`
      },
      {
        title: "Avaliação Neuropsicológica no Neurodesenvolvimento",
        content: `<p>A avaliação considera:</p>
        <ul>
          <li>Histórico desenvolvimental e escolar</li>
          <li>Funções cognitivas específicas afetadas</li>
          <li>Q.I. e perfil de habilidades/desabilitades</li>
          <li>Comorbidades (TDAH + dislexia, TEA + TDAH)</li>
        </ul>
        <div class="highlight">Objetivo: diagnóstico diferencial, orientação escolar e plano de intervenção/reabilitação.</div>`
      },
      {
        title: "Intervenção e Reabilitação",
        content: `<p><strong>TDAH:</strong> Psicoeducação, treino atencional, organização, medicação (psiquiatra).</p>
        <p><strong>Dislexia:</strong> Intervenção fonológica, reforço multissensorial.</p>
        <p><strong>TEA:</strong> ABA, terapia ocupacional, fonoaudiologia, suporte escolar.</p>
        <p>Neuroplasticidade permite ganhos com intervenção precoce e contínua.</p>`
      }
    ],
    quiz: [
      { q: "TDAH envolve disfunção principalmente em circuitos:", options: ["Occipitais de processamento visual","Fronto-estriatais de atenção","Auditivos do tronco encefálico","Olfativos do bulbo olfatório"], answer: 1, explain: "TDAH afeta circuitos fronto-estriatais e funções executivas." },
      { q: "Dislexia está associada a dificuldade em:", options: ["Cálculo e raciocínio numérico escolar","Decodificação e leitura de palavras","Coordenação motora global e postura","Regulação emocional básica social"], answer: 1, explain: "Dislexia = transtorno específico de aprendizagem da leitura." },
      { q: "Teoria da mente está comprometida em:", options: ["Transtorno de déficit de atenção","Transtorno do espectro autista","Dislexia do desenvolvimento","Discalculia isolada"], answer: 1, explain: "TEA envolve dificuldade em compreender estados mentais alheios." },
      { q: "Discalculia envolve dificuldade com:", options: ["Leitura e compreensão textual","Conceitos numéricos e cálculo","Produção articulada da fala","Coordenação motora grossa"], answer: 1, explain: "Discalculia = TEA específico de matemática." },
      { q: "Perfil neuropsicológico do TDAH inclui déficit de:", options: ["Visão periférica e contraste","Atenção sustentada e inibição","Audição tonal e localização","Percepção tátil discriminativa"], answer: 1, explain: "Desatenção, impulsividade e déficits executivos caracterizam o TDAH." }
    ],
    matching: [
      { term: "TDAH", match: "Desatenção e impulsividade" },
      { term: "Dislexia", match: "Dificuldade de leitura" },
      { term: "Discalculia", match: "Dificuldade com números" },
      { term: "TEA", match: "Déficits sociais e repetitividade" },
      { term: "Teoria da mente", match: "Compreender mentes alheias" }
    ],
    boss: {
      name: "Guardião do Desenvolvimento",
      desc: "Prove que entende os transtornos do neurodesenvolvimento!",
      questions: [
        { q: "Comorbidade comum entre TDAH e:", options: ["TEA e dislexia associados","Cegueira congênita total","Surdez congênita profunda","Paralisia cerebral grave"], answer: 0, explain: "TDAH frequentemente coocorre com TEA, dislexia e discalculia." },
        { q: "Dislexia envolve alterações em áreas:", options: ["Occipitais do hemisfério direito visual","Temporo-parietais do hemisfério esquerdo","Cerebelares de coordenação motora fina","Motoras primárias do membro superior"], answer: 1, explain: "Processamento fonológico depende de regiões temporo-parietais esquerdas." },
        { q: "Flexibilidade cognitiva reduzida é comum em:", options: ["Demência avançada de início tardio","Transtorno do espectro autista","Alexia pura sem agrafia associada","Surdez neurossensorial congênita"], answer: 1, explain: "TEA envolve rigidez cognitiva e dificuldade de alternância." },
        { q: "Intervenção precoce se baseia em:", options: ["Neuroplasticidade cerebral","Imobilização motora prolongada","Isolamento social total","Restrição alimentar severa"], answer: 0, explain: "Neuroplasticidade permite reorganização e ganhos com estimulação." },
        { q: "TEA caracteriza-se por déficits em:", options: ["Motricidade grossa exclusiva","Comunicação social e interação","Audição primária localizada","Visão periférica e contraste"], answer: 1, explain: "Déficits sociais e padrões repetitivos são critérios centrais do TEA." }
      ]
    }
  },
  {
    id: 6,
    title: "Quadros Neurológicos",
    icon: "🏥",
    color: "#ef4444",
    description: "Demências, epilepsia, AVE e neuroplasticidade",
    lessons: [
      {
        title: "Demências — Visão Geral",
        content: `<p><strong>Demência</strong> = síndrome de declínio cognitivo adquirido que interfere na funcionalidade.</p>
        <p>Tipos principais:</p>
        <ul>
          <li><strong>Alzheimer (DA)</strong> — memória, linguagem, funções visuoespaciais</li>
          <li><strong>Frontotemporal (DFT)</strong> — personalidade, linguagem, funções executivas</li>
          <li><strong>Vascular</strong> — padrão escalonado, relação com AVCs</li>
          <li><strong>Parkinson</strong> — demência subcortical, lentificação, rigidez cognitiva</li>
        </ul>`
      },
      {
        title: "Doença de Alzheimer",
        content: `<p>Demência mais comum. Início insidioso com <strong>comprometimento de memória episódica</strong>.</p>
        <p>Neuropatologia: placas de beta-amiloide e emaranhados neurofibrilares de tau.</p>
        <p>Progressão: memória → linguagem → funções visuoespaciais → funções executivas.</p>
        <div class="highlight">Atrofia inicial: hipocampo e córtex entorrinal. Diagnóstico: clínico + neuropsicológico + neuroimagem.</div>`
      },
      {
        title: "Epilepsia",
        content: `<p>Distúrbio neurológico caracterizado por crises recorrentes (descargas elétricas anormais).</p>
        <p><strong>Impacto neuropsicológico:</strong> Depende do foco — temporal (memória), frontal (executivo), etc.</p>
        <p>Epilepsia temporal mesial: comprometimento de memória verbal (hemisfério dominante).</p>
        <p>Interictal: possíveis déficits cognitivos entre crises; efeitos de medicação anticonvulsivante.</p>`
      },
      {
        title: "Acidente Vascular Encefálico (AVE)",
        content: `<p><strong>Isquêmico (85%):</strong> Obstrução de artéria → falta de oxigênio.</p>
        <p><strong>Hemorrágico (15%):</strong> Ruptura de vaso → sangramento.</p>
        <p>Sequelas dependem da <strong>localização e extensão</strong> da lesão:</p>
        <ul>
          <li>Frontal → executivo, comportamento</li>
          <li>Parietal → neglect, praxia, atenção espacial</li>
          <li>Temporal → memória, linguagem</li>
          <li>Occipital → déficits visuais</li>
        </ul>`
      },
      {
        title: "Neuroplasticidade",
        content: `<p><strong>Neuroplasticidade</strong> = capacidade do cérebro de se reorganizar estrutural e funcionalmente.</p>
        <ul>
          <li><strong>Desenvolvimento:</strong> Podagem sináptica, mielinização</li>
          <li><strong>Aprendizagem:</strong> Potenciação de longo prazo (LTP)</li>
          <li><strong>Lesão:</strong> Reorganização compensatória, neurogênese limitada</li>
          <li><strong>Envelhecimento:</strong> Plasticidade reduzida mas presente — "use it or lose it"</li>
        </ul>
        <div class="highlight">Base da reabilitação neuropsicológica: estimular circuitos alternativos e fortalecer reservas cognitivas.</div>`
      }
    ],
    quiz: [
      { q: "A Doença de Alzheimer inicia tipicamente com:", options: ["Mudança abrupta de personalidade","Comprometimento da memória episódica","Crises convulsivas generalizadas","Paralisia unilateral do corpo"], answer: 1, explain: "DA classicamente começa com memória episódica e aprendizado de novas informações." },
      { q: "Demência frontotemporal afeta precocemente:", options: ["Processamento visual complexo","Personalidade e funções executivas","Processamento auditivo primário","Motricidade fina dos dedos"], answer: 1, explain: "DFT compromete comportamento, personalidade e/ou linguagem precocemente." },
      { q: "AVE isquêmico resulta de:", options: ["Ruptura e sangramento vascular","Obstrução do fluxo sanguíneo","Trauma contuso craniano","Infecção do tecido nervoso"], answer: 1, explain: "Isquemia = falta de fluxo sanguíneo por obstrução." },
      { q: "Neuroplasticidade permite:", options: ["Apenas deterioração progressiva","Reorganização cerebral após lesão","Imobilidade permanente das sinapses","Perda total irreversível de função"], answer: 1, explain: "Plasticidade = capacidade de adaptação e reorganização neural." },
      { q: "Placas de beta-amiloide são marcadores de:", options: ["Transtorno de déficit de atenção","Doença de Alzheimer","Transtorno do espectro autista","Dislexia do desenvolvimento"], answer: 1, explain: "Placas amiloides e emaranhados de tau são hallmarks da DA." }
    ],
    matching: [
      { term: "Alzheimer", match: "Memória episódica inicial" },
      { term: "Demência Frontotemporal", match: "Alteração de personalidade" },
      { term: "AVE isquêmico", match: "Obstrução vascular" },
      { term: "Epilepsia temporal", match: "Déficit de memória verbal" },
      { term: "Neuroplasticidade", match: "Reorganização cerebral" }
    ],
    boss: {
      name: "Lorde Neurológico",
      desc: "O chefe final! Prove que domina todo o semestre!",
      questions: [
        { q: "Demência vascular caracteriza-se por:", options: ["Início lento e progressivo contínuo","Declínio em degraus com AVCs","Ocorrência exclusiva em jovens","Ausência de prejuízo cognitivo"], answer: 1, explain: "Demência vascular tem relação com AVCs múltiplos — declínio em degraus." },
        { q: "Reabilitação neuropsicológica se baseia em:", options: ["Repouso absoluto e isolamento","Neuroplasticidade e estimulação","Restrição total de estímulos","Imobilização prolongada do paciente"], answer: 1, explain: "Estimular circuitos e estratégias compensatórias aproveita plasticidade." },
        { q: "Em Parkinson, a demência é frequentemente do tipo:", options: ["Cortical com afasia precoce","Subcortical com lentificação","Puramente visual e espacial","Puramente auditiva e verbal"], answer: 1, explain: "Demência parkinsoniana tem perfil subcortical: lentificação, rigidez cognitiva." },
        { q: "Reserva cognitiva refere-se a:", options: ["Capacidade de compensar lesões","Impossibilidade de aprender novas tarefas","Ausência completa de neurônios","Determinação exclusivamente genética"], answer: 0, explain: "Maior reserva = mais capacidade de compensar dano cerebral." },
        { q: "A avaliação neuropsicológica em demência busca:", options: ["Medir apenas o quociente de inteligência","Mapear perfil cognitivo e diagnóstico","Avaliar somente o humor e afeto","Investigar exclusivamente a visão"], answer: 1, explain: "Mapeia déficits, diferencia tipos de demência e estabelece linha de base." }
      ]
    }
  }
];

const ACHIEVEMENTS = [
  { id: "first_lesson", icon: "📖", title: "Primeira Lição", desc: "Complete sua primeira lição", xp: 0 },
  { id: "first_quiz", icon: "✅", title: "Quizzer", desc: "Complete seu primeiro quiz", xp: 0 },
  { id: "first_boss", icon: "⚔️", title: "Caçador de Chefes", desc: "Derrote seu primeiro chefe", xp: 0 },
  { id: "unit1", icon: "🔬", title: "Neurocientista", desc: "Complete a Unidade 1", xp: 0 },
  { id: "unit3", icon: "💬", title: "Linguista Neural", desc: "Complete a Unidade 3", xp: 0 },
  { id: "unit6", icon: "🏆", title: "Mestre Neuro", desc: "Complete todas as 6 unidades", xp: 0 },
  { id: "streak5", icon: "🔥", title: "Em Chamas", desc: "Acerte 5 questões seguidas", xp: 0 },
  { id: "streak10", icon: "💎", title: "Imparável", desc: "Acerte 10 questões seguidas", xp: 0 },
  { id: "no_life_lost", icon: "🛡️", title: "Invencível", desc: "Complete um quiz sem perder vidas", xp: 0 },
  { id: "level5", icon: "⭐", title: "Veterano", desc: "Alcance o nível 5", xp: 0 },
  { id: "level10", icon: "👑", title: "Lenda Neural", desc: "Alcance o nível 10", xp: 0 },
  { id: "match_master", icon: "🔗", title: "Conector", desc: "Complete um desafio de associação", xp: 0 }
];

const FLASHCARDS = UNITS.flatMap(u =>
  u.lessons.map(l => ({ front: l.title, back: l.content.replace(/<[^>]+>/g, ' ').replace(/\s+/g, ' ').trim().substring(0, 200) + '...', unit: u.title }))
);
